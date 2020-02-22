<template>
  <div>
    <table border="0" style="width: 100%;background-color: #EBEEF5">
      <tbody>
      <tr>
        <td width="150">
          <img
            src="/static/logo3.png"
            style="margin-top: 10px;margin-left: 10px"
          />
          <button @click="resetPosts" style="display: none">统计发帖数</button>
        </td>
        <td>
          <div style="font-size: x-large;max-width: 100%">中国石凳联盟论坛</div>
        </td>

        <td style="text-align: right;margin-right: 20px">
            <span v-show="!loginState">
              <el-button type="text" @click="loginClick">登录</el-button>
              <el-button type="text" @click="loginClick">注册</el-button>
            </span>

          <span v-show="loginState">
              <el-dropdown @command="handleCommand">
                <span class="el-dropdown-link">
                  {{ uname }}<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item command="1">我的资料</el-dropdown-item>
                  <el-dropdown-item command="2">我的主题</el-dropdown-item>
                  <el-dropdown-item command="3">修改密码</el-dropdown-item>
                  <el-dropdown-item command="4">退出登录</el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </span>
        </td>
        <td width="20px"></td>
      </tr>
      <tr>
        <td></td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
    import LoginBox from "./login.vue"
    import ProfileBox from "./profile.vue"
    import PasswdBox from "./password.vue"
    import Vue from "vue"

    export default {
        name: "bbsHeader",
        data() {
            return {};
        },
        components: {
        },
        computed: {
            loginState() {
                return this.$store.state.loginState;
            },
            uid() {
                return this.$store.state.uid;
            },
            uname() {
                return this.$store.state.uname;
            }
        },
        methods: {
            resetPosts: function () {
                this.axios
                    .get("http://localhost:8081/resetPosts", {
                        params: {}
                    })
                    .then(response => {
                        this.tableData = JSON.parse(response.data);
                        this.loading = false;
                    });
            },
            handleCommand: function (command) {
                switch (command) {
                    case "1":
                        this.profileClick()
                        break;
                    case "2":
                        this.$router
                            .push({
                                name: "userThreads",
                                params: {uid: this.$store.getters.uid}
                            })
                            .catch(err => {
                                this.$message.error(err);
                            });
                        break;
                    case "3":
                        this.passwdClick()
                        break;
                    case "4":
                        let vm = this
                        this.axios
                            .get("http://localhost:8081/logout", {
                                withCredentials: true
                            })
                            .then(response => {
                                vm.delCookie("local");
                                vm.$store.commit("loginState", false);
                                vm.$message.success(response.data);
                            });
                }
            },
            passwdClick:function(){
                if (!document.getElementById("passwdBox")) {
                    const Passwd = Vue.extend(PasswdBox)
                    let instance = new Passwd({
                        propsData: {
                            root: this.$root,
                        }
                    })
                    document.body.appendChild(instance.$mount().$el)
                }
            },
            profileClick:function(){
                if (!document.getElementById("profileBox")) {
                    const profile = Vue.extend(ProfileBox)
                    let instance = new profile({
                        propsData: {
                            profileClose: null,
                            profileFinished: null,
                            root: this.$root,
                        }
                    })
                    document.body.appendChild(instance.$mount().$el)
                }
            },
            loginClick: function () {
                if (!document.getElementById("loginbox")) {
                    const loginbox = Vue.extend(LoginBox)
                    let instance = new loginbox({
                        propsData: {
                            loginClose: null, //无需返回值的弹窗
                            loginFinished: null,
                            root: this.$root,
                        }
                    })
                    document.body.appendChild(instance.$mount().$el)
                }
            }
        },
        mounted() {

        }
    };
</script>

<style scoped></style>
