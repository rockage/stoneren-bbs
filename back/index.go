package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"mysql_con"
	"regexp"
	"strconv"
	"strings"
)

//FillMain
func renderIndexMain(ctx iris.Context) {
	var sql string
	forumsId := ctx.FormValue("fid")
	page := ctx.FormValue("page")
	p1, _ := strconv.Atoi(page)
	t1 := p1*20 - 20

	startRec := strconv.Itoa(t1)
	stopRec := "20"

	if forumsId == "0" {
		sql = "select tid,author,subject,dateline,lastpost,lastposter,views,replies from pre_forum_thread ORDER BY dateline DESC limit " + startRec + "," + stopRec
	} else {
		sql = "select tid,author,subject,dateline,lastpost,lastposter,views,replies from pre_forum_thread where fid = " + forumsId + " ORDER BY dateline DESC limit " + startRec + "," + stopRec
	}
	var ok bool
	var rst []map[string]string

	rst, ok = mysql_con.Query(sql)
	if ok {
		b, err := json.Marshal(rst)
		if err == nil {
			_, _ = ctx.JSON(string(b)) //不返回错误代码，强行执行
		}
	}
}
func getTotalThreads(ctx iris.Context) {
	var sql string
	forumsId := ctx.FormValue("fid")
	if forumsId == "0" {
		sql = "explain select * from pre_forum_thread"
	} else {
		sql = "select  COUNT(1) as rows from pre_forum_thread where fid = " + forumsId
	}

	var rst []map[string]string
	rst, _ = mysql_con.Query(sql) //求出总行数
	_, _ = ctx.JSON(rst[0]["rows"])

}

//ThreadsView
func renderThreadsView(ctx iris.Context) {
	page := ctx.FormValue("page")
	tid := ctx.FormValue("tid")
	p1, _ := strconv.Atoi(page)
	t1 := p1*20 - 20
	startRec := strconv.Itoa(t1)
	stopRec := "20"

	var ok bool
	var rst []map[string]string
	rst, ok = mysql_con.Query("select pid,fid,tid,author,dateline,message from pre_forum_post where tid = " + tid + " ORDER BY dateline limit " + startRec + "," + stopRec)
	if ok {
		b, err := json.Marshal(rst)
		if err == nil {
			_, _ = ctx.JSON(string(b)) //不返回错误代码，强行执行
		}
	}
}
func getTotalPosts(ctx iris.Context) {
	posts := ctx.FormValue("tid")
	var rst []map[string]string
	rst, _ = mysql_con.Query("select COUNT(1) as rows from pre_forum_post where tid = " + posts) //求出总行数
	_, _ = ctx.JSON(rst[0]["rows"])
}

//forumView
func getForums(ctx iris.Context) {
	var rst []map[string]string
	var ok bool
	rst, ok = mysql_con.Query("select fid,name,status,displayorder from pre_forum_forum where status = 1 ORDER BY displayorder")
	if ok {
		b, err := json.Marshal(rst)
		if err == nil {
			_, _ = ctx.JSON(string(b))
		}
	}
}

//New Post
func setNewPost(ctx iris.Context) {
	fid := ctx.FormValue("fid")
	tid := ctx.FormValue("tid")
	uid := ctx.FormValue("uid")
	threadsTitle := ctx.FormValue("threadsTitle")
	postContens := ctx.FormValue("postContens")
	fmt.Println("fid = " + fid)
	fmt.Println("tid = " + tid)
	fmt.Println("uid = " + uid)
	fmt.Println("threadsTitle = " + threadsTitle)
	fmt.Println("postContens = " + postContens)
	// 105003
	postContens = strings.Replace(postContens, "'", "\\'", -1)
	postContens = strings.Replace(postContens, "</a>", "", -1)
	re, _ := regexp.Compile(`<a\s{1,}href(.+?)>`)
	postContens = re.ReplaceAllString(postContens, "")
	mysql_con.Exec("UPDATE pre_forum_post set message = '" + postContens + "' where pid = 105003")

}
