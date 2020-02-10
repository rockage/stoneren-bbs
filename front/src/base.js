//exports.install = function (Vue) {

export default {
  install(Vue) {

    Vue.prototype.dataInit = function () {
      const vm=this
      return new Promise(function (resolve) {
        vm.axios.get('http://localhost:8081/getForums', {})
          .then((response) => {
            let arr = JSON.parse(response.data)
            vm.GLOBAL.forumsData = arr
            resolve()
          })
      })
    }

    Vue.prototype.loginCheck = function () {
      const vm = this
      return new Promise(function (resolve) {
        vm.axios.get('http://localhost:8081/secret', {
          //携带cookie提交，mycookiesessionnameid是一个httponly cookie
          withCredentials: true
        }).then((response) => {
          resolve(response.data)
        })
      })


    }

    Vue.prototype.setCookie = function (c_name, value) {//全局函数2
      docCookies.setItem(c_name, value, "Tue, 06 Dec 2022 13:11:07 GMT", "/");
    }

    Vue.prototype.getCookie = function (c_name) {//全局函数3
      return docCookies.getItem(c_name)
    }

    Vue.prototype.delCookie = function (c_name) {//全局函数4
      docCookies.removeItem(c_name);
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
}

function isNull(e) {
  if (typeof (e) === "undefined") return true
  if (!e && typeof (e) != 'undefined' && e != 0) return true
  if (e !== e) return true
  if (String(e).length === 0) return true
  return false
}

let docCookies = {
  getItem: function (sKey) {
    return decodeURIComponent(document.cookie.replace(new RegExp("(?:(?:^|.*;)\\s*" + encodeURIComponent(sKey).replace(/[-.+*]/g, "\\$&") + "\\s*\\=\\s*([^;]*).*$)|^.*$"), "$1")) || null;
  },
  setItem: function (sKey, sValue, vEnd, sPath, sDomain, bSecure) {
    if (!sKey || /^(?:expires|max\-age|path|domain|secure)$/i.test(sKey)) {
      return false;
    }
    let sExpires = "";
    if (vEnd) {
      switch (vEnd.constructor) {
        case Number:
          sExpires = vEnd === Infinity ? "; expires=Fri, 31 Dec 9999 23:59:59 GMT" : "; max-age=" + vEnd;
          break;
        case String:
          sExpires = "; expires=" + vEnd;
          break;
        case Date:
          sExpires = "; expires=" + vEnd.toUTCString();
          break;
      }
    }
    document.cookie = encodeURIComponent(sKey) + "=" + encodeURIComponent(sValue) + sExpires + (sDomain ? "; domain=" + sDomain : "") + (sPath ? "; path=" + sPath : "") + (bSecure ? "; secure" : "");
    return true;
  },
  removeItem: function (sKey, sPath, sDomain) {
    if (!sKey || !this.hasItem(sKey)) {
      return false;
    }
    document.cookie = encodeURIComponent(sKey) + "=; expires=Thu, 01 Jan 1970 00:00:00 GMT" + (sDomain ? "; domain=" + sDomain : "") + (sPath ? "; path=" + sPath : "");
    return true;
  },
  hasItem: function (sKey) {
    return (new RegExp("(?:^|;\\s*)" + encodeURIComponent(sKey).replace(/[-.+*]/g, "\\$&") + "\\s*\\=")).test(document.cookie);
  },
  keys: /* optional method: you can safely remove it! */ function () {
    let aKeys = document.cookie.replace(/((?:^|\s*;)[^\=]+)(?=;|$)|^\s*|\s*(?:\=[^;]*)?(?:\1|$)/g, "").split(/\s*(?:\=[^;]*)?;\s*/);
    for (let nIdx = 0; nIdx < aKeys.length; nIdx++) {
      aKeys[nIdx] = decodeURIComponent(aKeys[nIdx]);
    }
    return aKeys;
  }
}
