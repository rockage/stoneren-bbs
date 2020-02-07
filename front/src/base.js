exports.install = function (Vue) {

  Vue.prototype.loginCheck = function (callback) {//全局函数1
    this.axios.get('http://localhost:8081/secret', {
      //携带cookie提交，mycookiesessionnameid是一个httponly cookie
      withCredentials: true
    }).then((response) => {
      callback(response.data)
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
    for (let i = 0; i < arrcookie.length; i++) {
      let arr = arrcookie[i].split("=")
      if (arr[0] == c_name) {
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
  Vue.prototype.getLocalTime = function (nS) { //返回yyyy-mm-dd HH:MM:SS
    if (isNull(nS)) nS = 1041350400
    let ds = String(new Date(parseInt(nS) * 1000).toISOString())
    ds = ds.replace(/(.+?)T(.+?).\d\d\dZ/, '$1 $2')
    return ds
  }

  Vue.prototype.getLocalDate = function (nS) { //返回yyyy-mm-dd
    if (isNull(nS)) nS = 1041350400 //2003-01-01 00:00:00
    let ds = String(new Date(parseInt(nS) * 1000).toISOString())
    ds = ds.replace(/(.+?)T.+?Z/, '$1')
    //if (ds < '2003-01-01') ds = '上古时期'
    return ds
  }
}

function isNull(e) {
  if (typeof (e) === "undefined") return true
  if (!e && typeof (e) != 'undefined' && e != 0) return true
  if (e !== e) return true
  if (String(e).length === 0) return true
  return false
}
