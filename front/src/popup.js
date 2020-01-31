import Vue from 'vue'
import Login from "./components/login.vue"

function LoginBox() {
  const LoginBox = Vue.extend(Login)
  let instance = new LoginBox()
  let LoginEl = instance.$mount().$el
  document.body.appendChild(LoginEl)
  Vue.nextTick(() => {
    instance.dialogVisible = true
  })
}

export {LoginBox}
