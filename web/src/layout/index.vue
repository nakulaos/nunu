<template>
<div class="app-wrapper">
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



</div>
</template>

<script setup lang="ts">
import {GlobalStore} from "@/stores";
import Sidebar from "@/layout/components/sidebar.vue";
import AppMain from "@/layout/components/app-main.vue";
import Rightbar from "@/layout/components/rightbar.vue";
import AppFooter from "@/layout/components/app-footer.vue";

const globalStore = GlobalStore();
</script>


<style lang="scss">
:root{
  --content-main:625px;
}

.app-wrapper{
  //background-color: blue;
  margin: 5px auto;
  .main-wrapper{
    //background-color: antiquewhite;
    display: flex;
    height: 90vh;
    overflow: auto;
    justify-content: center;

    .app-container{
      // 最大宽度
      max-width: var(--content-main);
      //background-color: blueviolet;
      width: 100%;
    }

  }
  .app-footer{
    height: 10vh;
  }


}





</style>