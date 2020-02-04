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
      index: '',
      label: '',
      icon: '',
      fid: '',
    },
    methods: {
      onCloseClick: function () {
        // 激活删除标签事件
        this.$emit("delete", this.text);
      },
      router_to: function () {

        console.log(this.fid)

        if (this.fid !== '0') {
          this.$router.push({name: 'forumsview', params: {fid: this.fid}}).catch(err => {
          })
        } else {
          this.$router.push({name: 'bbs'}).catch(err => {
          })
        }

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
