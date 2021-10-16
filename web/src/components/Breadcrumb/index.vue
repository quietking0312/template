<template>
  <el-breadcrumb class="app-breadcrumb" separator="/">
    <transition-group name="breadcrumb">
      <el-breadcrumb-item v-for="(item, index) in levelList" :key="item.path">
<!--        <svg-icon v-if="item.meta.icon" :icon-class="item.meta.icon" clas="icon-breadcrumb" />-->
        <span v-if="item.redirect === 'noredirect' || index==levelList.length-1" class="no-redirect">
          {{ generateTitle(item.meta.title) }}
        </span>
        <a v-else @click.prevent="handleLink(item)">
          {{ generateTitle(item.meta.title) }}
        </a>
      </el-breadcrumb-item>
    </transition-group>
  </el-breadcrumb>
</template>

<script lang="ts">
import {defineComponent, ref, watch} from "vue";
import {RouteLocationMatched, RouteLocationNormalizedLoaded, RouteRecordRaw, useRouter} from "vue-router";
import {compile} from "path-to-regexp";
import {generateTitle} from "@/utils/i18n";

export default defineComponent({
  name: "BreadcrumbWrap",
  setup() {
    const { currentRoute, push } = useRouter()
    const levelList =ref<RouteRecordRaw[]>([])
    function getBreadcrumb() {
      let matched: any[] = currentRoute.value.matched.filter((item: RouteLocationMatched) => item.meta && item.meta.title)
      const first = matched[0]
      if (!isDashboard(first)) {
        matched = [{path: '/dashboard', meta: { title: 'dashboard', icon: 'dashboard'}}].concat(matched)
      }
      levelList.value = matched.filter((item: RouteLocationMatched) => item.meta && item.meta.title && item.meta.breadcrumb !== false)
    }
    function isDashboard(route: RouteLocationMatched) {
      const name = route && route.name
      if (!name) {
        return false
      }
      return (name as any).trim().toLocaleLowerCase() === 'Dashboard'.toLocaleLowerCase()
    }
    function pathCompile(path: string): string {
      const { params } = currentRoute.value
      const toPath = compile(path)
      return toPath(params)
    }

    function handleLink(item: RouteRecordRaw): void {
      const { redirect, path } = item
      if (redirect) {
        push(redirect as string)
        return
      }
      push(pathCompile(path))
    }
    watch(() => currentRoute.value, (route: RouteLocationNormalizedLoaded) => {
      if (route.path.startsWith('/redirect/')) {
        return
      }
      getBreadcrumb()
    }, { immediate: true})
    return {
      levelList,
      handleLink,
      generateTitle
    }
  }
})
</script>

<style lang="less" scoped>
.app-breadcrumb {
  display: inline-block;
  font-size: 14px;
  margin-left: 10px;
  .no-redirect {
    color: #97a8be;
    cursor: text;
  }
  .icon-breadcrumb {
    color: #97a8be;
    margin-right: 8px;
  }
}
</style>
