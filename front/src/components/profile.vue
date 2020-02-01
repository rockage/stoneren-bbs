<template>
  <el-dialog
    title='我的资料'
    :modal="false"
    :visible.sync="dialogVisible"
    width="30%"
    :close-on-click-modal="false"
  >
    <el-card :body-style="{ padding: '10px' }" style="margin-bottom: 10px;height: 40px;">
      <div><span style="color:#909399;">用户名：{{uname}} / uid：{{uid}}</span></div>
    </el-card>
    <el-card :body-style="{ padding: '10px' }" style="margin-bottom: 10px;height: 40px;">
      <div><span style="color:#909399;">E-mail：{{email}} </span></div>
    </el-card>
    <el-card :body-style="{ padding: '10px' }" style="margin-bottom: 10px;height: 200px;">
      <table border="0" width="100%">
        <tr>
          <td width="50%" align="center">
            <div style="width:150px;height:150px">
              <vue-cropper autoCrop img="/static/avatar.png" ref="cropper" centerBox/>
            </div>
            <span style="color:#909399;font-size: x-small"> (鼠标滚轮缩放图片)</span>
          </td>
          <td width="50%">
            <el-button round size="small">上传<i class="el-icon-upload el-icon--right"></i></el-button>
            <br/><br/>
            <el-button round size="small">决定<i class="el-icon-check el-icon--right"></i></el-button>
          </td>
        </tr>
      </table>
    </el-card>




    <div>
      <el-select v-model="value" placeholder="性别" style="width: 100%;margin-bottom: 10px">
        <el-option
          v-for="item in options"
          :key="item.value"
          :label="item.label"
          :value="item.value">
        </el-option>
      </el-select>
    </div>

    <div class="block">
      <el-select v-model="value" placeholder="居住地" prefix-icon="el-icon-s-custom"
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
      <el-date-picker
        v-model="value"
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
              <el-button type="primary" @click="btnClick">保存</el-button>
              </span>
  </el-dialog>
</template>

<script>

  export default {

    name: "profile",
    data() {
      return {
        options: [{
          value: '选项1',
          label: '黄金糕'
        }, {
          value: '选项2',
          label: '双皮奶'
        }, {
          value: '选项3',
          label: '蚵仔煎'
        }, {
          value: '选项4',
          label: '龙须面'
        }, {
          value: '选项5',
          label: '北京烤鸭'
        }],
        value: '',
        mobilePhone: '',
        signature: '',
        email: '',
        dialogVisible: false,
        rootThis: '',
        uid: '',
        uname: '',


      }
    },
    computed: {
      /*
      uid: function () {
       // console.log(this.rootThis.$store.state.uid)
        //return this.rootThis.$store.state.uid
      },
      uname: function () {
        //return this.rootThis.$store.state.uname
      }

       */
    },
    methods: {
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

            } else {
              this.$message.error("资料读取失败")
            }
          })
      },
      btnClick: function () {

      }
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
