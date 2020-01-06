<template>
  <div class="myMenu">
    <el-menu-item :index="index">
      <i class="el-icon-star-on"></i>
      <span slot="title">{{label}}</span>
    </el-menu-item>
  </div>
</template>
<script>
  export default {
    name: "myMenu",
    props: {
      // 标签内容
      index:String,
      fid: String,
      label: String,
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
  .myMenu {
    margin: 0px;
    text-align: left;
  }
</style>
