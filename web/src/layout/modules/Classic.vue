<template>
  <div>
    <div id="sidebar__wrap" class="sidebar__wrap">
    </div>
    <div class="main__wrap">
      <el-scrollbar class="main__wrap--content">
        <div class="header__wrap">
        </div>
        <app-main />
      </el-scrollbar>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import {appStore} from "../../store/modules/app";
import AppMain from "../components/AppMain.vue";

export default defineComponent({
  name: "Classic",
  components: {AppMain},
  setup() {
    const layout = computed(() => appStore.layout)
    return {
      layout
    }
  }
})
</script>

<style lang="less" scoped>
@import "style";
</style>
