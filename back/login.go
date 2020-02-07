package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"mysql_con"
	"regexp"
)

// sessions.go
var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

func secret(ctx iris.Context) {
	// 检查用户是否已通过身份验证
	/*
		if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
			//ctx.StatusCode(iris.StatusForbidden)
			ctx.WriteString("false")
			return
		}
		// 打印秘密消息

		if uid == "" {
			ctx.WriteString("")
		} else {
			ctx.WriteString(uid)
		}

	*/
	session := sess.Start(ctx)
	auth, _ := session.GetBoolean("authenticated")

	if !auth {
		//ctx.StatusCode(iris.StatusForbidden)
		ctx.WriteString("")
		return
	}

	retJson := "{\"username\":\"" + session.GetString("username") + "\",\"uid\":\"" + session.GetString("uid") + "\"}"
	fmt.Println(retJson)
	ctx.WriteString(retJson)

}

// login
func login(ctx iris.Context) {
	session := sess.Start(ctx)
	// 在此处进行身份验证
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	var rst []map[string]string
	rst, _ = mysql_con.Query("select uid from pre_members where username = '" + username + "' and `password` = '" + password + "'")
	if rst != nil {
		session.Set("authenticated", true)
		session.Set("username", username)

		b, err := json.Marshal(rst)
		if err == nil {
			uid := string(b)
			reg := regexp.MustCompile(`^\[\{\"uid\":\"(\d{1,})\"\}\]$`) //uid是JSON字符串,需要正则一次取值
			match := reg.FindStringSubmatch(uid)
			session.Set("uid", match[1])
			_, _ = ctx.JSON(uid)
		}
	} else {
		ctx.Text("not found")
	}
	// 将用户设置为已验证

}

func logout(ctx iris.Context) {
	session := sess.Start(ctx)
	// 撤销用户身份验证
	session.Set("authenticated", false)
	ctx.WriteString("退出登录。")
}
