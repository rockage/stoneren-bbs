<template>
  <div class="bbs">
    <el-row>
      <div style="background-color: #cccccc;width: 100%;min-height: 40px;">
        <bbs-button @ShowDataTable="ShowDataTable"></bbs-button>

        <el-col :span="18" :push="0" style="text-align: right;margin-top: 5px;">
          <el-pagination
            small
            background
            @current-change="pageChange"
            layout="prev, pager, next"
            :total="this.totalPage"
            :page-size="20"
            :current-page="this.currentPage"
            style="text-align: right;"
          ></el-pagination>
        </el-col>
      </div>
    </el-row>
    <div id="post_location" style="back-groundcolor:red;width=100%;"></div>
    <div
      style="left:50%;top:50%;width: 100px;height: 100px;position: fixed;z-index: 99"
      v-loading="loading"
    ></div>
    <el-table :data="tableData" border style="width: 100%" v-show="dataShow">
      <el-table-column label="主题" min-width="60%">
        <template slot-scope="scope">
          <router-link
            :to="{ name: 'threadsview', params: { tid: scope.row.tid } }"
          >{{ scope.row.subject }}</router-link>
        </template>
      </el-table-column>
      <el-table-column label="作者" min-width="15%">
        <template slot-scope="scope">
          <span style="margin-left: 10px">
            <a
              href="javascript:void(0)"
              @click="userProfile({ uname: scope.row.author })"
            >{{ scope.row.author }}</a>
          </span>
          <br />
          <span style="margin-left: 10px">
            {{
            getLocalTime(scope.row.dateline)
            }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label="最后发表" min-width="15%">
        <template slot-scope="scope">
          <span style="margin-left: 10px">
            <a
              href="javascript:void(0)"
              @click="userProfile({ uname: scope.row.lastposter })"
            >{{ scope.row.lastposter }}</a>
          </span>
          <br />
          <span style="margin-left: 10px">
            {{
            getLocalTime(scope.row.lastpost)
            }}
          </span>
        </template>
      </el-table-column>

      <el-table-column label="回复/浏览" min-width="10%">
        <template slot-scope="scope">
          <span style="margin-left: 10px;font-size: 100%">
            {{
            scope.row.replies
            }}
          </span>
          <br />
          <span style="margin-left: 10px;color:#B3C0D1;">
            {{
            scope.row.views
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
  name: "bbs",
  data() {
    return {
      tableData: null,
      totalPage: 0,
      currentForum: 0,
      loading: false,
      postVisible: false,
      dataShow: true
    }
  },
  components: {
    bbsButton
  },
  computed: {
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
    ShowDataTable: function () {
      this.dataShow = !this.dataShow
    },
    userProfile: function (uname) {
      this.$userprofile({ uname: uname })
    },
    pageChange: function (val) {
      switch (this.$route.meta.rmode) {
        case "new":
          this.$router
            .push({ name: "new", params: { page: val } })
            .catch(err => { })
          break
        case "normal":
          this.$router
            .push({ name: "forumsview", params: { page: val } })
            .catch(err => { })
          break
        case "self":
          this.$router
            .push({ name: "userThreads", params: { page: val } })
            .catch(err => { })
          break
      }
    },
    renderMain: function () {
      this.loading = true
      this.axios
        .get("http://localhost:8081/renderIndexMain", {
          params: {
            page: this.currentPage,
            fid: this.$route.params.fid,
            uid: this.$route.params.uid,
            rmode: this.$route.meta.rmode
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
    console.log(this.$store.getters.rmode)
  },
  mounted: function () {
    this.renderMain()
  }
};
</script>
