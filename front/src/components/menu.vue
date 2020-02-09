<template>
  <div class="myMenu">
    <el-menu-item :index="this.index" @click="router_to">
      <i :class="icon"></i>
      <span slot="title">{{label}}</span>
    </el-menu-item>
    {{test}}
  </div>
</template>
<script>


  export default {
    name: "myMenu",
    props: {
      index: '',
      label: '',
      icon: '',
      fid: '',
      test: '',
    },
    methods: {
      onCloseClick: function () {
        // 激活删除标签事件
        this.$emit("delete", this.text);
      },
      router_to: function () {
        this.GLOBAL.fid = this.fid
        this.$router.push({name: 'forumsview', params: {fid: this.fid}}).catch(err => {
          console.log(err.name)
        })
      }
    }
  }
</script>
<style scoped>
  .myMenu {
    margin: 0px;
    text-align: left;
  }
</style>
