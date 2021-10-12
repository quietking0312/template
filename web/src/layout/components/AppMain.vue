<template>
  <section class="app-main">
    <router-view v-slot="{ Component, route }">
      <transition name="fade" mode="out-in" appear>
        <component :is="Component" :key="route.fullPath" />
      </transition>
    </router-view>
  </section>
</template>

<script lang="ts">
import {defineComponent} from "vue";

export default defineComponent({
  name: "AppMain",
  setup() {
    return {
    }
  }
})
</script>

<style scoped>

</style>
