<template>
  <div v-html="noticeData"></div>
</template>

<script setup lang="ts">
import { getNotice } from "@/api/notice";
import {ref} from "vue";

const noticeData = ref()
getNotice().then(res => {
  noticeData.value = res?.data
})
</script>

<style scoped>

</style>
