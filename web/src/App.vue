<template>
  <el-config-provider :locale="elementZHCNLocale">
    <router-view class="app" :class="{ grey__mode: greyMode }" />
  </el-config-provider>

</template>

<script setup lang="ts">
import {computed} from "vue";
import { useAppStore } from "@/store/modules/app";
const appStore = useAppStore()
import elementZHCNLocale from 'element-plus/lib/locale/lang/zh-cn'


const greyMode = computed(() => appStore.getGreyMode)
</script>

<style lang="less">
.size {
  min-width: @minWidth;
  width: 100;
  height: 100%;
}
html,body {
  .size;
    margin: 0;
    padding: 0;
  }
#app {
  .size;
    background: @appBg;
  }
.grey__mode {
    -webkit-filter: grayscale(100%);
    -moz-filter: grayscale(100%);
    -ms-filter: grayscale(100%);
    -o-filter: grayscale(100%);
    filter: grayscale(100%);
    filter: progid:DXImageTransform.Microsoft.BasicImage(grayscale=1);
  }
</style>
