package utils

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12/middleware/jwt"
)

var SigKey = []byte("qingfengliuyun")

type PlayLoad struct {
	Uid uint
}

/**
 * @deprecated 请改用 middware\auth.go
 */
func GenerateToken(uid uint) string {

	signer := jwt.NewSigner(jwt.HS256, SigKey, 600*time.Minute)
	claims := PlayLoad{Uid: uid}

	token, err := signer.Sign(claims)
	if err != nil {
		fmt.Println(err)

	}
	s := string(token)

	return s

}
