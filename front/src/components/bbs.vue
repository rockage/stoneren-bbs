<template>
  <div class="bbs">
    <el-main>
      <el-button type="primary" v-on:click="">发帖 <i class="el-icon-edit-outline"></i></el-button>
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
  export default {
    name: 'bbs',
    data() {
      return {
        tid: this.$route.params.fid,
        tableData: null,
        totalPage: 0,
        currentPage: 0,
        fid: 0,
        loading: false
      }
    },
    methods: {
      ccc: function () {
        this.$router.go(-1)
      },
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
        return;
      },
      handleCurrentChange(page) {
        this.currentPage = page
        this.setContextData("currentPage", this.currentPage)
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
          fid: this.fid
        })
          .then((response) => {
            this.totalPage = JSON.parse(response.data)
          })
      },
      getLocalTime: function (nS) {
        let d = new Date(parseInt(nS) * 1000).toLocaleString().replace(/:\d{1,2}$/, ' ')
        return d;
      }
    },
    created() {
      this.currentPage = this.getContextData("currentPage") || 1 //每次created的时候，总是刷新“上一页”

    },
    mounted() {


      if (typeof (this.$route.params.fid) == "undefined") {
        this.fid = 0
      } else {
        this.fid = this.$route.params.fid
      }

      this.getTotalThreads();
      this.renderMain(this.currentPage)
    },
    activated() {


    }
  }
</script>
