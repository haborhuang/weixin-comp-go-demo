package client

import (
	"net/http"
	"net/url"
)

type TemplatesResp struct {
	*ErrRes
	TmplList []*wxTmpl `json:"template_list"`
}

type wxTmpl struct {
	TmplId          string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

func (c *Client) GetPrivTemplates(token string) (*TemplatesResp, error) {
	var res *TemplatesResp
	query := url.Values{}
	query.Set("access_token", token)
	err := c.newResponse(
		c.newRequest().SubPath("template/get_all_private_template").Method(http.MethodGet).Query(query),
	).intoJson(&res)

	return res, err
}

type TmplData struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

type sendTmplMsgReq struct {
	ToUser      string `json:"touser"`
	TemplateId  string `json:"template_id"`
	Url         string `json:"url,omitempty"`
	MiniProgram *struct {
		AppId    string `json:"appid"`
		PagePath string `json:"pagepath,omitempty"`
	} `json:"miniprogram,omitempty"`
	Data map[string]*TmplData `json:"data"`
}

type SendTmplMsgResp struct {
	*ErrRes
	MsgId int64 `json:"msg_id"`
}

func (c *Client) SendTmplMsg(token, tmplId, to string, params map[string]*TmplData) (*SendTmplMsgResp, error) {
	var res *SendTmplMsgResp
	query := url.Values{}
	query.Set("access_token", token)
	err := c.newResponse(
		c.newRequest().SubPath("message/template/send").Method(http.MethodPost).Query(query).JsonBody(sendTmplMsgReq{
			ToUser: to,
			TemplateId: tmplId,
			Data: params,
		}),
	).intoJson(&res)

	return res, err
}
