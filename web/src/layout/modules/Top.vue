<template>
  <div :class="classObj" class="app__wrap">
    <!-- Top -->
    <div class="sidebar__wrap--Top">
      <div>
        <logo v-if="showLogo" :collapsed="collapsed" />
      </div>
      <div id="sidebar__wrap" class="sidebar__wrap--Top">
        <sider :layout="layout" mode="horizontal" />
      </div>
      <div>
        <div v-if="showScreenfull || showUserInfo" class="navbar__wrap--right">
          <screenfull v-if="showScreenfull" class="hover-container screenfull-container" />
          <user-info v-if="showUserInfo" class="hover-container user-container" />
        </div>
      </div>
    </div>
    <!-- Top -->
    <div class="main__wrap" :class="{'main__wrap--collapsed': collapsed}">
      <el-scrollbar class="main__wrap--content" :class="{
        'main__wrap--fixed--all': fixedHeader && showNavbar && showTags,
        'main__wrap--fixed--nav': fixedHeader && showNavbar && !showTags,
        'main__wrap--fixed--tags': fixedHeader && !showNavbar && showTags}">
        <div class="header__wrap" :class="{'header__wrap--fixed': fixedHeader,
        'header__wrap--collapsed': fixedHeader && collapsed}">
          <div v-if="showTags" id="tag-container" class="tags__wrap">
            <tags-view />
          </div>
        </div>
        <app-main />
      </el-scrollbar>
    </div>
  </div>
</template>

<script lang="ts">
import {computed, defineComponent} from "vue";
import {appStore} from "@/store/modules/app";
import Logo from "@/components/Logo/index.vue";
import Sider from "@/components/Sider/index.vue";
import TagsView from "@/components/TagsView/index.vue";
import AppMain from "@/layout/components/AppMain.vue";
import Screenfull from "@/components/Screenfull/index.vue";
import UserInfo from '@/components/UserInfo/index.vue'
export default defineComponent({
  name: "Top",
  components: {
    AppMain,
    TagsView,
    Sider,
    Logo,
    Screenfull,
    UserInfo
  },
  setup() {
    const layout = computed(() => appStore.layout)
    const collapsed = computed(() => appStore.collapsed)
    const showLogo = computed(() => appStore.showLogo)
    const showTags = computed(() => appStore.showTags)
    const showBreadcrumb = computed(() => appStore.showBreadcrumb)
    const showHamburger = computed(() => appStore.showHamburger)
    const showScreenfull = computed(() => appStore.showScreenfull)
    const showUserInfo = computed(() => appStore.showUserInfo)
    const showNavbar = computed(() => appStore.showNavbar)
    const fixedHeader = computed(() => appStore.fixedHeader)

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
      showBreadcrumb,
      showHamburger,
      showScreenfull,
      showUserInfo,
      showNavbar,
      fixedHeader,
      setCollapsed
    }
  }
})
</script>

<style lang="less" scoped>
@import "style.less";
</style>
