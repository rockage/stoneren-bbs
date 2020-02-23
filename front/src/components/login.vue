<template>
  <div class="div-center" ref="divCenter" id="loginbox">

    <table border="0" width="100%">
      <tbody>
      <tr>
        <td valign="top">
          <span style="  font-size: medium;color: #909399;margin-left:10px">登录</span>
        </td>
        <td align="right">
            <span style="margin-top:5px;" ref="divClose">
              <a href="javascript:void(0)" @click="close()">
                <i class="el-icon-close" style="font-size:150%;"></i>
              </a>
            </span>
        </td>
      </tr>
      </tbody></table>
    <br/>
          <el-input
            placeholder="用户名"
            prefix-icon="el-icon-s-custom"
            v-model="inputName"
            style="margin-left:5%;margin-bottom: 5px;width:90%;"
          >
          </el-input>
          <el-input
            placeholder="密码"
            prefix-icon="el-icon-lock"
            v-model="inputPasswd"
            show-password
            style="margin-left:5%;margin-bottom: 5px;width:90%;"
          >
          </el-input>
          <div class="info" style="margin-left:5%;margin-top:10px;margin-bottom: 5px;width:90%;">记住我：
          <el-switch
            v-model="autoLogin"
            active-color="#13ce66"
            inactive-color="#909399"
          >
          </el-switch></div>
          <br/>
          <div class="info" style="width: 90%;color: #dd6161">温馨提示：所有旧会员的密码已被重置为“<b>123456</b>”，请登录后自行修改，谢谢！</div>
    <br/>
    <el-button type="primary" @click="check" style="margin-left:10px;margin-top:10px;margin-bottom: 10px;">确 定</el-button>


  </div>
</template>

<script>
    export default {
        name: "login",
        props: {
            loginClose: Function,
            loginFinished: Function,
            root: Object //根实例
        },
        data() {
            return {
                inputName: "",
                inputPasswd: "",
                autoLogin: true,
            }
        },
        methods: {
            close: function () {
                this.$destroy() //销毁VUE
            },
            check: function () {
                const auto = this.autoLogin;
                const uname = this.inputName;
                const passwd = this.md5(this.inputPasswd);

                if (uname !== "" && passwd !== "") {
                    let vm = this;
                    this.axios
                        .get("login", {
                            withCredentials: true,
                            params: {
                                username: uname,
                                password: passwd
                            }
                        })
                        .then(response => {
                            if (response.status !== 200) {
                                this.$message.error("通讯失败，请检查网络。");
                                return;
                            }

                            if (response.data === "not found") {
                                this.$message.warning("用户名或密码错误。");
                                return;
                            }

                            const ret = JSON.parse(response.data);
                            if (ret[0].uid) {
                                vm.setCookie("local", auto + ";" + uname + ";" + passwd);
                                vm.root.$store.commit("loginState", true);
                                vm.root.$store.commit("uid", ret[0].uid);
                                vm.root.$store.commit("uname", uname);
                                vm.root.$store.commit("passwd", passwd);
                                vm.$message.success("欢迎，登录成功了。");
                                vm.close()
                            }
                        });
                } else {
                    this.$message.warning("用户名和密码不能为空。");
                }
            }
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
    position: fixed; /*定位*/
    border: 1px gray;
    background: #e4e7ed; /*设置一下背景*/
    z-index: 99;
    border-radius: 5px;
    text-align: center;

    max-height: 100%;
    text-align: left;
    max-width: 90%;
    min-height: 250px;
    left: 5%;
    top: 10%;

  }

  .info {
    font-size: smaller;
    color: #C0C4CC;
    margin-left: 10px
  }
</style>
