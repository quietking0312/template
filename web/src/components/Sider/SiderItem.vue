<template>
  <template v-if="!siderItem.meta?.hidden">
    <template v-if="hasOneShowingChild(siderItem.children, siderItem) && (!onlyOneChild.children || onlyOneChild.noShowingChildren) && !siderItem.meta?.alwaysShow">
      <el-menu-item :index="resolvePath(onlyOneChild.path)" :class="{'submenu-title-noDropdown': !isNest}">
        <item v-if="onlyOneChild.meta" :icon="onlyOneChild?.meta?.icon || siderItem?.meta?.icon" />
        <template #title>
          <span class="anticon-item">{{ generateTitle(onlyOneChild.meta.title) }}</span>
        </template>
      </el-menu-item>
    </template>

    <el-sub-menu v-else :popper-class="layout !== 'Top'? 'nest-popper-menu': 'top-popper-menu'" :index="resolvePath(siderItem.path)">
      <template #title>
        <item v-if="siderItem.meta" :icon="siderItem?.meta?.icon" :title="generateTitle(siderItem.meta.title)" />
      </template>
      <sider-item-com v-for="child in siderItem.children" :key="child.path" :is-nest="true" :item="child" :layout="layout"
          :base-path="resolvePath(child.path)"
      />
    </el-sub-menu>
  </template>
</template>

<script lang="ts">
export default {
  name: 'SiderItemCom'
}
</script>

<script lang="ts" setup>
import {computed, PropType, ref} from "vue";
import type {RouteRecordRaw} from "vue-router";
import {isExternal} from "@/utils/validate";
import Item from "@/components/Sider/Item.vue";
import {generateTitle} from "@/utils/i18n";

const props =  defineProps({
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
})

const onlyOneChild = ref<any>(null)

const siderItem:any = computed(() => props.item)

function hasOneShowingChild(children: RouteRecordRaw[] = [], parent: RouteRecordRaw): boolean {
  const showingChildren: RouteRecordRaw[] = children.filter((item: RouteRecordRaw) => {
    if (item.meta && item.meta.hidden) {
      return false
    } else {
      // 临时设置（如果只有一个显示子项，则将使用)
      onlyOneChild.value = item
      return true
    }
  })

  // 当只有一个子路由器时，默认情况下显示子路由器
  if (showingChildren.length === 1) {
    return true
  }

  // 如果没有要显示的子路由器，则显示父路由器
  if (showingChildren.length === 0) {
    onlyOneChild.value = { ...parent, path: '', noShowingChildren: true }
    return true
  }
  return false
}

function resolvePath(routePath: string): string {
  if (isExternal(routePath)) {
    return routePath
  }
  // path.resolve 方法无法使用, 临时解决方案
  return routePath? (props.basePath + '/' + routePath).replace("//", "/"): props.basePath
}

</script>

<style scoped>

</style>
