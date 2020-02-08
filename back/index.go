package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"mysql_con"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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

	switch ctx.FormValue("rmode") {
	case "new":
		sql = "select tid,author,subject,dateline,lastpost,lastposter,views,replies from pre_forum_thread ORDER BY dateline DESC limit " + startRec + "," + stopRec
	case "self":
		sql = "select tid,author,subject,dateline,lastpost,lastposter,views,replies from pre_forum_thread where authorid = " + forumsId + " ORDER BY dateline DESC limit " + startRec + "," + stopRec
	case "normal":
		sql = "select tid,author,subject,dateline,lastpost,lastposter,views,replies from pre_forum_thread where fid = " + forumsId + " ORDER BY dateline DESC limit " + startRec + "," + stopRec
	}
	fmt.Println(sql)
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

	switch ctx.FormValue("rmode") {
	case "new":
		sql = "explain select * from pre_forum_thread"
	case "self":
		sql = "select  COUNT(1) as rows from pre_forum_thread where authorid = " + forumsId
	case "normal":
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
	var sql string
	sql = "select pid,fid,tid,author,dateline,message,authorid," +
		"(select threads from pre_members where pre_members.uid = pre_forum_post.authorid) as threads," +
		"(select posts from pre_members where pre_members.uid = pre_forum_post.authorid) as postsA," + //PostA是为了后继查询level，直接用post无效
		"(select level from pre_level where pre_level.posts > postsA ORDER BY posts LIMIT 1) as level," +
		"(select avatar from pre_members where pre_members.uid = pre_forum_post.authorid) as avatar," +
		"(select regdate from pre_members where pre_members.uid = pre_forum_post.authorid) as regdate," +
		"(select lastvisited from pre_members where pre_members.uid = pre_forum_post.authorid) as lastvisited," +
		"(select signature from pre_members where pre_members.uid = pre_forum_post.authorid) as signature " +
		"from pre_forum_post where tid = " + tid + " ORDER BY dateline limit " + startRec + "," + stopRec

	rst, ok = mysql_con.Query(sql)
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
	if CheckLogin(ctx) == false {
		ctx.Text("login-error")
		return
	}
	tid := ctx.FormValue("tid")
	uid := ctx.FormValue("uid")
	threadsTitle := ctx.FormValue("threadsTitle")
	postContens := ctx.FormValue("postContens")
	fmt.Println("tid = " + tid)
	fmt.Println("uid = " + uid)
	fmt.Println("threadsTitle = " + threadsTitle)
	fmt.Println("postContens = " + postContens)
	// 105003
	return
	postContens = strings.Replace(postContens, "'", "\\'", -1)
	postContens = strings.Replace(postContens, "</a>", "", -1)
	re, _ := regexp.Compile(`<a\s{1,}href(.+?)>`) //删除全部原生超链接
	postContens = re.ReplaceAllString(postContens, "")
	re = regexp.MustCompile(`<img src="data:.+?;base64,(.+?)">`)
	for _, match := range re.FindAllStringSubmatch(postContens, -1) {
		data, _ := base64.StdEncoding.DecodeString(match[1])
		fileName := saveAttachment(data)
		src := "<img src=\"" + fileName + "\">"
		postContens = strings.Replace(postContens, match[0], src, -1)
	}
	mysql_con.Exec("UPDATE pre_forum_post set message = '" + postContens + "' where pid = 105003")
}

func saveAttachment(data []byte) string {
	dir := "../front/static/attachment/" + time.Now().Format("2006-01") + "/"
	myUuid := uuid.NewV4()
	_, err := os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, 0777)
	}
	fileName := dir + myUuid.String() + ".jpg"
	err = ioutil.WriteFile(fileName, data, 0666)
	fileName = strings.Replace(fileName, "../front", "", -1) //去掉../front，否则前端无法读取
	return fileName
}

// resetPosts
func resetPosts(ctx iris.Context) {
	var rst []map[string]string
	var rst2 []map[string]string
	var rst3 []map[string]string
	rst, _ = mysql_con.Query("SELECT uid,username from pre_members ORDER BY uid")

	fmt.Println(rst[0])
	if rst != nil {
		for _, v := range rst {
			uid := v["uid"]
			fmt.Println("- - - 正在处理：" + uid + "- - -")
			rst2, _ = mysql_con.Query("SELECT count(1) as tp from pre_forum_post WHERE authorid = " + uid)
			v2 := rst2[0]
			tp := v2["tp"]
			fmt.Println("发帖数：" + tp)
			mysql_con.Exec("UPDATE pre_members set posts = " + tp + " where uid = " + uid)
			rst3, _ = mysql_con.Query("SELECT count(1) as tf from pre_forum_thread WHERE authorid = " + uid)
			v3 := rst3[0]
			tf := v3["tf"]
			fmt.Println("主题数：" + tf)
			mysql_con.Exec("UPDATE pre_members set threads = " + tf + " where uid = " + uid)
		}
	} else {
	}

	fmt.Println()
}
