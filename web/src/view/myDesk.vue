<template>
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card gva-top-card">
        <div class="gva-top-card-left">
          <div class="gva-top-card-left-title">欢迎，{{ username }}，请开始一天的工作吧</div>
          <div class="gva-top-card-left-dot">{{ weatherInfo }}</div>
          <el-row class="my-8 w-[500px]">
            <el-col :span="8" :xs="24" :sm="8">
              <div class="flex items-center">
                <el-icon class="dashboard-icon">
                  <sort />
                </el-icon>
                今日流量 (1231231)
              </div>
            </el-col>
            <el-col :span="8" :xs="24" :sm="8">
              <div class="flex items-center">
                <el-icon class="dashboard-icon">
                  <avatar />
                </el-icon>
                总用户数 (24001)
              </div>
            </el-col>
            <el-col :span="8" :xs="24" :sm="8">
              <div class="flex items-center">
                <el-icon class="dashboard-icon">
                  <comment />
                </el-icon>
                好评率 (99%)
              </div>
            </el-col>
          </el-row>
          <div>
            <!-- <div class="gva-top-card-left-item">
              使用教学：
              <a style="color:#409EFF" target="view_window"
                href="https://www.bilibili.com/video/BV1Rg411u7xH/">https://www.bilibili.com/video/BV1Rg411u7xH</a>
            </div>
            <div class="gva-top-card-left-item">
              插件仓库：
              <a style="color:#409EFF" target="view_window"
                href="https://plugin.gin-vue-admin.com/#/layout/home">https://plugin.gin-vue-admin.com</a>
            </div> -->
          </div>
        </div>
        <img src="@/assets/dashboard.png" class="gva-top-card-right" alt>
      </div>
    </div>
    <div class="gva-card-box">
      <div class="gva-card quick-entrance">
        <div class="gva-card-title">快捷入口</div>
        <!-- <el-row :gutter="20"> -->
          <el-row>
          <!-- <el-col v-for="(card, key) in toolCards" :key="key" :span="4" :xs="8" class="quick-entrance-items"
            @click="toTarget(card.name)">
            <div class="quick-entrance-item">
              <div class="quick-entrance-item-icon" :style="{ backgroundColor: card.bg }">
                <el-icon>
                  <component :is="card.icon" :style="{ color: card.color }" />
                </el-icon>
              </div>
              <p>{{ card.label }}</p>
            </div>
          </el-col> -->
          <gdmap />
        </el-row>
      </div>
    </div>

    <!-- <div class="gva-card-box">
      <div class="gva-card">
        <div class="gva-card-title">数据统计</div>
        <div class="p-4"> -->
          <!-- <el-row :gutter="20"> -->
            <!-- <el-col :xs="24" :sm="18"> -->
            
              <!-- <echarts-line /> -->
           
            <!-- <el-col :xs="24" :sm="6"> -->
              <!-- <dashboard-table /> -->
            <!-- </el-col> -->

            <!-- <videoplay />
              <videoplay />
              <videoplay /> -->
              <!-- <el-carousel :interval="4000" type="card" height="600px">
                  <el-carousel-item v-for="item in 6" :key="item"> -->
                    <!-- <h3 text="2xl" justify="center">{{ item }}</h3> -->
                    <!-- <videoplay /> -->
                  <!-- </el-carousel-item> -->
                <!-- </el-carousel> -->
          <!-- </el-row> -->
        <!-- </div> -->
      <!-- </div> -->
    <!-- </div> -->

    <div class="gva-card">
        <div class="gva-card-title">数据统计</div>
        <div class="p-4">
          <el-row :gutter="20">
            <el-col
              :xs="24"
              :sm="18"
            >
              <echarts-line />
            </el-col>
            <el-col
              :xs="24"
              :sm="6"
            >
            <el-calendar v-model="value" />
            </el-col>
          </el-row>
        </div>
      </div>
  </div>
</template>
  
<script setup>
import EchartsLine from '@/view/dashboard/dashboardCharts/echartsLine.vue'
// import DashboardTable from '@/view/dashboard/dashboardTable/dashboardTable.vue'
// import xgPlayer from '@/view/mine/xgplay.vue'
import videoplay from '@/view/mine/video.vue'
import gdmap from '@/view/mine/map.vue'

import { useRouter } from 'vue-router'
import { useWeatherInfo } from '@/view/dashboard/weather.js'

import { ref } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
defineOptions({
  name: 'myDesk'
})

const weatherInfo = useWeatherInfo()

// 我的参数
const username = ref('')
const userStore = useUserStore()

// username.value = userStore.userInfo.authority.authorityName
username.value = userStore.userInfo.userName

const value = ref(new Date())

const router = useRouter()

const toTarget = (name) => {
  router.push({ name })
}

</script>
<style lang="scss" scoped>
.page {
  @apply p-0;

  .gva-card-box {
    @apply p-4;

    &+.gva-card-box {
      @apply pt-0;
    }
  }

  .gva-card {
    @apply box-border bg-white rounded h-auto px-6 py-8 overflow-hidden shadow-sm;

    .gva-card-title {
      @apply pb-5 border-t-0 border-l-0 border-r-0 border-b border-solid border-gray-100;
    }
  }

  .gva-top-card {
    @apply h-72 flex items-center justify-between text-gray-500;

    &-left {
      @apply h-full flex flex-col w-auto;

      &-title {
        @apply text-3xl text-gray-600;
      }

      &-dot {
        @apply mt-4 text-gray-600 text-lg;
      }

      &-item {
        +.gva-top-card-left-item {
          margin-top: 24px;
        }

        margin-top: 14px;
      }
    }

    &-right {
      height: 600px;
      width: 600px;
      margin-top: 28px;
    }
  }

  ::v-deep(.el-card__header) {
    @apply p-0 border-gray-200;
  }

  .card-header {
    @apply pb-5 border-b border-solid border-gray-200 border-t-0 border-l-0 border-r-0;
  }

  .quick-entrance-items {
    @apply flex items-center justify-center text-center text-gray-800;

    .quick-entrance-item {
      @apply px-8 py-6 flex items-center flex-col transition-all duration-100 ease-in-out rounded-lg cursor-pointer;

      &:hover {
        @apply shadow-lg;
      }

      &-icon {
        @apply flex items-center h-16 w-16 rounded-lg justify-center mx-0 my-auto text-2xl;
      }

      p {
        @apply mt-2.5;
      }
    }
  }
}

.dashboard-icon {
  @apply flex items-center text-xl mr-2 text-blue-400;
}

// 新增css
// .el-carousel__item h3 {
//   color: #475669;
//   opacity: 0.75;
//   line-height: 200px;
//   margin: 0;
//   text-align: center;
// }

// .el-carousel__item:nth-child(2n) {
//   background-color: #99a9bf;
// }

// .el-carousel__item:nth-child(2n + 1) {
//   background-color: #d3dce6;
// }

</style>
