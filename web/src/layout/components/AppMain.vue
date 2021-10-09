<template>
  <section class="app-main">
    <router-view v-slot="{ Component, route}">
      <transition name="fade" mode="out-in" appear>
        <keep-alive>
          <component :is="Comment" :key="route.fullPath" />
        </keep-alive>
      </transition>
    </router-view>
  </section>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "AppMain",
  setup() {

  }
})
</script>

<style scoped>

</style>
