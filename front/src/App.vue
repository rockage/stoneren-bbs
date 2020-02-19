<template>
  <div id="app">
    <el-row>
      <bbsHeader></bbsHeader>
    </el-row>
    <el-row>
      <el-col :span="24">
        <div>
          <el-menu
            :default-active="defaultActive"
            class="el-menu-demo"
            mode="horizontal"
            menu-trigger="click"
            background-color="#545c64"
            active-text-color="#ffd04b"
            text-color="#ffffff"
          >
            <el-submenu index="1">
              <template slot="title">
                <i class="el-icon-menu" />论坛
              </template>
              <div v-if="myMenus && myMenus.length > 0">
                <myMenu
                  is="myMenu"
                  v-for="(myMenu, index) in myMenus"
                  v-bind:fid="myMenu.fid"
                  v-bind:label="myMenu.label"
                  v-bind:index="myMenu.index"
                  v-bind:icon="myMenu.icon"
                  v-bind:key="myMenu.index"
                  v-on:remove="myMenu.splice(index, 1)"
                />
              </div>
            </el-submenu>
            <el-menu-item index="2" @click="router_to(2)">
              <i class="el-icon-time"></i>帖子
            </el-menu-item>
            <el-menu-item index="4" @click="router_to(3)">
              <i class="el-icon-user"></i>用户
            </el-menu-item>
          </el-menu>
        </div>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="24">
        <div style="margin-top: 0px;margin-bottom: 0px">
          <router-view :key="$route.fullPath" v-if="isRouterAlive"></router-view>
        </div>
      </el-col>
    </el-row>

    <el-row>
      <div style="text-align: center">
        <el-divider content-position="center">
          <a
            target="_blank"
            href="https://github.com/rockage/stoneren-bbs"
          >stoneren-bbs@github</a>
        </el-divider>
      </div>
    </el-row>
  </div>
</template>

<script>
import myMenu from "./components/menu";

export default {
  name: "App",
  data() {
    return {
      defaultActive: "2",
      dialogVisible: false,
      myMenus: [],
      isRouterAlive: true
    };
  },
  provide() {
    return {
      reload: this.reload
    };
  },
  methods: {
    reload: function() {
      this.isRouterAlive = false;
      this.$nextTick(function() {
        this.isRouterAlive = true;
      });
    },
    router_to: function(f) {
      switch (f) {
        case 2:
          this.$router
            .push({ name: "new", params: { page: 1 } })
            .catch(err => {});
          break;
        case 3:
          this.$router
            .push({ name: "usersview", params: { page: 1,rmode:'regdate' } })
            .catch(err => {});
          break;          
        default:
          break;
      }
    },
    autoLogin: async function() {
      const vm = this;
      let auto = false,
        uname = "",
        pass = "";
      if (vm.getCookie("local") !== null) {
        const arr = vm.getCookie("local").split(";");
        auto = arr[0];
        uname = arr[1];
        pass = arr[2];
      }
      if (auto === true || uname !== "" || pass !== "") {
        let data = await vm.loginCheck(); //先判断session能否登录
        if (data !== "") {
          vm.$store.commit("loginState", true);
          vm.$store.commit("uid", data.uid);
          vm.$store.commit("uname", data.username);
          vm.$store.commit("passwd", pass);
          console.log("session login");
        } else {
          vm.axios
            .get("http://localhost:8081/login", {
              //再使用cookie登录
              params: {
                withCredentials: true,
                username: uname,
                password: pass
              }
            })
            .then(response => {
              if (response.data !== "not found") {
                const ret = JSON.parse(response.data);
                vm.$store.commit("loginState", true);
                vm.$store.commit("uid", ret[0].uid);
                vm.$store.commit("uname", uname);
                vm.$store.commit("passwd", pass);
                console.log("cookie login");
              }
            });
        }
      } else {
        vm.$store.commit("loginState", false);
      }
    },

    addmyMenu: function(fid, label, index, icon) {},

    getForumsInfo: function(info) {
      let index = 1;
      for (let i = 0; i < info.length; i++) {
        this.myMenus.push({
          fid: info[i].fid,
          label: info[i].name,
          index: "1-" + String(index),
          icon: "el-icon-s-unfold",
          key: index
        });

        index++;
      }
    }
  },
  components: {
    myMenu
  },
  beforeMount() {},
  mounted: async function() {
    this.autoLogin();
    await this.$store.dispatch("setFsname");
    this.getForumsInfo(this.$store.getters.fsname);
  }
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

body {
  margin-top: 0px;
  padding: 0px;
}

#app {
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB",
    "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
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
</style>
