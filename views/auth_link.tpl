<html>
    <body>
{{if .preAuthCode}}
<a href="https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid={{.appId}}&pre_auth_code={{.preAuthCode}}&redirect_uri={{.redirectUrl}}">
点击此处进行授权
</a>
{{else}}
获取预授权码异常，请稍后重试
{{end}}
    </body>
</html>