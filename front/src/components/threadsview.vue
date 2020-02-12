<template>
  <div class="threadsview">
    <el-row>
      <div style="background-color: #cccccc;width: 100%;min-height: 40px;">

        <bbs-button></bbs-button>

        <el-col :span="18" :push="0" style="text-align: right;margin-top: 5px;">
          <el-pagination small background @current-change="handleCurrentChange" layout="prev, pager, next" :total=totalPosts
            :page-size="20">
          </el-pagination>
        </el-col>
      </div>
    </el-row>
    <el-row>
      <el-row>

        <el-col :span="24">
          <el-divider content-position="center"><span style="font-size: medium;color: black;font-family: 'Microsoft YaHei',sans-serif">主题：
              {{subject}} </span>
          </el-divider>

        </el-col>

      </el-row>
      <el-table :data="tableData" style="width: 100%;" :show-header="false" :border="true">

        <el-table-column label="" min-width="20%" className="my-cell">
          <template slot-scope="scope">
            <el-row style="margin-top: 20px" justify="center">
              <el-col :span="24">
                <img style="object-fit: fill;width: 60%;height: 60%;border-radius: 9px;" :src="(scope.row.avatar)?(scope.row.avatar):('/static/avatar.png')">
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
              <el-col :span="8" class="author-text" style="background-color: #EBEEF5;font-size: smaller;">武功等级：
              </el-col>
              <el-col :span="8" class="author-text" style="background-color: #EBEEF5;text-align: left;font-size: smaller;">{{scope.row.level}}
              </el-col>
            </el-row>

          </template>
        </el-table-column>
        <el-table-column min-width="80%" className="my-cell-2">
          <template slot-scope="scope">

            <el-row>
              <el-col :span="24">
                <div class="grid-content" style="color: #909399;">
                  <span>第{{scope.row.indexNum}}楼</span><span style="margin-right:15px;"></span><span>发表于：{{ getLocalTime(scope.row.dateline) }}</span>

                  <hr style="height:1px;border-width:0;color:#E4E7ED;background-color:#E4E7ED" />
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
                  <hr style="height:1px;border-width:0;color:#E4E7ED;background-color:#E4E7ED" v-show="msgHr(scope.row.signature)" />
                  <span style="margin-left: 5px;text-align: left;font-style: oblique;font-weight:bold;">{{ scope.row.signature}}</span>
                </div>
              </el-col>
            </el-row>
          </template>
          <el-row>

          </el-row>
        </el-table-column>
      </el-table>

      <div id="id2" style="min-width:20%;min-height:400px;float:left;background-color: white"></div>
      <div id="id3" style="min-width:80%;float:left;border:0px solid blue">
        <reply :tid="tid" :fid="fid" @replyFinished="renderMain($event)"></reply>
      </div>


    </el-row>


  </div>
</template>

<script>
  import bbsButton from './bbsButton'
  import reply from './reply'

  export default {
    name: 'threadsview',
    data() {
      return {
        tid: this.$route.params.tid,
        fid: 0,
        tableData: [],
        totalPosts: 0,
        messageBox: '',
        subject: '',
        currentPage: 1

      }
    },
    components: {
      bbsButton,
      reply,
    },
    methods: {

      userProfile: function(uname) {
        this.$userprofile({
          uname: uname
        })
      },
      msgHr: function(message) {
        return !!message
      },
      msgBox: function(message) { //为了在帖子内容很少的时候撑大回帖内容区
        if (message.length > 500) {
          return 'height: 100%;white-space: pre-line;'
        } else {
          return 'height: 500px;white-space: pre-line;'
        }
      },
      back: function() {

        this.$router.back();
      },
      up: function(str) {

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
        str = str.replace(/\[s:(\d+)\]/ig, function($1, $2) {
          return smilepath($2);
        });
        str = str.replace(/\[img\]([^\[]*)\[\/img\]/ig, '<img src="$1" border="0" />');
        str = str.replace(/\[url=([^\]]+)\]([^\[]+)\[\/url\]/ig, '<a href="$1">' + '$2' + '</a>');
        str = str.replace(/\[url\]([^\[]+)\[\/url\]/ig, '<a href="$1">' + '$1' + '</a>');
        return str;
      },
      handleCurrentChange(val) {
        this.showTitle = true
        this.currentPage = val
        this.renderMain(val)
      },
      renderMain: function(page) {
        console.log('renderMain page=',page)
        let vm = this
        this.axios.get('http://localhost:8081/renderThreadsView', {
            params: {
              page: page,
              tid: this.tid
            }
          })
          .then((response) => {
            let ret = JSON.parse(response.data)
            let array = [] //为返回数据组装一个原始数据不存在的index，即所谓帖子的 “楼层”
            ret.map(function(item, index) {
              array.push(
                Object.assign({}, item, {
                  indexNum: (index + 1) + ((vm.currentPage - 1) * 20) //楼层从1开始
                })
              )
            });

            console.log(array)

            vm.fid = array[0].fid
            vm.GLOBAL.fid = vm.fid
            //vm.fid是为了传递给子组件reply,
            //这里赋值GLOBAL.fid又赋值一次，是为了应付一种情况：在threadview状态下直接F5刷新，GLOBAL.fid会被重置
            vm.subject = array[0].subject
            vm.tableData = array
            for (let i = 0; i < vm.GLOBAL.forumsData.length; i++) {
              if (vm.fid === vm.GLOBAL.forumsData[i].fid) { //从论坛数组里找出相匹配的fid
                vm.GLOBAL.root.$store.state.forumsName = vm.GLOBAL.forumsData[i].name
              }
            }
          })
      },
      getTotalPosts: function() {
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
      getThreads: function() {

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
