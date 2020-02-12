package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"math"
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
	rMode := ctx.FormValue("rmode")
	p1, _ := strconv.Atoi(page)
	t1 := p1*20 - 20
	startRec := strconv.Itoa(t1)
	stopRec := "20"

	switch rMode {
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
			var ret [3]string
			ret[0] = string(b)
			ret[1], ret[2] = getTotalThreads(forumsId, rMode)
			_, _ = ctx.JSON(ret) //不返回错误代码
		}
	}
}
func getTotalThreads(forumsId string, rMode string) (string, string) {
	var sql string
	var sql2 string
	var forumsName string
	switch rMode {
	case "new":
		sql = "explain select * from pre_forum_thread" //最新
		sql2 = ""
	case "self":
		sql = "select  COUNT(1) as rows from pre_forum_thread where authorid = " + forumsId //指定版块
		sql2 = "select username as name from pre_members where uid = " + forumsId
	case "normal":
		sql = "select  COUNT(1) as rows from pre_forum_thread where fid = " + forumsId //只看自己
		sql2 = "select name from pre_forum_forum where fid = " + forumsId
	}
	var rst []map[string]string
	rst, _ = mysql_con.Query(sql) //求出总行数
	totalRow := rst[0]["rows"]

	if sql2 != "" {
		rst, _ = mysql_con.Query(sql2)
		forumsName = rst[0]["name"]
		if rMode == "self" {
			forumsName = rst[0]["name"] + "的帖子"
		}
	} else {
		forumsName = "最新的帖子"
	}
	return totalRow, forumsName
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
		"(select subject from pre_forum_thread where tid = " + tid + ") as subject," +
		"(select threads from pre_members where pre_members.uid = pre_forum_post.authorid) as threads," +
		"(select posts from pre_members where pre_members.uid = pre_forum_post.authorid) as postsA," + //PostA是为了后继查询level，直接用post无效
		"(select level from pre_level where pre_level.posts > postsA ORDER BY posts LIMIT 1) as level," +
		"(select avatar from pre_members where pre_members.uid = pre_forum_post.authorid) as avatar," +
		"(select regdate from pre_members where pre_members.uid = pre_forum_post.authorid) as regdate," +
		"(select lastvisited from pre_members where pre_members.uid = pre_forum_post.authorid) as lastvisited," +
		"(select signature from pre_members where pre_members.uid = pre_forum_post.authorid) as signature " +
		"from pre_forum_post where tid = " + tid + " ORDER BY dateline limit " + startRec + "," + stopRec

	fmt.Println(sql)

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
			fmt.Println(string(b))
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
	var sql string
	fid := ctx.FormValue("fid")
	tid := ctx.FormValue("tid")
	author := ctx.FormValue("uname")
	authorid := ctx.FormValue("uid")
	dateline := strconv.FormatInt(time.Now().Unix(), 10)
	message := ctx.FormValue("postContens")
	useip := ctx.RemoteAddr()
	threadsTitle := ctx.FormValue("threadsTitle")

	fmt.Println("fid = ", fid)
	fmt.Println("tid = ", tid)
	fmt.Println("author = " + author)
	fmt.Println("authorid = ", authorid)
	fmt.Println("dateline = " + dateline)
	fmt.Println("message = " + message)
	fmt.Println("useip = " + useip)
	fmt.Println("threadsTitle = " + threadsTitle)

	message = strings.Replace(message, "'", "\\'", -1)
	message = strings.Replace(message, "</a>", "", -1)
	re, _ := regexp.Compile(`<a\s{1,}href(.+?)>`) //删除全部原生超链接
	message = re.ReplaceAllString(message, "")
	re = regexp.MustCompile(`<img src="data:.+?;base64,(.+?)">`)
	for _, match := range re.FindAllStringSubmatch(message, -1) {
		data, _ := base64.StdEncoding.DecodeString(match[1])
		fileName := saveAttachment(data)
		src := "<img src=\"" + fileName + "\">"
		message = strings.Replace(message, match[0], src, -1)
	}

	if tid != "0" {
		sql = "INSERT INTO pre_forum_post ( fid, tid, author, authorid, dateline, useip, message) VALUES " +
			"(" + fid + ", " + tid + ", '" + author + "', " + authorid + ", '" + dateline + "', '" + useip + "', '" + message + "')"

		fmt.Println(sql)
		mysql_con.Exec(sql)
		//算出总页数，返回给前端直接显示最后回复的帖子：
		var rst []map[string]string
		rst, _ = mysql_con.Query("select COUNT(1) as rows from pre_forum_post where tid = " + tid) //求出总行数
		var i float64
		i, _ = strconv.ParseFloat(rst[0]["rows"], 32)
		page := math.Ceil(i / 20)
		res := strconv.FormatFloat(page, 'E', -1, 32)
		ctx.Text(res)

	}

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
