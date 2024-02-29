<template>
  <el-menu :router="true" :unique-opened="true" >
    <template v-for="subItem in menuList" :key="subItem.path">
      <!--    有子菜单项-->
      <el-sub-menu v-if="subItem?.children?.length>0" :index="subItem.path">
        <template #title>
          <el-icon >
            <component :is="subItem.icon"></component>
          </el-icon>
          <span>{{ $t('menu.'+subItem.titleKey as string) }}</span>
        </template>
        <nunu-menu :menu-list="subItem.children"></nunu-menu>
      </el-sub-menu>
      <!--没有子菜单-->
      <el-menu-item v-else :index="subItem.path">
        <el-icon v-if="subItem?.icon">
          <component :is="subItem.icon" style="width: 20px; height: 20px"></component>
        </el-icon>
        <template #title>
          <!--        有图标-->
          <span v-if="subItem?.icon">
          {{ $t('menu.'+subItem.titleKey as string) }}
        </span>
          <!--        没有图标-->
          <span v-else style="margin-left: 10px">
          {{ $t('menu.'+subItem.titleKey as string) }}
        </span>
        </template>
      </el-menu-item>
    </template>
  </el-menu>

</template>


<script lang="ts" setup>

export interface MenuList {
  titleKey: string,
  icon: string,
  path?: string,
  children?: MenuList[],
}


defineProps<{
  menuList: MenuList[],
}>()
</script>


<style lang="scss" scoped>

</style>