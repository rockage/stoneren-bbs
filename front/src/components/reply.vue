<template>
  <div class="edit_container" ref="editContainer">

    <span style="display: none">{{dummy}}</span>
    <quill-editor v-model="content" ref="myQuillEditor" :options="editorOption" @blur="onEditorBlur($event)"
                  @focus="onEditorFocus($event)"
                  @change="onEditorChange($event)" v-on:insertImage="insertImage($event)">
    </quill-editor>


    <input type="file" id="inputImg" @change="onInputImgChange($event)" style="display:none;">
    <img src="" id="myimg">

    <el-button v-on:click="saveHtml" style="text-align: center;margin-top: 10px" type="success" icon="el-icon-check">
      发表回复
    </el-button>

  </div>
</template>
<script>
    import store from '../store.js'

    const container = [
        ["bold", "italic"],
        ["blockquote", "code-block"],
        ["image", "rotate"]
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
                    },
                },
                content: ''
            }
        },
        methods: {
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
                }
            },
            saveHtml: function () {
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


                this.setContextData("lastReplyTime", new Date().getTime())
                param.append("tid", this.tid)
                param.append("fid", this.fid)
                param.append("uid", store.state.uid)
                param.append("uname", store.state.uname)
                param.append("threadsTitle", this.threadsTitle)
                param.append("postContens", this.content)
                this.axios.post('setNewPost', param)
                    .then((response) => {
                        if (response.data === 'login-error') {
                            vm.$message.error("登录后才可以发帖。")
                            return
                        } else {
                            myQuill.setText("")
                            this.$emit('replyFinished', response.data) //通知threadsview.vue自动跳转到最后一页
                            vm.$message.success("发帖成功。")
                        }
                    })
            },
            initRotateButton: function () { //在quill中新增一个旋转图片的按钮
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
            this.dummy = "999" //999是无意义的空渲染，在mounted阶段myQuill还未建造好，访问它会出错，只能将触发时机后移至updated阶段


        },
    }
</script>

<style>
  .edit_container {
    width: 100%;
    min-height: 300px;
    max-height: 100%;
    background-color: White

  }

  .ql-editor {
    width: 100%;
    min-height: 300px;
    max-height: 100%;
    background-color: White
  }
</style>
