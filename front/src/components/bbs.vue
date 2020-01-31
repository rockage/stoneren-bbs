<template>
  <div class="bbs">


    <el-main>
      <el-row :gutter="20">
        <el-col :span="1" v-show="loginState">
          <el-popover
            placement="bottom-start"
            title=""
            width="100%"
            trigger="manual"
            offset="0"
            :visible-arrow="false"
            v-model="postVisible">
            <post label="发帖"></post>
            <el-button slot="reference" @click="postVisible = !postVisible">发帖 <i class="el-icon-edit-outline"></i>
            </el-button>
          </el-popover>
        </el-col>
        <el-col :span="21" :push="1" style="text-align: right;">
          <el-pagination
            background
            @current-change="handleCurrentChange"
            layout="prev, pager, next"
            :total=this.totalPage
            :page-size="20"
            :current-page=this.currentPage
            style="text-align: right;"
          >
          </el-pagination>
        </el-col>
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
            <span style="margin-left: 10px">{{scope.row.author}}</span>
            <br>
            <span style="margin-left: 10px">{{ getLocalTime(scope.row.dateline) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="最后发表" min-width="15%">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{scope.row.lastposter}}</span>
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
      <el-pagination
        background
        @current-change="handleCurrentChange"
        layout="prev, pager, next"
        :total="1000"
        style="text-align: right;"
      >
      </el-pagination>
    </el-main>
  </div>
</template>

<script>
  let a = 1
  export default {
    name: 'bbs',
    data() {
      return {
        tid: this.$route.params.fid,
        tableData: null,
        totalPage: 0,
        currentPage: 0, //BUG：点击返回无法真正回到上一页，而总是返回第一页
        currentForum: 0, //BUG: 若水区翻页到100，返回一个不足100页的子论坛报错
        fid: 0,
        loading: false,
        postVisible: false,
      }
    },
    computed: {
      loginState() {
        return this.$store.state.loginState
      }
    },
    methods: {
      setContextData: function (key, value) {     //给sessionStorage存值
        if (typeof value == "string") {
          sessionStorage.setItem(key, value);
        } else {
          sessionStorage.setItem(key, JSON.stringify(value));
        }
      },
      getContextData: function (key) { // 从sessionStorage取值
        const str = sessionStorage.getItem(key);
        if (typeof str == "string") {
          try {
            return JSON.parse(str);
          } catch (e) {
            return str;
          }
        }
      },
      handleCurrentChange(page) {
        this.currentPage = page
        this.setContextData("currentPage", this.currentPage) //一旦currentPage发生改变，都存入session
        this.renderMain(page)
      },
      renderMain: function (page) {
        this.loading = true

        this.axios.get('http://localhost:8081/renderIndexMain', {
          params: {
            page: page,
            fid: this.fid,
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
          }
        })
          .then((response) => {
            this.totalPage = JSON.parse(response.data)
          })
      },
      getLocalTime: function (nS) {
        return new Date(parseInt(nS) * 1000).toLocaleString().replace(/:\d{1,2}$/, ' ')

      }
    },
    created: function () {
      if (String(this.$route.params.fid) === String(this.getContextData("currentForum"))) {
        //因为created在mounted前发生，created阶段先get页码，接下来在mounted阶段直接渲染该页，
        //否则点击back的时候会默认回到第1页
        this.currentPage = this.getContextData("currentPage") || 1 //如currentPage=0，默认为1
      } else {
        this.currentPage = 1 //如果fid发生了变化，则当前page强制设为1
      }
    },
    mounted: function () {

      if (this.$route.meta.renderMode === 0) {
        this.fid = '0' //最新帖子模式
      } else {
        this.fid = this.$route.params.fid //常规模式
      }
      this.setContextData("currentForum", this.fid) //将当前fid存入session
      this.getTotalThreads();
      this.renderMain(this.currentPage)



    },
  }
</script>
