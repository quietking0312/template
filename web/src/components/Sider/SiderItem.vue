<template>
  <template v-if="!item.meta?.hidden">
    <template v-if="hasOneShowingChild(item.children, item) && (!onlyOneChild.children || onlyOneChild.noShowingChildren) && item.meta?.alwayShow">
      <el-menu-item :index="resolvePath(onlyOneChild.path)" :class="{'submenu-title-noDropdown': !isNest}">
        <item v-if="onlyOneChild.meta" :icon="onlyOneChild.meta.icon || (item.meta && item.meta.icon)" />
        <template #title>
          <span class="anticon-item">{{ onlyOneChild.meta.title }}</span>
        </template>
      </el-menu-item>
    </template>
    <el-sub-menu v-else :popper-class="layout !== 'Top'? 'nest-popper-menu': 'top-popper-menu'" :index="resolvePath(item.path)">
      <template #title>
        <item v-if="item.meta" :icon="item.meta && icon.meta.icon" :title="item.meta.title" />
      </template>
      <sider-item v-for="child in item.children" :key="child.path" :is-nest="true" :item="child" :layout="layout" :base-path="resolvePath(child.path)" />
    </el-sub-menu>
  </template>
</template>

<script lang="ts">
import {defineComponent, PropType, ref} from "vue";
import {RouteRecordRaw} from "vue-router";
import path from 'path'
import Item from './Item.vue'
import {isExternal} from "@/utils/validate";

export default defineComponent({
  name: "SiderItem",
  components: { Item },
  props: {
    item: {
      type: Object as PropType<object>,
      required: true
    },
    isNest: {
      type: Boolean as PropType<boolean>,
      default: false
    },
    basePath: {
      type: String as PropType<string>,
      default: ''
    },
    layout: {
      type: String as PropType<string>,
      default: 'Classic'
    }
  },
  setup(props) {
    const onlyOneChild = ref<any>(null)

    function hasOneShowingChild(children: RouteRecordRaw[] = [], parent: RouteRecordRaw): boolean {

      const showingChildren: RouteRecordRaw[] = children.filter((item: RouteRecordRaw) => {
        if (item.meta && item.meta.hidden) {
          return false
        } else {
          onlyOneChild.value = item
          return true
        }
      })
      // 当只有一个子路由器时，默认情况下显示子路由器
      if (showingChildren.length === 1) {
        return true
      }
      //如果没有要显示的子路由器，则显示父路由器
      if (showingChildren.length === 0) {
        onlyOneChild.value = { ...parent, path: '', noShowingChildren: true}
        return true
      }
      return false
    }
    function resolvePath(routePath: string): string {
      if (isExternal(routePath)) {
        return routePath
      }
      return path.resolve(props.basePath, routePath)
    }
    return {
      onlyOneChild,
      hasOneShowingChild,
      resolvePath
    }
  }
});
</script>

<style scoped>

</style>
