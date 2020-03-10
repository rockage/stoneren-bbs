<template>
  <div style="z-index:99;position: sticky;top: 0;display: flex;justify-content:space-between;background: #DCDFE6">


    <el-button size="mini" type="primary" icon="el-icon-back" v-on:click="$router.back()"></el-button>

    <el-button icon="el-icon-edit-outline" size="mini" @click="popupPost()">发表新主题</el-button>


    <div style="font-size: medium;margin-top: 5px;margin-right: 5px">
      <i class="el-icon-location"></i>
      <a href="javascript:void(0)" @click="jump" style="font-size:small;">{{ forumsName }} </a>
    </div>


  </div>
</template>
<script>
    import Post from "../components/post.vue"

    export default {
        name: "bbsButton",
        inject: ["reload"],
        computed: {
            loginState() {
                return this.$store.getters.loginState
            },
            forumsName() {
                return this.$store.getters.fname
            }
        },
        methods: {
            jump: function () {
                switch (this.$store.getters.rmode) {
                    case "new":
                        this.$router
                            .push({name: "new", params: {page: 1, sortmode: this.$route.params.sortmode}})
                            .catch(err => {
                            })
                        break
                    case "normal":
                        this.$router
                            .push({
                                name: "forumsview",
                                params: {fid: this.$store.getters.fid, page: 1, sortmode: this.$route.params.sortmode}
                            })
                            .catch(err => {
                            })
                        break
                    case "self":
                        this.$router
                            .push({
                                name: "userThreads",
                                params: {
                                    uid: this.$store.getters.viewuid,
                                    page: 1,
                                    sortmode: this.$route.params.sortmode
                                }
                            })
                            .catch(err => {
                            })
                        break
                }
            },
            xclose: function (event) {
                //post.vue窗口已销毁，继续向上报
                this.$emit("ShowDataTable") //改变一次showData的true/false值
            },
            postFinished: function (event) {
                //post.vue完成，需要刷新页面
                this.reload() //页面刷新，从app.vue那里inject了一个reload方法过来
            },
            popupPost: function () {
                if (!document.getElementById("post_box")) {
                    //只允许弹出一个post窗口
                    const PostBox = Vue.extend(Post)
                    let instance = new PostBox({
                        propsData: {
                            xclose: this.xclose,
                            root: this.$root,
                            postFinished: this.postFinished
                        }
                    })
                    let PostEl = instance.$mount().$el
                    document.getElementById("post_location").appendChild(PostEl)
                    this.$emit("ShowDataTable")
                }
            }
        },
        mounted() {
        }
    };
</script>

<style scoped></style>
