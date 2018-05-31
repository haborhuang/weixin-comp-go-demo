package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"log"
	"encoding/xml"
	"encoding/json"

	"github.com/haborhuang/weixin-comp-go-demo/services/tmpl"
	"github.com/haborhuang/weixin-comp-go-demo/services/client"
)

type MsgController struct {
	beego.Controller
}

// @router /apps/:app/templates [get]
func (m *MsgController) GetTmpls() {
	res, err := tmpl.GetPrivTmpls(m.GetString(":app"))
	if err != nil {
		m.Ctx.WriteString(fmt.Sprintf("error: %v", err))
	} else {
		m.Data["json"] = res
		m.ServeJSON()
	}
}

type encryptedMsgEvent struct {
	ToUserName string `xml:"ToUserName"`
	Encrypt string `xml:"Encrypt"`
}

// @router /events/:app [post]
func (m *MsgController) HandleEvents() {
	defer func() {
		m.Ctx.WriteString("success")
	}()
	log.Println("Receive message event: ", string(m.Ctx.Input.RequestBody))

	var em encryptedMsgEvent
	if err := xml.Unmarshal(m.Ctx.Input.RequestBody, &em); nil != err {
		log.Println("Decode event error:", err)
		return
	}

	decrypted, err := WXMP.decrypt(em.Encrypt)
	if nil != err {
		log.Println("Fail to decrypt:", err)
		return
	}

	log.Println("Decrypted message event:", string(decrypted))
}

type sendTmplMsgReq struct {
	To string `json:"to"`
	Data map[string]*client.TmplData `json:"data"`
}

// @router /apps/:app/templates/:tmpl/message [post]
func (m *MsgController) SendTmplMsg() {
	var req sendTmplMsgReq
	var err error
	defer func() {
		if nil != err {
			m.Ctx.WriteString(fmt.Sprintf("%v", err))
		} else {
			m.Ctx.WriteString("success")
		}
	}()

	err = json.Unmarshal(m.Ctx.Input.RequestBody, &req)
	if nil != err {
		return
	}

	err = tmpl.SendTmplMsg(m.GetString(":app"), m.GetString(":tmpl"), req.To, req.Data)
}

// @router /apps/:app/token [get]
func (m *MsgController) GetToken() {
	token, err := tmpl.GetAppToken(m.GetString(":app"))
	if nil != err {
		m.Ctx.WriteString(fmt.Sprintf("%v", err))
	} else {
		m.Ctx.WriteString(token)
	}
}