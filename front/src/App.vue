<template>
  <div id="app">
    <el-container>
      <el-header style="text-align: left;background-color: #DCDFE6"><img src="/static/logo.png" style="margin-top: 10px;margin-left: 0px"></el-header>
      <el-container>
        <el-menu
          class="el-menu-vertical-demo"
          background-color="#303133"
          text-color="#fff"
          active-text-color="#ffd04b">
          <div v-if="myMenus && myMenus.length > 0">
            <myMenu
              is="myMenu"
              v-for="(myMenu,index) in myMenus"
              v-bind:key="myMenu.key"
              v-bind:label="myMenu.label"
              v-bind:icon="myMenu.icon"
              v-bind:index="myMenu.fid"
              v-bind:fid="myMenu.fid"
              v-on:remove="myMenu.splice(index, 1)"
            />
          </div>
          <div v-else>板块加载中...</div>
        </el-menu>
        <el-container>
          <router-view :key="$route.fullPath"></router-view>
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
        myMenus: [
          {
            key: 0,
            fid: '0',
            index:'0',
            label: '首页 - 最新帖子',
            icon: 'el-icon-s-home'
          }
        ]
      };
    },
    methods: {
      addmyMenu: function (key, label, icon, fid) {
        this.myMenus.push({
          key: key,
          index:key,
          label: label,
          icon: icon,
          fid: fid,
        })
      },
      getForums: function () {
        this.axios.get('http://localhost:8081/getForums', {})
          .then((response) => {
            let key = 1
            for (let i of  JSON.parse(response.data)) {
              this.addmyMenu(key, i['name'], 'el-icon-s-unfold', i['fid'])
              key++
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
