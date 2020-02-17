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

      <a href="javascript:void(0)" @click="jump" style="font-size:small;">{{
        forumsName
      }}</a>
    </el-col>
  </div>
</template>

<script>
import Post from "../components/post.vue";
import Vue from "vue";

export default {
  name: "bbsButton",
  inject: ["reload"],
  computed: {
    loginState() {
      return this.$store.getters.loginState;
    },
    forumsName() {
      return this.$store.getters.fname;
    }
  },
  methods: {
    jump: function() {
      console.log(this.$router);
      switch (this.$store.getters.rmode) {
        case "new":
          this.$router
            .push({ name: "new", params: { page: 1 } })
            .catch(err => {});
          break;
        case "normal":
          this.$router
            .push({
              name: "forumsview",
              params: { fid: this.$store.getters.fid, page: 1 }
            })
            .catch(err => {});
          break;
        case "self":
          this.$router
            .push({
              name: "userThreads",
              params: { uid: this.$store.getters.viewuid, page: 1 }
            })
            .catch(err => {});
          break;
      }
    },
    xclose: function(event) {
      //post.vue窗口已销毁，继续向上报
      this.$emit("ShowDataTable"); //改变一次showData的true/false值
    },
    postFinished: function(event) {
      //post.vue完成，需要刷新页面
      this.reload(); //页面刷新，从app.vue那里inject了一个reload方法过来
    },
    popupPost: function() {
      if (!document.getElementById("post_box")) {
        //只允许弹出一个post窗口
        const PostBox = Vue.extend(Post);
        let instance = new PostBox({
          propsData: {
            xclose: this.xclose,
            root: this.$root,
            postFinished: this.postFinished 
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
