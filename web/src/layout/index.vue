<template>
  <div class="app-wrapper">
    <component :is="layout" />
  </div>
</template>

<script setup lang="ts" name="Layout">
import {computed} from "vue";
import {appStore} from "@/store/modules/app";
import Classic from "./modules/Classic.vue"
import Top from "@/layout/modules/Top.vue";
import LeftTop from "@/layout/modules/LeftTop.vue";


const layout = computed(() => appStore.layout === 'Classic'? Classic : appStore.layout === 'Top'? Top:LeftTop)

</script>

<style lang="less" scoped>
.app-wrapper {
  position: relative;
  height: 100%;
  width: 100%;
}
</style>
