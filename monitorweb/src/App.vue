<template>
  <div id="app">
        <el-tabs v-model="activeName" @tab-click="handleClick" type="border-card" class="tabclass">
          <el-tab-pane name="first" >
            <span slot="label" style="font-size: 15px;font-weight: bold"><i class="el-icon-edit-outline" style="font-size: 25px"></i> 任务发布</span>
            <charts/>
          </el-tab-pane>
          <el-tab-pane name="second">
            <span slot="label" style="font-size: 15px;font-weight: bold"><i class="el-icon-document" style="font-size: 25px"></i> 日志查看</span>
            <LogViewComponent/>
          </el-tab-pane>
    <!--      <el-tab-pane label="角色管理" name="third">角色管理</el-tab-pane>-->
    <!--      <el-tab-pane label="定时任务补偿" name="fourth">定时任务补偿</el-tab-pane>-->
        </el-tabs>


<!--    <i class="el-icon-menu">菜单</i>-->
<!--    <el-menu-->
<!--        :default-active="this.$route.path"-->
<!--        router-->
<!--        class="el-menu-demo"-->
<!--        mode="horizontal"-->
<!--        background-color="#545c64"-->
<!--        text-color="#303133"-->
<!--        active-text-color="#ffd04b">-->

<!--      <el-menu-item :index="path.charts"><i class="el-icon-edit-outline"></i>mission</el-menu-item>-->
<!--      <el-menu-item :index="path.logView"><i class="el-icon-document"></i>log view</el-menu-item>-->
<!--      &lt;!&ndash;      <el-menu-item index="3"><a href="https://www.ele.me" target="_blank">订单管理</a></el-menu-item>&ndash;&gt;-->
<!--    </el-menu>-->

<!--    <router-view/>-->

  </div>
</template>

<script>
import Charts from './components/ChartsComponent.vue'
import LogViewComponent from "@/components/LogViewComponent";
export default {
  name: 'App',
  components:{
    Charts,
    LogViewComponent
  },
  data() {
    return {
      activeName: 'first',
      // logViewAgain:false,
      // chartsAgain:true,

      // path: {
      //   charts: "/charts",
      //   logView: "/logView"
      // }
    }
  },
  // created() {
  //   this.storeSomeValue()
  // },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    },


  }
}
</script>

<style>
#app {
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  /*text-align: center;*/
  /*color: #2c3e50;*/
}
.tabclass{
  background-color: #1b2945;
  color: #1b2945;
}
</style>
