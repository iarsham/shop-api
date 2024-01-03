package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/iarsham/shop-api/internal/common"
	"github.com/iarsham/shop-api/internal/db"
	"github.com/iarsham/shop-api/internal/utils"
	"github.com/kavenegar/kavenegar-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	ctx             = context.Background()
	kavenegarAPI    = os.Getenv("KAVENEGAR_SMS_API")
	kavenegarClient = kavenegar.New(kavenegarAPI)
	otpExpire       = os.Getenv("OTP_EXPIRE_MIN")
)

type OtpService struct {
	logs  *common.Logger
	db    *gorm.DB
	redis *redis.Client
}

func NewOTPService(logs *common.Logger) *OtpService {
	return &OtpService{
		logs:  logs,
		db:    db.GetDB(),
		redis: db.GetRedis(),
	}
}

func (o *OtpService) SendOTP(phone string, r *http.Request) {
	code := o.generateOTP()
	ip := utils.ClientIP(r)
	exp, _ := strconv.Atoi(otpExpire)
	message := fmt.Sprintf("Hi \nAuthentication Code is : \t %s", code)
	_, err := kavenegarClient.Message.Send("2000500666", []string{phone}, message, nil)
	if err != nil {
		o.logs.Warn(err.Error())
		return
	}
	o.redis.Set(ctx, ip, phone, time.Duration(exp)*time.Minute)
	o.redis.Set(ctx, phone, code, time.Duration(exp)*time.Minute)
	o.logs.Info(fmt.Sprintf("OTP code was send to : %s", phone))
}

func (o *OtpService) generateOTP() string {
	charSet := "0123456789"
	otp := make([]byte, 6)
	for i := range otp {
		otp[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(otp)
}

func (o *OtpService) IsOtpEqual(code string, r *http.Request) bool {
	phone, _ := o.redis.Get(ctx, utils.ClientIP(r)).Result()
	otp, _ := o.redis.Get(ctx, phone).Result()
	return otp == code
}

func (o *OtpService) IsOtpExpire(r *http.Request) bool {
	phone, _ := o.redis.Get(ctx, utils.ClientIP(r)).Result()
	_, err := o.redis.Get(ctx, phone).Result()
	return errors.Is(err, redis.Nil)
}
