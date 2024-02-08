import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path:"/",
      name:"home",
      meta:{
        title: "广场",
        keepAlive: true,
      },
      component: ()=>import("@/views/Home.vue")
    }
  ]
})

export default router
