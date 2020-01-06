<template>
  <div class="tag">
    <el-menu-item index="1">
      <i class="el-icon-menu"></i>
      <span slot="title">{{text}}</span>
    </el-menu-item>
  </div>
</template>
<script>
export default {
  name: "tag",
  props: {
    // 标签内容
    text: String
  },
  methods: {
    onCloseClick() {
      // 将删除标签事件暴露除去
      this.$emit("delete", this.text);
    }
  }
};
</script>
<style scoped>
.tag {
  display: inline-block;
  margin: 3px;
}
</style>
