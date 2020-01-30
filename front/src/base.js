exports.install = function (Vue) {
  Vue.prototype.loginCheck = function (autologin, username, password, callback) {//全局函数1

    console.log(username)
    console.log(password)

    this.axios.get('http://localhost:8081/login', {
      params: {
        username: username,
        password: password,
        autoLogin: this.autoLogin,
      }
    })
      .then((response) => {
        if (response.status !== 200) {
          this.$message.error('通讯失败，请检查网络。')
          callback(false)
        }
        if (response.data === 'not found') {
          this.$message.warning("请输入正确的用户名和密码。")
          callback(false)

        }
        const ret = JSON.parse(response.data)
        if (ret[0].uid) {
          this.$message.success("恭喜你，登录成功了。")
          callback(true)
        }
      })

  }

  Vue.prototype.setCookie = function (c_name, value, expiredays) {//全局函数2
    let exdate = new Date();
    exdate.setDate(exdate.getDate() + expiredays);
    document.cookie = c_name + "=" + escape(value) + ((expiredays == null) ? "" : ";expires=" + exdate.toUTCString());
  }

  Vue.prototype.getCookie = function (name) {//全局函数3
    let arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
    if (arr = document.cookie.match(reg))
      return (arr[2]);
    else
      return null;
  }

  Vue.prototype.delCookie = function (name) {//全局函数4
    let exp = new Date();
    exp.setTime(exp.getTime() - 1);
    let cval = this.getCookie(name);
    if (cval != null)
      document.cookie = name + "=" + cval + ";expires=" + exp.toUTCString();
  }
}
