<template>
  <div class="usersview">
    <el-row>
      <div style="background-color: #cccccc;width: 100%;min-height: 40px;">
        <section class="layout flex">
          <article>
            <div class="left">
              <div class="author-text" style="margin-top:5px;">
                <el-button size="mini" @click="sortChange">{{srotLabel}}</el-button>
              </div>
            </div>
            <div class="center">
              <el-col :span="24" :push="0" style="text-align: right;margin-top: 5px;">
                <el-pagination
                  small
                  @current-change="pageChange"
                  layout="prev, pager, next"
                  :total="this.totalPage"
                  :page-size="20"
                  :current-page="this.currentPage"
                  style="text-align: right;"
                ></el-pagination>
              </el-col>
            </div>
          </article>
        </section>
      </div>
    </el-row>
    <div
      style="left:50%;top:50%;width: 100px;height: 100px;position: fixed;z-index: 99"
      v-loading="loading"
    ></div>
    <el-table :data="tableData" border style="width: 100%" stripe>
      <el-table-column label="用户名">
        <template slot-scope="scope">
          <span class="info">
            <a
              href="javascript:void(0)"
              @click="userProfile({ uname: scope.row.username })"
            >{{ scope.row.username }}</a>
          </span>
        </template>
      </el-table-column>

      <el-table-column label="注册日期">
        <template slot-scope="scope">
          <span class="info">
            {{
            getLocalDate(scope.row.regdate)
            }}
          </span>
        </template>
      </el-table-column>

      <el-table-column label="发帖数">
        <template slot-scope="scope">
          <span style="margin-left: 10px;font-size: 100%">
            {{
            scope.row.posts
            }}
          </span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import bbsButton from "./bbsButton"

export default {
  name: "usersview",
  data() {
    return {
      tableData: null,
      totalPage: 0,
      loading: false,
      rmode: this.$route.params.rmode
    }
  },
  components: {
    bbsButton
  },
  computed: {
    srotLabel: function () {
      if (this.rmode === 'regdate') {
        return "排序：日期"
      } else {
        return "排序：发帖"
      }
    },
    currentPage: function () {
      let page = 0
      if (!this.$route.params.page) {
        page = 1
      } else {
        page = parseInt(this.$route.params.page)
      }
      return page
    }
  },
  methods: {
    sortChange: function () {
      if (this.rmode === 'regdate') {
        this.rmode = "posts"
      } else {
        this.rmode = "regdate"
      }
      this.pageChange(1)
    },
    userProfile: function (uname) {
      this.$userprofile({ uname: username })
    },
    pageChange: function (val) {

      console.log(this.rmode)
      this.$router
        .push({ name: "usersview", params: { page: val, rmode: this.rmode } })
        .catch(err => { })


    },
    renderMain: function () {
      this.loading = true
      this.axios
        .get("http://localhost:8081/users", {
          params: {
            page: this.currentPage,
            rmode: this.rmode
          }
        })
        .then(response => {
          this.tableData = JSON.parse(response.data[0])
          this.totalPage = parseInt(response.data[1])
          this.loading = false
        })
    }
  },

  mounted: function () {
    this.renderMain()
  }
};
</script>
<style scoped>
.info {
font-size: smaller;
color:#C0C4CC;
margin-left: 10px
}

.author-text {
  text-align: left;
  margin-left: 10px;
  color: #909399;
  font-size: x-small;
}

  .layout.flex article {
    display: flex;
  }

  .layout.flex .left {
    width:40%;
   }


  .layout.flex .center {
    flex: 1;
  }
  
</style>
