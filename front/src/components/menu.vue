<template>
  <div class="menu">
    <el-menu-item index="1">
      <i class="el-icon-menu"></i>
      <span slot="title">{{text}}</span>
    </el-menu-item>
  </div>
</template>
<script>
export default {
  name: "menu",
  props: {
        // 标签内容
    text: String
  },
  methods: {
    onCloseClick() {
      // 激活删除标签事件
      this.$emit("delete", this.text);
    }
  }
};
</script>
<style scoped>
.menu {
  display: inline-block;
  margin: 3px;
}
</style>
