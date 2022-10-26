package pogo

type BindAccPsw struct {
	Uuid         string `json:"uuid" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Captcha_code string `json:"captcha_code"`
	Captcha_key  string `json:"captcha_key"`
}
