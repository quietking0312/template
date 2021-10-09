<template>
  <div class="app-wrapper">
    <component :is="layout" />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from "vue";
import Classic from "@/layout/modules/Classic.vue";
import { appStore } from "@/store/modules/app";

export default defineComponent({
  name: 'Layout',
  components: {Classic},

  setup() {
    const layout = computed(() => appStore.layout)
    return {
      layout
    }
  }
})
</script>
