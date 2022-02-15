<template>
  <el-dropdown trigger="click" class="language-svg" @command="setLanguage">
    <div>
      <svg-icon icon-class="language" />
    </div>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item v-for="(v, k) in LangDict" :key="k" :disabled="lang === k" :command="k">
          {{ v }}
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script lang="ts">
import {computed, defineComponent} from "vue";
import {LangDict} from "@/lang";
import {useAppStore} from "@/store/modules/app";

export default defineComponent({
  name: "LangSelect",
  setup() {
    const appStore = useAppStore()
    const lang = computed(() => appStore.getLang)
    function setLanguage(lang: string) {
      appStore.SetLanguage(lang)
    }
    return {
      lang,
      LangDict,
      setLanguage
    }
  }
})
</script>

<style scoped>
.language-svg {
  display: inline-block;
  cursor: pointer;
  width: 20px;
  height: 20px;
  vertical-align: 10px;
}
</style>
