<template>
  <div>
    <el-card style="margin-left: 20px" shadow="hover">
      <div slot="header" class="clearfix">
        <span style="font-size: 15px;font-weight: bold">任务创建</span>
        <el-button style="float: right; padding: 3px 0" type="text" @click="formTips">填写说明</el-button>
        <el-drawer
            title="填写说明"
            :visible.sync="drawer"
            direction="rtl"
            >
          <span>填写说明</span>
        </el-drawer>
      </div>
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
        <el-form-item label="测试名称" prop="missionName">
          <el-input v-model="ruleForm.missionName"></el-input>
        </el-form-item>
        <el-form-item label="被测app" prop="apps">
          <el-select v-model="ruleForm.apps"
                     multiple placeholder="请选择要测试的app"
                     size="medium"
                     style="width: 300px"
                     clearable
                     filterable
          >
            <el-option

                v-for="app in apps"
                :key="app.value"
                :label="app.label"
                :value="app.value">
              <!--          <span style="float: left">{{ app.label }}</span>-->
              <!--          <span style="float: right; color: #8492a6; font-size: 13px">{{ app.value }}</span>-->

            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="测试线程">
          <el-input-number v-model="ruleForm.thread" :min="1" :max="100" label="测试线程" :step="1" size="small"
                           step-strictly></el-input-number>

          <el-popover
              placement="top-start"
              width="200"
              trigger="hover"
              content="数字越大代表测试流量越大，系统越容易出问题">
            <!--              <el-button slot="reference">hover 激活</el-button>-->
            <el-tag slot="reference" size="small" type="warning" style="margin-left: 20px" effect="plain">输入1-100的数字
            </el-tag>
          </el-popover>

          <!--        <el-input v-model="ruleForm.thread"></el-input>-->
        </el-form-item>

        <el-form-item label="随机事件数">
          <el-input-number v-model="ruleForm.operaNums" :min="10" :max="20000" label="测试线程" :step="10" size="small"
                           step-strictly></el-input-number>

          <el-popover
              placement="top-start"
              width="200"
              trigger="hover"
              content="数字越大代表每次测试的随机事件越多测试时长越长">
            <!--              <el-button slot="reference">hover 激活</el-button>-->
            <el-tag slot="reference" size="small" type="warning" style="margin-left: 20px" effect="plain">输入10-20000的数字
            </el-tag>
          </el-popover>

          <!--        <el-input v-model="ruleForm.thread"></el-input>-->
        </el-form-item>

        <el-form-item label="操作间隔(ms)">
          <el-input-number v-model="ruleForm.throehold" :min="300" :max="3000" label="测试线程" :step="100" size="small"
                           step-strictly></el-input-number>

          <el-popover
              placement="top-start"
              width="200"
              trigger="hover"
              content="数字越大，代表模拟操作越频繁">
            <!--              <el-button slot="reference">hover 激活</el-button>-->
            <el-tag slot="reference" size="small" type="warning" style="margin-left: 20px" effect="plain">
              输入300-3000的数字
            </el-tag>
          </el-popover>

          <!--        <el-input v-model="ruleForm.thread"></el-input>-->
        </el-form-item>


        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
          <el-button @click="resetForm('ruleForm')">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <p></p>
    <el-card class="box-card" style="margin-left: 20px;margin-top: 10px" shadow="hover">
      <div slot="header" class="clearfix">
        <span style="font-size: 15px;font-weight: bold">任务详情</span>
        <el-button type="text" style="margin-left: 5px" @click="showMissionsDetile">进度</el-button>
        <el-button style="float: right; padding: 3px 0" type="text" @click="clearAllMissions">清除任务</el-button>
      </div>
      <el-table
          v-loading="tableLoading"
          element-loading-text="拼命加载中"
          element-loading-spinner="el-icon-loading"
          element-loading-background="rgba(0, 0, 0, 0.8)"
          :data="missions"
          style="width: 100%">
        <el-table-column
            prop="id"
            label="序号"
            width="50">
        </el-table-column>
        <el-table-column
            prop="name"
            label="任务名称"
            width="130">
        </el-table-column>
        <el-table-column
            prop="startTime"
            label="开始时间"
            width="130">
        </el-table-column>
        <el-table-column
            prop="state"
            label="状态">
          <template slot-scope="scope">
            <el-tag
                effect="dark"
                style="width: 60px;text-align: center"
                :type="stateType"
                disable-transitions>{{ scope.row.state }}
            </el-tag>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <el-dialog title="任务详情" :visible.sync="dialogTableVisible">
      <el-table :data="missions">
        <el-table-column property="id" label="序号" width="100"></el-table-column>
        <el-table-column property="name" label="任务名称" width="200"></el-table-column>
        <el-table-column property="startTime" label="开始时间" width="200"></el-table-column>
        <el-table-column property="progress" label="进度" width="100"></el-table-column>
        <el-table-column property="sum" label="总数" width="100"></el-table-column>
        <el-table-column property="state" label="状态"></el-table-column>
      </el-table>
    </el-dialog>
  </div>

</template>

<script>
export default {
  name: "missionForm",
  data() {
    return {
      drawer:false,
      dialogTableVisible:false,
      tableLoading:'',
      wspath: "ws://" + window.location.href.split("//")[1].split(":")[0] + ":8089/missionStatews",
      // wspath:'ws://''/ws',

      socket: "",
      stateType:"warning",
      androidID: '',
      apps: [],
      ruleForm: {
        missionName: '',
        apps: [],
        thread: '10',
        operaNums:'100',
        throehold: '1000',
        // date1: '',
        // date2: '',
        // delivery: false,
        // type: [],
        // resource: '',
        // desc: ''
      },
      missions: [],
      rules: {
        missionName: [
          {required: true, message: '请输入测试名称', trigger: 'blur'},
          {min: 1, max: 20, message: '长度在 1 到 20 个字符', trigger: 'change'}
        ],
        apps: [
          {required: true, message: '请选择测试apps', trigger: 'blur'}
        ],
        // thread: [
        //   {required: true, message: '请输入测试线程数', trigger: 'blur'},
        //   {min: 1, max: 50, message: '输入 1 到 50的数字', trigger: 'change'}
        // ],
        // date1: [
        //   {type: 'date', required: true, message: '请选择日期', trigger: 'change'}
        // ],
        // date2: [
        //   {type: 'date', required: true, message: '请选择时间', trigger: 'change'}
        // ],
        // type: [
        //   {type: 'array', required: true, message: '请至少选择一个活动性质', trigger: 'change'}
        // ],
        // resource: [
        //   {required: true, message: '请选择活动资源', trigger: 'change'}
        // ],
        // desc: [
        //   {required: true, message: '请填写活动形式', trigger: 'blur'}
        // ]
      }
    }
  },
  // watch:{
  //   androidID(val, oldVal) =>
  // },


  created() {
    this.getAppList()
  },
  methods: {
    showMissionsDetile(){
      this.dialogTableVisible=true
    },
    submitForm(formName) {
      if (this.missions.length === 1) {
        this.$message.warning("单任务系统！")
        return
      }
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.tableLoading = true
          console.log(this.ruleForm)
          const formData = {
            name: this.ruleForm.missionName,
            apps: this.ruleForm.apps,
            thread: this.ruleForm.thread,
            operanums:this.ruleForm.operaNums,
            interval: this.ruleForm.throehold,
          }
          // let conf = {headers:{
          //     "Content-Type": "application/json",
          //     // "Access-Control-Allow-Origin":"*",
          //   }}
          this.$axios.post(
              "/commitMission",
              // this.$qs.stringify(formData),
              JSON.stringify(formData),
              // conf

          ).then(res => {
            console.log(res)
            if(res.data.msg==="running"){
              this.$message.warning("上一任务未结束！")
              return
            }
            this.$message.success("提交成功！")
            this.missions.push({id:1,name:this.ruleForm.missionName,state:"",startTime:"",progress:0,sum:this.ruleForm.thread*this.ruleForm.apps.length,deviceoffline:false})

            this.startWs()
          }).catch(err => {
            console.log(err)
          })

        } else {
          this.$message.error("请检查参数！")
          return false;
        }
      });
    },
    resetForm(formName) {
      this.ruleForm.thread = 10
      this.ruleForm.throehold = 1000
      this.$refs[formName].resetFields();
    },
    formTips() {
      this.drawer = true
    },
    clearAllMissions() {
      if(this.missions[0].state==="running"){
        this.$message.warning("当前任务未结束!")
        return
      }
        this.missions = []
    },
    startWs() {
      if (typeof (WebSocket) === "undefined") {
        this.$message.error("您的浏览器不支持socket")

      } else {
        this.socket = new WebSocket(this.wspath)
        this.socket.onopen = this.open
        this.socket.onerror = this.error
        this.socket.onmessage = this.getMessage
      }
    },
    open: function () {
      console.log("socket2连接成功")
    },
    error: function () {
      console.log("连接错误")
    },
    getMessage: function (msg) {
      // if (msg.msg === "offline") {
      //   this.$message({
      //         message: "设备掉线",
      //         duration: 0,
      //       }
      //   )
      //   return
      // }
      let data = JSON.parse(msg.data)
      if (data.deviceoffline){
        this.stateType="error"
        this.missions[0].state = "offline"
        this.$message({
          message:"设备掉线！",
          type:"error",
          duration:0,
          showClose:true,
        })
        return
      }
      this.missions[0].progress = data.progress
      this.missions[0].state = data.state
      this.missions[0].startTime = data.startTime
      if (data.state==="finished"){
        this.stateType = "success"
        this.$message({
          message:"测试完毕！",
          type:"success",
          showClose:true,
          duration:10000,
            })
      }else {
        this.stateType = "warning"
      }
      this.tableLoading=false


    },
    send: function () {
      this.socket.send(this.ruleForm)
    },
    close: function () {
      console.log("socket已经关闭")
    },

    getAppList() {

      this.$axios.get(
          '/getAppList',
      ).then(res => {
        console.log(res)
        if (res.data.msg === "offline") {

          this.$message.error("设备掉线")
          return
        }
        this.$message.success("连接成功！")
        let apps = res.data.packages
        this.androidID = res.data.androidID
        for (let i = 0; i < apps.length; i++) {
          this.apps.push({label: apps[i], value: apps[i]})
        }
      }).catch(err => {
        console.log(err)
        this.$message.error("连接后台失败，请刷新页面")
      })
    }
  }
}
</script>

<style scoped>

</style>