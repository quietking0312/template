<script setup lang="ts">
import {computed, nextTick, ref, unref, watch} from "vue";
import Iconify from '@purge-icons/generated'

const props = defineProps({
  icon: {
    type: String,
    default: () => ''
  },
  size: {
    type: Number,
    default: () => 16
  },
  color: {
    type: String,
    default: () => ''
  }
})

const elRef = ref<HTMLElement | null >(null)

const isLocal = computed(() => props.icon.startsWith('svg-icon:'))
const symbolId = computed(() => {
  return unref(isLocal) ? `#icon-${props.icon.split('svg-icon:')[1]}` : props.icon
})
const getIconifyStyle = computed(() => {
  const { color, size } = props
  return {
    fontSize: `${size}px`,
    color
  }
})

const updateIcon = async (icon: string) => {

  if (unref(isLocal)) return

  const el = unref(elRef)
  if (!el) return

  await nextTick()

  if (!icon) return
  const svg = Iconify.renderSVG(icon, {})
  if (svg) {
    el.textContent = ''
    el.appendChild(svg)
  } else {
    const span = document.createElement('span')
    span.className = 'iconify'
    span.dataset.icon = icon
    el.textContent = ''
    el.appendChild(span)
  }
}

watch(
    () => props.icon,
    (icon: string) => {
      updateIcon(icon)
    }
)

</script>


<template>
  <el-icon class="el-icon"  :size="size" :color="color">
    <svg v-if="isLocal" aria-hidden="true" >
      <use :xlink:href="symbolId" />
    </svg>
    <span v-else ref="elRef" :class="$attrs.class" :style="getIconifyStyle">
      <span class="iconify" :data-icon="symbolId"></span>
    </span>
  </el-icon>

</template>


<style scoped>
.svg-icon {
  width: 1em;
  height: 1em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
</style>
