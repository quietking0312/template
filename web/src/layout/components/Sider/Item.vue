<template>
  <i v-if="icon && icon.includes('el-icon')" :class="[icon, 'sub-el-icon', 'anticon']" />
  <Icon v-else-if="icon" :icon="icon" class="anticon" />
  <slot name="title">
    <span class="anticon-item">{{ title }}</span>
  </slot>
</template>

<script setup lang="ts">
import {PropType} from "vue";
import Icon from "@/components/Icon/Icon.vue";

defineProps({
  icon: {
    type: String as PropType<string>,
    default: ''
  },
  title: {
    type: String as PropType<string>,
    default: ''
  }
})
</script>

<style lang="less" scoped>
.anticon-item {
  opacity: 1;
  transition: opacity .3s cubic-bezier(.645,.045,.355,1),width .3s cubic-bezier(.645,.045,.355,1);
}
</style>
