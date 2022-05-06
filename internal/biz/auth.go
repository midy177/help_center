package biz

import (
	"github.com/dgrijalva/jwt-go"
	"help_center/internal/conf"
	"net/http"
	"time"
)

func AdminLogin(d *LoginData) (int, *BaseJson) {
	if d.Name == conf.GetCfg().Admin.User && d.Password == conf.GetCfg().Admin.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = conf.GetCfg().Admin.User
		exp := time.Now().Add(time.Hour * 6).Unix()
		claims["exp"] = exp
		// Generate encoded token and send it as response.
		secret := conf.GetCfg().Jwt.Secret
		tk, err := token.SignedString([]byte(secret)) //密钥
		if err != nil {
			return http.StatusUnauthorized, &BaseJson{Code: 0, Data: err.Error()}
		}
		return http.StatusOK, &BaseJson{Code: 1, Data: tk}
	} else {
		return http.StatusUnauthorized, &BaseJson{Code: 0, Data: "用户名或密码错误"}
	}
}
