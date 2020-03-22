<template>
  <div id="app">
    <title>{{bbstitle}}</title>

      <bbsHeader></bbsHeader>




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
                <i class="el-icon-menu"/>论坛
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
              <i class="el-icon-time"></i>新帖
            </el-menu-item>
            <el-menu-item index="4" @click="router_to(3)">
              <i class="el-icon-user"></i>用户
            </el-menu-item>
          </el-menu>
        </div>


        <div style="margin-top: 0px;margin-bottom: 0px">
          <router-view :key="$route.fullPath" v-if="isRouterAlive"></router-view>
        </div>

      <div style="text-align: center">
        <el-divider content-position="center">
          <a
            target="_blank"
            href="https://github.com/rockage/stoneren-bbs"
          >stoneren-bbs@github</a>
        </el-divider>
      </div>


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
                isRouterAlive: true,
            };
        },
        computed: {
            bbstitle() {
                return this.$store.getters.bbstitle
            },
        },
        provide() {
            return {
                reload: this.reload
            };
        },
        methods: {
            reload: function () {
                this.isRouterAlive = false;
                this.$nextTick(function () {
                    this.isRouterAlive = true;
                });
            },
            router_to: function (f) {
                switch (f) {
                    case 2:
                        this.$router
                            .push({name: "new", params: {page: 1, rmode: 'new', sortmode: 'date'}})
                        break;
                    case 3:
                        this.$router
                            .push({name: "usersview", params: {page: 1, sortmode: 'date'}})
                        break;
                    default:
                        break;
                }
            },
            autoLogin: async function () {
                const vm = this
                let auto = false
                if (vm.getCookie("auto") !== null) {
                    auto = vm.getCookie("auto")
                }
                if (auto) {
                    vm.axios.get('secret', {
                        withCredentials: true //携带cookie提交，mycookiesessionnameid是一个httponly cookie
                    }).then((response) => {
                        if (response.data === 'token check fail') {
                            vm.$store.commit("loginState", false)
                            vm.$store.commit("uid", '')
                            vm.$store.commit("uname", '')
                        } else {
                            let ret = response.data.split("|")
                            vm.$store.commit("loginState", true)
                            vm.$store.commit("uid", ret[1])
                            vm.$store.commit("uname", ret[0])
                        }
                    })
                }
            },
            getForumsInfo: function (info) {
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
        beforeMount() {
        },
        mounted: async function () {
            this.autoLogin();
            await this.$store.dispatch("setFsname");
            this.getForumsInfo(this.$store.getters.fsname);
            this.$store.commit("bbstitle", "石凳BBS");
        }
    };
</script>

<style>
  .footer {
    display: flex;
    align-items: center;
    justify-content: center;
    text-align: justify;
    width: 100%;
    height: 50px;
    background: #99a9bf;
    color: #fff;
  }

  body {
    margin: 0;
    padding: 0;
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
