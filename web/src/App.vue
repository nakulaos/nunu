<template>
    <el-config-provider :locale="i18nLocale" :button="buttonConfig" :message="messageConfig">
      <router-view/>
    </el-config-provider>
</template>


<script lang="ts" setup>
import {GlobalStore} from "@/stores";
import {computed, reactive} from "vue";
import zhCN from 'element-plus/es/locale/lang/zh-cn'
import en from 'element-plus/es/locale/lang/en'
import de from 'element-plus/es/locale/lang/de'
import fr from 'element-plus/es/locale/lang/fr'
import zhTW from 'element-plus/es/locale/lang/zh-tw'


const globalStore = GlobalStore()

const i18nLocale = computed(()=>{
  if(globalStore.language === 'zh') return zhCN
  if(globalStore.language==="en") return en
  if(globalStore.language==="de")return de
  if(globalStore.language==="fr")return fr
  if(globalStore.language==="zhTW")return zhTW
  return en
})

const buttonConfig = reactive({
  // 自动插入空格
  autoInsertSpace: false,
})

const messageConfig = reactive({
  // 最大信息数
  max: 3,
})

</script>


<style  lang="scss">

</style>
