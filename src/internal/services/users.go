package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/iarsham/shop-api/internal/utils"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/models"
)

var ctx = context.Background()

type UserService struct {
	logs  *common.Logger
	db    *gorm.DB
	redis *redis.Client
	token *TokenService
}

func NewUserService(logs *common.Logger) *UserService {
	return &UserService{
		logs:  logs,
		db:    db.GetDB(),
		redis: db.GetRedis(),
		token: NewTokenService(logs),
	}
}

func (s *UserService) RegisterByPhone(req *dto.RegisterRequest) error {
	bytePass := []byte(utils.GeneratePassword())
	hashPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	user := models.Users{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Password:  string(hashPass),
	}

	err := s.db.Create(&user).Error
	if err != nil {
		common.LogError(s.logs, err)
		return err
	}

	return nil
}

func (s *UserService) UserExistsByPhone(phone string) bool {
	var user models.Users
	err := s.db.Where("phone=?", phone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		common.LogWarning(s.logs, err)
		return false
	}
	return true
}

func (s *UserService) VerifyUser(phone string) (*dto.TokenDto, error) {
	var user models.Users
	s.db.Where("phone=?", phone).First(&user)
	if !user.IsActive {
		user.IsActive = true
		if err := s.db.Save(&user).Error; err != nil {
			common.LogWarning(s.logs, err)
			return nil, err
		}
	}
	token, err := s.token.GenerateToken(strconv.Itoa(int(user.ID)), phone)
	if err != nil {
		return nil, err
	}
	s.redis.Del(ctx, phone)
	return token, nil
}

func (s *UserService) CheckOtpEqual(phone, code string) bool {
	otp, _ := s.redis.Get(ctx, phone).Result()
	return otp == code
}

func (s *UserService) CheckOtpExpire(phone string) bool {
	_, err := s.redis.Get(ctx, phone).Result()
	return errors.Is(err, redis.Nil)
}

func (s *UserService) SendOTP(phone string) {
	client := twilio.NewRestClient()
	code := s.generateOTP()

	params := &openapi.CreateMessageParams{}
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetTo(phone)
	params.SetBody(fmt.Sprintf("Hi\nVerification code is %s", code))

	_, err := client.Api.CreateMessage(params)
	common.LogError(s.logs, err)

	err = s.redis.Set(ctx, phone, code, 10*time.Minute).Err()
	common.LogError(s.logs, err)

	s.logs.Info(fmt.Sprintf("OTP code was send to : %s", phone))
}

func (s *UserService) generateOTP() string {
	charSet := "0123456789"
	otp := make([]byte, 6)

	for i := range otp {
		otp[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(otp)
}

func (s *UserService) GetUserByID(id string) (*models.Users, error) {
	var user models.Users
	err := s.db.Where("id=?", id).First(&user).Error
	switch {
	case err == nil:
		return &user, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("user not found")
	default:
		return nil, err
	}
}

func (s *UserService) UpdateUserByID(id string, data *dto.UpdateUserRequest) (*models.Users, error) {
	user, err := s.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	user.FirstName = data.FirstName
	user.LastName = data.LastName
	err = s.db.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetClaims(token string) (map[string]any, error) {
	return s.token.GetClaims(token)
}

func (s *UserService) NewAccessToken(userID, phone string) string {
	token, _ := s.token.GenerateAccessToken(userID, phone)
	return token
}
