import Vue from 'vue'
import Login from "./components/login.vue"
import Password from "./components/password.vue"
import UserProfile from "./components/userprofile.vue"
import Post from "./components/post.vue"
import ProFile from "./components/profile.vue"

//以下所有通过动态方法new新建的Vue实例是不能访问运行中的$root的，在动态实例中，$root就等于它自身的this，
//如需访问运行中真正的根实例，需要一个全局变量来传递。
function LoginBox() {
  const LoginBox = Vue.extend(Login)
  let instance = new LoginBox()
  let LoginEl = instance.$mount().$el
  document.body.appendChild(LoginEl)
  Vue.nextTick(() => {
    instance.dialogVisible = true
  })
}


function ProFileBox() {
  const ProFileBox = Vue.extend(ProFile)
  let instance = new ProFileBox()
  let ProFileEl = instance.$mount().$el
  document.body.appendChild(ProFileEl)
  Vue.nextTick(() => {
    instance.dialogVisible = true
  })
}


function PostBox() {
  const PostBox = Vue.extend(Post)
  let instance = new PostBox()
  let PostEl = instance.$mount().$el
  document.getElementById("post_location").appendChild(PostEl) //在指定挂载点挂载
  Vue.nextTick(() => {
//    instance.dialogVisible = true
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

function UserProfileBox(data) {
  const UserProfileBox = Vue.extend(UserProfile)
  let instance = new UserProfileBox({data})
  let UserProfileEl = instance.$mount().$el
  document.body.appendChild(UserProfileEl)
  Vue.nextTick(() => {
    instance.dialogVisible = false
  })
}


export {LoginBox, PasswordBox, ProFileBox, PostBox, UserProfileBox}
