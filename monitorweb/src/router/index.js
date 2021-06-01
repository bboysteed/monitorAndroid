import Vue from 'vue'
import VueRouter from "vue-router"
// import ChartsComponent from "@/components/ChartsComponent";
// import LogViewComponent from "@/components/LogViewComponent";
Vue.use(VueRouter)

export default new VueRouter({
    routes: [
        {
            path:"/",
            // redirect: '/charts',
            name: 'index',
            meta:{
                // 页面标题title
                title: '测试任务发布监控平台'
            }

        },
        // {
        //     path:"/charts",
        //     component:ChartsComponent
        // },
        // {
        //     path: "/logView",
        //     component:LogViewComponent
        // }
    ]
    }
)
