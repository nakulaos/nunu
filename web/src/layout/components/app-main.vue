<template>
  <router-view v-slot="{ Component , route }" class="content-wrapper" :key="key" >
<!--    key的作用去保证缓存组件内容更新-->
    <Transition mode="out-in" name="fade-transition">
<!--      缓存组件-->
      <keep-alive>
        <component
            :is="Component"
            v-if="$route.meta.keepAlive"
            :key="route.path"
        />
      </keep-alive>

    </Transition>
    <Transition mode="out-in" name="fade-transition">
<!--      组件不缓存-->
      <component
          :is="Component"
          v-if="!$route.meta.keepAlive"
      />
    </Transition>

  </router-view>
</template>
<script lang="ts" setup>
import {computed} from "vue";
import {useRoute} from "vue-router";

const key = computed(()=>{
  return useRoute().path+Math.random()
})
</script>

<style lang="scss" scoped>
.content-wrapper{
  width: 100%;
  margin: 0 auto;
  background-color: coral;
}
</style>