<template>
  <el-dialog
    title="用户登录"
    :modal="false"
    :visible.sync="dialogVisible"
    width="30%"
    :before-close="handleClose">
    <span>这是一段信息</span>
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
          dialogVisible: true,
        };
      },

    }
</script>

<style scoped>

</style>
