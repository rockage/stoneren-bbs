<template>
  <div>
    <el-dialog
      title='我的资料'
      :modal="false"
      :visible.sync="dialogVisible"
      v-if="dialogVisible"
      width="500px"
      :close-on-click-modal="false"
      @close="handleClose()"
    >
      <el-card :body-style="{ padding: '10px' }" style="margin-bottom: 10px;height: 40px;">
        <div><span style="color:#909399;">用户名：{{uname}} / uid：{{uid}}</span></div>
      </el-card>
      <el-card :body-style="{ padding: '10px' }" style="margin-bottom: 10px;height: 40px;">
        <div><span style="color:#909399;">E-mail：{{email}} </span></div>
      </el-card>
      <el-card :body-style="{ padding: '10px' }" style="margin-bottom: 10px;height: 200px;">
        <table border="0">
          <tr>
            <td align="center" style="max-width: 160px;min-width: 160px">
              <div style="width:150px;height:150px;margin-top: 10px">
                <vue-cropper autoCrop :img="avatarSrc" ref="cropper" centerBox/>
              </div>
              <span style="color:#909399;font-size: x-small"> (鼠标滚轮缩放图片)</span>
            </td>
            <td align="center" style="max-width: 100px;min-width: 100px">
              <el-button round size="small" @click="clickFile()" style="margin:10px 0">上传<i
                class="el-icon-picture-outline el-icon--right"></i></el-button>
              <br/><br/>
              <el-button round size="small" @click="getCropData()" style="margin:10px 0">使用<i
                class="el-icon-check el-icon--right"></i></el-button>
            </td>
            <td style="max-width: 160px;min-width: 160px">
              <el-image id="imgResult" :src="avatar"
                        style="width: 150px;height: 150px;margin-top: 0px;border-radius: 9px;">
                <div slot="error" class="image-slot"
                     style="margin-left:10px;color:#909399;font-size: x-small;text-align: center;line-height:150px;">
                  (头像预览)
                </div>
              </el-image>
            </td>
          </tr>
        </table>
      </el-card>
      <div>
        <el-select v-model="gender"
                   placeholder="性别"
                   style="width: 100%;margin-bottom: 10px">
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </div>
      <div class="block">
        <el-select v-model="location" placeholder="省份" prefix-icon="el-icon-s-custom"
                   style="width: 100%;margin-bottom: 10px">
          <el-option
            v-for="item in options2"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </div>
      <div class="block">
        <el-date-picker
          v-model="born"
          type="date"
          placeholder="出生日期" style="width: 100%;margin-bottom: 10px">
        </el-date-picker>
      </div>
      <div class="block">
        <el-input
          placeholder="手机"
          prefix-icon="el-icon-mobile"
          v-model="mobilePhone" style="width: 100%;margin-bottom: 10px">
        </el-input>
      </div>
      <div class="block">
        <el-input
          placeholder="个性签名"
          prefix-icon="el-icon-edit-outline"
          v-model="signature" style="width: 100%;margin-bottom: 10px">
        </el-input>
      </div>
      <span slot="footer" class="dialog-footer">
              <el-button type="primary" @click="setProfile">保存</el-button>
              </span>
    </el-dialog>
  </div>
</template>

<script>

  export default {
    name: "profile",
    data() {
      return {
        dialogVisible: false,
        options: [],
        options2: [],
        gender: '',
        location: '',
        born: '',
        mobilePhone: '',
        signature: '',
        email: '',
        avatar: '',
        avatarSrc: '',
        imgFile: '',
      }
    },
    computed: {
      passwd() {
        return this.GLOBAL.root.$store.state.passwd
      },
      uid() {
        return this.GLOBAL.root.$store.state.uid
      },
      uname() {
        return this.GLOBAL.root.$store.state.uname
      },
    },
    methods: {
      handleClose: function () {

      },
      getCropData: function () {
        this.$refs.cropper.getCropData(data => {
          this.avatar = data
        })
      },
      clickFile: function () {
        let e = document.createElement("input") //激活文件选择框
        let vm = this
        e.type = "file"
        e.accept = "image/png, image/jpeg"
        e.addEventListener('change', function () {
          vm.loadImg(this.files[0])
        })
        e.click()
      },
      loadImg: function (file) {
        let vm = this
        if (window.FileReader) {
          let fr = new FileReader()
          fr.readAsDataURL(file) //开始加载文件
          fr.onload = function (e) {
            vm.avatarSrc = e.target.result
          }
        }
      },
      getProfile: function () {
        let vm = this
        this.axios.get('http://localhost:8081/data/getGender', {}).then((response) => {
          this.options = JSON.parse(response.data)
          //Next:
          this.axios.get('http://localhost:8081/data/getLocation', {}).then((response) => {
            this.options2 = JSON.parse(response.data)
            //Next:
            this.axios.get('http://localhost:8081/getProfile', {
              params: {
                uid: vm.uid,
              }
            }).then((response) => {
              if (response.data !== 'error') {
                const ret = JSON.parse(response.data)
                vm.email = ret[0].email
                vm.avatarSrc = ret[0].avatar
                vm.avatar = ret[0].avatar
                vm.gender = ret[0].gender
                vm.location = ret[0].location
                vm.born = this.getLocalTime(ret[0].born)
                vm.mobilePhone = ret[0].mobile
                vm.signature = ret[0].signature
              } else {
                this.$message.error("资料读取失败")
              }
            })

          })
        })
      },
      getLocalTime: function (nS) {
        let ds = new Date(parseInt(nS) * 1000).toDateString().replace(/:\d{1,2}$/, ' ')
        let d = new Date(Date.parse(ds))

        let month = '' + (d.getMonth() + 1)
        let day = '' + d.getDate()
        let year = d.getFullYear()
        if (month.length < 2) month = '0' + month;
        if (day.length < 2) day = '0' + day;
        return [year, month, day].join('-');
      },
      setProfile: function () {
        let d = new Date(this.born), timestamp = Date.parse(d)
        timestamp = String(timestamp / 1000)
        if (timestamp === 'NaN') timestamp = null
        let param = new URLSearchParams() //axios如不采用URLSearchParams后端无法收到post请求
        param.append("uid", this.uid)
        param.append("password", this.passwd)
        param.append("avatar", this.avatar)
        param.append("gender", this.gender)
        param.append("location", this.location)
        param.append("born", timestamp)
        param.append("mobilePhone", this.mobilePhone)
        param.append("signature", this.signature)
        this.axios.post('http://localhost:8081/setProfile', param)
          .then((response) => {
            this.$message.success(response.data)
          })
      },
    },
    mounted() {
      this.root = this.GLOBAL.root
      this.getProfile()
    }
  }
</script>

<style scoped>

</style>
