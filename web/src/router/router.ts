import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'

import {GlobalStore} from "@/stores";
import {Layout} from "@/router/constant";
import homeRouter from "@/router/modules/home";


export const routes : RouteRecordRaw[] = [
    homeRouter,

]

const router = createRouter({
  history: createWebHistory('/'),
  routes: routes as RouteRecordRaw[],

})




export default router
