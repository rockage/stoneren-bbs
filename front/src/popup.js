import Vue from 'vue'
import Login from "./components/login.vue"
import Password from "./components/password.vue"
import Profile from "./components/profile.vue"
import UserProfile from "./components/userprofile.vue"

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

function ProfileBox() {
  const ProfileBox = Vue.extend(Profile)
  let instance = new ProfileBox()
  let ProfileEl = instance.$mount().$el
  document.body.appendChild(ProfileEl)
  Vue.nextTick(() => {
    instance.dialogVisible = true
  })
}

function UserProfileBox(data) {
  const UserProfileBox = Vue.extend(UserProfile)
  let instance = new UserProfileBox({data})
  let UserProfileEl = instance.$mount().$el
  document.body.appendChild(UserProfileEl)
  Vue.nextTick(() => {
    instance.dialogVisible = true
  })
}


export {LoginBox, PasswordBox, ProfileBox, UserProfileBox}