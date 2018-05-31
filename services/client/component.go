package client

import (
	"github.com/haborhuang/weixin-comp-go-demo/services/common"
	"net/http"
	"net/url"
)

type ComponentTokenRes struct {
	*ErrRes
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int64  `json:"expires_in"`
}

type componentTokenReq struct {
	ComponentAppid        string `json:"component_appid"`
	ComponentAppsecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}

func (c *Client) GetCompAccessToken(ticket string) (*ComponentTokenRes, error) {
	var res *ComponentTokenRes

	err := c.newResponse(
		c.newRequest().SubPath("component/api_component_token").Method(http.MethodPost).JsonBody(componentTokenReq{
			ComponentAppid:        common.GetCompAppId(),
			ComponentAppsecret:    common.GetCompAppSecret(),
			ComponentVerifyTicket: ticket,
		}),
	).intoJson(&res)

	return res, err
}

type preAuthCodeReq struct {
	ComponentAppid string `json:"component_appid"`
}

type PreAuthCodeRes struct {
	*ErrRes
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c *Client) GetPreAuthCode(compAccessToken string) (*PreAuthCodeRes, error) {
	var res *PreAuthCodeRes
	query := url.Values{}
	query.Set("component_access_token", compAccessToken)

	err := c.newResponse(
		c.newRequest().SubPath("component/api_create_preauthcode").Method(http.MethodPost).JsonBody(preAuthCodeReq{
			ComponentAppid: common.GetCompAppId(),
		}).Query(query),
	).intoJson(&res)

	return res, err
}

type queryAuthReq struct {
	ComponentAppid    string `json:"component_appid"`
	AuthorizationCode string `json:"authorization_code"`
}

type QueryAuthRes struct {
	*ErrRes
	AuthorizationInfo authInfo `json:"authorization_info"`
}

type authInfo struct {
	Appid string `json:"authorizer_appid"`
	AuthToken
	FuncInfo []*funcInfo `json:"func_info"`
}

type funcInfo struct {
	FuncscopeCategory struct {
		Id int64 `json:"id"`
	} `json:"funcscope_category"`
}

func (c *Client) QueryAuth(compAccessToken, authCode string) (*QueryAuthRes, error) {
	var res *QueryAuthRes
	query := url.Values{}
	query.Set("component_access_token", compAccessToken)

	err := c.newResponse(
		c.newRequest().SubPath("component/api_query_auth").Method(http.MethodPost).Query(query).JsonBody(queryAuthReq{
			ComponentAppid:    common.GetCompAppId(),
			AuthorizationCode: authCode,
		}),
	).intoJson(&res)

	return res, err
}

type authTokenReq struct {
	RefreshToken   string `json:"authorizer_refresh_token"`
	Appid          string `json:"authorizer_appid"`
	ComponentAppid string `json:"component_appid"`
}

type AuthTokenRes struct {
	*ErrRes
	AuthToken
}

type AuthToken struct {
	AccessToken  string `json:"authorizer_access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"authorizer_refresh_token"`
}

func (c *Client) GetAuthToken(compAccessToken, appId, refreshToken string) (*AuthTokenRes, error) {
	var res *AuthTokenRes
	query := url.Values{}
	query.Set("component_access_token", compAccessToken)

	err := c.newResponse(
		c.newRequest().SubPath("component/api_authorizer_token").Method(http.MethodPost).Query(query).JsonBody(authTokenReq{
			ComponentAppid: common.GetCompAppId(),
			Appid:          appId,
			RefreshToken:   refreshToken,
		}),
	).intoJson(&res)

	return res, err
}
