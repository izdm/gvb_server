package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
	"gvb_server/global"
	"time"
)

func GenToken(user JwtPayload) (string, error) {
	// 获取密钥，确保是 []byte 类型
	MySecret := []byte(global.Config.Jwt.Secret)

	// 创建自定义的 Claims 结构体
	claim := CustomClaims{
		JwtPayload: user,
		StandardClaims: jwt.StandardClaims{
			// 设置过期时间，转换为 Unix 时间戳（秒）
			ExpiresAt: &jwt.Time{Time: time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))}, // 设置为 *Time 类型
			Issuer:    global.Config.Jwt.Issuer,                                                              // 设置签发人
		},
	}

	// 使用 HMAC SHA256 算法创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// 签名并返回 token 字符串
	signedToken, err := token.SignedString(MySecret)
	if err != nil {
		// 错误处理
		return "", nil
	}
	return signedToken, nil
}
