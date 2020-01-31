import Vue from 'vue'
import Login from "./components/login.vue"
import Password from "./components/password.vue"

function LoginBox() {
  const LoginBox = Vue.extend(Login)
  let instance = new LoginBox()
  let LoginEl = instance.$mount().$el
  document.body.appendChild(LoginEl)
  Vue.nextTick(() => {
    instance.dialogVisible = true
  })
}

function PasswordBox() {
  const PasswordBox = Vue.extend(Password)
  let instance = new PasswordBox()
  let LoginEl = instance.$mount().$el
  document.body.appendChild(LoginEl)
  Vue.nextTick(() => {
    instance.dialogVisible = true
  })
}

export {LoginBox, PasswordBox}
