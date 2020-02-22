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
	var id string
	var sort string
	fid := ctx.FormValue("fid")
	uid := ctx.FormValue("uid")
	page := ctx.FormValue("page")
	rMode := ctx.FormValue("rmode")
	sortmode := ctx.FormValue("sortmode")
	p1, _ := strconv.Atoi(page)
	t1 := p1*20 - 20
	startRec := strconv.Itoa(t1)
	stopRec := "20"

	fmt.Println(fid)
	fmt.Println(uid)
	fmt.Println(page)
	fmt.Println(rMode)

	switch sortmode {
	case "date":
		sort = "lastpost DESC"
	case "posts":
		sort = "replies DESC"
	}

	switch rMode {
	case "new":
		id = ""
		sql = "select tid,author,subject,dateline,lastpost,lastposter,views,replies from pre_forum_thread ORDER BY " + sort + " limit " + startRec + "," + stopRec
	case "self":
		id = uid
		sql = "select tid,author,subject,dateline,lastpost,lastposter,views,replies from pre_forum_thread where authorid = " + uid + " ORDER BY " + sort + " limit " + startRec + "," + stopRec
	case "normal":
		id = fid
		sql = "select tid,author,subject,dateline,lastpost,lastposter,views,replies from pre_forum_thread where fid = " + fid + " ORDER BY " + sort + " limit " + startRec + "," + stopRec
	}

	var ok bool
	var rst []map[string]string
	rst, ok = mysql_con.Query(sql)
	if ok {
		b, err := json.Marshal(rst)
		if err == nil {
			var returnValue [3]string
			returnValue[0] = string(b)
			returnValue[1] = getTotalThreads(id, rMode)
			_, _ = ctx.JSON(returnValue) //不返回错误代码
		}
	}
}
func getTotalThreads(id string, rMode string) string {
	var sql string
	switch rMode {
	case "new":
		sql = "explain select * from pre_forum_thread" //最新,不需要指定id
	case "self":
		sql = "select  COUNT(1) as rows from pre_forum_thread where authorid = " + id //指定版块(fid)
	case "normal":
		sql = "select  COUNT(1) as rows from pre_forum_thread where fid = " + id //只看自己(uid)
	}
	var rst []map[string]string
	rst, _ = mysql_con.Query(sql) //求出总行数
	totalRow := rst[0]["rows"]
	return totalRow
}

//forumView
func users(ctx iris.Context) {
	var sql string
	page := ctx.FormValue("page")
	sortmode := ctx.FormValue("sortmode")
	p1, _ := strconv.Atoi(page)
	t1 := p1*20 - 20
	startRec := strconv.Itoa(t1)
	stopRec := "20"
	var rst []map[string]string
	var ok bool

	fmt.Println(sortmode)

	switch sortmode {
	case "date":
		sql = "select uid,username,regdate,posts from pre_members" + " ORDER BY regdate limit " + startRec + "," + stopRec
	case "posts":
		sql = "select uid,username,regdate,posts from pre_members" + " ORDER BY posts desc limit " + startRec + "," + stopRec
	}
	var returnValue [2]string
	rst, ok = mysql_con.Query(sql)
	if ok {
		b, err := json.Marshal(rst)
		if err == nil {
			returnValue[0] = string(b)
			sql = "explain select * from pre_members" //获得总用户数
			rst, _ = mysql_con.Query(sql)
			returnValue[1] = rst[0]["rows"]
			_, _ = ctx.JSON(returnValue)
		}
	}
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
	var sql string
	var rst []map[string]string
	fid := ctx.FormValue("fid")
	tid := ctx.FormValue("tid")
	author := ctx.FormValue("uname")
	authorid := ctx.FormValue("uid")
	dateline := strconv.FormatInt(time.Now().Unix(), 10)
	message := ctx.FormValue("postContens")
	useip := ctx.RemoteAddr()
	subject := ctx.FormValue("threadsTitle")

	message = strings.Replace(message, "'", "\\'", -1)
	message = strings.Replace(message, "</a>", "", -1)
	re, _ := regexp.Compile(`<a\s{1,}href(.+?)>`) //删除全部原生超链接
	message = re.ReplaceAllString(message, "")

	re = regexp.MustCompile(`<img.+?src="data:.+?;base64,(.+?)">`) //用户上传的图片
	for _, match := range re.FindAllStringSubmatch(message, -1) {
		data, _ := base64.StdEncoding.DecodeString(match[1])
		fileName := saveAttachment(data)
		src := "<img style=\"max-width:100%;\" src=\"" + fileName + "\" >"
		message = strings.Replace(message, match[0], src, -1)
	}

	reg0 := regexp.MustCompile(`(<img.+?)width\s{0,}=\s{0,}"\d{1,}"(.+?>)`) //去掉width
	message = reg0.ReplaceAllString(message, "$1$2")

	reg1 := regexp.MustCompile(`(<img.+?)height\s{0,}=\s{0,}"\d{1,}"(.+?>)`) //去掉height
	message = reg1.ReplaceAllString(message, "$1$2")

	reg2 := regexp.MustCompile(`(<img.+?)|style\s{0,}=\s{0,}".+?"(.+?>)`) //去掉style
	message = reg2.ReplaceAllString(message, "$1$2")

	reg3 := regexp.MustCompile(`<img(.+?>)`)
	message = reg3.ReplaceAllString(message, `<img style="max-width:100%;"$1`) //加上style, 前端的容器设为width=100%，图片设为max-width=100%，可实现自适应屏幕

	rst, _ = mysql_con.Query("SELECT posts,threads from pre_members where uid = " + authorid) //求出总行数
	posts, _ := strconv.Atoi(rst[0]["posts"])
	threads, _ := strconv.Atoi(rst[0]["threads"]) //先算出现有posts和threads
	posts++
	threads++

	if tid != "0" { //回帖
		sql = "INSERT INTO pre_forum_post ( fid, tid, author, authorid, dateline, useip, message) VALUES " +
			"(" + fid + ", " + tid + ", '" + author + "', " + authorid + ", '" + dateline + "', '" + useip + "', '" + message + "')"
		mysql_con.Exec(sql)
		//算出总页数，返回给前端直接显示最后回复的帖子：
		rst, _ = mysql_con.Query("select COUNT(1) as rows from pre_forum_post where tid = " + tid) //求出总行数
		var i float64
		i, _ = strconv.ParseFloat(rst[0]["rows"], 32)
		page := math.Ceil(i / 20)
		res := strconv.FormatFloat(page, 'E', -1, 32)
		mysql_con.Exec("update pre_members set posts = " + strconv.Itoa(posts) + " WHERE uid = " + authorid) //写入自增后的posts
		ctx.Text(res)
	} else { //开新主题
		sql = "INSERT INTO pre_forum_thread ( fid, author, authorid, subject, dateline, lastpost, lastposter, views, replies) VALUES " +
			"(" + fid + ", '" + author + "', " + authorid + ", '" + subject + "', '" + dateline + "', '" + dateline + "', '" + author + "', 1, 1)"
		insID := mysql_con.Exec(sql) // 新增主题
		sql = "INSERT INTO pre_forum_post ( fid, tid, author, authorid, dateline, useip, message) VALUES " +
			"(" + fid + ", " + insID + ", '" + author + "', " + authorid + ", '" + dateline + "', '" + useip + "', '" + message + "')"
		mysql_con.Exec(sql)                                                                                                                // 新增第一条回复
		sql = "update pre_members set posts = " + strconv.Itoa(posts) + ",threads = " + strconv.Itoa(threads) + " WHERE uid = " + authorid //写入自增后的posts,threads
		mysql_con.Exec(sql)

		ctx.Text("success")

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
