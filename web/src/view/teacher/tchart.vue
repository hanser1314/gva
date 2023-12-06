<template>
    <div  ref="main" style="width: 100%; height: 400px">
    </div>

</template>

<script setup>
import * as echarts from 'echarts'
import { ref,onMounted } from 'vue'

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
      text: '教师职称分布情况'
    },
    tooltip: {
      trigger: 'item'
    },
    legend: {
      top: '5%',
      left: 'center',
      // doesn't perfectly work with our tricks, disable it
      selectedMode: false
    },
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: ['40%', '70%'],
        center: ['50%', '70%'],
        // adjust the start angle
        startAngle: 180,
        label: {
          show: true,
          formatter(param) {
            // correct the percentage
            return param.name + ' (' + param.percent * 2 + '%)';
          }
        },
        data: [
          { value: 1048, name: '助教' },
          { value: 527, name: '讲师' },
          { value: 335, name: '副教授' },
          { value: 190, name: '教授' },
          // { value: 484, name: 'Union Ads' },
          // { value: 300, name: 'Video Ads' },
          {
            // make an record to fill the bottom 50%
            value: 1048 + 527 + 335 + 190,
            itemStyle: {
              // stop the chart from rendering this piece
              color: 'none',
              decal: {
                symbol: 'none'
              }
            },
            label: {
              show: false
            }
          }
        ]
      }
    ]
  }
// 使用刚指定的配置项和数据显示图表。
  myChart.setOption(option)
}


</script>


<style>


</style>
