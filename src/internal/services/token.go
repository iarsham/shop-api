package services

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"

	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
)

var (
	secretKey        = os.Getenv("SECRET_KEY")
	accessExpireEnv  = os.Getenv("ACCESS_TOKEN_EXPIRE_MIN")
	refreshExpireEnv = os.Getenv("REFRESH_TOKEN_EXPIRE_DAY")
)

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

func (t *TokenService) GenerateAccessToken(userID string, phone string) (string, error) {
	accessExp, _ := strconv.Atoi(accessExpireEnv)
	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"phone":   phone,
		"exp":     time.Now().Add(time.Minute * time.Duration(accessExp)).Unix(),
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (t *TokenService) GenerateRefreshToken(userID string) (string, error) {
	refreshExp, _ := strconv.Atoi(refreshExpireEnv)
	refreshClaims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24 * time.Duration(refreshExp)).Unix(),
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return refreshToken, nil
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
	claimMap := make(map[string]any)
	verifyToken, err := t.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, errors.New("claims properties not found")
}
