<template>
  <div class="threadsview">
    <el-row>
      <div style="background-color: #cccccc;width: 100%;min-height: 40px;">
        <bbs-button @ShowDataTable="ShowDataTable"></bbs-button>

        <el-col :span="18" :push="0" style="text-align: right;margin-top: 5px;">
          <el-pagination
            small
            background
            @current-change="handleCurrentChange"
            layout="prev, pager, next"
            :total="totalPosts"
            :page-size="20"
          >
          </el-pagination>
        </el-col>
      </div>
    </el-row>
    <div id="post_location" style="back-groundcolor:red;width=100%;"></div>
    <div v-show="dataShow">
      <el-row>
        <el-row>
          <el-col :span="24">
            <el-divider content-position="center"
              ><span
                style="font-size: medium;color: black;font-family: 'Microsoft YaHei',sans-serif"
                >主题： {{ subject }}
              </span>
            </el-divider>
          </el-col>
        </el-row>
        <el-table
          :data="tableData"
          style="width: 100%;"
          :show-header="false"
          :border="true"
        >
          <el-table-column label="" min-width="20%" className="my-cell">
            <template slot-scope="scope">
              <el-row style="margin-top: 20px" justify="center">
                <el-col :span="24">
                  <img
                    style="object-fit: fill;width: 60%;height: 60%;border-radius: 9px;"
                    :src="
                      scope.row.avatar ? scope.row.avatar : '/static/avatar.png'
                    "
                  />
                </el-col>
              </el-row>
              <el-row type="flex" justify="center">
                <el-col :span="15" style="background-color: #EBEEF5">
                  <span
                    class="author-text"
                    style="font-size: small;font-weight: bold"
                  >
                    <a
                      href="javascript:void(0)"
                      @click="userProfile({ uname: scope.row.author })"
                      >{{ scope.row.author }}</a
                    >
                  </span>
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="12" class="author-text">回复帖子：</el-col>
                <el-col
                  :span="12"
                  class="author-text"
                  style="text-align: left"
                  >{{ scope.row.postsA }}</el-col
                >
              </el-row>
              <el-row>
                <el-col :span="12" class="author-text">发表主题：</el-col>
                <el-col
                  :span="12"
                  class="author-text"
                  style="text-align: left"
                  >{{ scope.row.threads }}</el-col
                >
              </el-row>
              <el-row>
                <el-col :span="12" class="author-text">注册日期：</el-col>
                <el-col :span="12" class="author-text" style="text-align: left"
                  >{{ getLocalDate(scope.row.regdate) }}
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="12" class="author-text">最后登录：</el-col>
                <el-col :span="12" class="author-text" style="text-align: left"
                  >{{ getLocalDate(scope.row.lastvisited) }}
                </el-col>
              </el-row>
              <el-row type="flex" justify="center">
                <el-col
                  :span="8"
                  class="author-text"
                  style="background-color: #EBEEF5;font-size: smaller;"
                  >武功等级：
                </el-col>
                <el-col
                  :span="8"
                  class="author-text"
                  style="background-color: #EBEEF5;text-align: left;font-size: smaller;"
                  >{{ scope.row.level }}
                </el-col>
              </el-row>
            </template>
          </el-table-column>
          <el-table-column min-width="80%" className="my-cell-2">
            <template slot-scope="scope">
              <el-row>
                <el-col :span="24">
                  <div class="grid-content" style="color: #909399;">
                    <span>第{{ scope.row.indexNum }}楼</span
                    ><span style="margin-right:15px;"></span
                    ><span>发表于：{{ getLocalTime(scope.row.dateline) }}</span>

                    <hr
                      style="height:1px;border-width:0;color:#E4E7ED;background-color:#E4E7ED"
                    />
                  </div>
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="24">
                  <div
                    class="grid-content"
                    style="min-height:500px;max-height:100%;width:100%;"
                    v-html="up(scope.row.message)"
                  ></div>
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="24">
                  <div class="grid-content">
                    <hr
                      style="height:1px;border-width:0;color:#E4E7ED;background-color:#E4E7ED"
                      v-show="msgHr(scope.row.signature)"
                    />
                    <span
                      style="margin-left: 5px;text-align: left;font-style: oblique;font-weight:bold;"
                      >{{ scope.row.signature }}</span
                    >
                  </div>
                </el-col>
              </el-row>
            </template>
            <el-row> </el-row>
          </el-table-column>
        </el-table>
      </el-row>
      <table border="0" style="width:100%">
        <td style="width:20%"></td>
        <td style="width:80%">
          <reply
            :tid="tid"
            :fid="fid"
            @replyFinished="renderMain($event)"
          ></reply>
        </td>
      </table>
    </div>
  </div>
</template>

<script>
import bbsButton from "./bbsButton";
import reply from "./reply";

export default {
  name: "threadsview",
  data() {
    return {
      tableData: [],
      totalPosts: 0,
      messageBox: "",
      subject: "",
      currentPage: 1,
      dataShow: true
    };
  },
  components: {
    bbsButton,
    reply
  },
  computed: {
    fid: function() {
      return this.$store.getters.fid;
    },
    tid: function() {
      return this.$route.params.tid;
    }
  },
  methods: {
    ShowDataTable: function() {
      this.dataShow = !this.dataShow;
    },
    userProfile: function(uname) {
      this.$userprofile({
        uname: uname
      });
    },
    msgHr: function(message) {
      return !!message;
    },
    up: function(str) {
      return str;
      str = str.replace(/\[i(=s)\]/gi, "[i]");
      str = str.replace(/</gi, "&lt;");
      str = str.replace(/>/gi, "&gt;");
      //str = str.replace(/\n/ig, '<br />');
      //str = str.replace(/\[code\](.+?)\[\/code\]/ig, function ($1, $2) {
      //  return phpcode($2);
      //});
      str = str.replace(/\[hr\]/gi, "<hr />");
      str = str.replace(/\[\/(size|color|font|backcolor)\]/gi, "</font>");
      str = str.replace(/\[(sub|sup|u|i|strike|b|blockquote|li)\]/gi, "<$1>");
      str = str.replace(
        /\[\/(sub|sup|u|i|strike|b|blockquote|li)\]/gi,
        "</$1>"
      );
      str = str.replace(/\[\/align\]/gi, "</p>");
      str = str.replace(/\[(\/)?h([1-6])\]/gi, "<$1h$2>");
      str = str.replace(
        /\[align=(left|center|right|justify)\]/gi,
        '<p align="$1">'
      );
      str = str.replace(/\[size=(\d+?)\]/gi, '<font size="$1">');
      str = str.replace(/\[color=([^\[\<]+?)\]/gi, '<font color="$1">');
      str = str.replace(
        /\[backcolor=([^\[\<]+?)\]/gi,
        '<font style="background-color:$1">'
      );
      str = str.replace(/\[font=([^\[\<]+?)\]/gi, '<font face="$1">');
      str = str.replace(
        /\[list=(a|A|1)\](.+?)\[\/list\]/gi,
        '<ol type="$1">$2</ol>'
      );
      str = str.replace(/\[(\/)?list\]/gi, "<$1ul>");
      str = str.replace(/\[s:(\d+)\]/gi, function($1, $2) {
        return smilepath($2);
      });
      str = str.replace(
        /\[img\]([^\[]*)\[\/img\]/gi,
        '<img src="$1" border="0" />'
      );
      str = str.replace(
        /\[url=([^\]]+)\]([^\[]+)\[\/url\]/gi,
        '<a href="$1">' + "$2" + "</a>"
      );
      str = str.replace(
        /\[url\]([^\[]+)\[\/url\]/gi,
        '<a href="$1">' + "$1" + "</a>"
      );
      return str;
    },
    handleCurrentChange(val) {
      this.currentPage = val;
      this.renderMain(val);
    },
    renderMain: function(page) {
      this.currentPage = page;
      let vm = this;
      this.axios
        .get("http://localhost:8081/renderThreadsView", {
          params: {
            page: page,
            tid: this.tid
          }
        })
        .then(response => {
          let ret = JSON.parse(response.data);
          let array = []; //为返回数据组装一个原始数据不存在的index，即所谓帖子的 “楼层”
          ret.map(function(item, index) {
            array.push(
              Object.assign({}, item, {
                indexNum: index + 1 + (vm.currentPage - 1) * 20 //楼层从1开始
              })
            );
          });
          vm.subject = array[0].subject;
          vm.tableData = array;
        });
    },
    getTotalPosts: function() {
      this.axios
        .get("http://localhost:8081/getTotalPosts", {
          params: {
            tid: this.tid
          }
        })
        .then(response => {
          this.totalPosts = JSON.parse(response.data);
          this.renderMain(1);
        });
    },
    getThreads: function() {}
  },
  mounted() {
    this.getTotalPosts();
  }
};
</script>

<style>
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

.author-text {
  text-align: right;
  color: #909399;
  font-size: x-small;
}
</style>
