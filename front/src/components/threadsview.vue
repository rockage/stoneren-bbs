<template>
  <div class="threadsview">
    <p>{{totalPosts}}</p>
    <el-main>
      <el-pagination
        background
        @current-change="handleCurrentChange"
        layout="prev, pager, next"
        :total=totalPosts
        :page-size="20"
        style="text-align: right;">
      </el-pagination>
      <el-table
        :data="tableData"
        style="width: 100%">
        <el-table-column
          label="发帖ID"
          min-width="30%">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{scope.row.author}}</span>
            <br>
            <span style="margin-left: 10px">{{ getLocalTime(scope.row.dateline) }}</span>
          </template>
        </el-table-column>
        <el-table-column
           label="内容"
          min-width="70%">
          <template slot-scope="scope">
            <div style="margin-left: 10px;white-space: pre-line;">{{scope.row.message}}</div>
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
        tableData:[],
        totalPosts: 0,
      }
    },
    methods: {
      handleCurrentChange(val) {
        this.renderMain(val)
      },
      renderMain: function (page) {
       // this.loading = true
        this.axios.get('http://localhost:8081/renderThreadsView', {
          params: {
            page: page,
            tid: this.tid
          }
        })
          .then((response) => {
            this.tableData = JSON.parse(response.data)
            console.log(JSON.parse(response.data))
            //this.loading = false
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
      getLocalTime: function (nS) {
        let d = new Date(parseInt(nS) * 1000).toLocaleString().replace(/:\d{1,2}$/, ' ')
        return d;
      }
    },
    mounted() {
      this.getTotalPosts()
    }
  }
</script>

<style scoped>
  .el-table .cell {
    white-space: pre-line;
  }

</style>
