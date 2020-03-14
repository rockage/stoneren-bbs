<template>
  <div class="threadsview">
    <bbs-button @ShowDataTable="ShowDataTable"></bbs-button>
    <el-row>
      <div
        style="background-color: #cccccc;width: 100%;min-height: 30px;display: flex;justify-content:center;align-items:center;">
        <el-pagination
          small
          @current-change="handleCurrentChange"
          layout="prev, pager, next"
          :total="totalPosts"
          :page-size="20"
          :current-page="currentPage"
        ></el-pagination>
      </div>
    </el-row>
    <div id="post_location"></div>
    <div v-show="dataShow">
      <el-row>
        <el-row>
          <el-col :span="24">
            <span
              style="font-size: small;color: black;font-family: 'Microsoft YaHei',sans-serif"
            >主题： {{ subject }}</span>
          </el-col>
        </el-row>
        <div
          v-loading="loading"
        ></div>
        <el-table :data="tableData" style="width: 100%;" :show-header="false" :border="true" stripe v-show="!loading">
          <el-table-column label class="my-cell">
            <template slot-scope="scope">
              <el-row>
                <el-col :span="24">
                  <div style="color: #909399;margin-bottom:5px;text-align: left;background: aliceblue">
                    <span class="info" style="margin-left: 0px">第{{ scope.row.indexNum }}楼</span>
                    <span class="info">发表于：{{ getLocalTime(scope.row.dateline) }}</span>
                    <a class="info" href="javascript:void(0)" @click="editPost(scope.row.pid)"
                       v-show="authorid===scope.row.authorid">编辑</a>
                  </div>
                </el-col>
              </el-row>

              <div style="display: flex;justify-content:flex-start;">

                <img style="object-fit: fill;border-radius: 50%;" width="50" height="50"
                     :src="scope.row.avatar ? scope.row.avatar : '/static/avatar.png'"/>

                <div>
                  <a href="javascript:void(0)" style="font-size: medium;font-weight: bold;margin-left: 25px"
                     @click="userProfile({ uname: scope.row.author })">{{ scope.row.author}}
                  </a><BR/>

                  <span class="info">回帖：{{ scope.row.postsA }}</span>
                  <span class="info">等级：{{ scope.row.level }}</span>
                </div>
              </div>

              <el-row>

                <el-col :span="24">
                  <div
                    class="grid-content"
                    style="min-height:100px;max-height:100%;width:100%;margin-top:5px;white-space: pre-line;"
                    v-html="up(scope.row.message)"
                  ></div>
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="24">
                  <div style="text-align: right;">
                    <span
                      style="font-style: oblique;font-weight:bold;color: #909399"
                    >{{ scope.row.signature }}</span>
                  </div>
                </el-col>
              </el-row>
            </template>
          </el-table-column>
        </el-table>
      </el-row>
      <div
        style="background-color: #cccccc;width: 100%;min-height: 30px;display: flex;justify-content:center;align-items:center;">
        <el-pagination
          small
          @current-change="handleCurrentChange"
          layout="prev, pager, next"
          :total="totalPosts"
          :page-size="20"
          :current-page="currentPage"
        ></el-pagination>
      </div>
      <reply :tid="tid" :fid="fid" @replyFinished="renderMain($event)"></reply>
    </div>
  </div>
</template>

<script>
  import bbsButton from "./bbsButton"
  import reply from "./reply"
  import Post from "../components/post.vue"

  export default {
    name: "threadsview",
    data() {
      return {
        tableData: [],
        totalPosts: 0,
        messageBox: "",
        subject: "",
        currentPage: 1,
        loading: false,
        dataShow: true
      }
    },
    components: {
      bbsButton,
      reply
    },
    computed: {
      fid: function () {
        return this.$store.getters.fid
      },
      tid: function () {
        return this.$route.params.tid
      },
      authorid: function () {
        return this.$store.getters.uid
      },
    },
    methods: {
      editPost: function (pid) {
        if (!document.getElementById("post_box")) {
          const PostBox = Vue.extend(Post)
          let instance = new PostBox({
            propsData: {
              xclose: this.xclose,
              root: this.$root,
              reply: true,
              pid: pid,
              postFinished: this.postFinished
            }
          })
          let PostEl = instance.$mount().$el
          document.getElementById("post_location").appendChild(PostEl)
        }
      },
      ShowDataTable: function () {
        this.dataShow = !this.dataShow
      },
      userProfile: function (uname) {
        this.$userprofile(uname)
      },
      msgHr: function (message) {
        return !!message
      },
      up: function (str) {
        return str
        str = str.replace(/\[i(=s)\]/gi, "[i]")
        str = str.replace(/</gi, "&lt;")
        str = str.replace(/>/gi, "&gt;")
        //str = str.replace(/\n/ig, '<br />');
        //str = str.replace(/\[code\](.+?)\[\/code\]/ig, function ($1, $2) {
        //  return phpcode($2);
        //});
        str = str.replace(/\[hr\]/gi, "<hr />")
        str = str.replace(/\[\/(size|color|font|backcolor)\]/gi, "</font>")
        str = str.replace(/\[(sub|sup|u|i|strike|b|blockquote|li)\]/gi, "<$1>")
        str = str.replace(
          /\[\/(sub|sup|u|i|strike|b|blockquote|li)\]/gi,
          "</$1>"
        )
        str = str.replace(/\[\/align\]/gi, "</p>")
        str = str.replace(/\[(\/)?h([1-6])\]/gi, "<$1h$2>")
        str = str.replace(
          /\[align=(left|center|right|justify)\]/gi,
          '<p align="$1">'
        )
        str = str.replace(/\[size=(\d+?)\]/gi, '<font size="$1">')
        str = str.replace(/\[color=([^\[\<]+?)\]/gi, '<font color="$1">')
        str = str.replace(
          /\[backcolor=([^\[\<]+?)\]/gi,
          '<font style="background-color:$1">'
        )
        str = str.replace(/\[font=([^\[\<]+?)\]/gi, '<font face="$1">')
        str = str.replace(
          /\[list=(a|A|1)\](.+?)\[\/list\]/gi,
          '<ol type="$1">$2</ol>'
        )
        str = str.replace(/\[(\/)?list\]/gi, "<$1ul>")
        str = str.replace(/\[s:(\d+)\]/gi, function ($1, $2) {
          return smilepath($2)
        })
        str = str.replace(
          /\[img\]([^\[]*)\[\/img\]/gi,
          '<img src="$1" border="0" />'
        )
        str = str.replace(
          /\[url=([^\]]+)\]([^\[]+)\[\/url\]/gi,
          '<a href="$1">' + "$2" + "</a>"
        )
        str = str.replace(
          /\[url\]([^\[]+)\[\/url\]/gi,
          '<a href="$1">' + "$1" + "</a>"
        )
        return str
      },
      handleCurrentChange(val) {
        this.currentPage = val
        this.renderMain(val)

      },
      renderMain: function (page) {
        window.scrollTo(0, 0)
        this.currentPage = page
        let vm = this
        this.axios
          .get("renderThreadsView", {
            params: {
              page: page,
              tid: this.tid
            }
          })
          .then(response => {
            let ret = JSON.parse(response.data)
            let array = [] //为返回数据组装一个原始数据不存在的index，即所谓帖子的 “楼层”
            ret.map(function (item, index) {
              array.push(
                Object.assign({}, item, {
                  indexNum: index + 1 + (vm.currentPage - 1) * 20 //楼层从1开始
                })
              )
            })
            vm.subject = array[0].subject
            vm.tableData = array
            vm.$store.commit("bbstitle", vm.subject)
            vm.loading = false
          })
      },
      getTotalPosts: function () {
        this.axios
          .get("getTotalPosts", {
            params: {
              tid: this.tid
            }
          })
          .then(response => {
            this.totalPosts = JSON.parse(response.data)
            this.renderMain(1)
          })
      },
      getThreads: function () {
      }
    },
    mounted() {
      this.loading = true
      this.getTotalPosts()
    }
  };
</script>

<style scoped>
  .el-table .cell {
    white-space: pre-line;
  }

  .el-table .my-cell {
    height: 300px;
    text-align: center;
    vertical-align: top;
  }

  .el-table .my-cell-2 {
    text-align: left;
    vertical-align: top;
  }

  .el-table .cell {
    white-space: pre-line;
  }

  .info {
    font-size: smaller;
    color: #909399;
    margin-left: 25px;
    max-width: 90%;
    min-width: 90%;
  }

</style>
