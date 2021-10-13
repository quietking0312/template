<template>
  <div class="app-wrapper">
    <component :is="layout" />
  </div>
</template>

<script lang="ts">
import {defineComponent, computed} from "vue";
import {appStore} from "@/store/modules/app";
import Classic from "./modules/Classic.vue"

export default defineComponent({
  name: "Layout",
  components: { Classic },
  setup() {
    const layout = computed(() => appStore.layout)
    return {
      layout
    }
  }
})
</script>

<style lang="less" scoped>
.app-wrapper {
  position: relative;
  height: 100%;
  width: 100%;
}
</style>
