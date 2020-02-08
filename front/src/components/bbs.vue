<template>
  <div class="bbs">
    <el-main>
      <el-row>
        <el-col :span="1" v-show="loginState">
          <div>
            <el-button type="primary" icon="el-icon-back" circle v-on:click="$router.back()"></el-button>
          </div>
        </el-col>
        <div style="background-color: #2c3e50">
          <el-col :span="1" v-show="loginState">
            <el-popover
              placement="bottom-start"
              title=""
              trigger="manual"
              offset="0"
              :visible-arrow="false"
              v-model="postVisible">
              <post label=""></post>
              <el-button type="primary" icon="el-icon-edit" circle slot="reference" @click="postVisible = !postVisible">
              </el-button>
            </el-popover>
          </el-col>
        </div>
        <button @click="openpost">post</button>
        <div style="background-color: #2c3e50">
          <el-col :span="20" :push="1" style="text-align: right;">
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
      <el-table :data="tableData" style="width: 100%">
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

    </el-main>
    <post :isShow="postShow" @postClose="postShow= $event"></post>
  </div>

</template>

<script>
  import Post from './post.vue'

  export default {
    name: 'bbs',
    data() {
      return {
        postShow:false,
        tid: this.$route.params.fid,
        tableData: null,
        totalPage: 0,
        currentForum: 0,
        fid: '',
        loading: false,
        postVisible: false,
        rMode: '',
      }
    },
    components: {
      'post': Post,
    },
    computed: {
      loginState() {
        return this.$store.state.loginState
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
      openpost:function(){
        this.postShow = true
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
            this.tableData = JSON.parse(response.data)
            this.loading = false
          })
      },
      getTotalThreads: function () {
        this.axios.get('http://localhost:8081/getTotalThreads', {
          params: {
            fid: this.fid,
            rmode: this.$route.meta.renderMode,
          }
        })
          .then((response) => {
            this.totalPage = JSON.parse(response.data)
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
          this.fid = this.$route.params.fid //常规模式
          if (String(this.$route.params.fid) !== String(this.getContextData("currentForum"))) this.currentPage = 1
          break
        case "self":
          this.fid = this.$route.params.uid//只显示我的主题
          break
      }
    },
    mounted: function () {
      this.setContextData("currentForum", this.fid) //将当前fid存入session
      this.setContextData("currentMode", this.$route.meta.renderMode) //将当前rMode存入session
      this.getTotalThreads();

      this.renderMain(this.currentPage)

    },
  }
</script>
