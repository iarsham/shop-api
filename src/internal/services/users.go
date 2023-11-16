package services

import (
	"errors"
	"fmt"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/models"
	"github.com/iarsham/shop-api/internal/utils"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
)

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

func (s *UserService) RegisterLoginByPhone(req *dto.RegisterRequest) error {
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(utils.GeneratePassword()), bcrypt.DefaultCost)
	user := models.Users{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Password:  string(hashPass),
	}
	query := s.db.Create(&user)
	if query.RowsAffected == 1 {
		s.logs.Info(fmt.Sprintf("New User Created. Phone : %s", req.Phone))
	}
	if query.Error != nil {
		s.logs.Warn(query.Error.Error())
		return query.Error
	}
	return nil
}

func (s *UserService) GetUserByPhone(phone string) (*models.Users, bool) {
	var user models.Users
	err := s.db.Where("phone=?", phone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		common.LogWarning(s.logs, err)
		return nil, false
	}
	return &user, true
}

func (s *UserService) GetUserByID(id string) (*models.Users, bool) {
	var user models.Users
	err := s.db.Where("id=?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		common.LogWarning(s.logs, err)
		return nil, false
	}
	return &user, true
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
