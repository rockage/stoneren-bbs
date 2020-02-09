<template>
  <div class="bbs">
    <el-row>
      <div style="background-color: #cccccc;width: 100%;min-height: 40px;">

        <el-col :span="1" v-show="loginState" style="margin-top: 5px;margin-left: 5px">
          <el-button type="primary" size="mini" icon="el-icon-back" circle v-on:click="$router.back()"></el-button>
        </el-col>
        <el-col :span="1" v-show="loginState" style="margin-top: 5px">

          <el-button type="primary"
                     icon="el-icon-plus"
                     size="mini"
                     circle
                     @click="popupPost()"></el-button>
        </el-col>
        <div id="post_location"></div>
        <el-col :span="3" v-show="loginState" style="margin-top: 10px;margin-left: 0px">
          <span style="font-size: medium;"><i class="el-icon-location"></i></span>
          <span style="color:#606266;font-size: small;font-weight: normal;">{{fname}}</span>
        </el-col>
        <el-col :span="18" :push="0" style="text-align: right;margin-top: 5px;">
          <el-pagination
            small
            background
            @current-change="renderMain"
            layout="prev, pager, next"
            :total=this.totalPage
            :page-size="20"
            :current-page=this.currentPage
            style="text-align: right;"
          >
          </el-pagination>
        </el-col>

      </div>
    </el-row>

    <div style="left:50%;top:50%;width: 100px;height: 100px;position: fixed;z-index: 99"
         v-loading="loading"></div>
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column label="主题" min-width="60%">
        <template slot-scope="scope">
          <router-link :to="{name:'threadsview',params:{ tid:scope.row.tid }}"> {{scope.row.subject}}</router-link>
        </template>
      </el-table-column>
      <el-table-column label="作者" min-width="15%">
        <template slot-scope="scope">
          <span style="margin-left: 10px"><a href="javascript:void(0)" @click="userProfile(scope.row.author)">{{scope.row.author}}</a></span>
          <br>
          <span style="margin-left: 10px">{{ getLocalTime(scope.row.dateline) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="最后发表" min-width="15%">
        <template slot-scope="scope">
          <span style="margin-left: 10px"><a href="javascript:void(0)" @click="userProfile(scope.row.lastposter)">{{scope.row.lastposter}}</a></span>
          <br>
          <span style="margin-left: 10px">{{ getLocalTime(scope.row.lastpost) }}</span>
        </template>
      </el-table-column>

      <el-table-column label="回复/浏览" min-width="10%">
        <template slot-scope="scope">
          <span style="margin-left: 10px;font-size: 100%">{{scope.row.replies}}</span>
          <br>
          <span style="margin-left: 10px;color:#B3C0D1;">{{scope.row.views}}</span>
        </template>
      </el-table-column>
    </el-table>


  </div>

</template>

<script>
  import store from '../store'

  export default {
    name: 'bbs',
    store,
    data() {
      return {
        postShow: true,
        tid: this.$route.params.fid,
        tableData: null,
        totalPage: 0,
        currentForum: 0,
        fid: '',
        loading: false,
        postVisible: false,
        rMode: '',
        fname: '',
      }
    },
    components: {},

    computed: {
      loginState() {
        return store.state.loginState
      },
      currentPage: { //将currentPage放入computed的目的：当访问它的时候自动从session中取值，设置它的时候自动存入session
        get() {
          const str = sessionStorage.getItem("currentPage");
          if (typeof str == "string") {
            try {
              return JSON.parse(str)
            } catch (e) {
              return str
            }
          }
        },
        set(v) {
          if (typeof v == "string") {
            sessionStorage.setItem("currentPage", v)
          } else {
            sessionStorage.setItem("currentPage", JSON.stringify(v))
          }
        }
      }
    },
    methods: {
      popupPost: function () {
        if (!document.getElementById("post_box")) this.$post() //只允许弹出一个post窗口
      },
      userProfile: function (uname) {
        this.$userprofile({uname: uname})

      },
      setContextData: function (key, value) {     //sessionStorage存值
        if (typeof value == "string") {
          sessionStorage.setItem(key, value);
        } else {
          sessionStorage.setItem(key, JSON.stringify(value));
        }
      },
      getContextData: function (key) { // sessionStorage取值
        const str = sessionStorage.getItem(key);
        if (typeof str == "string") {
          try {
            return JSON.parse(str) || 1 // ||1的意思：如果是0则为1
          } catch (e) {
            return str || 1
          }
        }
      },
      renderMain: function (page) {
        this.currentPage = page
        this.loading = true
        this.axios.get('http://localhost:8081/renderIndexMain', {
          params: {
            page: page,
            fid: this.fid,
            rmode: this.$route.meta.renderMode,
          }
        })
          .then((response) => {
            this.tableData = JSON.parse(response.data[0])
            this.totalPage = parseInt(response.data[1])
            this.fname = response.data[2]
            this.loading = false
          })
      },
    },
    created: function () {
      //页码重置规则：1) rMode模式转换；2) 论坛间fid切换；
      if (this.$route.meta.renderMode !== this.getContextData("currentMode")) this.currentPage = 1

      switch (this.$route.meta.renderMode) {
        case "new":
          this.fid = 0 //最新帖子模式，不指定论坛
          break
        case "normal":
          this.fid = this.$route.params.fid //常规模式,不用设置fname，而是从远端返回
          if (String(this.$route.params.fid) !== String(this.getContextData("currentForum"))) this.currentPage = 1
          break
        case "self":
          this.fid = this.$route.params.uid//只显示我的主题
          break
      }
    },
    mounted: function () {
      this.setContextData("currentForum", this.fid)//将当前fid存入session
      this.setContextData("currentMode", this.$route.meta.renderMode) //将当前rMode存入session
      this.renderMain(this.currentPage)
    },
  }
</script>
