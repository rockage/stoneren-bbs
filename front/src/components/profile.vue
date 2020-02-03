<template>
  <el-dialog
    title='我的资料'
    :modal="false"
    :visible.sync="dialogVisible"
    width="500px"
    :close-on-click-modal="false"
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
            <el-image id="imgResult" :src="avatar" style="width: 150px;height: 150px;margin-top: 0px;">
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
      <el-select v-model="gender" placeholder="性别" style="width: 100%;margin-bottom: 10px">
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
</template>

<script>
  export default {
    name: "profile",
    data() {
      return {
        options: [
          {value: '1', label: '男'},
          {value: '2', label: '女'},
          {value: '3', label: '不告诉你'}
        ],
        options2: [
          {value: '20', label: '广东省'},
          {value: '1', label: '北京市'},
          {value: '2', label: '天津市'},
          {value: '3', label: '上海市'},
          {value: '4', label: '重庆市'},
          {value: '5', label: '河北省'},
          {value: '6', label: '山西省'},
          {value: '7', label: '台湾省'},
          {value: '8', label: '辽宁省'},
          {value: '9', label: '吉林省'},
          {value: '10', label: '黑龙江省'},
          {value: '11', label: '江苏省'},
          {value: '12', label: '浙江省'},
          {value: '13', label: '安徽省'},
          {value: '14', label: '福建省'},
          {value: '15', label: '江西省'},
          {value: '16', label: '山东省'},
          {value: '17', label: '河南省'},
          {value: '18', label: '湖北省'},
          {value: '19', label: '湖南省'},
          {value: '21', label: '甘肃省'},
          {value: '22', label: '四川省'},
          {value: '23', label: '贵州省'},
          {value: '24', label: '海南省'},
          {value: '25', label: '云南省'},
          {value: '26', label: '青海省'},
          {value: '27', label: '陕西省'},
          {value: '28', label: '广西壮族自治区'},
          {value: '29', label: '西藏自治区'},
          {value: '30', label: '宁夏回族自治区'},
          {value: '31', label: '新疆维吾尔自治区'},
          {value: '32', label: '内蒙古自治区'},
          {value: '33', label: '澳门特别行政区'},
          {value: '34', label: '香港特别行政区'},
          {value: '35', label: '海外'}
        ],
        gender: '',
        location: '',
        born: '',
        mobilePhone: '',
        signature: '',
        email: '',
        dialogVisible: false,
        rootThis: '',
        uid: '',
        uname: '',
        avatar: '',
        avatarSrc: '',
        imgFile: '',
      }
    },
    methods: {
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
        let root = this.rootThis
        let vm = this
        this.axios.get('http://localhost:8081/getProfile', {
          params: {
            uid: root.$store.state.uid,
          }
        })
          .then((response) => {
            if (response.data !== 'error') {
              this.$message.success("资料读取成功")
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
        param.append("uid", this.rootThis.$store.state.uid)
        param.append("password", this.getCookie('password'))
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
      this.rootThis = this.GLOBAL.globalThis
      this.uid = this.rootThis.$store.state.uid
      this.uname = this.rootThis.$store.state.uname
      this.getProfile()


    }
  }


</script>

<style scoped>

</style>
