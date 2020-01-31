<template>
  <el-dialog
    title='用户登录'
    :modal="false"
    :visible.sync="dialogVisible"
    width="30%"

  >
    <el-input
      placeholder="用户名"
      prefix-icon="el-icon-s-custom"
      v-model="inputName" style="margin-bottom: 5px">
    </el-input>

    <el-input
      placeholder="密码"
      prefix-icon="el-icon-lock"
      v-model="inputPasswd" show-password style="margin-bottom: 15px">
    </el-input>
    <span>自动登录：</span>
    <el-switch
      v-model="autoLogin"
      active-color="#13ce66"
      inactive-color="#909399">
    </el-switch>

    <span slot="footer" class="dialog-footer">
              <el-button type="primary" @click="check">确 定</el-button>
              </span>
  </el-dialog>
</template>

<script>

  export default {
    name: "login",
    data() {
      return {
        dialogVisible: false,
        msg: '',
        inputName: '',
        inputPasswd: '',
        autoLogin: true,
        rootThis: '',
      };
    },
    methods: {
      init: function (options) {
        this.msg = options.msg
      },
      check: function () {

        const a = this.autoLogin
        const b = this.inputName
        const c = this.md5(this.inputPasswd)

        if (b !== '' && c !== '') {
          this.setCookie('autologin', a)
          this.setCookie('username', b)
          this.setCookie('password', c)
          let vm = this
          let root = this.rootThis
          this.loginCheck(false, b, c, function (r, uid) {
            if (r) {
              root.$store.commit('setLoginState', r)
              root.$store.commit('setUid', uid)
              root.$store.commit('setUname', b)
              vm.dialogVisible = false
            }
          })
        } else {
          this.$message.warning("用户名和密码不能为空。")
        }

      }
    },
    mounted() {
      this.rootThis = this.GLOBAL.globalThis

    }
  }
</script>

<style scoped>

</style>
