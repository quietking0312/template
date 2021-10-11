<template>
  <div :class="classObj" class="app__wrap">
    <div id="sidebar__wrap" class="sidebar__wrap" :class="{'sidebar__wrap--collapsed': collapsed}">
      <!--侧边栏-->
      <sider :layout="layout" mode="vertical" />
    </div>
    <div class="main__wrap" :class="{'main__wrap--collapsed': collapsed}">
      <el-scrollbar class="main__wrap--content" :class="{'main__wrap--fixed--all': true, 'main__wrap--fixed--nav': true, 'main__wrap--fixed-tags': true}">
        <div class="hear__wrap">
          <!--head-->
          <div class="navbar__wrap">
          </div>
          <div class="tags__wrap">
            <!-- tags -->
          </div>
        </div>
        <app-main />
      </el-scrollbar>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from "vue";
import {appStore} from "@/store/modules/app";
import AppMain from "@/layout/components/AppMain.vue";
import Sider from '@c/Sider/index.vue'

export default defineComponent({
  name: "Classic",
  components: { AppMain, Sider },
  setup() {
    const layout = computed(() => appStore.layout)
    const collapsed = computed(() => appStore.collapsed)
    const showLogo = computed(() => appStore.showLogo)
    const classObj = computed(() => {
      const obj: {[key: string]: boolean} = {}
      obj[`app__wrap--${layout.value}`] = true
      return obj
    })

    return {
      classObj,
      layout,
      collapsed,
      showLogo
    }
  }
})
</script>

<style lang="less" scoped>
@import "./style.less";
</style>
