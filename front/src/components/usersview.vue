<template>
  <div class="usersview">
    <el-row>
      <div style="background-color: #cccccc;width: 100%;min-height: 40px;">
        <section class="layout flex">
          <article>
            <div class="left">
              <div class="author-text" style="margin-top:13px;">
                <a href="javascript:void(0)" @click="sortChange">{{srotLabel}}</a>
              </div>
            </div>
            <div class="center" style="margin-top: 7px;">

              <el-pagination
                small
                @current-change="pageChange"
                layout="prev, pager, next"
                :total="this.totalPage"
                :page-size="20"
                :current-page="currentPage"
                style="text-align: right;"
                :pager-count="5"
              ></el-pagination>

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

      <el-table-column label="注册/最后">
        <template slot-scope="scope">
          <div style="display: flex;flex-direction: column;">
            <div class="info" > {{ getLocalDate(scope.row.regdate) }}</div>
            <div class="info" > {{ getLocalDate(scope.row.lastvisited) }}</div>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="发帖数">
        <template slot-scope="scope">
          <span style="margin-left: 10px;font-size: 100%">
            {{ scope.row.posts }}
          </span>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      small
      @current-change="pageChange"
      layout="prev, pager, next"
      :total="this.totalPage"
      :page-size="20"
      :current-page="currentPage"
      style="text-align: right;"
      :pager-count="5"
    ></el-pagination>
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
                sortmode: this.$route.params.sortmode
            }
        },
        components: {
            bbsButton
        },
        computed: {
            srotLabel: function () {
                switch (this.sortmode) {
                    case "date":
                        return "排序：注册时间"
                    case "last":
                        return "排序：最后访问"
                    case "posts":
                        return "排序：发帖总数"
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
                switch (this.sortmode) {
                    case "date":
                        this.sortmode = "last"
                        break
                    case "last":
                        this.sortmode = "posts"
                        break
                    case "posts":
                        this.sortmode = "date"
                        break
                }
                this.pageChange(1)
            },
            userProfile: function (uname) {
                this.$userprofile(uname)
            },
            pageChange: function (val) {
                this.$router
                    .push({name: "usersview", params: {page: val, sortmode: this.sortmode}})

            },
            renderMain: function () {
                this.loading = true
                this.axios
                    .get("users", {
                        params: {
                            page: this.currentPage,
                            sortmode: this.sortmode
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
    color: #C0C4CC;
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
    width: 40%;
  }


  .layout.flex .center {
    flex: 1;
  }

</style>
