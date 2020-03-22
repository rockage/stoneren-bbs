<template>

  <div class="div-center" id="registerBox">

    <div style="display: flex;justify-content:space-between;margin:0px;background: #a0cfff;">
      <div class="title"><span style="font-size: medium"><i class="el-icon-user-solid"></i></span> 登录/注册</div>
      <div>
        <a href="javascript:void(0)" @click="close()">
          <i class="el-icon-close" style="font-size:150%"></i>
        </a>
      </div>
    </div>

    <div style="display: flex; flex-direction: column ;margin: 10px;padding: 10px">

      <el-input class="elinput" v-model="loginUsername" prefix-icon="el-icon-user-solid" placeholder="用户名"></el-input>
      <el-input class="elinput" v-model="LoginPassword" show-password prefix-icon="el-icon-lock" placeholder="密码"></el-input>


      <div class="title">记住我：
        <el-switch v-model="autoLogin" active-color="#13ce66" inactive-color="#909399"></el-switch>
      </div>
      <div class="title" style="color: #dd6161;margin:20px">温馨提示：所有老会员的密码已被重置为“<b>123456</b>”，请登录后自行修改！</div>
      <el-button type="primary" @click="submit" style="margin-top: 30px">确 定</el-button>
    </div>


  </div>
</template>

<script>

    export default {
        name: "zhuce",
        props: {root: Object},
        components: {},
        data() {
            return {
                loginUsername: '',
                LoginPassword: '',
                autoLogin: true
            }
        },
        methods: {
            submit: function () {
                if (this.loginUsername === "" || this.LoginPassword === "") {
                    this.$message.warning("用户名和密码不能为空。")
                    return
                }
                this.axios
                    .get("userExist", {
                        withCredentials: true,
                        params: {
                            username: this.loginUsername
                        }
                    })
                    .then(response => {
                        if (response.data === "sucess") {
                            this.userLogin()
                        } else {
                            this.userRegister()
                        }
                    })
            },
            userRegister: function () {
                this.axios
                    .get("userZhuce", {
                        withCredentials: true,
                        params: {
                            username: this.loginUsername,
                            password: this.md5(this.LoginPassword)
                        }
                    })
                    .then(response => {
                        let vm = this
                        if (response.data) {
                            const uid = response.data
                            vm.setCookie("auto", this.loginUsername)
                            vm.root.$store.commit("loginState", true)
                            vm.root.$store.commit("uid", uid)
                            vm.root.$store.commit("uname", this.loginUsername)
                            this.infoBox({
                                title: '注册成功',
                                message: '恭喜, 注册成功了!\n\n用户名：' + this.loginUsername + '\n密  码：' + this.LoginPassword,
                                root: this.root
                            })
                            vm.close()
                        }
                    })
            },
            userLogin: function () {
                let vm = this
                this.axios
                    .get("login", {
                        withCredentials: true,
                        params: {
                            username: this.loginUsername,
                            password: this.md5(this.LoginPassword)
                        }
                    })
                    .then(response => {
                        if (response.status !== 200) {
                            vm.$message.error("通讯失败，请检查网络。")
                            return
                        }
                        if (response.data === "not found") {
                            vm.$message.warning("用户名已经被注册或密码错误。")
                            return
                        }
                        const uid = response.data
                        vm.setCookie("auto", this.loginUsername)
                        vm.root.$store.commit("loginState", true)
                        vm.root.$store.commit("uid", uid)
                        vm.root.$store.commit("uname", this.loginUsername)
                        vm.$message.success("欢迎，登录成功了。")
                        vm.close()
                    })
            },
            close: function () {
                this.$destroy() //销毁VUE
            },
        },
        updated() {

        },
        mounted() {


        },
        destroyed() {
            document.body.removeChild(this.$el) //销毁DOM
        }
    }
</script>

<style scoped>

  .div-center {
    border: 1px gray;
    background: #f0f0f0;
    z-index: 99;
    border-radius: 5px;
    position: fixed;
    top: 20%;
    left: 5%;
    height: 50%;
    width: 90%;
    overflow: hidden;
  }

  .elinput {
    margin-bottom: 10px;
  }

  .title {
    font-size: small;
    color: #606266;
    margin-left: 5px;
  }
</style>
