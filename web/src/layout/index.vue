<template>
  <div class="app-wrapper">
    <component :is="layout" />
  </div>
</template>

<script setup lang="ts" name="Layout">
import {computed} from "vue";
import { useAppStore } from "@/store/modules/app";
const appStore = useAppStore()
import Classic from "./modules/Classic.vue"
import Top from "@/layout/modules/Top.vue";
import LeftTop from "@/layout/modules/LeftTop.vue";


const layout = computed(() => appStore.getLayout === 'Classic'? Classic : appStore.getLayout === 'Top'? Top:LeftTop)

</script>

<style lang="less" scoped>
.app-wrapper {
  position: relative;
  height: 100%;
  width: 100%;
}
</style>
