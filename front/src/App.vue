<template>
  <div id="app">

    <el-container>


      <el-header>中国石凳联盟</el-header>
      <el-container>
        <el-aside width="330px">


          <el-row class="tac">
            <el-col :span="12">
              <el-menu
                default-active="2"
                class="el-menu-vertical-demo"
                style="width: 200px"
              >

                <el-menu-item index="1">
                  <i class="el-icon-menu"></i>
                  <span slot="title">新帖</span>
                </el-menu-item>

              </el-menu>
            </el-col>
          </el-row>

          <div>
            <h3>标签管理</h3>
            <!-- 标签容器已经被初始化而且内容多余0条，则显示内容 -->
            <div v-if="tags && tags.length > 0">
              <tag v-for="tag in tags" :text="tag" :key="tag" @delete="onTagDelete"/>
            </div>
            <!-- 否则，标签只是被初始化了，但是没有数据 -->
            <div v-else-if="tags">没有标签了</div>
            <!-- 标签容器还没被初始化 -->
            <div v-else>标签加载中...</div>
            <button @click="addTag">添加标签</button>
          </div>


          <button v-on:click="ccc()">BACK</button>

        </el-aside>
        <el-container>
          <router-view/>
          <el-footer>
            <a href="https://github.com/rockage/stoneren-bbs">https://github.com/rockage/stoneren-bbs</a>
          </el-footer>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>


<script>
  import myMenu from "./components/menu"
  import Tag from './components/tag.vue';

  export default {
    name: 'App',
    data() {
      return {
        tags: ['hello','hi'],
      }
    },
    methods: {
      addTag: function () {
        this.tags.push("标签" + this.tags.length);
      },
      onTagDelete: function (tag) {
        this.tags.splice(this.tags.indexOf(tag), 1);
      },
    },
    components: {
      Tag
    }
  }


</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: left;
    color: #2c3e50;
    margin-top: 0px;
  }


  a:link {
    text-decoration: none;
    color: #303133;
  }

  a:visited {
    text-decoration: none;
    color: #303133;
  }

  a:hover {
    text-decoration: none;
    color: #409EFF;
  }

  .el-header, .el-footer {
    background-color: #B3C0D1;
    color: #333;
    text-align: center;
    line-height: 60px;
  }

  .el-aside {
    background-color: #D3DCE6;
    color: #333;
    text-align: center;
    line-height: 200px;
  }
</style>
