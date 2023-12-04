<template>
  <div>
    <video-player :options="videoOptions" />
  </div>
</template>

<script>
import VideoPlayer from '@/view/mine/VideoPlayer.vue'
export default {
  name: 'VideoExample',
  components: {
    VideoPlayer
  },
  data() {
    return {
      videoOptions: {
        autoplay: false,
        controls: true,
        poster: "https://gova.oss-cn-shenzhen.aliyuncs.com/1311148.jpg",
        height: "400px",
        width: "700px",
        sources: [
          {
            src: 'https://gova.oss-cn-shenzhen.aliyuncs.com/1.mp4',
            type: 'video/mp4',
          }
        ]
      }
    };
  }
};
</script>
