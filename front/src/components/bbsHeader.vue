<template>
  <el-row>
    <el-col :span="24">
      <div class="div-one">
        <img src="/static/logo2.png" style="margin-top: 10px;margin-left: 10px">
        <div class="div-two">
            <span v-show="!loginState">
            <el-button type="text" @click="btnClick">登录</el-button>
            <el-button type="text" @click="btnClick">注册</el-button>
            </span>

          <span v-show="loginState">
            <el-dropdown @command="handleCommand">
              <i class="el-icon-setting" style="margin-right: 15px"></i>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item command="1">我的资料</el-dropdown-item>
                <el-dropdown-item command="2">我的主题</el-dropdown-item>
                <el-dropdown-item command="3">修改密码</el-dropdown-item>
                <el-dropdown-item command="4">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
            <span style="color:#909399;font-size: small;">{{uname}} / {{uid}}</span>
              </span>


        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script>
  export default {
    name: "bbsHeader",
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
      handleCommand: function (command) {
         switch (command) {
          case '1':
            this.$profile()
            break
          case '2':
            break
          case '3':
            this.$password()
            break
          case '4':
            this.delCookie('username')
            this.delCookie('password')
            this.delCookie('autologin')
            this.$store.commit('setLoginState', false)
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
