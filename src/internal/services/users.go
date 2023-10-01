package services

import (
	"errors"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/dto"
	"github.com/iarsham/shop-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	logs *common.Logger
	db   *gorm.DB
}

func NewUserService(logs *common.Logger) *UserService {
	return &UserService{
		logs: logs,
		db:   db.GetDB(),
	}
}

func (s *UserService) RegisterByPhone(req *dto.RegisterRequest) error {
	if exists := s.userExistsByPhone(req.Phone); exists {
		return errors.New("user with this phone already exists")
	}

	bytePass := []byte(common.GeneratePassword())
	hashPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	user := models.Users{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Password:  string(hashPass),
	}

	err := s.db.Create(&user).Error
	if err != nil {
		s.logs.Warn(err.Error())
		return err
	}

	return nil
}

func (s *UserService) userExistsByPhone(phone string) bool {
	var user models.Users
	err := s.db.Where("phone=?", phone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		s.logs.Warn(err.Error())
		return false
	}
	return true
}
