<template>
  <div class="threadsview">

    <el-main>
      <el-row>
        <el-col :span="1">
          <el-button type="primary" icon="el-icon-back" circle v-on:click="$router.back()"></el-button>
        </el-col>
        <el-col :span="22" :push="1" style="text-align: right;">
          <el-pagination
            small
            background
            @current-change="handleCurrentChange"
            layout="prev, pager, next"
            :total=totalPosts
            :page-size="20"
          >
          </el-pagination>
        </el-col>
      </el-row>


      <el-table
        :data="tableData"
        style="width: 100%;"
        :show-header="false"
        :border="true"
      >
        <el-table-column
          label=""
          min-width="20%"
          className="my-cell"
        >
          <template slot-scope="scope">
            <el-row style="margin-top: 20px" justify="center">
              <el-col :span="24">
                <img style="object-fit: fill;width: 60%;height: 60%;border-radius: 9px;"
                     :src="(scope.row.avatar)?(scope.row.avatar):('/static/avatar.png')">
              </el-col>
            </el-row>
            <el-row type="flex" justify="center">
              <el-col :span="15" style="background-color: #EBEEF5">
                <span class="author-text" style="font-size: small;font-weight: bold">
                  <a href="javascript:void(0)" @click="userProfile(scope.row.author)"> {{scope.row.author}}</a>
                </span>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="12" class="author-text">回复帖子：</el-col>
              <el-col :span="12" class="author-text" style="text-align: left">{{scope.row.postsA}}</el-col>
            </el-row>
            <el-row>
              <el-col :span="12" class="author-text">发表主题：</el-col>
              <el-col :span="12" class="author-text" style="text-align: left">{{scope.row.threads}}</el-col>
            </el-row>
            <el-row>
              <el-col :span="12" class="author-text">注册日期：</el-col>
              <el-col :span="12" class="author-text" style="text-align: left">{{getLocalDate(scope.row.regdate)}}
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="12" class="author-text">最后登录：</el-col>
              <el-col :span="12" class="author-text" style="text-align: left">{{getLocalDate(scope.row.lastvisited)}}
              </el-col>
            </el-row>
            <el-row type="flex" justify="center">
              <el-col :span="8" class="author-text"
                      style="background-color: #EBEEF5;font-size: smaller;">武功等级：
              </el-col>
              <el-col :span="8" class="author-text"
                      style="background-color: #EBEEF5;text-align: left;font-size: smaller;">{{scope.row.level}}
              </el-col>
            </el-row>

          </template>
        </el-table-column>
        <el-table-column
          min-width="85%"
          className="my-cell-2"
        >
          <template slot-scope="scope">
            <el-row>
              <el-col :span="24">
                <div class="grid-content">
                  <span style="margin-left: 5px;text-align: left">发表于：{{ getLocalTime(scope.row.dateline) }}</span>
                  <hr style="height:1px;border-width:0;color:#E4E7ED;background-color:#E4E7ED"/>
                </div>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="24">

                <div class="grid-content" :style="msgBox(scope.row.message)" v-html="up(scope.row.message)"></div>

              </el-col>
            </el-row>
            <el-row>
              <el-col :span="24">
                <div class="grid-content">
                  <hr style="height:1px;border-width:0;color:#E4E7ED;background-color:#E4E7ED"
                      v-show="msgHr(scope.row.signature)"/>
                  <span style="margin-left: 5px;text-align: left;font-style: oblique;font-weight:bold;">{{ scope.row.signature}}</span>
                </div>
              </el-col>
            </el-row>

          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        background
        @current-change="handleCurrentChange"
        layout="prev, pager, next"
        :total=totalPosts
        :page-size="20"
        style="text-align: right;">
      </el-pagination>
    </el-main>
  </div>
</template>

<script>
  export default {
    name: 'threadsview',
    data() {
      return {
        tid: this.$route.params.tid,
        tableData: [],
        totalPosts: 0,
        messageBox: '',
      }
    },
    methods: {
      userProfile: function (uname) {
        this.$userprofile({uname: uname})
      },
      msgHr: function (message) {
        return !!message
      },
      msgBox: function (message) { //为了在帖子内容很少的时候撑大回帖内容区
        if (message.length > 500) {
          return 'height: 100%;white-space: pre-line;'
        } else {
          return 'height: 500px;white-space: pre-line;'
        }
      },
      back: function () {

        this.$router.back();
      },
      up: function (str) {

        return str

        str = str.replace(/\[i(=s)\]/ig, '[i]');
        str = str.replace(/</ig, '&lt;');
        str = str.replace(/>/ig, '&gt;');
        //str = str.replace(/\n/ig, '<br />');
        //str = str.replace(/\[code\](.+?)\[\/code\]/ig, function ($1, $2) {
        //  return phpcode($2);
        //});
        str = str.replace(/\[hr\]/ig, '<hr />');
        str = str.replace(/\[\/(size|color|font|backcolor)\]/ig, '</font>');
        str = str.replace(/\[(sub|sup|u|i|strike|b|blockquote|li)\]/ig, '<$1>');
        str = str.replace(/\[\/(sub|sup|u|i|strike|b|blockquote|li)\]/ig, '</$1>');
        str = str.replace(/\[\/align\]/ig, '</p>');
        str = str.replace(/\[(\/)?h([1-6])\]/ig, '<$1h$2>');
        str = str.replace(/\[align=(left|center|right|justify)\]/ig, '<p align="$1">');
        str = str.replace(/\[size=(\d+?)\]/ig, '<font size="$1">');
        str = str.replace(/\[color=([^\[\<]+?)\]/ig, '<font color="$1">');
        str = str.replace(/\[backcolor=([^\[\<]+?)\]/ig, '<font style="background-color:$1">');
        str = str.replace(/\[font=([^\[\<]+?)\]/ig, '<font face="$1">');
        str = str.replace(/\[list=(a|A|1)\](.+?)\[\/list\]/ig, '<ol type="$1">$2</ol>');
        str = str.replace(/\[(\/)?list\]/ig, '<$1ul>');
        str = str.replace(/\[s:(\d+)\]/ig, function ($1, $2) {
          return smilepath($2);
        });
        str = str.replace(/\[img\]([^\[]*)\[\/img\]/ig, '<img src="$1" border="0" />');
        str = str.replace(/\[url=([^\]]+)\]([^\[]+)\[\/url\]/ig, '<a href="$1">' + '$2' + '</a>');
        str = str.replace(/\[url\]([^\[]+)\[\/url\]/ig, '<a href="$1">' + '$1' + '</a>');
        return str;
      },
      handleCurrentChange(val) {
        this.renderMain(val)
      },
      renderMain: function (page) {
        this.axios.get('http://localhost:8081/renderThreadsView', {
          params: {
            page: page,
            tid: this.tid
          }
        })
          .then((response) => {
            this.tableData = JSON.parse(response.data)
          })
      },
      getTotalPosts: function () {
        this.axios.get('http://localhost:8081/getTotalPosts', {
          params: {
            tid: this.tid
          }
        })
          .then((response) => {
            this.totalPosts = JSON.parse(response.data)
            this.renderMain(1)
          })
      },
      getThreads: function () {

      }
    },
    mounted() {
      this.getTotalPosts()
    }
  }
</script>

<style>
  .el-table .cell {
    white-space: pre-line;
  }

  .el-table .my-cell {
    height: 300px;
    text-align: center;
    vertical-align: top
  }

  .el-table .my-cell-2 {
    text-align: left;
    vertical-align: top
  }

  .author-text {
    text-align: right;
    color: #909399;
    font-size: x-small;
  }

</style>
