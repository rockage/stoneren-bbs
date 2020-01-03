package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"mysql_con"
	"strconv"
)

//FillMain
func renderIndexMain(ctx iris.Context) {
	page := ctx.FormValue("page")
	p1, _ := strconv.Atoi(page)
	t1 := p1*20 - 20
	startRec := strconv.Itoa(t1)
	stopRec := "20"

	var ok bool
	var rst []map[string]string

	rst, ok = mysql_con.Query("select tid,author,subject,dateline,lastpost,lastposter,views,replies from pre_forum_thread ORDER BY dateline DESC limit " + startRec + "," + stopRec)
	if ok {
		b, err := json.Marshal(rst)
		if err == nil {
			_, _ = ctx.JSON(string(b)) //不返回错误代码，强行执行
		}
	}
}
func getTotalThreads(ctx iris.Context) {
	var rst []map[string]string
	rst, _ = mysql_con.Query("explain select * from pre_forum_thread") //求出总行数
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

	rst, ok = mysql_con.Query("select pid,fid,tid,author,dateline,message from pre_forum_post where tid = " + tid + " ORDER BY dateline DESC limit " + startRec + "," + stopRec)
	if ok {
		b, err := json.Marshal(rst)
		if err == nil {
			fmt.Println(string(b))
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
