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
            <breadcrumb-wrap v-if="showBreadcrumb" id="breadcrumb-container" />
            <div v-if="showScreenfull || showUserInfo" class="navbar__wrap--right">
              <screenfull v-if="showScreenfull" class="hover-containeer screenfull-container" />
              <lang-select v-if="showLanguage" class="hover-containeer language-container" />
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

<script setup lang="ts">
import { computed } from "vue";
import { useAppStore } from "@/store/modules/app";
import AppMain from "@/layout/components/AppMain.vue";
import Sider from "../components/Sider/index.vue";
import BreadcrumbWrap from "../components/Breadcrumb/index.vue";
import Hamburger from '../components/Hamburger/index.vue';
import Screenfull from "../components/Screenfull/index.vue";
import Logo from "../components/Logo/index.vue";
import UserInfo from "../components/UserInfo/index.vue";
import TagsView from "../components/TagsView/index.vue";
import LangSelect from "@/components/LangSelect/index.vue";
const appStore = useAppStore()

const layout = computed(() => appStore.getLayout)
const collapsed = computed(() => appStore.getCollapsed)
const showLogo = computed(() => appStore.getShowLogo)
const showTags = computed(() => appStore.getShowTags)
const showBreadcrumb = computed(() => appStore.getShowBreadcrumb)
const showHamburger = computed(() => appStore.getShowHamburger)
const showScreenfull = computed(() => appStore.getShowScreenfull)
const showUserInfo = computed(() => appStore.getShowUserInfo)
const showNavbar = computed(() => appStore.getShowNavbar)
const showLanguage = computed(() => appStore.getShowLanguage)
const fixedHeader = computed(() => appStore.getFixedHeader)
const classObj = computed(() => {
  const obj = {}
  obj[`app__wrap--${layout.value}`] = true
  return obj
})

function setCollapsed(collapsed: boolean): void {
  appStore.SetCollapsed(collapsed)
}

</script>

<style lang="less" scoped>
@import "style.less";
</style>
