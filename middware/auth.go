package middware

import (
	"time"

	"github.com/iris-contrib/middleware/jwt"
)

var SigKey = []byte("qingfengliuyun")

func GenerateToken(uid uint) string {
	// claims := PlayLoad{Uid: uid}
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Uid": uid,
		// 签发人
		"iss": "iris",
		// 签发时间
		"iat": time.Now().Unix(),
		// 设定过期时间，便于测试，设置1分钟过期
		"exp": time.Now().Add(7 * 24 * 60 * time.Minute * time.Duration(1)).Unix(),
	})
	tokenString, _ := token.SignedString(SigKey)

	return tokenString
}
