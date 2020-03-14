<template>
  <div class="div-center" @mousedown="move" ref="divCenter" id="post_box">
    <!-- post_box为了不重复开启 -->

    <div style="display: flex;justify-content:flex-end;margin-top:5px;" ref="divClose">
      <a href="javascript:void(0)" @click="close()">
        <i class="el-icon-close" style="font-size:150%;"></i>
      </a>
    </div>

    <div v-show="!reply">
      <div style="align-items: center;display: flex;justify-content: space-between;">
        <div class="div-sendto">发表到：</div>

        <el-select style="width: 100%;" v-model="value" placeholder="请选择">
          <el-option
            v-for="item in options"
            :key="item.key"
            :label="item.name"
            :value="item.fid"
          ></el-option>
        </el-select>
      </div>
      <el-input placeholder="请输入标题" v-model="threadsTitle">
        <template slot="prepend">主 题：</template>
      </el-input>
    </div>
    <div id="quill-container">
      <span style="display: none">{{ dummy }}</span>
      <quillEditor
        v-model="content"
        ref="myQuillEditor"
        :options="editorOption"
        v-on:insertImage="insertImage($event)"
      ></quillEditor>
      <input type="file" id="inputImg" @change="onInputImgChange($event)" style="display:none;"/>
      <img src id="myimg"/>
    </div>
  </div>
</template>

<script>


  const container = [
    [{'size': [false, 'large', 'huge']}],
    ["code-block", 'link', "image", "rotate", "post"],
  ]
  let myQuill
  let vm

  import {quillEditor} from 'vue-quill-editor'

  export default {
    name: "post",
    props: {
      xclose: Function,
      postFinished: Function,
      root: Object,
      reply: Boolean,
      pid: Number
    },
    components: {
      quillEditor,
    },
    data() {
      return {
        value: "",
        options: [],
        dummy: "",
        positionX: 0,
        positionY: 0,
        threadsTitle: "",
        rotateImg: "",
        editorOption: {
          placeholder: "",
          modules: {
            toolbar: {
              container: container,
              handlers: {
                // 拦截image的click事件
                image: function () {
                  this.title = "图片"
                  let c = document.getElementById("inputImg")
                  c.click()
                },
                rotate: function () {
                  vm.rotateImage()
                },
                'post': function () {
                  vm.saveHtml()
                },
              }
            },
            history: {
              delay: 1000,
              maxStack: 50,
              userOnly: false
            },
          }
        },
        content: ""
      }
    },
    methods: {
      move(e) {
        let odiv = e.target //获取目标元素
        //算出鼠标相对元素的位置
        let disX = e.clientX - odiv.offsetLeft
        let disY = e.clientY - odiv.offsetTop
        document.onmousemove = e => {
          //鼠标按下并移动的事件
          //用鼠标的位置减去鼠标相对元素的位置，得到元素的位置
          let left = e.clientX - disX
          let top = e.clientY - disY
          //绑定元素位置到positionX和positionY上面
          //this.positionX = top;
          //this.positionY = left;
          //移动当前元素
          odiv.style.left = left + "px"
          odiv.style.top = top + "px"
        }
        document.onmouseup = e => {
          document.onmousemove = null
          document.onmouseup = null
        }
      },
      close: function () {
        if (!this.reply) {
          this.xclose() //触发父组件的回调
        }

        this.$destroy() //销毁VUE
      },
      handleClick: function (evt) {
        if (
          evt.target &&
          evt.target.tagName &&
          evt.target.tagName.toUpperCase() === "IMG"
        ) {
          //此处判断在编辑区点击的对象是否图片
          this.rotateImg = evt.target
        }
      },
      onInputImgChange: function () {
        let file = document.getElementById("inputImg").files[0]
        let vm = this
        if (window.FileReader) {
          let fr = new FileReader()
          fr.readAsDataURL(file) //开始加载文件
          fr.onload = function (e) {
            //文件加载结束，this = e.target
            vm.transImage(e.target.result, 1024, 0, function (data) { //将图片宽度设为1024，高度自适应
              let length = myQuill.selection.savedRange.index //用全局quill来访问quill实例
              myQuill.insertEmbed(length, "image", data)
              myQuill.setSelection(length + 1) // 调整光标到最后
            })
          }
        } else {
          alert("暂不支持FileReader")
        }
      },
      rotateImage: function () {
        if (vm.rotateImg.src === "") {
          return
        }
        let cvs = document.createElement("canvas")
        let ctx = cvs.getContext("2d")
        let image = new Image()
        image.src = vm.rotateImg.src
        image.onload = function () {
          let w = image.naturalWidth
          let h = image.naturalHeight
          let degrees = 90
          ctx.save()
          let x
          let y
          const c = w //w,h互换
          w = h
          h = c
          cvs.width = w
          cvs.height = h
          x = 0
          y = -w
          ctx.rotate(degrees * (Math.PI / 180))
          ctx.drawImage(image, x, y)
          ctx.restore()
          vm.rotateImg.src = cvs.toDataURL("image/jpeg")
        }
      },
      saveHtml: function () {
        let param = new URLSearchParams() //axios如不采用URLSearchParams后端无法收到post请求
        const vm = this

        if (this.threadsTitle.length < 3) {
          vm.$message.error("帖子标题至少3个字符!")
          return
        }

        if (this.threadsTitle.length > 30) {
          vm.$message.error("帖子标题太长了!")
          return
        }

        if (this.content.length < 10) {
          vm.$message.error("发帖内容至少3个字符!")
          return
        }
        let endTime = new Date().getTime()
        let beginTime = this.getContextData("lastReplyTime")
        if (endTime - beginTime < 8000) {
          vm.$message.error("发帖间隔时间不得低于8秒!")
          return
        }
        this.setContextData("lastReplyTime", new Date().getTime())

        param.append("tid", 0)
        param.append("fid", vm.root.$store.getters.fid)
        param.append("uid", vm.root.$store.getters.uid)
        param.append("uname", vm.root.$store.getters.uname)
        param.append("threadsTitle", vm.threadsTitle)
        param.append("postContens", vm.content)
        this.axios
          .post("setNewPost", param)
          .then(response => {
            if (response.data === "login-error") {
              vm.$message.error("登录后才可以发帖。")
              return
            } else {
              myQuill.setText("")
              vm.$message.success("主题发表成功。")

              this.postFinished() //触发父组件的回调,通知bbsButton.vue发帖完成
              this.$destroy() //销毁VUE
            }
          })
      },
      initRotateButton: function () {
        //在quill中新增一个旋转图片的按钮
        const sourceEditorButton = document.querySelector(".ql-rotate") //每一个quill功能按钮都会自动加上一个ql-用以区别，如ql-image/ql-link等
        sourceEditorButton.style.cssText = "border:0px"
        sourceEditorButton.innerHTML = "<img src='/static/rotate.png'>"
        sourceEditorButton.title = "旋转图片"

        const postButton = document.querySelector('.ql-post') //每一个quill功能按钮都会自动加上一个ql-用以区别，如ql-image/ql-link等
        postButton.style.cssText = "border:0px"
        postButton.innerHTML = "<img src='/static/post.png'>"
        postButton.title = "发帖"
      },
      loadPostContens: function () {
        this.axios
          .get("getPost", {
            params: {
              pid: this.pid
            }
          })
          .then(response => {
            let ret = JSON.parse(response.data)
            this.content = ret[0].message
          })
      }
    },
    updated() {
      this.initRotateButton()
    },
    mounted() {
      vm = this
      myQuill = this.$refs.myQuillEditor.quill //全局quill实例
      myQuill.root.addEventListener("click", this.handleClick, false) //为全局quill创建一个根监听
      this.dummy = "999" //在mounted阶段myQuill还未建造好，只能将触发时机后移至updated阶段
      this.value = String(this.root.$store.getters.fid)
      this.options = this.root.$store.getters.fsname

      if (this.reply) { //帖子编辑
        this.loadPostContens()
      }

    },
    destroyed() {
      document.getElementById("post_location").removeChild(this.$el) //销毁DOM
    }
  };
</script>

<style>
  .edit_container {
    font-family: "Avenir", Helvetica, Arial, sans-serif;
    background-color: White;
  }


  .div-sendto {
    font-size: small;
    min-width: 88px;
    height: 38px;
    background-color: #f2f6fc;
    color: #909399;
    border-radius: 5px;
    align-items: center;
    display: flex;
    justify-content: center;

  }

  .ql-editor {
    background-color: White;
    width: 100%;
    height: 500px;
    overflow-y: scroll;
  }

  .div-center {
    border: 1px gray;
    background: #e4e7ed; /*设置一下背景*/
    z-index: 99;
    border-radius: 5px;
    position: fixed;
    top: 50%;
    left: 50%;
    width: 90%;
    height: 680px;
    -webkit-transform: translate(-50%, -50%);
    transform: translate(-50%, -50%);
  }

  .ql-clipboard {
    position: fixed;
    display: none;
    left: 50%;
    top: 50%;
  }
</style>
