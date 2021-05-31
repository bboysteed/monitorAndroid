import Vue from 'vue'
import VueRouter from "vue-router"
import ChartsComponent from "@/components/ChartsComponent";
import LogViewComponent from "@/components/LogViewComponent";
Vue.use(VueRouter)

export default new VueRouter({
    routes: [
        {
            path:"/",
            redirect: '/charts',

        },
        {
            path:"/charts",
            component:ChartsComponent
        },
        {
            path: "/logView",
            component:LogViewComponent
        }
    ]
    }
)
