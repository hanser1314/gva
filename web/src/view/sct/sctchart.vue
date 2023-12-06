<template>
  <div class="gva-search-box" ref="main" style="width: 100%; height: 800px"></div>
</template>

<script setup>
import { ref, onMounted } from 'vue'


import * as echarts from 'echarts'
const main = ref() // 使用ref创建虚拟DOM引用，使用时用main.value


onMounted(
    () => {
      init()
    }
)

const init = async() => {


// 基于准备好的dom，初始化echarts实例

  var myChart = echarts.init(main.value)

// 指定图表的配置项和数据

  var option = {
    title: {
      text: '各学院选课人数'
    },
    tooltip: {
      trigger: 'axis'
    },
    legend: {
      data: ['计算机学院', '文学院', '体育学院', '音乐学院', '生物科学学院']
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    toolbox: {
      feature: {
        saveAsImage: {}
      }
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: ['JAVA', 'C++', '大学物理', '高等数学', '马克思主义原理', '企业创新', '线性代数']
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        name: '计算机学院',
        type: 'line',
        stack: 'Total',
        data: [120, 132, 101, 134, 90, 230, 210]
      },
      {
        name: '文学院',
        type: 'line',
        stack: 'Total',
        data: [220, 182, 191, 234, 290, 330, 310]
      },
      {
        name: '体育学院',
        type: 'line',
        stack: 'Total',
        data: [150, 232, 201, 154, 190, 330, 410]
      },
      {
        name: '音乐学院',
        type: 'line',
        stack: 'Total',
        data: [320, 332, 301, 334, 390, 330, 320]
      },
      {
        name: '生物科学学院',
        type: 'line',
        stack: 'Total',
        data: [820, 932, 901, 934, 1290, 1330, 1320]
      }
    ]
  }
  // 使用刚指定的配置项和数据显示图表。
  myChart.setOption(option)
}

</script>

<style>


</style>
