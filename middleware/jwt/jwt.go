package jwt

import (
	"alger/common/global"
	"alger/common/response"
	redisUtil "alger/common/utils/redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/jsonx"
	"net/http"
)

const (
	TOKEN       = "IEC-Token"
	Lang        = "accept-language"
	ProductCode = "IEC-ProductCode"
)

// ParseToken jwt解析
func ParseToken(next http.Handler, accessSecret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("IEC-Token")
		if token == "" {
			response.JwtFailWithMessageId(w, "请重新登录")
			return
		}
		j := NewJWT(accessSecret)
		//校验token
		validToken := j.verificationToken(token)
		//parse 解析token包含的信息
		claims := j.parse(validToken)
		bytes, err := jsonx.Marshal(claims)
		if err != nil {
			global.Logger.Error("jwt字符串化失败" + err.Error())
			response.FailWithMessageId(w, err.Error())
			return
		}
		r.Header.Set("IEC-claims", string(bytes))
		next.ServeHTTP(w, r)
	},
	)
}

/**
 * @Description: 微信验证token
 * @param next
 * @param accessSecret
 * @return http.Handler
 */
func WxParseToken(next http.Handler, accessSecret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("IEC-Token")
		if token == "" {
			response.JwtFailWithMessageId(w, "请重新登录")
			return
		}
		j := NewJWT(accessSecret)
		//校验token
		validToken := j.verificationToken(token)
		//parse 解析token包含的信息
		claims := j.parse(validToken)
		//获取token是否存在
		redisKey := redisUtil.CreateKey(false, "")("user", "login", claims.OpenId)
		exists, err := global.Redis.Exists(redisKey)
		if err != nil {
			response.JwtFailWithMessageId(w, "请重新登录")
			return
		}
		//如果存在直接返回key里面储存的token
		if !exists {
			response.JwtFailWithMessageId(w, "请重新登录")
			return
		}
		bytes, err := jsonx.Marshal(claims)
		if err != nil {
			global.Logger.Error("jwt字符串化失败" + err.Error())
			response.FailWithMessageId(w, err.Error())
			return
		}
		r.Header.Set("IEC-claims", string(bytes))
		next.ServeHTTP(w, r)
	},
	)
}

type Claims struct {
	Exp        int64
	Iat        int64
	UserId     int64
	Username   string
	TenantCode string
	Lang       string
	jwt.StandardClaims
	OpenId string
}

type JWT struct {
	SigningKey []byte
}

func NewJWT(secret string) *JWT {
	return &JWT{
		[]byte(secret),
	}
}

// 校验token
func (j *JWT) verificationToken(tokenString string) *jwt.Token {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	return token
}

// parse 解析 token
func (j *JWT) parse(token *jwt.Token) *Claims {
	claims, _ := token.Claims.(*Claims)
	return claims
}
