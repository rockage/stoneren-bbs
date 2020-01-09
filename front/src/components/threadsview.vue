<template>
  <div class="threadsview">

    <el-main>
      <el-row :gutter="20">
        <el-col :span="1">
          <el-button type="primary" v-on:click="back"><i class="el-icon-back"></i>返回</el-button>
        </el-col>
        <el-col :span="21" :push="1" style="text-align: right;">
          <el-pagination
            background
            @current-change="handleCurrentChange"
            layout="prev, pager, next"
            :total=totalPosts
            :page-size="20"
            >
          </el-pagination>
        </el-col>
      </el-row>


      <el-table
        :data="tableData"
        style="width: 100%;"
        :show-header="false"
      >

        <el-table-column
          label=""
          min-width="15%"
          className="my-cell"
        >
          <template slot-scope="scope">
            <div class="block" style="margin-top: 20px;">
              <el-avatar shape="square" :size="50" src="/static/avatar.png"></el-avatar>
            </div>

            <span style="margin-left: 10px">{{scope.row.author}}</span>
            <br>
            <span style="margin-left: 10px">{{ getLocalTime(scope.row.dateline) }}</span>
            <div class="block" style="margin-bottom: 50px;"></div>
          </template>
        </el-table-column>
        <el-table-column
          min-width="85%"
          className="my-cell-2"
        >
          <template slot-scope="scope">
            <div style="margin-left: 10px;white-space: pre-line;" v-html="up(scope.row.message)"></div>
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
        tableData: [],
        totalPosts: 0,
      }
    },
    methods: {
      back: function () {
        this.$router.back();
      },
      up: function (str) {
        str = str.replace(/\[i(=s)\]/ig, '[i]');
        str = str.replace(/</ig, '&lt;');
        str = str.replace(/>/ig, '&gt;');
        //str = str.replace(/\n/ig, '<br />');
        //str = str.replace(/\[code\](.+?)\[\/code\]/ig, function ($1, $2) {
        //  return phpcode($2);
        //});
        str = str.replace(/\[hr\]/ig, '<hr />');
        str = str.replace(/\[\/(size|color|font|backcolor)\]/ig, '</font>');
        str = str.replace(/\[(sub|sup|u|i|strike|b|blockquote|li)\]/ig, '<$1>');
        str = str.replace(/\[\/(sub|sup|u|i|strike|b|blockquote|li)\]/ig, '</$1>');
        str = str.replace(/\[\/align\]/ig, '</p>');
        str = str.replace(/\[(\/)?h([1-6])\]/ig, '<$1h$2>');
        str = str.replace(/\[align=(left|center|right|justify)\]/ig, '<p align="$1">');
        str = str.replace(/\[size=(\d+?)\]/ig, '<font size="$1">');
        str = str.replace(/\[color=([^\[\<]+?)\]/ig, '<font color="$1">');
        str = str.replace(/\[backcolor=([^\[\<]+?)\]/ig, '<font style="background-color:$1">');
        str = str.replace(/\[font=([^\[\<]+?)\]/ig, '<font face="$1">');
        str = str.replace(/\[list=(a|A|1)\](.+?)\[\/list\]/ig, '<ol type="$1">$2</ol>');
        str = str.replace(/\[(\/)?list\]/ig, '<$1ul>');
        str = str.replace(/\[s:(\d+)\]/ig, function ($1, $2) {
          return smilepath($2);
        });
        str = str.replace(/\[img\]([^\[]*)\[\/img\]/ig, '<img src="$1" border="0" />');
        str = str.replace(/\[url=([^\]]+)\]([^\[]+)\[\/url\]/ig, '<a href="$1">' + '$2' + '</a>');
        str = str.replace(/\[url\]([^\[]+)\[\/url\]/ig, '<a href="$1">' + '$1' + '</a>');
        return str;
      },
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

<style>
  .el-table .cell {
    white-space: pre-line;
  }

  .el-table .my-cell {
    height: 300px;
    text-align: center;
    vertical-align: top
  }

  .el-table .my-cell-2 {
    text-align: left;
    vertical-align: top
  }


</style>
