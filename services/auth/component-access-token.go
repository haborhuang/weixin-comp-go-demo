package auth

import (
	"github.com/haborhuang/weixin-comp-go-demo/services/client"

	"time"
	"log"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

var wxClient *client.Client = client.GetClient()

func init() {
	go refreshcompAccessToken()
}

func refreshcompAccessToken() {
	for {
		func (){
			after := 10 * time.Second
			defer func() {
				time.Sleep(after)
			}()

			t, err := getAndDecodeCompAccessToken()
			if nil != err {
				log.Println("Lookup component access token error:", err)
				return
			}

			now := time.Now().UTC()
			if t.ExpiredAt - now.Unix() > 600 {
				after = 5 * time.Minute
				return
			}

			// Will expire in 10 minutes
			// Refresh the token

			// get ticket
			ticket, err := GetVerifyTicket()
			if nil != err {
				if err == redis.Nil {
					after = time.Minute
				}
				log.Println("Get verify ticket error:", err)
				return
			}
			if ticket == "" {
				log.Println("Empty verify ticket")
				return
			}

			// call api
			res, err := wxClient.GetCompAccessToken(ticket)
			if nil != err {
				log.Println("Call api_component_token api error:", err)
				return
			}

			if res.ErrRes != nil {
				log.Println("API api_component_token response error:", res.ErrRes)
				return
			}

			// save token
			data, _ := json.Marshal(compAccessToken{
				Token: res.ComponentAccessToken,
				ExpiredAt: now.Unix() + res.ExpiresIn,
			})
			redisClient.Set(compAccessTokenKey, string(data), 0)
			after = 5 * time.Minute
		}()
	}
}

type compAccessToken struct {
	Token string `json:"token"`
	ExpiredAt int64 `json:"expired_at"`
}

const compAccessTokenKey = keyPrefix + "comp_access_token"

func GetCompAccessToken() (token string, err error) {
	t, err := getAndDecodeCompAccessToken()
	if nil != t {
		token = t.Token
	}
	return
}

func getAndDecodeCompAccessToken() (*compAccessToken, error) {
	data, err := redisClient.Get(compAccessTokenKey).Result()
	if nil != err && err != redis.Nil {
		return nil, fmt.Errorf("Lookup '%s' from redis error: %v", compAccessTokenKey, err)
	}

	if err == redis.Nil {
		return &compAccessToken{}, nil
	}

	var t *compAccessToken
	json.Unmarshal([]byte(data), &t)

	return t, nil
}