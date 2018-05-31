package auth

import "fmt"

func GetPreAuthCode() (code string, err error) {
	token, err := GetCompAccessToken()
	if nil != err {
		return "", fmt.Errorf("Get component access token error: %v", err)
	}

	res, err := wxClient.GetPreAuthCode(token)
	if nil != err {
		return "", fmt.Errorf("Call api_create_preauthcode api error: %v", err)
	}

	if res.ErrRes != nil {
		return "", res.ErrRes
	}

	return res.PreAuthCode, nil
}