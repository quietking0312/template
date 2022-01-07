<template>
  <section class="app-main">
    <router-view v-slot="{ Component, route }">
      <transition name="fade" mode="out-in" appear>
        <keep-alive :include="getCaches">
          <div>
            <component :is="Component" :key="route.fullPath" />
          </div>
        </keep-alive>
      </transition>
    </router-view>
  </section>
</template>

<script setup lang="ts">
import {defineComponent} from "vue";
import {useCache} from "@/components/ParentView/useCache";
const { getCaches } = useCache(true)
</script>

<style lang="less" scoped>
.app-main {
// overflow: hidden;
  padding: 20px;
}
</style>
