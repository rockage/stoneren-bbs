<template>
  <div class="div-center" ref="divCenter" id="profilebox">

    <div style="display: flex;justify-content:space-between">
      <div class="left"> 用户简介
      </div>
      <div class="right">
        <a href="javascript:void(0)" @click="close()">
          <i class="el-icon-close" style="font-size:150%;"></i>
        </a>
      </div>
    </div>
    <hr style="width: 90%;">

    <div style="text-align: center;">
      <el-image :src="avatar" style="width: 130px;height: 130px;margin-bottom: 10px;border-radius: 9px;">
      </el-image>
    </div>

    <div style="white-space: pre;text-align: left;" class="info">
      {{ message }}
    </div>


    <el-button type="primary" @click="viewThreads()" style="margin-bottom: 10px;margin-left: 10px;">浏览{{ genderName
      }}的全部主题 <i class="el-icon-search"></i>
    </el-button>

  </div>

</template>

<script>
    export default {
        name: "userprofile",
        props: {
            uname: String,
            root: Object //根实例
        },
        data() {
            return {
                message: "",
                uid: "",
                gender: "",
                genderName: "",
                location: "",
                born: "",
                mobilePhone: "",
                signature: "",
                email: "",
                avatar: "",
                posts: "",
                threads: "",
                lastvisited: "",
                regdate: "",
                level: "",
                options: [],
                options2: [],
            }
        },
        methods: {
            close: function () {
                this.$destroy() //销毁VUE
            },
            viewThreads: function () {
                this.root.$store.commit("fname", this.uname);
                this.root.$store.commit("viewuid", this.uid); //为了bbsButton保存这个id
                this.root.$router
                    .push({
                        name: "userThreads",
                        params: {uid: this.uid, sortmode: 'date', page: 1}
                    })
                    .catch(err => {
                        //this.$message.error(err);
                    });
                this.close()
            },
            getProfile: function () {
                let vm = this;
                //Next
                this.axios
                    .get("data/getGender", {})
                    .then(response => {
                        this.options = JSON.parse(response.data);
                        //Next:
                        this.axios
                            .get("data/getLocation", {})
                            .then(response => {
                                this.options2 = JSON.parse(response.data);
                                //Next:
                                this.axios
                                    .get("getUserProfile", {
                                        params: {
                                            uname: this.uname
                                        }
                                    })
                                    .then(response => {
                                        if (response.data !== "error") {
                                            const ret = JSON.parse(response.data);
                                            vm.uid = ret[0].uid;
                                            vm.gender = ret[0].gender;
                                            vm.location = ret[0].location;
                                            vm.born = this.getLocalDate(ret[0].born);
                                            vm.mobilePhone = ret[0].mobile;
                                            vm.signature = ret[0].signature;
                                            vm.email = ret[0].email;
                                            if (isNull(ret[0].avatar)) {
                                                vm.avatar = "/static/avatar.png";
                                            } else {
                                                vm.avatar = ret[0].avatar;
                                            }
                                            vm.posts = ret[0].posts;
                                            vm.threads = ret[0].threads;
                                            vm.lastvisited = this.getLocalDate(ret[0].lastvisited);
                                            vm.regdate = this.getLocalDate(ret[0].regdate);
                                            //Next:
                                            this.axios
                                                .get("data/getLevel", {
                                                    params: {
                                                        posts: vm.posts
                                                    }
                                                })
                                                .then(response => {
                                                    let ret = JSON.parse(response.data);
                                                    vm.level = ret[0].level;
                                                    let msg;
                                                    msg = vm.uname + "，";
                                                    if (isNull(vm.gender)) {
                                                        msg += "性别不详，";
                                                    } else {
                                                        msg += "性别" + findX(vm.options, vm.gender) + "，";
                                                    }
                                                    let g;
                                                    switch (vm.gender) {
                                                        case "1":
                                                            g = "他";
                                                            break;
                                                        case "2":
                                                            g = "她";
                                                            break;
                                                        default:
                                                            g = "它";
                                                    }
                                                    vm.genderName = g;
                                                    if (isNull(vm.location)) {
                                                        msg += "居无定所，";
                                                    } else {
                                                        msg += findX(vm.options2, vm.location) + "人士，";
                                                    }
                                                    let age =
                                                        parseInt(String(new Date(Date.now()).getFullYear())) -
                                                        parseInt(String(new Date(vm.born).getFullYear()));
                                                    let joined =
                                                        parseInt(String(new Date(Date.now()).getFullYear())) -
                                                        parseInt(String(new Date(vm.regdate).getFullYear()));
                                                    if (isNull(vm.born)) {
                                                        msg += "年龄不详，";
                                                    } else {
                                                        msg += "现年" + age + "岁。\r\n";
                                                    }
                                                    if (joined < 1) {
                                                        msg +=
                                                            g +
                                                            "在" +
                                                            vm.regdate +
                                                            "加入本坛，迄今仍不足一年。\r\n";
                                                    } else {
                                                        msg +=
                                                            g +
                                                            "在" +
                                                            vm.regdate +
                                                            "加入本坛，迄今已过去了" +
                                                            joined +
                                                            "年。\r\n";
                                                    }
                                                    if (!isNull(vm.lastvisited)) {
                                                        msg +=
                                                            "最后一次返回论坛的时间是" +
                                                            vm.lastvisited +
                                                            "。\r\n\r\n";
                                                    }
                                                    msg += "  " +
                                                        g +
                                                        "总共发表过" +
                                                        vm.posts +
                                                        "篇帖子，" +
                                                        vm.threads +
                                                        "个主题。\r\n";
                                                    msg +=
                                                        "武功等级似乎已经达到了" +
                                                        vm.level +
                                                        "的境界。\r\n\r\n";
                                                    if (!isNull(vm.mobilePhone))
                                                        msg += "手机：" + vm.mobilePhone + "\r\n";
                                                    if (!isNull(vm.signature))
                                                        msg += "签名：" + vm.signature;
                                                    this.message = msg;
                                                    this.loading = false;
                                                });
                                        } else {
                                            this.loading = false;
                                            this.dialogVisible = false;

                                        }
                                    });
                            });
                    });
            },
            getLocalTime: function (nS) {
                let ds = new Date(parseInt(nS) * 1000)
                    .toDateString()
                    .replace(/:\d{1,2}$/, " ");
                let d = new Date(Date.parse(ds));
                let month = "" + (d.getMonth() + 1);
                let day = "" + d.getDate();
                let year = d.getFullYear();
                if (month.length < 2) month = "0" + month;
                if (day.length < 2) day = "0" + day;
                return [year, month, day].join("-");
            }
        },
        mounted() {
            this.getProfile();
        },
        destroyed() {
            document.body.removeChild(this.$el) //销毁DOM
        }
    }

    function isNull(e) {
        if (typeof e === "undefined") return true;
        if (!e && typeof e != "undefined" && e != 0) return true;
        if (e !== e) return true;
        if (String(e).length === 0) return true;
        return false;
    }

    function findX(arr, keyword) {
        for (let i = 0; i < arr.length; ++i) {
            if (arr[i].value == keyword) {
                return arr[i].label;
            }
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
    width: 90%;
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
