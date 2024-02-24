import {RouteRecordRaw} from "vue-router";
import {Layout} from "@/router/constant";

const homeRouter:RouteRecordRaw = {
    path:"/",
    component:Layout,
    redirect: '/',
    meta:{
        keepAlive:true,
        title:'menu-home',
    },
    children:[
        {
            path:'/',
            name:'home',
            component: ()=>import('@/views/home/index.vue')

        }
    ]
}

export default homeRouter
