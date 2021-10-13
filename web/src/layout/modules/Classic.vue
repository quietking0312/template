<template>
  <div :class="classObj" class="app__wrap">
    <div id="sidebar__wrap" class="sidebar__wrap" :class="{'sidebar__wrap--collapsed': collapsed}">
      <sider :layout="layout" mode="vertical" />
    </div>
    <div class="main__wrap" :class="{'main__wrap--collapsed': collapsed}">
      <el-scrollbar class="main__wrap--content">
        <div class="header__wrap">
        </div>
        <app-main />
      </el-scrollbar>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent, computed} from "vue";
import {appStore} from "@/store/modules/app";
import AppMain from "@/layout/components/AppMain.vue";
import Sider from "@/components/Sider/index.vue";

export default defineComponent({
  name: "Classic",
  components: {Sider, AppMain},
  setup() {
    const layout = computed(() => appStore.layout)
    const collapsed = computed(() => appStore.collapsed)
    const showLogo = computed(() => appStore.showLogo)
    const showTags = computed(() => appStore.showTags)

    const classObj = computed(() => {
      const obj = {}
      obj[`app__wrap--${layout.value}`] = true
      return obj
    })

    function setCollapsed(collapsed: boolean): void {
      appStore.SetCollapsed(collapsed)
    }

    return {
      classObj,
      layout,
      collapsed,
      showLogo,
      showTags,
      setCollapsed
    }
  }
})
</script>

<style lang="less" scoped>
@import "style.less";
</style>
