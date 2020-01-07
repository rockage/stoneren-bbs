<template>
  <div id="app">
    <button @click="getForums">get标签</button>
    <button @click="addmyMenu">add</button>
    <el-container>
      <el-header>中国石凳联盟</el-header>
      <el-container>
        <el-menu style="text-align: left;" >
          <div v-if="myMenus && myMenus.length > 0">
            <myMenu
              is="myMenu"
              v-for="(myMenu,index) in myMenus"
              v-bind:key="myMenu.key"
              v-bind:index="myMenu.index"
              v-bind:label="myMenu.label"
              v-bind:icon="myMenu.icon"
              v-bind:fid="myMenu.fid"
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
        myMenus: [
          {
            key: 0,
            fid: '0',
            label: '《首页》 - 最新帖子',
            icon: 'el-icon-s-home'
          }
        ]
      };
    },
    methods: {
      addmyMenu: function (key, index, label, icon, fid) {
        this.myMenus.push({
          index: index,
          label: label,
          icon: icon,
          fid: fid,
        })
      },
      getForums: function () {
        this.axios.get('http://localhost:8081/getForums', {})
          .then((response) => {
            let key = 1
            let index = ''
            for (let i of  JSON.parse(response.data)) {
              index = '/forums/view/' + i['fid']
              this.addmyMenu(key, index, i['name'], 'el-icon-star-off', i['fid'])
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
