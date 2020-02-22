import Vue from 'vue'
import ProfileBox from "./components/userprofile.vue"

//以下所有通过动态方法new新建的Vue实例都不能访问$root, 需创建时传递

function UserProfileBox(data) {

  if (!this.$root.$store.getters.loginState) {
    this.$message.warning("登录后才允许查看用户资料。")
    return
  }

  if (!document.getElementById("ProfileBox")) {
    const UserProfile = Vue.extend(ProfileBox)
    let instance = new UserProfile({
      propsData: {
        root: this.$root,
        uname: data.uname,
      }
    })
    document.body.appendChild(instance.$mount().$el)
  }
}


export {UserProfileBox}
