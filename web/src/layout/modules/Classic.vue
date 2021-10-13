<template>
  <div :class="classObj" class="app__wrap">
    <div id="sidebar__wrap" class="sidebar__wrap" :class="{'sidebar__wrap--collapsed': collapsed}">
      <logo v-if="showLogo && layout === 'Classic'" :collapsed="collapsed" />
      <sider :layout="layout" mode="vertical" />
    </div>
    <div class="main__wrap" :class="{'main__wrap--collapsed': collapsed}">
      <el-scrollbar class="main__wrap--content"
                    :class="{'main__wrap--fixed--all': fixedHeader && showNavbar && showTags,
                    'main__wrap--fixed--nav': fixedHeader && showNavbar && !showTags,
                    'main__wrap--fixed--tags': fixedHeader && !showNavbar && showTags}">
        <div class="header__wrap" :class="{'header__wrap--fixed': fixedHeader,
        'header__wrap--collapsed': fixedHeader && collapsed}">
          <div v-if="showNavbar" class="navbar__wrap">
            <hamburger v-if="showHamburger" id="hamburger-container" :collapsed="collapsed" class="hover-container" @toggleClick="setCollapsed" />
            <breadcrumb v-if="showBreadcrumb" id="breadcrumb-container" />
            <div v-if="showScreenfull|| showUserInfo" class="navbar__wrap--right">
              <screenfull v-if="showScreenfull" class="hover-containeer screenfull-container" />
              <user-info v-if="showUserInfo" class="hover-container user-container" />
            </div>
          </div>
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
import {defineComponent, computed} from "vue";
import {appStore} from "@/store/modules/app";
import AppMain from "@/layout/components/AppMain.vue";
import Sider from "@/components/Sider/index.vue";
import Breadcrumb from "@/components/Breadcrumb/index.vue";
import Hamburger from '@/components/Hamburger/index.vue';
import Screenfull from "@/components/Screenfull/index.vue";
import Logo from "@/components/Logo/index.vue";
import UserInfo from "@/components/UserInfo/index.vue";
import TagsView from "@/components/TagsView/index.vue";
export default defineComponent({
  name: "Classic",
  components: {TagsView, UserInfo, Logo, Screenfull, Breadcrumb, Sider, AppMain, Hamburger},
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
