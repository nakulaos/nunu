<template>
<div class="sidebar-wrap">
  <div class="logo-wrap">
    <el-image class="logo-img" :src="NUNU" style="width: 100px; height: 33px" @click="goHome" :fit="'cover'"/>
  </div>

  <div class="menu-wrap">
    <nunu-menu :menu-list="menuList"></nunu-menu>
  </div>

  <div class="user-wrap" v-if="globalStore.$state.userInfo.id>0">

  </div>







</div>
</template>


<script setup lang="ts">
import NUNU from '@/assets/image/nunu.png';
import {GlobalStore} from "@/stores";
import {computed} from "vue";
import NunuMenu from "@/layout/components/nunu-menu.vue";

const globalStore = GlobalStore()

const menuList = computed(()=>{
  const list = [
      {
        icon:"HomeFilled",
        titleKey: "home",
        path:"/",
      },
    {
      icon: "ChatRound",
      titleKey:"topic",
      path:"/topic"
    }
    ]
  return globalStore.$state.userInfo.id>0?list:
      [
        {
          icon:"HomeFilled",
          titleKey: "home",
          path:"/",
        },
        {
          icon: "ChatRound",
          titleKey:"topic",
          path:"/topic"
        }
      ]
})

const goHome = ()=>{
  console.log("img")

}
</script>
<style scoped lang="scss">
.sidebar-wrap::-webkit-scrollbar{
  // 隐藏滚动条
  width: 0;
  height: 0;
}

.sidebar-wrap{
  width: 200px;
  height: 100%;
  position: fixed;
  right: calc(50% + var(--content-main) / 2 + 10px);
  padding: 8px 5px;
  box-sizing: border-box;
  overflow: auto;
  max-height: 90vh;


  .logo-wrap{
    margin:5px 0 0 0 ;
    position: relative;
    right: 8px;

    .logo-img{
      //margin-left: 24px;
      
      &:hover{
        cursor: pointer;
      }
    }



  }
}



</style>