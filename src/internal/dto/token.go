package dto

type TokenDto struct {
	AccessToken   string `json:"access-token"`
	RefreshToken  string `json:"refresh-token"`
	AccessExpire  int64  `json:"access-expire"`
	RefreshExpire int64  `json:"refresh-expire"`
}
