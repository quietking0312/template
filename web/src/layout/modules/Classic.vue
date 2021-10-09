<template>
  <div>

  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from "vue";
import {appStore} from "@/store/modules/app";

export default defineComponent({
  name: "Classic",
  components: {},
  setup() {
    const layout = computed(() => appStore.layout)
    const collapsed = computed(() => appStore.collapsed)
    return {
      layout,
      collapsed
    }
  }
})
</script>

<style scoped>

</style>
