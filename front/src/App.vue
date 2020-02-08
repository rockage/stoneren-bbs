<template>
  <div id="app">
    <bbsHeader></bbsHeader>
    <el-row style="margin-top: 0px">
      <el-col :span="3">
        <div>
          <el-menu
            class="el-menu-vertical-demo"
            background-color="#303133"
            text-color="#fff"
            active-text-color="#ffd04b">
            <div v-if="myMenus && myMenus.length > 0">
              <myMenu
                :test="test"
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
  import myMenu from "./components/menu"
  import store from './store'

  export default {
    name: "App",
    store: store,
    data() {
      return {
        test: 'Rockage',
        dialogVisible: false,
        myMenus: [
          {
            fid: '0',
            label: '最新帖子',
            index: '0',
            icon: 'el-icon-s-home'
          }
        ]
      };
    },
    methods: {
      autoLogin: function () {
        let auto = false, uname = '', pass = ''
        if (this.getCookie('local') !== null) {
          const arr = this.getCookie('local').split(";")
          auto = arr[0]
          uname = arr[1]
          pass = arr[2]
        }
        if (auto === true || uname !== '' || pass !== '') {
          this.loginCheck(function (data) { //先判断session能否登录
            if (data !== "") {
              store.commit('setLoginState', true)
              store.commit('setUid', data.uid)
              store.commit('setUname', data.username)
              store.commit('setPasswd', pass)
              console.log('session login')
            } else {
              this.axios.get('http://localhost:8081/login', { //再使用cookie登录
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
