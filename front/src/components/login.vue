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
    <span>记住我：</span>
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
      };
    },
    methods: {
      init: function (options) {
        this.msg = options.msg
      },
      check: function () {

        const auto = this.autoLogin
        const uname = this.inputName
        const passwd = this.md5(this.inputPasswd)

        if (uname !== '' && passwd !== '') {

          let vm = this
          this.axios.get('http://localhost:8081/login', {
            withCredentials: true,
            params: {
              username: uname,
              password: passwd,
            }
          })
            .then((response) => {
              if (response.status !== 200) {
                this.$message.error('通讯失败，请检查网络。')
                return
              }

              if (response.data === 'not found') {
                this.$message.warning("用户名或密码错误。")
                return
              }

              const ret = JSON.parse(response.data)
              if (ret[0].uid) {

                vm.setCookie('local', auto + ";" + uname + ";" + passwd)
                this.GLOBAL.root.$store.commit('setLoginState', true)
                this.GLOBAL.root.$store.commit('setUid', ret[0].uid)
                this.GLOBAL.root.$store.commit('setUname', uname)
                this.GLOBAL.root.$store.commit('setPasswd', passwd)
                vm.$message.success("恭喜你，登录成功了。")
                vm.dialogVisible = false
              }
            })
        } else {
          this.$message.warning("用户名和密码不能为空。")
        }
      }
    },
    mounted() {

    }
  }
</script>

<style scoped>

</style>
