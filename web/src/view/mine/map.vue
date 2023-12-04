<template>
    <!-- 百度地图  -->
    <el-card class="card">
        <div id="bai-du-map" class="map">
        </div>
    </el-card>

    <!-- 技术支持和联系方式  -->
</template>
 
<script>
import AMapLoader from "@amap/amap-jsapi-loader";
window._AMapSecurityConfig = {
    // 设置安全密钥
    securityJsCode: 'fc18084e55a86c5bc15182c70eed674d',
}


export default {
    props: ["iptId"],
    data() {
        return {
            map: null,
            mouseTool: null,
            overlays: [],
            auto: null,
            placeSearch: null,
        }
    },
    methods: {
        initMap() {
            AMapLoader.load({
                "key": "f5cb72f847dc939fa337ab9f6a219455", // 申请好的Web端开发者Key，首次调用 load 时必填
                "version": "2.0",   // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
                "plugins": ["AMap.Scale", "AMap.ToolBar", "AMap.ControlBar", "AMap.MouseTool", "AMap.MapType", "AMap.HawkEye"],           // 需要使用的的插件列表，如比例尺'AMap.Scale'等
                // "plugins": [],           // 需要使用的的插件列表，如比例尺'AMap.Scale'等
            }).then((AMap) => {
                this.map = new AMap.Map('bai-du-map', {
                    viewMode: "2D",  //  是否为3D地图模式
                    zoom: 18,  //设置地图显示的缩放级别
                    center: [112.495173, 23.107184], //中心点坐标  肇庆学院
                    resizeEnable: true
                });
                this.map.addControl(new AMap.Scale());   // 添加左下角的比例尺
                let toolBar = new AMap.ToolBar({
                    position: {
                        // bottom: '210px',
                        // right: '35px'
                        top: '210px',
                        left: '35px'
                    }
                });
                let controlBar = new AMap.ControlBar({
                    position: {
                        // bottom: '280px',
                        // right: '10px',
                        top: '280px',
                        left: '10px'
                    },
                });
                this.map.addControl(toolBar);   // 添加右下角的放大缩小
                this.map.addControl(controlBar);   // 方向盘
                this.map.addControl(new AMap.MapType());   // 添加右上角的标准图和卫星图  和路况
                this.map.addControl(new AMap.HawkEye());   // 添加地图放大镜
            }).catch(e => {
                console.log(e);
            });
        },
    },
    mounted() {
        this.initMap();
    },
};
</script>


<style>
.map {
    height: 768px;
    width: 1600px;
}

.card {
    width: 1640px;
    height: 800px;
}
</style>
 