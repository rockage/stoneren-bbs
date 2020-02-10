<template>
  <div id="app">
    <bbsHeader></bbsHeader>
    <el-row>
      <el-col :span="24">
        <div>
          <el-menu :default-active='defaultActive'
                   class="el-menu-demo"
                   mode="horizontal"
                   menu-trigger="click"
                   background-color="#545c64"
                   active-text-color="#ffd04b"
                   text-color="#ffffff">
            <el-submenu index="1">
              <template slot="title"><i class="el-icon-menu"/>版块</template>
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
            </el-submenu>
            <el-menu-item index="2" @click="router_to(2)"><i class="el-icon-time"></i>最新</el-menu-item>
            <el-menu-item index="3" @click="router_to(3)"><i class="el-icon-sunny"></i>最热</el-menu-item>
            <el-menu-item index="4" @click="router_to(4)"><i class="el-icon-user"></i>用户</el-menu-item>
            <el-menu-item index="5" @click="router_to(5)"><i class="el-icon-s-data"></i>排行榜</el-menu-item>
          </el-menu>

        </div>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="24">
        <div style="margin-top: 0px;margin-bottom: 0px">
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
  import myMenu from "./components/menu"
  import store from './store'

  export default {
    name: "App",
    store: store,
    data() {
      return {
        defaultActive: '2',
        dialogVisible: false,
        myMenus: []
      };
    },
    methods: {
      router_to: function (f) {
        switch (f) {
          case 2:
            this.$router.push({name: 'bbs'}).catch(err => {
              console.log(err.name)
            })
            break
          default:
            break
        }
      },
      autoLogin: function () {
        const vm = this
        let auto = false, uname = '', pass = ''
        if (vm.getCookie('local') !== null) {
          const arr = vm.getCookie('local').split(";")
          auto = arr[0]
          uname = arr[1]
          pass = arr[2]
        }
        if (auto === true || uname !== '' || pass !== '') {
          vm.loginCheck(function (data) { //先判断session能否登录
            if (data !== "") {
              store.commit('setLoginState', true)
              store.commit('setUid', data.uid)
              store.commit('setUname', data.username)
              store.commit('setPasswd', pass)
              console.log('session login')
            } else {
              vm.axios.get('http://localhost:8081/login', { //再使用cookie登录
                params: {
                  withCredentials: true,
                  username: uname,
                  password: pass,
                }
              })
                .then((response) => {
                  if (response.data !== 'not found') {
                    const ret = JSON.parse(response.data)
                    store.commit('setLoginState', true)
                    store.commit('setUid', ret[0].uid)
                    store.commit('setUname', uname)
                    store.commit('setPasswd', pass)
                    console.log('cookie login')
                  }
                })
            }
          })
        } else {
          store.commit('setLoginState', false)
        }
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

      },
    },
    components: {
      myMenu
    },
    beforeMount() {

    },
    mounted() {
      this.autoLogin()
      this.getForums()
    },
  }
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


</style>
