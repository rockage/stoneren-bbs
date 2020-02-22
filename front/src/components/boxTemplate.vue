<template>
  <div class="div-center" ref="divCenter" id="profilebox">

    <div style="display: flex;justify-content:space-between">
      <div class="left"> 我的资料
      </div>
      <div class="right">
        <a href="javascript:void(0)" @click="close()">
          <i class="el-icon-close" style="font-size:150%;"></i>
        </a>
      </div>
    </div>
    <hr style="width: 90%;">


  </div>
</template>

<script>
    export default {
        name: "box",
        props: {
            profileClose: Function,
            profileFinished: Function,
            root: Object //根实例
        },
        data() {
            return {}
        },
        methods: {
            close: function () {
                this.profileClose() //触发父组件的回调
                this.$destroy() //销毁VUE
            },
            done: function () {
                this.profileFinished() //触发父组件的回调
                this.$destroy() //销毁VUE
            }

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
    position: fixed; /*定位*/
    border: 1px gray;
    background: #e4e7ed; /*设置一下背景*/
    z-index: 99;
    border-radius: 5px;
    max-height: 100%;
    text-align: left;
    max-width: 90%;
    min-width: 330px;
    min-height: 300px;
    left: 5%;
    top: 10%;
  }

  .info {
    font-size: smaller;
    color: #606266;
    margin-left: 10px;
    margin-bottom: 10px;
    max-width: 90%;
    min-width: 90%;
  }

  .flex {
    display: flex;
    justify-content: space-between
  }
  .left {
    flex: 1;
    font-size: medium;
    color: #909399;
    margin-top: 5px;
    margin-left: 5px;
  }
  .right {
    margin-top: 5px;
    margin-right: 5px;
  }
</style>

<!-- 调用方法：
    import Box from "./login.vue"
    import Vue from "vue"

            xClose: function (event) {
            },
            xFinished: function (event) {
            },
            boxTest: function () {
                if (!document.getElementById("box")) {
                    const aBox = Vue.extend(Box)
                    let instance = new aBox({
                        propsData: {
                            xClose: this.xClose,
                            xFinished: this.xFinished,
                            root: this.$root,
                        }
                    })
                    document.body.appendChild(instance.$mount().$el)
                }
            },

-->
