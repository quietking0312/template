<template>
  <i v-if="icon && icon.includes('el-icon')" :class="[icon, 'sub-el-icon', 'anticon']" />
  <svg-icon v-else-if="icon" :icon-class="icon" class="anticon" />
  <slot name="title">
    <span class="anticon-item">{{ title }}</span>
  </slot>
</template>

<script lang="ts">
import {defineComponent, PropType} from "vue";
import SvgIcon from "@/components/SvgIcon/index.vue";

export default defineComponent({
  name: "Item",
  components: {SvgIcon},
  props: {
    icon: {
      type: String as PropType<string>,
      default: ''
    },
    title: {
      type: String as PropType<string>,
      default: ''
    }
  }
})
</script>

<style lang="less" scoped>
.anticon-item {
  opacity: 1;
  transition: opacity .3s cubic-bezier(.645,.045,.355,1),width .3s cubic-bezier(.645,.045,.355,1);
}
</style>
