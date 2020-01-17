<template>
  <div class="edit_container">
    <el-input
      placeholder="帖子标题"
      v-model="threadsTitle">
    </el-input>
    <quill-editor
      v-model="content"
      ref="myQuillEditor"
      :options="editorOption"
      @blur="onEditorBlur($event)" @focus="onEditorFocus($event)"
      @change="onEditorChange($event)"
      v-on:insertImage="insertImage($event)"
    >
    </quill-editor>
    <el-button v-on:click="saveHtml">保存</el-button>

    <input type="file" id="inputImg" @change="onInputImgChange($event)" style="display:none;">
    <img src="" id="myimg">

  </div>
</template>


<script>


  //import {ImageDrop} from 'quill-image-drop-module'
  import ImageResize from 'quill-image-resize-module'

  //Quill.register('modules/imageDrop', ImageDrop)
  Quill.register('modules/imageResize', ImageResize)

  const container = [
    ['bold', 'italic', 'underline', 'strike'],
    ['blockquote', 'code-block'],
    [{'indent': '-1'}, {'indent': '+1'}],
    [{'size': ['small', false, 'large', 'huge']}],
    [{'font': []}],
    [{'color': []}, {'background': []}],
    [{'align': []}],
    ['clean'],
    ['link', 'image', 'rotate', 'video'],
  ]
  let myQuill //增加一个全局quill，代表当前quill实例
  let Vue

  export default {
    name: "post",
    data() {
      return {
        threadsTitle: '',
        rotateImg: null,
        editorOption: {
          placeholder: '输入内容...',
          modules: {
            toolbar: {
              container: container,
              handlers: {
                // 拦截image的click事件
                'image': function () {
                  let c = document.getElementById("inputImg")
                  c.click()
                },
                'rotate': function () {
                  Vue.rotateImage()
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
      handleClick: function (evt) {
        if (evt.target && evt.target.tagName && evt.target.tagName.toUpperCase() === 'IMG') {
          this.rotateImg = evt.target
        }
      },
      onEditorBlur: function () { // 失去焦点事件
      },
      onEditorFocus: function () { // 获得焦点事件
      },
      onEditorChange: function () { // 内容改变事件
      },
      onInputImgChange: function () {
        let file = document.getElementById("inputImg").files[0]
        if (window.FileReader) {
          let fr = new FileReader()
          fr.readAsDataURL(file) //开始加载文件
          fr.onload = function (e) { //文件加载结束，this = e.target
            let originImg = new Image()
            originImg.src = e.target.result //开始将结果（base64格式）加载到image
            originImg.onload = function () { //原始图片加载结束，this = 原始图片
              let w_old = this.width, h_old = this.height
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
                canvas.getContext("2d").drawImage(this, 0, 0, w_old, h_old, 0, 0, w_new, h_new)  //超宽图强制设为1024
              }
              let src = canvas.toDataURL("image/jpeg")
              let length = myQuill.selection.savedRange.index //用全局quill来访问quill实例
              myQuill.insertEmbed(length, 'image', src)
              myQuill.setSelection(length + 1)               // 调整光标到最后
            }
          }
        } else {
          alert('暂不支持FileReader')
        }
      },

      rotateImage: function () {
        let cvs = document.createElement("canvas");
        let ctx = cvs.getContext("2d")
        let image = new Image()
        image.src = Vue.rotateImg.src
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
          Vue.rotateImg.src = cvs.toDataURL("image/jpeg")
          myQuill.imageResize.hide()
        }
      },
      saveHtml: function () {
        let param = new URLSearchParams() //axios如不采用URLSearchParams后端无法收到post请求
        param.append("fid", this.fid)
        param.append("tid", this.tid)
        param.append("uid", this.uid)
        param.append("threadsTitle", this.threadsTitle)
        param.append("postContens", this.content)
        const regex = /<img src="(.+?)">/g;
        let m = regex.exec(this.content)


        this.axios.post('http://localhost:8081/setNewPost', param)
          .then((response) => {
            this.totalPage = JSON.parse(response.data)
          })
      },
      initRotateButton: function () {      //在quill中新增一个旋转图片的按钮
        const sourceEditorButton = document.querySelector('.ql-rotate');
        sourceEditorButton.style.cssText = "border:0px"
        sourceEditorButton.innerHTML = "<img src='/static/rotate.png'>"
        sourceEditorButton.title = "旋转图片"
      },
    }
    ,
    mounted() {
      this.initRotateButton() //初始化自定义按钮
      Vue = this
      myQuill = this.$refs.myQuillEditor.quill //全局quill实例
      myQuill.root.addEventListener('click', this.handleClick, false) //为全局quill创建一个根监听
      let quillWidth = document.documentElement.clientWidth * 80 / 100
      let quillHeight = document.documentElement.clientHeight * 70 / 100
      myQuill.container.style.width = `${quillWidth}px`
      myQuill.container.style.height = `${quillHeight}px`
    }
    ,
  }

</script>

<style>
  .ql-editor {


  }
</style>
