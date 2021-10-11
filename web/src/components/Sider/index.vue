<template>
  <div :class="{'has-logo': showLogo && layout==='Classic', 'sidebar-container--Top': layout === 'Top'}" class="sidebar-container">
    <el-scrollbar>
      <el-menu :default-active="activeMenu" :collapse="collapsed" :unique-opened="false" :mode="mode" @select="selectMenu">
        <sider-item v-for="route in routers" :key="route.path" :item="route" :layout="layout" :base-path="route.path" />
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script lang="ts">

import {computed, defineComponent, PropType} from "vue";
import {RouteRecordRaw, useRouter} from "vue-router";
import { permissionStore } from "@/store/modules/permission";
import {appStore} from "@/store/modules/app";
import SiderItem from "@/components/Sider/SiderItem.vue";
import {isExternal} from "@/utils/validate";

export default defineComponent( {
  name: "Sider",
  components: {SiderItem},
  props: {
    layout: {
      type: String as PropType<string>,
      default: 'Classic'
    },
    mode: {
      type: String as PropType<'horizontal' | 'vertical'>,
      default: 'vertical'
    }
  },
  setup() {
    const { currentRoute, push } = useRouter()
    const routers = computed((): RouteRecordRaw[] => {
      return permissionStore.routers
    })

    const activeMenu = computed(() => {
      const { meta, path } = currentRoute.value
      return path
    })

    const collapsed = computed(() => appStore.collapsed)
    const showLogo = computed(() => appStore.showLogo)

    function selectMenu(path: string) {
      if (isExternal(path)) {
        window.open(path)
      } else {
        push(path)
      }
    }

    return {
      routers,
      activeMenu,
      collapsed,
      showLogo,
      selectMenu
    }
  }
})
</script>

<style lang="less" scoped>
.sidebar-container {
  height: 100%;
    @{deep}(.svg-icon) {
  margin-right: 16px;
}
    @{deep}(.el-scrollbar) {
  width: 100%;
  height: 100%;
  .el-scrollbar__wrap {
    overflow: scroll;
    overflow-x: hidden;
    .el-menu {
      width: 100%;
      border: none;
    }
  }
}
}
.has-logo {
  height: calc(~"100% - @{topSiderHeight}");
}

.sidebar-container--Top {
    @{deep}(.el-scrollbar) {
  width: 100%;
  height: 100%;
  .el-scrollbar__wrap {
    overflow: scroll;
    overflow-x: hidden;
    .el-scrollbar__view {
      height: @topSiderHeight;
    }
    .el-menu {
      width: auto;
      height: 100%;
      border: none;
    }
  }
}
}
</style>
