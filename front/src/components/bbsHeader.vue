<template>
  <div>
    <table border="0" style="width: 100%;background-color: #EBEEF5">
      <tbody>
      <tr>
        <td width="150">
          <img src="/static/logo3.png" style="margin-top: 10px;margin-left: 10px">
          <button @click="ccc" style="display: none">统计发帖数</button>

        </td>
        <td><span style="font-size: xx-large">中国石凳联盟论坛</span>
          <button @click="login">login</button>
          <button @click="logout">logout</button>
          <button @click="secret">secret</button>
          <button @click="savecookie">savecookie</button>
          <button @click="readcookie">readcookie</button>
          <button @click="readpasswd">readpasswd</button>
        </td>

        <td style="text-align: right;margin-right: 20px">

      <span v-show="!loginState">
            <el-button type="text" @click="btnClick">登录</el-button>
            <el-button type="text" @click="btnClick">注册</el-button>
            </span>

          <span v-show="loginState">
            <el-dropdown @command="handleCommand">
              <span class="el-dropdown-link">
                {{uname}}<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item command="1">我的资料</el-dropdown-item>
                <el-dropdown-item command="2">我的主题</el-dropdown-item>
                <el-dropdown-item command="3">修改密码</el-dropdown-item>
                <el-dropdown-item command="4">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
      </span>
        </td>
        <td width="20px"></td>
      </tr>
      <tr>
        <td>

        </td>
      </tr>
      </tbody>
    </table>

    <profile :isShow="profileShow" @profileClose="profileShow=$event"></profile>
  </div>
</template>

<script>
  import Profile from "./profile.vue"

  export default {
    name: "bbsHeader",
    data() {
      return {
        profileShow: false,
        rootThis: '',
      };
    },
    components: {
      Profile
    },
    computed: {
      loginState() {
        return this.$store.state.loginState
      },
      uid() {
        return this.$store.state.uid
      },
      uname() {
        return this.$store.state.uname
      },
    },
    methods: {
      closeProfile: function () {
        this.profileShow = false

      },
      readpasswd: function () {

        console.log(root.$store.state.password)

      },
      readcookie: function () {
        console.log(this.getCookie('local'))
      },
      savecookie: function () {
        this.setCookie('local', this.getCookie('true;rockage;e10adc3949ba59abbe56e057f20f883e'))
      },

      login: function () {
        this.axios.get('http://localhost:8081/login', {
          withCredentials: true,
          params: {
            username: 'rockage',
            password: 'e10adc3949ba59abbe56e057f20f883e',
          }
        })
          .then((response) => {
            console.log(response.data)
          })

      },
      logout: function () {
        this.axios.get('http://localhost:8081/logout', {withCredentials: true}).then((response) => {
          console.log(response.data)
        })
      },
      secret: function () {
        this.axios.get('http://localhost:8081/secret', {
          withCredentials: true,
        }).then((response) => {
          console.log(response.data)
        })
      },


      ccc: function () {
        this.axios.get('http://localhost:8081/resetPosts', {
          params: {}
        })
          .then((response) => {
            this.tableData = JSON.parse(response.data)
            this.loading = false
          })
      },
      handleCommand: function (command) {
        switch (command) {
          case '1':
            this.profileShow = true
            break
          case '2':
            this.$router.push(
              {
                name: 'userThreads',
                params: {uid: this.$store.state.uid}
              }).catch(err => {

            })
            break
          case '3':
            this.$password()
            break
          case '4':
            this.axios.get('http://localhost:8081/logout', {
              withCredentials: true
            })
              .then((response) => {
                this.delCookie('local')
                this.$store.commit('setLoginState', false)
                this.$message.success(response.data)
              })
        }
      },
      btnClick: function () {
        this.$login()
      },
    },
    mounted() {

    }
  }
</script>

<style scoped>

</style>
