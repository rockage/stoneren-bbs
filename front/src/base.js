exports.install = function (Vue) {
  Vue.prototype.loginCheck = function (silence, username, password, callback) {//全局函数1

    this.axios.get('http://localhost:8081/login', {
      params: {
        username: username,
        password: password,
        autoLogin: this.autoLogin,
      }
    })
      .then((response) => {
        if (response.status !== 200) {

          if (!silence) { //自动登录不做信息提示
            this.$message.error('通讯失败，请检查网络。')
          }

          callback(false)
        }
        if (response.data === 'not found') {
          if (!silence) {
            this.$message.warning("请输入正确的用户名和密码。")
          }
          callback(false)

        }
        const ret = JSON.parse(response.data)
        if (ret[0].uid) {
          if (!silence) {
            this.$message.success("恭喜你，登录成功了。")
          }
          callback(true, ret[0].uid)
        }
      })

  }

  Vue.prototype.setCookie = function (c_name, value) {//全局函数2
    let expiredays = 30
    let exdate = new Date();
    exdate.setDate(exdate.getDate() + expiredays);
    document.cookie = c_name + "=" + escape(value) + ";expires=" + exdate.toUTCString() + ";path=/"
  }

  Vue.prototype.getCookie = function (c_name) {//全局函数3
    let strcookie = document.cookie
    let arrcookie = strcookie.split("; ")
    //遍历匹配
    for ( let i = 0; i < arrcookie.length; i++) {
      let arr = arrcookie[i].split("=")
      if (arr[0] == c_name){
        return arr[1]
      }
    }
    return null
  }

  Vue.prototype.delCookie = function (c_name) {//全局函数4
    let exdate = new Date();
    exdate.setTime(exdate.getTime() - 1);
    document.cookie = c_name + "=null;expires:-1;path=/"
  }
}
