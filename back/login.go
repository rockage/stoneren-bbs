package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"mysql_con"
	"strconv"
	"time"
)

// sessions.go
var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

func CheckLogin(ctx iris.Context) bool {
	session := sess.Start(ctx)
	auth, _ := session.GetBoolean("authenticated")
	if !auth {
		return false
	}
	return true
}
func secret(ctx iris.Context) {
	// 检查用户是否已通过身份验证
	session := sess.Start(ctx)
	auth, _ := session.GetBoolean("authenticated")

	if !auth {
		//ctx.StatusCode(iris.StatusForbidden)
		_, _ = ctx.WriteString("token check fail")
		return
	}

	_, _ = ctx.WriteString(session.GetString("username") + "|" + session.GetString("uid"))
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
		uid := rst[0]["uid"]
		session.Set("authenticated", true)
		session.Set("username", username)
		session.Set("uid", uid)
		unixTime := strconv.FormatInt(time.Now().Unix(), 10)
		mysql_con.Exec("update pre_members set lastvisited = " + unixTime + " where uid =" + rst[0]["uid"]) //写入最后登录时间
		_, _ = ctx.WriteString(uid)

	} else {
		_, _ = ctx.Text("not found")
	}
	// 将用户设置为已验证

}

func logout(ctx iris.Context) {
	session := sess.Start(ctx)
	// 撤销用户身份验证
	session.Set("authenticated", false)
	_, _ = ctx.WriteString("退出登录。")
}
