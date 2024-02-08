<template>
  <div class="sidebar-wrap">
    <div class="logo-wrap">
<!--logo区域-->
    </div>
    <a-menu >
      <a-menu-item key="0" :style="{ padding: 0, marginRight: '38px' }" disabled>
      </a-menu-item>
      <a-menu-item key="1">Home</a-menu-item>
      <a-menu-item key="2">Solution</a-menu-item>
      <a-menu-item key="3">Cloud Service</a-menu-item>
      <a-menu-item key="4">Cooperation</a-menu-item>
    </a-menu>
    <div class="user-wrap">


    </div>


  </div>
</template>
<script setup lang="ts">

</script>
<style scoped lang="scss">
.sidebar-wrap{
  width: 200px;
  height: 100vh;
  background-color: antiquewhite;


}

</style>