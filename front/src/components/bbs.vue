<template>
  <div class="bbs">
    <bbs-button></bbs-button>

    <div style="background-color: #EBEEF5;display: flex;justify-content: space-between;align-items:center;">


      <div class="info" style="font-size: x-small">
        <a href="javascript:void(0)" @click="sortChange">{{srotLabel}}</a>
      </div>

      <div>
        <el-pagination
          small
          @current-change="pageChange"
          layout="prev, pager, next"
          :total="this.totalPage"
          :page-size="20"
          :current-page="this.currentPage"
          :pager-count="5"
        ></el-pagination>
      </div>


    </div>

    <div id="post_location"></div>
    <el-table :data="tableData" border style="width: 100%" :show-header="false" stripe>
      <el-table-column label="主题" min-width="60%">
        <template slot-scope="scope">
          <router-link
            :to="{ name: 'threadsview', params: { tid: scope.row.tid } }"
          >{{ scope.row.subject }}
          </router-link>
          <div>
            <div style="display: flex; flex-direction: row;">
              <div style="min-width: 120px">
                <span class="info">楼主：<a href="javascript:void(0)" @click="userProfile({ uname: scope.row.author })">{{scope.row.author}} </a></span>
              </div>
              <div style="min-width: 100px">
                <span class="info">{{ getLocalDate(scope.row.dateline) }}</span>
              </div>
            </div>
            <div style="display: flex; flex-direction: row;">
              <div style="min-width: 120px">
                <span class="info">回复：<a href="javascript:void(0)"
                                         @click="userProfile({ uname: scope.row.lastposter })">{{scope.row.lastposter}}</a></span>
              </div>
              <div style="min-width: 100px">
                <span class="info">{{ getLocalDate(scope.row.lastpost) }}</span>
              </div>
            </div>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="回复/浏览" min-width="15%">
        <template slot-scope="scope">
          <div class="info" style="display: flex; flex-direction: column ;margin: 0px;padding: 0px">
            <div> {{ scope.row.views }}</div>
            <div>{{ scope.row.replies }}</div>
          </div>
        </template>
      </el-table-column>
    </el-table>
    <div style="text-align: right;margin-bottom: 30px">
      <el-pagination
        small
        @current-change="pageChange"
        layout="prev, pager, next"
        :total="this.totalPage"
        :page-size="20"
        :current-page="this.currentPage"
        :pager-count="5"
      ></el-pagination>

    </div>
  </div>
</template>

<script>
    import bbsButton from "./bbsButton"

    export default {
        name: "bbs",
        data() {
            return {
                tableData: null,
                totalPage: 0,
                currentForum: 0,
                loading: false,
                postVisible: false,
                sortmode: this.$route.params.sortmode
            }
        },
        components: {
            bbsButton
        },
        computed: {
            srotLabel: function () {
                if (this.sortmode === 'date') {
                    return "排序方式：时间"
                } else {
                    return "排序方式：回贴量"
                }
            },
            fid: function () {
                return this.$store.getters.fid
            },
            uid: function () {
                return this.$store.getters.uid
            },
            rmode: function () {
                return this.$store.getters.rmode
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
                if (this.sortmode === 'date') {
                    this.sortmode = "posts"
                } else {
                    this.sortmode = "date"
                }
                this.pageChange(1)
            },
            userProfile: function (uname) {
                this.$userprofile(uname)
            },
            pageChange: function (val) {
                switch (this.$route.meta.rmode) {
                    case "new":
                        this.$router
                            .push({name: "new", params: {page: val, sortmode: this.sortmode}})
                        break
                    case "normal":
                        this.$router
                            .push({name: "forumsview", params: {page: val, sortmode: this.sortmode}})
                        break
                    case "self":
                        this.$router
                            .push({name: "userThreads", params: {page: val, sortmode: this.sortmode}})
                        break
                }
            },
            renderMain: function () {
                this.loading = true

                //默认排序：
                if (!this.sortmode) this.sortmode = 'date'
                let rmode = this.$route.meta.rmode
                if (!this.$route.meta.rmode) rmode = 'new'

                this.axios
                    .get("renderIndexMain", {
                        params: {
                            page: this.currentPage,
                            fid: this.$route.params.fid,
                            uid: this.$route.params.uid,
                            rmode: rmode,
                            sortmode: this.sortmode,
                        }
                    })
                    .then(response => {
                        this.tableData = JSON.parse(response.data[0])
                        this.totalPage = parseInt(response.data[1])
                        this.loading = false
                    })
            }
        },
        created: function () {
            this.$store.commit("rmode", this.$route.meta.rmode) //写入store

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

</style>
