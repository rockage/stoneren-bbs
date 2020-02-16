import Vue from 'vue'
import Login from "./components/login.vue"
import Password from "./components/password.vue"
import UserProfile from "./components/userprofile.vue"
import ProFile from "./components/profile.vue"
//以下所有通过动态方法new新建的Vue实例都不能访问$root, 需创建时传递

function LoginBox() {
  const LoginBox = Vue.extend(Login)
  let data = { root: this.$root }
  let instance = new LoginBox({ data })
  let LoginEl = instance.$mount().$el
  document.body.appendChild(LoginEl)
  Vue.nextTick(() => {
    instance.dialogVisible = true
  })
}


function ProFileBox() {
  const ProFileBox = Vue.extend(ProFile)
  let data = { root: this.$root }
  let instance = new ProFileBox({ data })
  let ProFileEl = instance.$mount().$el
  document.body.appendChild(ProFileEl)
  Vue.nextTick(() => {
    instance.dialogVisible = true
  })
}


function PasswordBox() {
  const PasswordBox = Vue.extend(Password)
  let data = { root: this.$root }
  let instance = new PasswordBox({ data })
  let LoginEl = instance.$mount().$el
  document.body.appendChild(LoginEl)
  Vue.nextTick(() => {
    instance.dialogVisible = true
  })
}

function UserProfileBox(datas) {
  const UserProfileBox = Vue.extend(UserProfile)
  let data = {}
  data = datas.uname
  data.root = this.$root
  let instance = new UserProfileBox({ data })
  let UserProfileEl = instance.$mount().$el
  document.body.appendChild(UserProfileEl)
  Vue.nextTick(() => {
    instance.dialogVisible = false
  })
}


export { LoginBox, PasswordBox, ProFileBox, UserProfileBox }
