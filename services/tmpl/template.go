package tmpl

import (
	"github.com/haborhuang/weixin-comp-go-demo/services/client"
	"github.com/haborhuang/weixin-comp-go-demo/services/auth"

	"fmt"
	"log"
)

func GetPrivTmpls(appId string) (*client.TemplatesResp, error) {
	tok, err := auth.GetAccessToken(appId)
	if nil != err {
		return nil, fmt.Errorf("Get access token error: %v", err)
	}

	return client.GetClient().GetPrivTemplates(tok)
}

func SendTmplMsg(appId, tmplId, to string, params map[string]*client.TmplData) error {
	token, err := auth.GetAccessToken(appId)
	if nil != err {
		return fmt.Errorf("Get access token error: %v", err)
	}

	res, err := client.GetClient().SendTmplMsg(token, tmplId, to, params)
	if nil != err {
		return fmt.Errorf("Call send api error: %v", err)
	}
	if err := res.Err(); nil != err {
		return fmt.Errorf("Api response error: %v", err)
	}

	log.Println("Sent message %d", res.MsgId)

	return nil
}

func GetAppToken(appId string) (string, error) {
	return auth.GetAccessToken(appId)
}