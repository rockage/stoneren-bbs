<template>
  <div class="div-center" ref="divCenter" id="passwdBox">

    <div style="display: flex;justify-content:space-between">
      <div class="left"> 修改密码
      </div>
      <div class="right">
        <a href="javascript:void(0)" @click="close()">
          <i class="el-icon-close" style="font-size:150%;"></i>
        </a>
      </div>
    </div>
    <hr style="width: 98%;">
    <div style="text-align: center">
      <el-input
        placeholder="当前密码"
        prefix-icon="el-icon-s-custom"
        v-model="oldPasswd"
        show-password
        style="margin-bottom: 15px;width: 90%"
      >
      </el-input>

      <el-input
        placeholder="新密码"
        prefix-icon="el-icon-lock"
        v-model="newPasswd1"
        show-password
        style="margin-bottom: 15px;width: 90%"
      >
      </el-input>

      <el-input
        placeholder="确认新密码"
        prefix-icon="el-icon-lock"
        v-model="newPasswd2"
        show-password
        style="margin-bottom: 15px;width: 90%"
      >
      </el-input>
    </div>
    <div style="text-align: left;margin-left: 5%">
      <el-button type="primary" @click="btnClick">提交</el-button>
    </div>

  </div>
</template>

<script>
    export default {
        name: "password",
        props: {
            root: Object //根实例
        },
        data() {
            return {
                oldPasswd: "",
                newPasswd1: "",
                newPasswd2: "",
            }
        },
        methods: {
            close: function () {
                this.$destroy() //销毁VUE
            },
            btnClick: function () {
                let vm = this;

                if (
                    this.oldPasswd === "" ||
                    this.newPasswd1 === "" ||
                    this.newPasswd2 === ""
                ) {
                    this.$message.warning("输入不完整，请重新输入。");
                    return;
                }

                if (this.newPasswd1.length < 6) {
                    this.$message.warning("密码至少6位以上，请重新输入。");
                    return;
                }

                if (this.newPasswd1 !== this.newPasswd2) {
                    this.$message.warning("两次输入的密码不同，请重新输入。");
                    return;
                }

                this.axios
                    .get("setNewPasswd", {
                        params: {
                            uid: vm.root.$store.getters.uid,
                            username: vm.root.$store.getters.uname,
                            oldpasswd: vm.md5(vm.oldPasswd),
                            newpasswd: vm.md5(vm.newPasswd1)
                        }
                    })
                    .then(response => {
                        if (response.data !== "error") {
                            vm.$message.success("修改成功，用新密码重新登录。");
                            vm.close()
                            this.axios
                                .get("logout", {
                                    withCredentials: true
                                })
                                .then(response => {
                                    vm.delCookie("local");
                                    vm.root.$store.commit("loginState", false);
                                });
                            vm.dialogVisible = false;
                        } else {
                            vm.$message.error("密码修改失败，请确认当前密码。");
                        }
                    });
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
    max-height: 100%;
    text-align: left;
    max-width: 90%;
    min-width: 330px;
    min-height: 300px;
    left: 5%;
    top: 10%;
  }

  .info {
    font-size: smaller;
    color: #606266;
    margin-left: 10px;
    margin-bottom: 10px;
    max-width: 90%;
    min-width: 90%;
  }

  .flex {
    display: flex;
    justify-content: space-between
  }

  .left {
    flex: 1;
    font-size: medium;
    color: #909399;
    margin-top: 5px;
    margin-left: 5px;
  }

  .right {
    margin-top: 5px;
    margin-right: 5px;
  }
</style>
