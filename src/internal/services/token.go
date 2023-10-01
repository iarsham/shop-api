package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"

	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/dto"
)

var secretKey = os.Getenv("SECRET_KEY")

type TokenService struct {
	logs *common.Logger
	db   *gorm.DB
}

func (t *TokenService) NewTokenService(logs *common.Logger) *TokenService {
	return &TokenService{
		logs: logs,
		db:   db.GetDB(),
	}
}

func (t *TokenService) GenerateToken(userID, phone string) (*dto.TokenDto, error) {
	var err error
	token := &dto.TokenDto{}
	accessClaims := jwt.MapClaims{}

	accessClaims["user_id"] = userID
	accessClaims["phone"] = phone
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	token.AccessToken, err = accessToken.SignedString([]byte(secretKey))
	token.AccessExpire = time.Now().Add(30 * time.Minute).Unix()
	if err != nil {
		return nil, err
	}
	refreshClaims := jwt.MapClaims{}
	refreshClaims["user_id"] = userID
	refreshClaims["phone"] = phone
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	token.RefreshExpire = time.Now().Add(time.Hour * 72).Unix()
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
