<template>
  <el-dialog
    title='人物小传'
    :modal="false"
    :visible.sync="dialogVisible"
    v-if="dialogVisible"
    width="500px"
    :close-on-click-modal="false">
    <div style="text-align: center;">
      <el-image :src="avatar" style="width: 150px;height: 150px;margin-bottom: 10px;border-radius: 9px;">
      </el-image>
      <div style="left:50%;top:50%;width: 100px;height: 100px;z-index: 99;position:fixed;" v-loading="loading"></div>
    </div>
    <el-card :body-style="{ padding: '10px' }" style="margin-bottom: 10px;height: 80%;">
      <div style="white-space: pre;line-height: 50px;font-family: 'Microsoft YaHei',sans-serif"><span
        style="color:#909399;">{{message}}</span>
      </div>
    </el-card>
    <div style="text-align: right">
      <el-button @click="viewThreads()">浏览{{genderName}}的主题 <i class="el-icon-notebook-1"></i></el-button>
      <el-button @click="dialogVisible=false">关闭 <i class="el-icon-circle-close"></i></el-button>
    </div>


  </el-dialog>
</template>

<script>
  export default {
    name: "userprofile",
    data() {
      return {
        message: '',
        dialogVisible: false,
        rootThis: '',
        uid: '',
        gender: '',
        genderName: '',
        location: '',
        born: '',
        mobilePhone: '',
        signature: '',
        email: '',
        uname: '',
        avatar: '',
        posts: '',
        threads: '',
        lastvisited: '',
        regdate: '',
        level: '',
        options: [],
        options2: [],
        loading: false,
      }
    },
    methods: {
      viewThreads: function () {
        this.rootThis.$router.push(
          {
            name: 'userThreads',
            params: {uid: this.uid}
          }).catch(err => {
          this.$message.error(err)
        })
        this.dialogVisible = false
      },
      getProfile: function () {
        let vm = this
        //Next
        this.axios.get('http://localhost:8081/data/getGender', {}).then((response) => {
          this.options = JSON.parse(response.data)
          //Next:
          this.axios.get('http://localhost:8081/data/getLocation', {}).then((response) => {
            this.options2 = JSON.parse(response.data)
            //Next:
            this.axios.get('http://localhost:8081/getUserProfile', {
              params: {
                uname: this.uname,
              }
            }).then((response) => {
              if (response.data !== 'error') {
                this.dialogVisible = true
                const ret = JSON.parse(response.data)
                vm.uid = ret[0].uid
                vm.gender = ret[0].gender
                vm.location = ret[0].location
                vm.born = this.getLocalDate(ret[0].born)
                vm.mobilePhone = ret[0].mobile
                vm.signature = ret[0].signature
                vm.email = ret[0].email
                if (isNull(ret[0].avatar)) {
                  vm.avatar = '/static/avatar.png'
                } else {
                  vm.avatar = ret[0].avatar
                }
                vm.posts = ret[0].posts
                vm.threads = ret[0].threads
                vm.lastvisited = this.getLocalDate(ret[0].lastvisited)
                vm.regdate = this.getLocalDate(ret[0].regdate)
                //Next:
                this.axios.get('http://localhost:8081/data/getLevel', {
                  params: {
                    posts: vm.posts,
                  }
                }).then((response) => {
                  let ret = JSON.parse(response.data)
                  vm.level = ret[0].level
                  let msg
                  msg = vm.uname + '，'
                  if (isNull(vm.gender)) {
                    msg += '性别不详，'
                  } else {
                    msg += '性别' + findX(vm.options, vm.gender) + '，'

                  }
                  let g
                  switch (vm.gender) {
                    case "1":
                      g = '他'
                      break
                    case "2":
                      g = '她'
                      break
                    default:
                      g = '它'
                  }
                  vm.genderName = g
                  if (isNull(vm.location)) {
                    msg += '居无定所，'
                  } else {
                    msg += findX(vm.options2, vm.location) + '人士，'
                  }
                  let age = parseInt(String(new Date(Date.now()).getFullYear())) - parseInt(String(new Date(vm.born).getFullYear()))
                  let joined = parseInt(String(new Date(Date.now()).getFullYear())) - parseInt(String(new Date(vm.regdate).getFullYear()))
                  if (isNull(vm.born)) {
                    msg += '年龄不详，'
                  } else {
                    msg += '现年' + age + '岁。\r\n'
                  }
                  if (joined < 1) {
                    msg += g + '在' + vm.regdate + '加入本坛，迄今仍不足一年。\r\n'
                  } else {
                    msg += g + '在' + vm.regdate + '加入本坛，迄今已过去了' + joined + '年。\r\n'
                  }
                  if (!isNull(vm.lastvisited)) {
                    msg += g + '最后一次返回本坛的时间是' + vm.lastvisited + '。\r\n'
                  }
                  msg += g + '总共发表过' + vm.posts + '篇帖子，' + vm.threads + "个主题。\r\n"
                  msg += g + '的武功等级似乎达到了' + vm.level + '的境界。\r\n'
                  if (!isNull(vm.mobilePhone)) msg += '飞鸽传书：' + vm.mobilePhone + '\r\n'
                  if (!isNull(vm.signature)) msg += '人生格言：' + vm.signature
                  this.message = msg
                  this.loading = false
                })
              } else {
                this.loading = false
                this.dialogVisible = false
                this.$login()

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
      this.loading = true
      this.getProfile()
      this.rootThis = this.GLOBAL.globalThis
    }
  }

  function isNull(e) {
    if (typeof (e) === "undefined") return true
    if (!e && typeof (e) != 'undefined' && e != 0) return true
    if (e !== e) return true
    if (String(e).length === 0) return true
    return false
  }

  function findX(arr, keyword) {
    for (let i = 0; i < arr.length; ++i) {
      if (arr[i].value == keyword) {
        return arr[i].label
      }
    }
  }

</script>

<style scoped>

</style>
