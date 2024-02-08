<template>
  <el-button @click="test">
    click
  </el-button>
</template>


<script lang="ts" setup>
import {GlobalStore} from "@/stores";

function test() {
  globalStore.initTest("aaaa")

}
const globalStore = GlobalStore()

</script>


<style  lang="scss">

</style>
