<template>
  <div class="myMenu">
    <el-menu-item :index="this.index" @click="router_to">
      <i :class="icon"></i>
      <span slot="title">{{label}}</span>
    </el-menu-item>
  </div>
</template>
<script>


  export default {
    name: "myMenu",
    props: {
      index:String,
      label: String,
      icon: String,
      fid: String,
    },
    methods: {
      onCloseClick: function () {
        // 激活删除标签事件
        this.$emit("delete", this.text);
      },
      router_to: function () {

        this.$router.push({name: 'forumsview', params: {fid: this.fid}})

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
