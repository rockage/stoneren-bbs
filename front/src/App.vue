<template>
  <div id="app">
    <button @click="getForums">get标签</button>
    <button @click="addmyMenu">add</button>
    <el-container>
      <el-header>中国石凳联盟</el-header>
      <el-container>
        <el-menu class="el-menu-vertical-demo" style="text-align: left;">
          <el-menu-item index="1">
            <i class="el-icon-s-home"></i>
            <span slot="title"> 《首页》 - 最新帖子</span>

          </el-menu-item>
            <div v-if="myMenus && myMenus.length > 0">
              <myMenu
                is="myMenu"
                v-for="(myMenu,index) in myMenus"
                v-bind:key="myMenu.index"
                v-bind:fid="myMenu.fid"
                v-bind:label="myMenu.label"
                v-on:remove="myMenu.splice(index, 1)"
              />
            </div>
            <div v-else>板块加载中...</div>
        </el-menu>

        <el-container>
          <router-view/>
          <el-footer>
            <a
              href="https://github.com/rockage/stoneren-bbs"
            >https://github.com/rockage/stoneren-bbs</a>
          </el-footer>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>


<script>
  import myMenu from "./components/menu";

  export default {
    name: "App",
    data() {
      return {
        myMenus: []
      };
    },
    methods: {
      addmyMenu: function (index, fid, label) {
        this.myMenus.push({
          index: index,
          fid: fid,
          label: label
        })
      },
      getForums: function () {
        this.axios.get('http://localhost:8081/getForums', {})
          .then((response) => {
            let index = 2
            for (let i of  JSON.parse(response.data)) {
              this.addmyMenu(index, i['fid'], i['name'])
              index++
            }
          })
      },
    },
    components: {
      myMenu
    },
    mounted() {
      this.getForums();
    },
  };
</script>

<style>
  #app {
    font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: left;
    color: #2c3e50;
    margin-top: 0px;
  }

  a:link {
    text-decoration: none;
    color: #303133;
  }

  a:visited {
    text-decoration: none;
    color: #303133;
  }

  a:hover {
    text-decoration: none;
    color: #409eff;
  }

  .el-header,
  .el-footer {
    background-color: #b3c0d1;
    color: #333;
    text-align: center;
    line-height: 60px;
  }

  .el-aside {
    background-color: #000000;
    color: #333;
    text-align: center;
    line-height: 100px;
    width: 500px;
  }
</style>
