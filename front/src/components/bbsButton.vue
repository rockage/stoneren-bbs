<template>
  <div>
    <el-col :span="1" style="margin-top: 5px;margin-left: 5px">
      <el-button type="primary" size="mini" icon="el-icon-back" circle v-on:click="$router.back()"></el-button>
    </el-col>
    <el-col :span="1" v-show="loginState" style="margin-top: 5px">
      <el-button type="primary"
                 icon="el-icon-plus"
                 size="mini"
                 circle
                 @click="popupPost()"></el-button>
      <div id="post_location"></div>
    </el-col>
    <el-col :span="3" style="margin-top: 10px;margin-left: 0px">
      <span style="font-size: medium;"><i class="el-icon-location"></i></span>
      <span style="color:#606266;font-size: small;font-weight: normal;">{{forumsName}}</span>
    </el-col>
  </div>
</template>

<script>
  import store from '../store'

  export default {
    name: "bbsButton",
    store,
    computed: {
      loginState() {
        return store.state.loginState
      },
      forumsName(){
        return store.state.forumsName
      }
    },
    methods:{
      popupPost: function () {
        if (!document.getElementById("post_box")) this.$post() //只允许弹出一个post窗口
      },
    },mounted() {

    }
  }
</script>

<style scoped>

</style>
