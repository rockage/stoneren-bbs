<template>
  <el-dialog
    title="密码修改"
    :modal="false"
    :visible.sync="dialogVisible"
    width="30%"
  >
    <el-input
      placeholder="当前密码"
      prefix-icon="el-icon-s-custom"
      v-model="oldPasswd"
      show-password
      style="margin-bottom: 15px"
    >
    </el-input>

    <el-input
      placeholder="新密码"
      prefix-icon="el-icon-lock"
      v-model="newPasswd1"
      show-password
      style="margin-bottom: 15px"
    >
    </el-input>

    <el-input
      placeholder="确认新密码"
      prefix-icon="el-icon-lock"
      v-model="newPasswd2"
      show-password
      style="margin-bottom: 15px"
    >
    </el-input>

    <span slot="footer" class="dialog-footer">
      <el-button type="primary" @click="btnClick">提交</el-button>
    </span>
  </el-dialog>
</template>

<script>
export default {
  name: "password",
  data() {
    return {
      dialogVisible: false,
      oldPasswd: "",
      newPasswd1: "",
      newPasswd2: "",
      root: ""
    };
  },
  methods: {
    btnClick: function() {
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
        .get("http://localhost:8081/setNewPasswd", {
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
            this.axios
              .get("http://localhost:8081/logout", {
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
  mounted() {}
};
</script>

<style scoped></style>
