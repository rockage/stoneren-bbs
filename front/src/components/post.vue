<template>
  <div class="div-center" @mousedown="move" ref="divCenter" id="post_box">
    <!-- post_box为了不重复开启 -->
    <div style="margin-top: 10px;" ref="divClose">
      <a href="javascript:void(0)" @click="close()"
        ><i class="el-icon-close" style="font-size:150%;"></i
      ></a>
    </div>

    <el-select v-model="value" placeholder="请选择" style="width:100%;font-size:smaller;color:#C0C4CC">
      <el-option
        v-for="item in options"
        :key="item.key"
        :label="item.name"
        :value="item.fid"
      >
      </el-option>
    </el-select>

    <el-input placeholder="帖子标题" v-model="threadsTitle"> </el-input>

    <div class="edit_container" ref="editContainer">
      <el-row>
        <span style="display: none">{{ dummy }}</span>
        <quill-editor
          v-model="content"
          ref="myQuillEditor"
          :options="editorOption"
          v-on:insertImage="insertImage($event)"
        >
        </quill-editor>
        <div style="text-align:left;margin-top:10px;">
          <el-button v-on:click="saveHtml" type="success" icon="el-icon-check"
            >发表主题</el-button
          >
        </div>
        <input
          type="file"
          id="inputImg"
          @change="onInputImgChange($event)"
          style="display:none;"
        />
        <img src="" id="myimg" />
      </el-row>
    </div>
  </div>
</template>

<script>
//import {ImageDrop} from 'quill-image-drop-module'
//Quill.register('modules/imageDrop', ImageDrop)
import ImageResize from "quill-image-resize-module";
Quill.register("modules/imageResize", ImageResize);
const container = [
  ["bold", "italic"],
  ["blockquote", "code-block"],
  [{ size: ["small", false, "large", "huge"] }],
  [{ color: [] }, { background: [] }],
  [{ align: [] }],
  ["link", "image", "rotate", "video"]
];
let myQuill; //用来访问quill实例
let vm;
export default {
  name: "post",
  props: {
    xclose: Function,
    root: Object
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
              image: function() {
                let c = document.getElementById("inputImg");
                c.click();
              },
              rotate: function() {
                vm.rotateImage();
              }
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
              backgroundColor: "black",
              border: "none",
              color: "white"
            },
            modules: ["Resize", "DisplaySize", "Toolbar"]
          }
        }
      },
      content: ""
    };
  },
  methods: {
    move(e) {
      let odiv = e.target; //获取目标元素
      //算出鼠标相对元素的位置
      let disX = e.clientX - odiv.offsetLeft;
      let disY = e.clientY - odiv.offsetTop;
      document.onmousemove = e => {
        //鼠标按下并移动的事件
        //用鼠标的位置减去鼠标相对元素的位置，得到元素的位置
        let left = e.clientX - disX;
        let top = e.clientY - disY;
        //绑定元素位置到positionX和positionY上面
        //this.positionX = top;
        //this.positionY = left;
        //移动当前元素
        odiv.style.left = left + "px";
        odiv.style.top = top + "px";
      };
      document.onmouseup = e => {
        document.onmousemove = null;
        document.onmouseup = null;
      };
    },
    close: function() {
      this.xclose(); //触发父组件的回调
      this.$destroy(); //销毁VUE
    },
    handleClick: function(evt) {
      if (
        evt.target &&
        evt.target.tagName &&
        evt.target.tagName.toUpperCase() === "IMG"
      ) {
        //此处判断在编辑区点击的对象是否图片
        this.rotateImg = evt.target;
      }
    },
    onInputImgChange: function() {
      let file = document.getElementById("inputImg").files[0];
      if (window.FileReader) {
        let fr = new FileReader();
        fr.readAsDataURL(file); //开始加载文件
        fr.onload = function(e) {
          //文件加载结束，this = e.target
          let originImg = new Image();
          originImg.src = e.target.result; //开始将结果（base64格式）加载到image
          originImg.onload = function() {
            //原始图片加载结束，this = 原始图片
            let w_old = this.width,
              h_old = this.height;
            const w_new = 1024;
            let h_new = (w_new * parseInt(h_old)) / parseInt(w_old);
            let canvas = document.createElement("canvas");
            canvas.width = w_new;
            canvas.height = h_new;
            if (w_old <= 1024) {
              canvas.width = w_old;
              canvas.height = h_old;
              canvas
                .getContext("2d")
                .drawImage(this, 0, 0, w_old, h_old, 0, 0, w_old, h_old); //宽度小于1024不做处理
            } else {
              canvas.width = w_new;
              canvas.height = h_new;
              canvas
                .getContext("2d")
                .drawImage(this, 0, 0, w_old, h_old, 0, 0, w_new, h_new); //超宽图强制设为1024
            }
            let src = canvas.toDataURL("image/jpeg");
            let length = myQuill.selection.savedRange.index; //用全局quill来访问quill实例
            myQuill.insertEmbed(length, "image", src);
            myQuill.setSelection(length + 1); // 调整光标到最后
          };
        };
      } else {
        alert("暂不支持FileReader");
      }
    },
    rotateImage: function() {
      if (vm.rotateImg.src === "") {
        return;
      }
      let cvs = document.createElement("canvas");
      let ctx = cvs.getContext("2d");
      let image = new Image();
      image.src = vm.rotateImg.src;
      image.onload = function() {
        let w = image.naturalWidth;
        let h = image.naturalHeight;
        let degrees = 90;
        ctx.save();
        let x;
        let y;
        const c = w; //w,h互换
        w = h;
        h = c;
        cvs.width = w;
        cvs.height = h;
        x = 0;
        y = -w;
        ctx.rotate(degrees * (Math.PI / 180));
        ctx.drawImage(image, x, y);
        ctx.restore();
        vm.rotateImg.src = cvs.toDataURL("image/jpeg");
        myQuill.imageResize.hide();
      };
    },
    saveHtml: function() {
      let param = new URLSearchParams(); //axios如不采用URLSearchParams后端无法收到post请求
      const vm = this;

      if (this.threadsTitle.length < 10) {
        vm.$message.error("帖子标题至少3个字符!");
        return;
      }

      if (this.content.length < 10) {
        vm.$message.error("发帖内容至少3个字符!");
        return;
      }
      let endTime = new Date().getTime();
      let beginTime = this.getContextData("lastReplyTime");
      if (endTime - beginTime < 8000) {
        vm.$message.error("发帖间隔时间不得低于8秒!");
        return;
      }
      this.setContextData("lastReplyTime", new Date().getTime());

      console.log(this.root.$store.getters.fid);
      console.log(this.root.$store.getters.uid);
      console.log(this.threadsTitle);
      console.log(this.content);

      return;
      param.append("tid", 0);
      param.append("fid", this.root.$store.getters.fid);
      param.append("uid", this.root.$store.getters.uid);
      param.append("uname", store.state.uname);
      param.append("threadsTitle", this.threadsTitle);
      param.append("postContens", this.content);
      this.axios
        .post("http://localhost:8081/setNewPost", param)
        .then(response => {
          if (response.data === "login-error") {
            vm.$login();
            return;
          } else {
            myQuill.setText("");
            vm.$message.success("主题发表成功。");
          }
        });
    },
    initRotateButton: function() {
      //在quill中新增一个旋转图片的按钮
      const sourceEditorButton = document.querySelector(".ql-rotate"); //每一个quill功能按钮都会自动加上一个ql-用以区别，如ql-image/ql-link等
      sourceEditorButton.style.cssText = "border:0px";
      sourceEditorButton.innerHTML = "<img src='/static/rotate.png'>";
      sourceEditorButton.title = "旋转图片";
    }
  },
  updated() {
    this.initRotateButton();
  },
  mounted() {
    vm = this;
    myQuill = this.$refs.myQuillEditor.quill; //全局quill实例
    myQuill.root.addEventListener("click", this.handleClick, false); //为全局quill创建一个根监听
    /* 定义弹窗宽度(因是动态创建的窗口，没有父元素，所以无法直接使用%号)：
    let div_center = this.$refs.divCenter;
    let w = (document.documentElement.clientWidth * 50) / 100;
    let h = (document.documentElement.clientheight * 50) / 100;
    div_center.style.width = `${w}px`;*/
    this.dummy = "999"; //在mounted阶段myQuill还未建造好，只能将触发时机后移至updated阶段

    for (let i = 0; i < this.root.$store.getters.fsname.length; i++) {
      this.options.push({
        key: i,
        name: "发布到：" + this.root.$store.getters.fsname[i].name,
        fid: this.root.$store.getters.fsname[i].fid
      });
      this.value = this.root.$store.getters.fid;
    }
  },
  destroyed() {
    document.getElementById("post_location").removeChild(this.$el); //销毁DOM
  }
};
</script>

<style>
.edit_container {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  background-color: White;
}

.ql-editor {
  background-color: White;
  max-width: 100%;
  min-height: 300px;
}

.div-center {
  position: static; /*定位*/
  border: 1px gray;
  background: #e4e7ed; /*设置一下背景*/
  z-index: 99;
  border-radius: 5px;
  margin-top: 10px;
  margin-left: 10px;
  margin-bottom: 8l0px;
  max-height: 100%;
  text-align: right;
  max-width: 100%;
  min-height: 300px;
}
</style>
