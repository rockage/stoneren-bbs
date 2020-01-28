exports.install = function (Vue) {
  Vue.prototype.loginCheck = function (callback) {//全局函数1
    setTimeout(() => {

      this.$store.commit('loginStateChanged', false)
      callback(true)
    }, 2000)
  }

  Vue.prototype.text2 = function () {//全局函数2
    alert('执行成功2')
  }
}
