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
      this.getTotalThreads();
    }
  }
</script>

<script>
  function htmltoubb(){
    str = pattern(document.getElementById("Hsource").value);
    document.getElementById("Uresult").value=str;
  }

  function pattern(str){
//str = str.replace(/(\r\n|\n|\r)/ig, '');
    str = str.replace(/<br[^>]*>/ig,'\n');
    str = str.replace(/<p[^>\/]*\/>/ig,'\n');
//str = str.replace(/\[code\](.+?)\[\/code\]/ig, function($1, $2) {return phpcode($2);});
    str = str.replace(/\son[\w]{3,16}\s?=\s*([\'\"]).+?\1/ig,'');
    str = str.replace(/<hr[^>]*>/ig,'[hr]');
    str = str.replace(/<(sub|sup|u|strike|b|i|pre)>/ig,'[$1]');
    str = str.replace(/<\/(sub|sup|u|strike|b|i|pre)>/ig,'[/$1]');
    str = str.replace(/<(\/)?strong>/ig,'[$1b]');
    str = str.replace(/<(\/)?em>/ig,'[$1i]');
    str = str.replace(/<(\/)?blockquote([^>]*)>/ig,'[$1blockquote]');
    str = str.replace(/<img[^>]*smile=\"(\d+)\"[^>]*>/ig,'[s:$1]');
    str = str.replace(/<img[^>]*src=[\'\"\s]*([^\s\'\"]+)[^>]*>/ig,'[img]'+'$1'+'[/img]');
    str = str.replace(/<a[^>]*href=[\'\"\s]*([^\s\'\"]*)[^>]*>(.+@)<\/a>/ig,'[url=$1]'+'$2'+'[/url]');
//str = str.replace(/<h([1-6]+)([^>]*)>(.*?)<\/h\1>/ig,function($1,$2,$3,$4){return h($3,$4,$2);});
    str = str.replace(/<[^>]*?>/ig, '');
    str = str.replace(/&amp;/ig, '&');
    str = str.replace(/&lt;/ig, '<');
    str = str.replace(/&gt;/ig, '>');
    return str;
  }
  function up(str){
    str = str.replace(/</ig,'&lt;');
    str = str.replace(/>/ig,'&gt;');
    str = str.replace(/\n/ig,'<br />');
    str = str.replace(/\[code\](.+?)\[\/code\]/ig, function($1, $2) {return phpcode($2);});
    str = str.replace(/\[hr\]/ig,'<hr />');
    str = str.replace(/\[\/(size|color|font|backcolor)\]/ig,'</font>');
    str = str.replace(/\[(sub|sup|u|i|strike|b|blockquote|li)\]/ig,'<$1>');
    str = str.replace(/\[\/(sub|sup|u|i|strike|b|blockquote|li)\]/ig,'</$1>');
    str = str.replace(/\[\/align\]/ig,'</p>');
    str = str.replace(/\[(\/)?h([1-6])\]/ig,'<$1h$2>');
    str = str.replace(/\[align=(left|center|right|justify)\]/ig,'<p align="$1">');
    str = str.replace(/\[size=(\d+?)\]/ig,'<font size="$1">');
    str = str.replace(/\[color=([^\[\<]+?)\]/ig, '<font color="$1">');
    str = str.replace(/\[backcolor=([^\[\<]+?)\]/ig, '<font style="background-color:$1">');
    str = str.replace(/\[font=([^\[\<]+?)\]/ig, '<font face="$1">');
    str = str.replace(/\[list=(a|A|1)\](.+?)\[\/list\]/ig,'<ol type="$1">$2</ol>');
    str = str.replace(/\[(\/)?list\]/ig,'<$1ul>');
    str = str.replace(/\[s:(\d+)\]/ig,function($1,$2){ return smilepath($2);});
    str = str.replace(/\[img\]([^\[]*)\[\/img\]/ig,'<img src="$1" border="0" />');
    str = str.replace(/\[url=([^\]]+)\]([^\[]+)\[\/url\]/ig, '<a href="$1">'+'$2'+'</a>');
    str = str.replace(/\[url\]([^\[]+)\[\/url\]/ig, '<a href="$1">'+'$1'+'</a>');
    return str;
  }
  function ubbtohtml(){
    str = up(document.getElementById("Usource").value);
    document.getElementById("Hresult").value=str;
  }
</script>
