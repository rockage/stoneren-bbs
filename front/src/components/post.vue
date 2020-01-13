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
    <input type="file" id="inputimg" style="display: none">
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
                    let c = document.getElementById("inputimg");
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
      saveHtml: function (event) {

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
          imgFile.readAsDataURL(document.getElementById('inputimg').files[0]);
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
