<template>
  <div class="div-center" ref="divCenter" id="profileBox">

    <div style="display: flex;justify-content:space-between">
      <div style="text-align: left;margin-top:10px;margin-left: 10px;">
        <el-button type="primary" size="small" @click="setProfile">保存</el-button>
      </div>
      <div class="right">
        <a href="javascript:void(0)" @click="close()">
          <i class="el-icon-close" style="font-size:150%;"></i>
        </a>
      </div>
    </div>
    <hr style="width: 100%;">

    <div class="info">用户名：{{ uname }} / uid：{{ uid }}</div>
    <div class="info">E-mail：{{ email }}</div>

    <div style="text-align: center">

      <div style="display: flex;justify-content:center; align-items:center;">
        <div style="width:150px;height:150px;margin-top: 10px;margin-left: 5%;margin-right: 5%;">
          <vue-cropper
            autoCrop
            :img="avatarSrc"
            ref="cropper"
            centerBox/>
        </div>
        <div>
          <el-button
            round
            size="small"
            @click="clickFile()"
            style="margin:10px 0">从相册中选择<i class="el-icon-picture-outline el-icon--right"></i>
          </el-button>
        </div>
      </div>
      <hr style="width: 90%;">

      <div style="display: flex;justify-content:center; ">
        <div style="width: 150px">
          <el-select v-model="gender" placeholder="性别" class="info">
            <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
        </div>
        <div style="width: 150px">
          <el-select v-model="location" placeholder="省份" prefix-icon="el-icon-s-custom" class="info">
            <el-option v-for="item in options2" :key="item.value" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
        </div>
      </div>
      <div style="display: flex;justify-content:center; ">
        <div style="width: 150px">
        <el-date-picker v-model="born" type="date" placeholder="出生日期" class="info">
        </el-date-picker>
      </div>
        <div style="width: 150px">
        <el-input placeholder="手机" prefix-icon="el-icon-mobile" v-model="mobilePhone" class="info">
        </el-input>
      </div>
      </div>

      <div style="width: 96%;margin-left: 7px;">
        <el-input placeholder="个性签名" prefix-icon="el-icon-edit-outline" v-model="signature" class="info">
        </el-input>
      </div>




    </div>
  </div>
</template>

<script>
    export default {
        name: "profile",
        props: {
            profileClose: Function,
            profileFinished: Function,
            root: Object //根实例
        },
        computed: {
            passwd() {
                return this.root.$store.getters.passwd;
            },
            uid() {
                return this.root.$store.getters.uid;
            },
            uname() {
                return this.root.$store.getters.uname;
            }
        },
        data() {
            return {
                options: [],
                options2: [],
                gender: "",
                location: "",
                born: "",
                mobilePhone: "",
                signature: "",
                email: "",
                avatar: "",
                avatarSrc: "",
                imgFile: "",
            }
        },
        methods: {
            close: function () {
                this.$destroy() //销毁VUE
            },
            getCropData: function () {
                this.$refs.cropper.getCropData(data => {
                    this.avatar = data;
                });
            },
            clickFile: function () {
                let e = document.createElement("input"); //激活文件选择框
                let vm = this;
                e.type = "file";
                e.accept = "image/png, image/jpeg";
                e.addEventListener("change", function () {
                    vm.loadImg(this.files[0]);
                });
                e.click();
            },
            loadImg: function (file) {
                let vm = this;
                if (window.FileReader) {
                    let fr = new FileReader();
                    fr.readAsDataURL(file); //开始加载文件
                    fr.onload = function (e) {
                        vm.avatarSrc = e.target.result;
                    };
                }
            },
            getProfile: function () {
                let vm = this;
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
                                    .get("getProfile", {
                                        params: {
                                            uid: vm.uid
                                        }
                                    })
                                    .then(response => {
                                        if (response.data !== "error") {
                                            const ret = JSON.parse(response.data);
                                            vm.email = ret[0].email;
                                            vm.avatarSrc = ret[0].avatar;
                                            vm.avatar = ret[0].avatar;
                                            vm.gender = ret[0].gender;
                                            vm.location = ret[0].location;
                                            vm.born = this.getLocalTime(ret[0].born);
                                            vm.mobilePhone = ret[0].mobile;
                                            vm.signature = ret[0].signature;
                                        } else {
                                            this.$message.error("资料读取失败");
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
            },
            setProfile: function () {
                let vm = this

                if (vm.mobilePhone.length > 20) {
                    vm.$message.error("手机号码太长了!")
                    return
                }

                if (vm.signature.length > 30) {
                    vm.$message.error("个性签名太长了!")
                    return
                }

                this.$refs.cropper.getCropData(data => {
                    this.transImage(data, 130, 130, function (data) {
                        let d = new Date(vm.born)
                        let timestamp = Date.parse(d);
                        timestamp = String(timestamp / 1000);
                        if (timestamp === "NaN") timestamp = null;
                        let param = new URLSearchParams(); //axios如不采用URLSearchParams后端无法收到post请求
                        param.append("uid", vm.uid);
                        param.append("password", vm.passwd);
                        param.append("avatar", data);
                        param.append("gender", vm.gender);
                        param.append("location", vm.location);
                        param.append("born", timestamp);
                        param.append("mobilePhone", vm.mobilePhone);
                        param.append("signature", vm.signature);
                        vm.axios
                            .post("setProfile", param)
                            .then(response => {
                                vm.$message.success(response.data);
                                vm.close()
                            });
                    })

                })
            }
        },
        mounted() {
            this.getProfile();
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
