<template>
  <div>
    <el-col :span="1" style="margin-top: 5px;margin-left: 5px">
      <el-button
        type="primary"
        size="mini"
        icon="el-icon-back"
        circle
        v-on:click="$router.back()"
      ></el-button>
    </el-col>
    <el-col :span="1" v-show="loginState" style="margin-top: 5px">
      <el-button
        type="primary"
        icon="el-icon-plus"
        size="mini"
        circle
        @click="popupPost()"
      ></el-button>
    </el-col>
    <el-col :span="3" style="margin-top: 10px;margin-left: 0px">
      <span style="font-size: medium;"><i class="el-icon-location"></i></span>
      <span style="color:#606266;font-size: small;font-weight: normal;">{{
        forumsName
      }}</span>
    </el-col>
  </div>
</template>

<script>
import Post from "../components/post.vue";
import Vue from "vue";

export default {
  name: "bbsButton",
  computed: {
    loginState() {
      return this.$store.getters.loginState;
    },
    forumsName() {
      return this.$store.getters.fname;
    }
  },
  methods: {
    xclose: function(event) {
      //post窗口已销毁，继续向上报
      this.$emit("ShowDataTable"); //改变一次showData的true/false值
    },
    popupPost: function() {
      if (!document.getElementById("post_box")) {
        //只允许弹出一个post窗口
        const PostBox = Vue.extend(Post);
        console.log(this.$root)
        let instance = new PostBox({
          propsData: {
            xclose: this.xclose,
            root: this.$root
          }
        });
        let PostEl = instance.$mount().$el;
        document.getElementById("post_location").appendChild(PostEl);
        this.$emit("ShowDataTable");
      }
    }
  },
  mounted() {}
};
</script>

<style scoped></style>
