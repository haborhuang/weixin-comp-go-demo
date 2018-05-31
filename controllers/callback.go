package controllers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/astaxie/beego"
	"github.com/haborhuang/weixin-comp-go-demo/services/auth"
	"github.com/haborhuang/weixin-comp-go-demo/services/common"
)

type CallbackController struct {
	beego.Controller
}

type PostVerifyTicket struct {
	AppId                 string `xml:"AppId"`
	CreatedTime           int64  `xml:"CreatedTime"`
	InfoType              string `xml:"InfoType"`
	ComponentVerifyTicket string `xml:"ComponentVerifyTicket"`
}

type EncryptedTicket struct {
	AppId   string `xml:"AppId"`
	Encrypt string `xml:"Encrypt"`
}

// @router /events [post]
func (c *CallbackController) Post() {
	defer func() {
		c.Ctx.WriteString("success")
	}()

	timestamp, nonce, signature := c.GetString("timestamp"), c.GetString("nonce"), c.GetString("msg_signature")
	log.Println("Receive auth event: ", string(c.Ctx.Input.RequestBody))
	log.Println("encrypt_type:", c.GetString("encrypt_type"))
	log.Println("msg_signature:", signature)
	log.Println("timestamp:", timestamp)
	log.Println("nonce:", nonce)

	var et EncryptedTicket
	if err := xml.Unmarshal(c.Ctx.Input.RequestBody, &et); nil != err {
		log.Println("Decode event error:", err)
		return
	}

	if !WXMP.verifySign(timestamp, nonce, et.Encrypt, signature) {
		log.Println("Fail to verify signature")
		return
	}

	decrypted, err := WXMP.decrypt(et.Encrypt)
	if nil != err {
		log.Println("Fail to decrypt:", err)
		return
	}

	var t PostVerifyTicket
	if err := xml.Unmarshal(decrypted, &t); nil != err {
		log.Println("Decode ticket error:", err)
		return
	}

	if t.InfoType == "component_verify_ticket" {
		log.Println("Ticket:", t.ComponentVerifyTicket)
		auth.SetVerifyTicket(t.ComponentVerifyTicket)
	}
}

// @router /auth_link [get]
func (c *CallbackController) AuthLink() {
	code, err := auth.GetPreAuthCode()
	if nil != err {
		log.Println("Get pre auth code error:", err)
	}

	c.TplName = "auth_link.tpl"
	c.Data["preAuthCode"] = code
	c.Data["appId"] = common.GetCompAppId()
	c.Data["redirectUrl"] = "https://api.wizardcloud.cn/wx/auth/auth_cb"
	c.Render()
}

// @router /auth_cb [get]
func (c *CallbackController) AuthCallback() {
	log.Println("auth_code:", c.GetString("auth_code"))
	log.Println("expires_in:", c.GetString("expires_in"))

	if err := auth.QueryAuth(c.GetString("auth_code")); nil != err {
		log.Println("Query auth info error: %v", err)
	}
	c.Ctx.WriteString("success")
}

var WXMP *wxMP
var b64Encoding = base64.StdEncoding

func init() {
	sKey := "joasjogiaIasoigjwog83jsoigajw134jiojaohwiut"
	aesKey, err := b64Encoding.DecodeString(sKey + "=")
	if nil != err {
		panic(fmt.Sprintf("Fail to parse aes key from symmetric_key: %v", err))
	}

	WXMP = &wxMP{
		token: "urnvmeofhekdliej",
		sKey:  aesKey,
	}
}

type wxMP struct {
	token string
	sKey  []byte
}

func (w *wxMP) verifySign(timestamp, nonce, encryped, signature string) bool {
	s := []string{w.token, timestamp, nonce, encryped}
	sort.Strings(s)
	check := sha1.Sum([]byte(strings.Join(s, "")))

	return hex.EncodeToString(check[:]) == signature
}

func (w *wxMP) decrypt(encryped string) ([]byte, error) {
	data, err := b64Encoding.DecodeString(encryped)
	if nil != err {
		return nil, fmt.Errorf("Base64 decode cipher error: %v", err)
	}

	// decrypt
	block, err := aes.NewCipher(w.sKey)
	iv := w.sKey[:aes.BlockSize]
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(data, data)

	// decode data
	buf := bytes.NewBuffer(data[16:20])
	var l int32
	binary.Read(buf, binary.BigEndian, &l)

	return data[20 : l+20], nil
}
