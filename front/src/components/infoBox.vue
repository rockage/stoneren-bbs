<template>

  <div class="div-center" id="registerBox">

    <div style="display: flex;justify-content:space-between;margin:0px;background: #a0cfff;">
      <div class="title"><span style="font-size: medium"><i class="el-icon-chat-line-round"></i></span> {{title}}</div>
      <div>
        <a href="javascript:void(0)" @click="close()">
          <i class="el-icon-close" style="font-size:150%"></i>
        </a>
      </div>
    </div>

    <div style="display: flex; flex-direction: column ;margin: 10px;padding: 10px">

      <div class="title" style="white-space: pre-line;">{{message}}</div>
      <el-button @click="close" style="margin-top: 50px">好的</el-button>
    </div>


  </div>
</template>

<script>
    export default {
        name: "infoBox",
        props: {
            root: Object,
            title: String,
            message: String,
        },
        components: {},
        data() {
            return {}
        },
        methods: {
            close: function () {
                this.$destroy() //销毁VUE
            },
        },
        updated() {

        },
        mounted() {

        },
        destroyed() {
            document.body.removeChild(this.$el) //销毁DOM
        }
    }
</script>

<style scoped>

  .div-center {
    border: 1px gray;
    background: #f0f0f0;
    z-index: 99;
    border-radius: 5px;
    position: fixed;
    top: 20%;
    left: 5%;
    height: 20%;
    width: 90%;
    overflow: hidden;
  }

  .elinput {
    margin-bottom: 10px;
  }

  .title {
    font-size: small;
    color: #606266;
    margin-left: 5px;
  }
</style>
