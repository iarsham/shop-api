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
	"net/http"
)

type UserService struct {
	logs  *common.Logger
	db    *gorm.DB
	redis *redis.Client
}

func NewUserService(logs *common.Logger) *UserService {
	return &UserService{
		logs:  logs,
		db:    db.GetDB(),
		redis: db.GetRedis(),
	}
}

func (s *UserService) RegisterLoginByPhone(req *dto.RegisterLoginRequest) error {
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

func (s *UserService) GetUserByIP(r *http.Request) (*models.Users, error) {
	var user models.Users
	phone, _ := s.redis.Get(ctx, utils.ClientIP(r)).Result()
	err := s.db.Where("phone=?", phone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		common.LogWarning(s.logs, err)
		return nil, err
	}
	return &user, nil
}

func (s *UserService) ActivateUser(r *http.Request) error {
	ip := utils.ClientIP(r)
	phone, _ := s.redis.Get(ctx, ip).Result()
	user, _ := s.GetUserByPhone(phone)
	if !user.IsActive {
		user.IsActive = true
		if err := s.db.Save(&user).Error; err != nil {
			common.LogWarning(s.logs, err)
			return err
		}
	}
	s.redis.Del(ctx, ip, phone)
	return nil
}

func (s *UserService) UpdateUserByID(id string, data *dto.UpdateUserRequest) (*models.Users, error) {
	user, _ := s.GetUserByID(id)
	user.FirstName = data.FirstName
	user.LastName = data.LastName
	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
