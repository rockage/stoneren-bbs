<template>
  <div class="myMenu">
    <el-menu-item :index="index" @click="router_to">
      <i :class="icon"></i>
      <span slot="title" >{{label}}</span>
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
      icon:String
    },
    methods: {
      onCloseClick:function() {
        // 激活删除标签事件
        this.$emit("delete", this.text);
      },
      router_to:function(){
        console.log(this.icon)
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
