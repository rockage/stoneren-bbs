<template>
  <div class="div-center" @mousedown="move" ref="divCenter" id="post_box">

    <div class="edit_container" ref="editContainer">
      <el-row style="margin-top: 20px;">
        <div style="
          border: 1px solid #cccccc;">
          <table border="0" style="width: 100%">

            <td style="width: 85%">
              <el-input
                placeholder="帖子标题"
                v-model="threadsTitle">
              </el-input>
            </td>
            <td style="width: 3%"></td>
            <td style="width: 6%">
              <el-button v-on:click="saveHtml" style="text-align: right;" type="success" icon="el-icon-check"
                         circle size="mini"></el-button>
            </td>
            <td style="width: 6%">
              <el-button v-on:click="close" style="text-align: right;" type="warning" icon="el-icon-close"
                         circle size="mini"></el-button>
            </td>
          </table>
        </div>
      </el-row>
      <el-row>
        <span style="display: none">{{dummy}}</span>
        <quill-editor
          v-model="content"
          ref="myQuillEditor"
          :options="editorOption"
          @blur="onEditorBlur($event)" @focus="onEditorFocus($event)"
          @change="onEditorChange($event)"
          v-on:insertImage="insertImage($event)">
        </quill-editor>
        <input type="file" id="inputImg" @change="onInputImgChange($event)" style="display:none;">
        <img src="" id="myimg">
      </el-row>
    </div>
  </div>
</template>

<script>
  //import {ImageDrop} from 'quill-image-drop-module'
  //Quill.register('modules/imageDrop', ImageDrop)

  import ImageResize from 'quill-image-resize-module'

  Quill.register('modules/imageResize', ImageResize)
  const container = [
    ['bold', 'italic'],
    ['blockquote', 'code-block'],
    [{'size': ['small', false, 'large', 'huge']}],
    [{'color': []}, {'background': []}],
    [{'align': []}],
    ['link', 'image', 'rotate', 'video'],
  ]
  let myQuill //增加一个全局quill，代表当前quill实例
  let vm
  export default {
    name: "post",
    props: ['isShow',],
    data() {
      return {
        dummy: '',
        positionX: 0,
        positionY: 0,
        currentIsShow: this.isShow,
        tid: '',
        rootThis: '',
        threadsTitle: '',
        rotateImg: '',
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
    watch: {
      isShow(val) {
        this.currentIsShow = val
      }
    },
    methods: {
      move(e) {
        let odiv = e.target;        //获取目标元素
        //算出鼠标相对元素的位置
        let disX = e.clientX - odiv.offsetLeft;
        let disY = e.clientY - odiv.offsetTop;
        document.onmousemove = (e) => {       //鼠标按下并移动的事件
          //用鼠标的位置减去鼠标相对元素的位置，得到元素的位置
          let left = e.clientX - disX;
          let top = e.clientY - disY;
          //绑定元素位置到positionX和positionY上面
          //this.positionX = top;
          //this.positionY = left;
          //移动当前元素
          odiv.style.left = left + 'px';
          odiv.style.top = top + 'px';
        };
        document.onmouseup = (e) => {
          document.onmousemove = null;
          document.onmouseup = null;
        };
      },
      close: function () {

        this.$destroy() //销毁VUE

      },
      handleClick: function (evt) {
        if (evt.target && evt.target.tagName && evt.target.tagName.toUpperCase() === 'IMG') {
          //此处判断在编辑区点击的对象是否图片
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
        if (vm.rotateImg.src === '') {
          return
        }
        let cvs = document.createElement("canvas");
        let ctx = cvs.getContext("2d")
        let image = new Image()
        image.src = vm.rotateImg.src
        image.onload = function () {
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
      saveHtml: function () {
        let param = new URLSearchParams() //axios如不采用URLSearchParams后端无法收到post请求
        const vm = this
        param.append("tid", this.tid)
        param.append("uid", this.rootThis.$store.state.uid)
        param.append("threadsTitle", this.threadsTitle)
        param.append("postContens", this.content)
        this.axios.post('http://localhost:8081/setNewPost', param)
          .then((response) => {
            if (response.data === 'login-error') {
              vm.$login()
              return
            } else {
              vm.$message.success("发帖成功。")
            }
          })
      },
      initRotateButton: function () {      //在quill中新增一个旋转图片的按钮
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
      let div_center = this.$refs.divCenter
      let edit_container = this.$refs.editContainer
      let w = document.documentElement.clientWidth * 50 / 100
      let h = document.documentElement.clientHeight * 50 / 100
      myQuill.container.style.width = `${w}px`
      myQuill.container.style.height = `${h}px`
      div_center.style.width = `${w}px`
      div_center.style.height = `${h}px`
      edit_container.style.width = `${w}px`
      edit_container.style.height = `${h}px`

      this.rootThis = this.GLOBAL.globalThis
      this.dummy = "999"
      //999是无意义的空渲染，在mounted阶段myQuill还未建造好，访问它会出错，
      //只好将触发时机后移至updated阶段
    },
    destroyed() {
      document.getElementById("post_location").removeChild(this.$el) //销毁DOM
    }
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

  .div-center {
    position: fixed; /*定位*/
    border: 1px gray;
    background: #E4E7ED; /*设置一下背景*/
    z-index: 99;
    border-radius: 5px;
    margin-top: 10px;
    margin-left: 10px;
  }

</style>
