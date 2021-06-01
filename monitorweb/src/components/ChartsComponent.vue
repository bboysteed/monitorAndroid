<template>
  <div>
    <el-row type="flex" class="row-bg">
      <el-col :span="18">
        <div id="cpu" :style="{width: '1400px', height: '400px'}"></div>
        <p/>
        <div id="ram" :style="{width: '1400px', height: '400px'}"></div>
      </el-col>
      <el-col :span="6">
        <missionForm/>
      </el-col>
    </el-row>

  </div>

</template>

<script>

import * as echarts from 'echarts'
import missionForm from "@/components/missionForm";

export default {
  components: {
    missionForm,
  },
  data() {
    return {
      loading:'',
      wspath: "ws://" + window.location.href.split("//")[1].split(":")[0] + ":8089/ws",
      socket: "",
      cpuData: [],
      cpuTemperatureData: [],
      ramData: [],
      date: []
    }
  },
  mounted() {
    this.loading = this.$loading({
      lock: true,
      text: '正在连接...',
      spinner: 'el-icon-loading',
      background: 'rgba(0, 0, 0, 0.7)'
    })


    // 初始化
    console.log(this.$store.state.apis)
    console.log(this.wspath)
    this.initWS()
    this.initChart()
  },
  methods: {
    initWS: function () {
      if (typeof (WebSocket) === "undefined") {
        alert("您的浏览器不支持socket")
      } else {

        this.socket = new WebSocket(this.wspath)
        this.socket.onopen = this.open
        this.socket.onerror = this.error
        this.socket.onmessage = this.getMessage
      }
    },
    open: function () {
      console.log("socket1连接成功")
      this.loading.close()
    },
    error: function () {
      console.log("连接错误")
    },
    getMessage: function (msg) {
      let data = JSON.parse(msg.data)
      if (data.msg === "offline") {
        this.$message({
          showClose: true,
          message: "设备掉线",
          duration: 0,
          type: "error",
          offset:40,
        })
        return
      }

      this.date.push((new Date()).toLocaleDateString() + " " + (new Date()).toLocaleTimeString())
      let rate = JSON.parse(msg.data)
      // console.log(rate.cpuInfo.all_rate)
      // console.log(rate.ramInfo.rate)
      this.cpuData.push(rate.cpuInfo.all_rate)
      this.ramData.push(rate.ramInfo.rate)
      this.cpuTemperatureData.push(rate.cpuInfo.temperature)
      // console.log(this.cpuData)
      // console.log(this.date)
      this.cpuChart.setOption({
        xAxis: {
          data: this.date,
        },
        series: [{
          name: 'cpu使用率',
          data: this.cpuData
        },
          {
            name: 'cpu温度',
            data: this.cpuTemperatureData
          }]

      })
      this.ramChart.setOption({
        xAxis: {
          data: this.date,
        },
        series: {
          name: 'ram使用率',
          data: this.ramData
        },
      })

    },
    send: function () {
      this.socket.send()
    },
    close: function () {
      console.log("socket已经关闭")
    },
    initChart: function () {
      this.cpuChart = echarts.init(document.getElementById('cpu'), 'dark')
      this.ramChart = echarts.init(document.getElementById('ram'), 'dark')
      this.cpuChart.setOption({
        title: {
          // left: 'center',
          text: 'cpu使用信息'
        },
        tooltip: {
          trigger: 'axis',
          // formatter: function (params) {
          //   console.log(params)
          //
          //   let rateParams = params[0]
          //   let temperatureParams = params[1]
          //   // console.log(params.name+":"+params.value)
          //   return rateParams.name+"-->"+rateParams.value+"\n"+temperatureParams.name+"-->"+temperatureParams.value
          // },
          axisPointer: {
            animation: false
          }
        },
        legend: {
          //color:'#fff',
          data: ['cpu使用率', "cpu温度"],

        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          data: []
        },
        yAxis: {},
        toolbox: {
          right: 10,
          feature: {
            dataZoom: {
              yAxisIndex: 'none'
            },
            restore: {},
            saveAsImage: {}
          }
        },
        dataZoom: [{
          type: 'inside'
        }],
        visualMap: {
          top: 50,
          right: 10,
          pieces: [{
            gt: 0,
            lte: 50,
            color: '#93CE07'
          },
            {
              gt: 50,
              lte: 80,
              color: '#FC7D02'
            }, {
              gt: 80,
              lte: 100,
              color: '#FD0100'
            }],
          outOfRange: {
            color: '#999'
          }
        },
        series: [{
          name: 'cpu使用率',
          type: 'line',
          smooth: 'true',
          symbol: 'none',

          //stack:'a',
          // areaStyle:{
          //   normal:{}
          // },

          data: this.cpuData,
          markLine: {
            silent: true,
            lineStyle: {
              color: '#333'
            },
            data: [{
              yAxis: 90
            }]
          }
        },
          {
            name: 'cpu温度',
            type: 'line',
            smooth: 'true',
            symbol: 'none',
            //stack:'a',
            // areaStyle:{
            //   normal:{}
            // },

            data: this.cpuTemperatureData,
            // markLine: {
            //   silent: true,
            //   lineStyle: {
            //     color: '#333'
            //   },
            //   data: [{
            //     yAxis: 90
            //   }]
            // }
          },
        ]
      })
      this.ramChart.setOption({
        title: {
          // left: 'center',
          text: 'ram使用信息'
        },
        tooltip: {
          trigger: 'axis',
          // formatter: function (params) {
          //   // console.log(params)
          //
          //   params = params[0];
          //   // console.log(params.name+":"+params.value)
          //   return params.name+"-->"+params.value
          // },
          axisPointer: {
            animation: false
          }
        },
        legend: {
          data: ['ram使用率']
        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          data: []
        },
        yAxis: {},
        toolbox: {
          right: 10,
          feature: {
            dataZoom: {
              yAxisIndex: 'none'
            },
            restore: {},
            saveAsImage: {}
          }
        },
        dataZoom: [{
          type: 'inside'
        }],
        visualMap: {
          top: 50,
          right: 10,
          pieces: [{
            gt: 0,
            lte: 50,
            color: '#93CE07'
          }, {
            gt: 50,
            lte: 80,
            color: '#FC7D02'
          }, {
            gt: 80,
            lte: 100,
            color: '#FD0100'
          }],
          outOfRange: {
            color: '#999'
          }
        },
        series: [{
          name: 'ram使用率',
          type: 'line',
          smooth: 'true',
          symbol: 'none',
          stack: 'a',
          areaStyle: {
            normal: {}
          },
          data: this.ramData
        }]
      })
    },
    // getWsAdr: function (){
    //   this.$axios.get('/getWsAdr')
    //       .then(function (response) {
    //         console.log(response)
    //         this.wspath = "ws://"+response.data+"/ws"
    //       })
    //       .catch(function (error) {
    //         console.log(error)
    //       });
    // }
  },

  destroyed() {
    // 销毁监听
    this.socket.onclose = this.close
  }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
