<template>
  <el-dialog
    :title=msg
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
      prefix-icon="el-icon-s-order"
      v-model="inputPasswd" show-password style="margin-bottom: 15px">
    </el-input>
    <span>自动登录：</span>
    <el-switch
      v-model="autoLogin"
      active-color="#13ce66"
      inactive-color="#909399">
    </el-switch>

    <span slot="footer" class="dialog-footer">
              <el-button @click="dialogVisible = false">取 消</el-button>
              <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
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
        autoLogin: false,
      };
    },
    methods: {
      init: function (options) {
        this.msg = options.msg
      },
    }

  }
</script>

<style scoped>

</style>
