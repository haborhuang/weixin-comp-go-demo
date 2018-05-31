package auth

import (
	"fmt"
	"time"
	"log"
	"encoding/json"

	"github.com/haborhuang/weixin-comp-go-demo/services/client"
)

func authRefreshTokenKey(appId string) string {
	return keyPrefix + "authorizer_refresh_token." + appId
}

func authAccessTokenKey(appId string) string {
	return keyPrefix + "authorizer_access_token." + appId
}

type authAccessToken struct {
	Token string `json:"token"`
	ExpiredAt int64 `json:"expired_at"`
}

func QueryAuth(authCode string) error {
	tok, err := GetCompAccessToken()
	if nil != err {
		return fmt.Errorf("Get component access token error: %v", err)
	}
	if tok == "" {
		return fmt.Errorf("Empty component access token")
	}

	res, err := wxClient.QueryAuth(tok, authCode)
	if nil != err {
		return fmt.Errorf("Call api_query_auth api error: %v", err)
	}
	if res.ErrRes != nil {
		return res.ErrRes
	}

	log.Printf("Got func info of app %s:\n", res.AuthorizationInfo.Appid)
	for _, f := range res.AuthorizationInfo.FuncInfo {
		log.Println("  funcscope_category:", f.FuncscopeCategory.Id)
	}

	return saveAuthKeyInfo(res.AuthorizationInfo.Appid, &res.AuthorizationInfo.AuthToken)
}

func saveAuthKeyInfo(appId string, at *client.AuthToken) error {
	if err := redisClient.Set(authRefreshTokenKey(appId), at.RefreshToken, 0).
		Err(); nil != err {

		return fmt.Errorf("Save refresh token error: %v", err)
	}

	data, _ := json.Marshal(authAccessToken{
		Token: at.AccessToken,
		ExpiredAt: time.Now().UTC().Unix() + at.ExpiresIn,
	})

	redisClient.Set(authAccessTokenKey(appId), string(data), 0)
	return nil
}

func GetAccessToken(appId string) (string, error) {
	ckey := authAccessTokenKey(appId)
	data, err := redisClient.Get(ckey).Result()
	if nil != err {
		return "", fmt.Errorf("Lookup %s from redis error: %v", ckey, err)
	}

	var at authAccessToken
	json.Unmarshal([]byte(data), &at)

	token := at.Token
	if time.Now().UTC().Unix() >= at.ExpiredAt {
		ckey := authRefreshTokenKey(appId)
		refreshToken, err := redisClient.Get(ckey).Result()
		if nil != err {
			return "", fmt.Errorf("Lookup %s from redis error: %v", ckey, err)
		}

		compToken, err := GetCompAccessToken()
		if nil != err {
			return "", fmt.Errorf("Get component access token error: %v", err)
		}

		res, err := wxClient.GetAuthToken(compToken, appId, refreshToken)
		if nil != err {
			return "", fmt.Errorf("Call api_authorizer_token api error: %v", err)
		}
		if res.ErrRes != nil {
			return "", res.ErrRes
		}

		token = res.AccessToken

		go func() {
			if err := saveAuthKeyInfo(appId, &res.AuthToken); nil != err {
				log.Println("save auth key info error:", err)
			}
		}()
	}

	return token, nil
}
