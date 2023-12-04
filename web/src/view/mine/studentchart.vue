<template>
  <div ref="main" style="width: 100%; height: 400px"></div>

</template>

<script setup>


import * as echarts from 'echarts'
import { ref, onMounted } from 'vue'
import { getStudentByDept } from '@/api/student.js'


const main = ref() // 使用ref创建虚拟DOM引用，使用时用main.value



onMounted(
  () => {

init()
  }
)

function init() {
  // 基于准备好的dom，初始化echarts实例
  var myChart = echarts.init(main.value)
  
  // 指定图表的配置项和数据
  var option = {
    title: {
      text: 'ECharts 入门示例'
    },
    tooltip: {},
    legend: {
      data: ['销量']
    },
    xAxis: {
      data: ['计算机学院', '羊毛衫', '雪纺衫', '裤子', '高跟鞋', '袜子']
    },
    yAxis: {},
    series: [
      {
        name: '销量',
        type: 'bar',
        data: [1,2,3,4,5,6]
      }
    ]
  }
  // 使用刚指定的配置项和数据显示图表。
  myChart.setOption(option)
}

</script>

