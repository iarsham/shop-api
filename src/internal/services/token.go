package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"os"
	"time"

	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/dto"
)

var secretKey = os.Getenv("SECRET_KEY")

type TokenService struct {
	logs *common.Logger
	db   *gorm.DB
}

func NewTokenService(logs *common.Logger) *TokenService {
	return &TokenService{
		logs: logs,
		db:   db.GetDB(),
	}
}

func (t *TokenService) GenerateToken(userID, phone string) (*dto.TokenDto, error) {
	var err error
	token := &dto.TokenDto{}

	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"phone":   phone,
		"exp":     time.Now().Add(time.Minute * 30).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	token.AccessToken, err = accessToken.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"phone":   phone,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	token.RefreshToken, err = refreshToken.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (t *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected error")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (t *TokenService) GetClaims(token string) (map[string]any, error) {
	var claimMap map[string]any
	verifyToken, err := t.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claims[k] = v
		}
		return claimMap, nil
	}
	return nil, errors.New("claims properties not found")
}
