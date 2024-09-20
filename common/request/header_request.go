package request

import (
	"alger/common/global"
	"alger/middleware/jwt"
	"encoding/json"
	"net/http"
)

const CLAIMS = "IEC-claims"
const Lang = "zh-CN"

// GetToken 获取用户ID从toKen里面解析
func GetToken(r *http.Request) (*jwt.Claims, error) {
	claims := r.Header.Get(CLAIMS)
	var userInfo jwt.Claims
	err := json.Unmarshal([]byte(claims), &userInfo)
	if err != nil {
		global.Logger.Error("token parse err:", err)
		return nil, err
	}
	return &userInfo, nil
}

// 从headers request
func GetLanguage(r *http.Request) string {
	lang := r.Header.Get("lang")
	if lang == "" {
		return Lang
	}
	return lang
}

func GetUserId(r *http.Request) int64 {
	claims := r.Header.Get(CLAIMS)
	var userInfo jwt.Claims
	err := json.Unmarshal([]byte(claims), &userInfo)
	if err != nil {
		return 0
	}

	return userInfo.UserId
}

// 获取用户租户code从toKen里面解析
func GetUserTenantCode(r *http.Request) string {
	claims := r.Header.Get(CLAIMS)
	var userInfo jwt.Claims
	err := json.Unmarshal([]byte(claims), &userInfo)
	if err != nil {
		return ""
	}

	return userInfo.TenantCode
}

func GetProductCode(r *http.Request) string {
	productCode := r.Header.Get("IEC-ProductCode")
	if productCode == "" {
		return ""
	}
	return productCode
}

func GetPhone(r *http.Request) string {
	claims := r.Header.Get(CLAIMS)
	var userInfo jwt.Claims
	err := json.Unmarshal([]byte(claims), &userInfo)
	if err != nil {
		return ""
	}

	return userInfo.Username
}
