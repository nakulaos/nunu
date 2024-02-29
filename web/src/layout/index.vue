<template>
<div class="app-wrapper">
  <div class="header"></div>
  <div class="main-wrapper">
    <!--      侧边栏-->
    <div class="app-sidebar" v-if="globalStore.$state.desktopModelShow">

      <sidebar></sidebar>
    </div>

    <div class="app-container">
<!--      主内容区-->
      <app-main></app-main>
    </div>
    <div class="app-rightbar">
<!--      右侧边栏-->
      <rightbar></rightbar>
    </div>
  </div>
  <div class="app-footer">
    <app-footer></app-footer>
  </div>
  <auth-card></auth-card>



</div>
</template>

<script setup lang="ts">
import {GlobalStore} from "@/stores";
import Sidebar from "@/layout/components/sidebar.vue";
import AppMain from "@/layout/components/app-main.vue";
import Rightbar from "@/layout/components/rightbar.vue";
import AppFooter from "@/layout/components/app-footer.vue";
import AuthCard from "@/components/auth-card/auth-card.vue";

const globalStore = GlobalStore();
</script>


<style lang="scss">
:root{
  --content-main:625px;
}

.app-wrapper{
  //background-color: blue;
  .header{
    height: 1vh;
    width: 100%;
    -webkit-app-region: drag;
  }

  .main-wrapper{
    //background-color: antiquewhite;
    display: flex;
    height: 94vh;
    overflow: auto;
    justify-content: center;
    padding: 10px 0;

    .app-container{
      // 最大宽度
      max-width: var(--content-main);
      //background-color: blueviolet;
      width: 100%;

    }
  }
  .main-wrapper::-webkit-scrollbar {
     display: none;
   }
  .app-footer{
    height: 5vh;
  }
}

.app-wrapper::-webkit-scrollbar {
  display: none;
}




</style>