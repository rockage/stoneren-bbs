<template>
  <div class="edit_container">
    <el-input
      placeholder="标题"
      v-model="threadsTitle">
    </el-input>
    <quill-editor
      v-model="content"
      ref="myQuillEditor"
      :options="editorOption"
      @blur="onEditorBlur($event)" @focus="onEditorFocus($event)"
      @change="onEditorChange($event)">
    </quill-editor>
    <el-button v-on:click="saveHtml">保存</el-button>

    <input type="file" id="inputImg" @change="onInputImgChange($event)">

    <select id="myselect">
      <option value="1">webp格式</option>
      <option value="2">jpeg格式</option>
      <option value="3">png格式</option>
    </select>
    <button id="start" @click="xxx">开始转换</button>
    <p>预览：</p>
    <img id="imgShow" src="" alt="">
    <img id="img2" src="" alt="">

  </div>
</template>


<script>


  import {ImageDrop} from 'quill-image-drop-module'
  import ImageResize from 'quill-image-resize-module'

  Quill.register('modules/imageDrop', ImageDrop)
  Quill.register('modules/imageResize', ImageResize)

  const container = [
    ['bold', 'italic', 'underline', 'strike'],
    ['blockquote', 'code-block'],
    [{'header': 1}, {'header': 2}],
    [{'list': 'ordered'}, {'list': 'bullet'}],
    [{'script': 'sub'}, {'script': 'super'}],
    [{'indent': '-1'}, {'indent': '+1'}],
    [{'direction': 'rtl'}],
    [{'size': ['small', false, 'large', 'huge']}],
    [{'header': [1, 2, 3, 4, 5, 6, false]}],
    [{'font': []}],
    [{'color': []}, {'background': []}],
    [{'align': []}],
    ['clean'],
    ['link', 'image', 'video']
  ]

  let fileBtn = document.getElementById('inputImg')

  export default {
    name: "post",
    data() {
      return {
        threadsTitle: '',
        editorOption: {
          placeholder: '输入内容...',
          modules: {
            toolbar: {
              container: container,
              handlers: {
                // 拦截image的click事件
                'image': function () {
                  this.quill.insertText(0, 'Hello Quill', 'bold', true);
                  let c = document.getElementById("inputImg");
                  c.click()
                }
              }

            },
            history: {
              delay: 1000,
              maxStack: 50,
              userOnly: false
            },
            imageDrop: true,
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
      onEditorReady: function (editor) { // 准备编辑器
      },
      onEditorBlur: function () { // 失去焦点事件
      },
      onEditorFocus: function () { // 获得焦点事件
      },
      onEditorChange: function () { // 内容改变事件
      },
      iii: function () { // 内容改变事件
        console.log("i am in iii")

      },
      onInputImgChange: function (event) {
        console.log("in")
        let file = document.getElementById("inputImg").files[0]

        if (window.FileReader) {

          let fr = new FileReader()
          fr.readAsDataURL(file) //开始加载文件
          fr.onload = function (e) { //文件加载结束，this = e.target
            //返回FileReader对象：事件，状态，属性，结果等

            let originImg = new Image()
            originImg.src = e.target.result //开始将结果（base64格式）加载到image

            originImg.onload = function () { //原始图片加载结束，this = 原始图片
              let w_old = this.width, h_old = this.height
              const w_new = 1024
              let h_new = (w_new * parseInt(h_old)) / parseInt(w_old)



              let canvas = document.createElement("canvas")
              canvas.width = w_new
              canvas.height = h_new

              if (w_old<=1024) {
                canvas.width = w_old
                canvas.height = h_old
                canvas.getContext("2d").drawImage(this, 0, 0, w_old, h_old, 0, 0, w_old, h_old) //宽度小于1024不做处理
              }else{
                canvas.width = w_new
                canvas.height = h_new
                canvas.getContext("2d").drawImage(this, 0, 0, w_old, h_old, 0, 0, w_new, h_new)  //超宽图强制设为1024
              }

              let imgshow = document.getElementById("imgShow")
              let src = canvas.toDataURL("image/jpeg")
              imgshow.setAttribute('src', src)



              // 获取光标所在位置
                // let length = this.quill.getSelection()
              // 插入图片，res为服务器返回的图片链接地址
              //this. quill.insertEmbed(0, 'image', imgshow)
              // 调整光标到最后
              //this.quill.setSelection(length + 1)

              this.$emit('myQuillEditor')






            }


          }
        } else {
          alert('暂不支持FileReader')
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

        let img2 = document.getElementById("img2")


        img2.setAttribute('src', m[1]);

        this.yyy(img2)


        return
        this.axios.post('http://localhost:8081/setNewPost', param)
          .then((response) => {
            this.totalPage = JSON.parse(response.data)
          })
      },
      imgToCanvas: function (image) {
        let canvas = document.createElement("canvas");
        canvas.width = image.width;
        canvas.height = image.height;
        canvas.getContext("2d").drawImage(image, 0, 0);
        return canvas;
      },
      //canvas转换为image
      canvasToImg: function (canvas) {
        let array = ["image/webp", "image/jpeg", "image/png"],
          type = document.getElementById('myselect').value - 1;
        let src = canvas.toDataURL(array[type]);
        return src;
      },
      //获取图片信息
      getImg: function getImg(fn) {
        let imgFile = new FileReader();
        try {
          imgFile.onload = function (e) {
            let image = new Image();
            image.src = e.target.result; //base64数据
            image.onload = function () {
              fn(image);
            }
          }
          imgFile.readAsDataURL(document.getElementById('inputImg').files[0]);
        } catch (e) {
          console.log("请上传图片！" + e);
        }
      },
      xxx: function () {
        this.getImg(function (image) {
          console.log(image)
          let canvas = document.createElement("canvas");
          canvas.width = image.width;
          canvas.height = image.height;
          canvas.getContext("2d").drawImage(image, 0, 0);

          let imgshow = document.getElementById("imgShow")

          let array = ["image/webp", "image/jpeg", "image/png"],
            type = document.getElementById('myselect').value - 1;
          let src = canvas.toDataURL(array[type]);


          imgshow.setAttribute('src', src);
        });
      },
      yyy: function (image) {
        console.log(image)
        let canvas = document.createElement("canvas");
        //canvas.width = image.width;
        //canvas.height = image.height;
        let scale = image.width / image.height;


        canvas.getContext("2d").drawImage(image, 0, 0);

        let imgshow = document.getElementById("imgShow")

        let array = ["image/webp", "image/jpeg", "image/png"],
          type = document.getElementById('myselect').value - 1;
        let src = canvas.toDataURL(array[type]);


        imgshow.setAttribute('src', src);
      },
      zzz: function () {
        console.log("hi")
      }
    },
    mounted() {


    },
  }

</script>

<style>
  .ql-editor {
    height: 600px;
    width: 100%;
  }
</style>
