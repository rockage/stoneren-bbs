<template>
  <el-dialog
    title='密码修改'
    :modal="false"
    :visible.sync="dialogVisible"
    width="30%"
  >
    <el-input
      placeholder="当前密码"
      prefix-icon="el-icon-s-custom"
      v-model="oldPasswd" show-password style="margin-bottom: 15px">
    </el-input>

    <el-input
      placeholder="新密码"
      prefix-icon="el-icon-lock"
      v-model="newPasswd1" show-password style="margin-bottom: 15px">
    </el-input>

    <el-input
      placeholder="确认新密码"
      prefix-icon="el-icon-lock"
      v-model="newPasswd2" show-password style="margin-bottom: 15px">
    </el-input>


    <span slot="footer" class="dialog-footer">
              <el-button type="primary" @click="check">确 定</el-button>
              </span>
  </el-dialog>
</template>

<script>
    export default {
        name: "password",
      data() {
        return {
          oldPasswd: '',
          newPasswd1: '',
          newPasswd2: '',
          rootThis: '',
        };
      },
    }
</script>

<style scoped>

</style>
