<template>
  <div id="app">
    <el-row>
      <el-col :span="24">
        <div class="div-one">
          <img src="/static/logo2.png" style="margin-top: 10px;margin-left: 10px">
          <div class="div-two">
            <el-button type="text" @click="dialogVisible = true">登录</el-button>
            <button @click="btnClick">注册</button>
          </div>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="3">
        <div class="grid-content bg-purple">
          <el-menu
            class="el-menu-vertical-demo"
            background-color="#303133"
            text-color="#fff"
            active-text-color="#ffd04b">
            <div v-if="myMenus && myMenus.length > 0">
              <myMenu
                is="myMenu"
                v-for="(myMenu,index) in myMenus"
                v-bind:fid="myMenu.fid"
                v-bind:label="myMenu.label"
                v-bind:index="myMenu.index"
                v-bind:icon="myMenu.icon"
                v-bind:key="myMenu.index"
                v-on:remove="myMenu.splice(index, 1)"
              />
            </div>
            <div v-else>板块加载中...</div>
          </el-menu>

        </div>
      </el-col>
      <el-col :span="21">
        <div class="grid-content bg-purple-light">
          <router-view :key="$route.fullPath"></router-view>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <div class="footer">
          <a href="https://github.com/rockage/stoneren-bbs"
          >https://github.com/rockage/stoneren-bbs</a>
        </div>
      </el-col>
    </el-row>


  </div>
</template>


<script>
  import myMenu from "./components/menu";

  export default {
    name: "App",
    data() {
      return {
        dialogVisible: false,
        myMenus: [
          {
            fid: '0',
            label: '首页 - 最新帖子',
            index: '0',
            icon: 'el-icon-s-home'
          }
        ]
      };
    },
    methods: {
      btnClick: function () {

        this.$login(
          {
            msg: '用户登录'
          }
        )
      },
      handleClose(done) {
        this.$confirm('确认关闭？')
          .then(_ => {
            done();
          })
          .catch(_ => {
          });
      },
      addmyMenu: function (fid, label, index, icon) {
        this.myMenus.push({
          fid: fid,
          label: label,
          index: index,
          icon: icon,
          key: index,
        })
      },
      getForums: function () {
        this.axios.get('http://localhost:8081/getForums', {})
          .then((response) => {
            let index = 1
            for (let i of  JSON.parse(response.data)) {
              this.addmyMenu(i['fid'], i['name'], String(index), 'el-icon-s-unfold')
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

  .footer {
    /*flex 布局*/
    display: flex;
    /*实现垂直居中*/
    align-items: center;
    /*实现水平居中*/
    justify-content: center;

    text-align: justify;
    width: 100%;
    height: 50px;
    background: #99a9bf;
    margin: 0 auto;
    color: #fff;
  }

  .div-one {
    width: 100%;
    height: 70px;
    position: relative;
  / / border: 1 px solid #ff871e;
    background: #F2F6FC
  }

  .div-two {
    position: absolute;
    left: 85%;
    right: 0;
    top: 0;
    bottom: 10px;
    margin: auto;
    width: 150px;
    height: 20px;
  / / border: 1 px solid #ff871e;
    background: #F2F6FC
  }

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
    color: #ffffff;
    text-align: center;
    line-height: 50px;
    width: 300px;
  }

  .el-row {
    margin-bottom: 20px;

    &
    :last-child {
      margin-bottom: 0;
    }

  }

  .el-col {
    border-radius: 4px;
  }

  .bg-purple-dark {
    background: #99a9bf;
  }

  .bg-purple {
    background: #d3dce6;
  }

  .bg-purple-light {
    background: #e5e9f2;
  }

  .grid-content {
    border-radius: 4px;
    min-height: 36px;
  }

  .row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
  }

</style>
