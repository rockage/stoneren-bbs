<template>
  <div class="edit_container" ref="editContainer">

    <span style="display: none">{{dummy}}</span>
    <quill-editor v-model="content" ref="myQuillEditor" :options="editorOption" @blur="onEditorBlur($event)" @focus="onEditorFocus($event)"
      @change="onEditorChange($event)" v-on:insertImage="insertImage($event)">
    </quill-editor>

    <input type="file" id="inputImg" @change="onInputImgChange($event)" style="display:none;">
    <img src="" id="myimg">

    <el-button v-on:click="saveHtml" style="text-align: center;margin-top: 10px" type="success" icon="el-icon-check">发表回复</el-button>

  </div>
</template>
<script>
  import ImageResize from 'quill-image-resize-module'
  import store from '../store.js'

  Quill.register('modules/imageResize', ImageResize)
  const container = [
    ['bold', 'italic'],
    ['blockquote', 'code-block'],
    [{
      'size': ['small', false, 'large', 'huge']
    }],
    [{
      'color': []
    }, {
      'background': []
    }],
    [{
      'align': []
    }],
    ['link', 'image', 'rotate', 'video'],
  ]
  let myQuill //增加一个全局quill，代表当前quill实例
  let vm
  export default {
    name: "reply",
    store,
    props: ['tid', 'fid'],
    data() {
      return {
        value: '',
        options: [],
        dummy: '',
        positionX: 0,
        positionY: 0,
        threadsTitle: '',
        rotateImg: '',
        editorOption: {
          placeholder: '输入回帖内容...',
          modules: {
            toolbar: {
              container: container,
              handlers: {
                // 拦截image的click事件
                'image': function() {
                  let c = document.getElementById("inputImg")
                  c.click()
                },
                'rotate': function() {
                  vm.rotateImage()
                },
              }
            },
            history: {
              delay: 1000,
              maxStack: 50,
              userOnly: false
            },
            //imageDrop: true,
            imageResize: {
              displayStyles: {
                backgroundColor: 'black',
                border: 'none',
                color: 'white'
              },
              modules: ['Resize', 'DisplaySize', 'Toolbar']
            },
          },
        },
        content: ''
      }
    },
    methods: {
      handleClick: function(evt) {
        if (evt.target && evt.target.tagName && evt.target.tagName.toUpperCase() === 'IMG') {
          //此处判断在编辑区点击的对象是否图片
          this.rotateImg = evt.target
        }
      },
      onEditorBlur: function() { // 失去焦点事件
      },
      onEditorFocus: function() { // 获得焦点事件
      },
      onEditorChange: function() { // 内容改变事件
      },
      onInputImgChange: function() {
        let file = document.getElementById("inputImg").files[0]
        if (window.FileReader) {
          let fr = new FileReader()
          fr.readAsDataURL(file) //开始加载文件
          fr.onload = function(e) { //文件加载结束，this = e.target
            let originImg = new Image()
            originImg.src = e.target.result //开始将结果（base64格式）加载到image
            originImg.onload = function() { //原始图片加载结束，this = 原始图片
              let w_old = this.width,
                h_old = this.height
              const w_new = 1024
              let h_new = (w_new * parseInt(h_old)) / parseInt(w_old)
              let canvas = document.createElement("canvas")
              canvas.width = w_new
              canvas.height = h_new
              if (w_old <= 1024) {
                canvas.width = w_old
                canvas.height = h_old
                canvas.getContext("2d").drawImage(this, 0, 0, w_old, h_old, 0, 0, w_old, h_old) //宽度小于1024不做处理
              } else {
                canvas.width = w_new
                canvas.height = h_new
                canvas.getContext("2d").drawImage(this, 0, 0, w_old, h_old, 0, 0, w_new, h_new) //超宽图强制设为1024
              }
              let src = canvas.toDataURL("image/jpeg")
              let length = myQuill.selection.savedRange.index //用全局quill来访问quill实例
              myQuill.insertEmbed(length, 'image', src)
              myQuill.setSelection(length + 1) // 调整光标到最后
            }
          }
        } else {
          alert('暂不支持FileReader')
        }
      },
      rotateImage: function() {
        if (vm.rotateImg.src === '') {
          return
        }
        let cvs = document.createElement("canvas");
        let ctx = cvs.getContext("2d")
        let image = new Image()
        image.src = vm.rotateImg.src
        image.onload = function() {
          let w = image.naturalWidth
          let h = image.naturalHeight;
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
          myQuill.imageResize.hide()
        }
      },
      saveHtml: function() {
        let param = new URLSearchParams() //axios如不采用URLSearchParams后端无法收到post请求
        const vm = this
        if (this.content.length < 10) {
          vm.$message.error("发帖内容至少3个字符!")
          return
        }
        let endTime = new Date().getTime()
        let beginTime = this.getContextData("lastReplyTime")
        if ((endTime - beginTime) < 8000) {
          vm.$message.error("发帖间隔时间不得低于8秒!")
          return
        }

        console.log(this.content.length)


        this.setContextData("lastReplyTime", new Date().getTime())
        param.append("tid", this.tid)
        param.append("fid", this.fid)
        param.append("uid", store.state.uid)
        param.append("uname", store.state.uname)
        param.append("threadsTitle", this.threadsTitle)
        param.append("postContens", this.content)
        this.axios.post('http://localhost:8081/setNewPost', param)
          .then((response) => {
            if (response.data === 'login-error') {
              vm.$login()
              return
            } else {
              myQuill.setText("")
              this.$emit('replyFinished', response.data) //此处的resonse.data是计算好的分页值，回传给renderMain自动跳转到最后一页
              vm.$message.success("发帖成功。")
            }
          })
      },
      initRotateButton: function() { //在quill中新增一个旋转图片的按钮
        const sourceEditorButton = document.querySelector('.ql-rotate') //每一个quill功能按钮都会自动加上一个ql-用以区别，如ql-image/ql-link等
        sourceEditorButton.style.cssText = "border:0px"
        sourceEditorButton.innerHTML = "<img src='/static/rotate.png'>"
        sourceEditorButton.title = "旋转图片"
      },
    },
    updated() {
      this.initRotateButton()
    },
    mounted() {
      vm = this
      myQuill = this.$refs.myQuillEditor.quill //全局quill实例
      myQuill.root.addEventListener('click', this.handleClick, false) //为全局quill创建一个根监听
      //元素定位：
      let edit_container = this.$refs.editContainer
      myQuill.container.style.width = `100%`
      myQuill.container.style.height = `300px`
      edit_container.style.width = `100%`
      edit_container.style.height = `300px`
      this.dummy = "999" //999是无意义的空渲染，在mounted阶段myQuill还未建造好，访问它会出错，只能将触发时机后移至updated阶段
      this.setContextData("lastReplyTime", new Date().getTime())







    },
  }
</script>

<style>
  .edit_container {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    background-color: White
  }

  .ql-editor {
    background-color: White
  }
</style>
