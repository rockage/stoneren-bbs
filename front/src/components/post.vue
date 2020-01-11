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
  </div>

</template>

<script>
  import VueQuillEditor, {Quill} from 'vue-quill-editor'
  import {ImageDrop} from 'quill-image-drop-module'
  import ImageResize from 'quill-image-resize-module'

  Quill.register('modules/imageDrop', ImageDrop)
  Quill.register('modules/imageResize', ImageResize)


  export default {
    name: "post",
    data() {
      return {
        editorOption: {
          placeholder: '输入内容...',
          modules: {
            toolbar: [
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
            ],
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
            }
          }
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

        /* 预览
        let div = document.createElement("div")
        div.innerHTML = this.content
        document.body.appendChild(div)
        */

        this.axios.post('http://localhost:8081/setNewPost', param)
          .then((response) => {
            this.totalPage = JSON.parse(response.data)
          })
      },
    },
  }

</script>

<style>
  .ql-editor {
    height: 600px;
    width: 100%;
  }
</style>
