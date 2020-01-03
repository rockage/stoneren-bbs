<template>
  <div class="bbs">
    <el-main>
      <el-pagination
        background
        @current-change="handleCurrentChange"
        layout="prev, pager, next"
        :total=totalPage
        :page-size="20"
        style="text-align: right;"
      >
      </el-pagination>
      <div style="left:50%;top:50%;width: 100px;height: 100px;position: fixed;z-index: 99"
           v-loading="loading"></div>
      <el-table :data="tableData" style="width: 100%">
        <el-table-column label="主题" min-width="60%">
          <template slot-scope="scope">
            <router-link :to="{path:'threads/view/'+scope.row.tid}"> {{scope.row.subject}}</router-link>
          </template>
        </el-table-column>
        <el-table-column label="作者" min-width="15%">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{scope.row.author}}</span>+
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
        tid: 112,
        tableData: null,
        totalPage: 0,
        loading: false
      }
    },
    methods: {
      ccc: function () {

      },
      handleCurrentChange(val) {
        this.renderMain(val)
      },
      renderMain: function (page) {
        this.loading = true
        this.axios.get('http://localhost:8081/renderIndexMain', {
          params: {
            page: page
          }
        })
          .then((response) => {
            this.tableData = JSON.parse(response.data)
            this.loading = false
          })
      },
      getTotalThreads: function () {
        this.axios.get('http://localhost:8081/getTotalThreads', {})
          .then((response) => {
            this.totalPage = JSON.parse(response.data)
            this.renderMain(1)
          })
      },
      getLocalTime: function (nS) {
        let d = new Date(parseInt(nS) * 1000).toLocaleString().replace(/:\d{1,2}$/, ' ')
        return d;
      }
    },
    mounted() {
      this.getTotalThreads()
    }
  }
</script>

<style scoped>
  h1, h2 {
    font-weight: normal;
  }

  ul {
    list-style-type: none;
    padding: 0;
  }

  li {
    display: inline-block;
    margin: 0 10px;
  }

  a {
    color: #2c3e50;
    text-decoration: none;
  }

  a:hover {
    text-decoration: underline;
  }


  .el-main {
    background-color: #E9EEF3;
    color: #333;
    text-align: center;
    line-height: 30px;
  }

  body > .el-container {
    margin-bottom: 40px;
  }

  .el-container:nth-child(5) .el-aside,
  .el-container:nth-child(6) .el-aside {
    line-height: 260px;
  }

  .el-container:nth-child(7) .el-aside {
    line-height: 320px;
  }

</style>
