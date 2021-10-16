<template>
  <div>
    <div>hello world</div>
    <span>{{ count }}</span>
    <el-button @click="onClickBtn">点击</el-button>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { PROTO_MESSAGE } from "@/proto/message"
import { fetch } from "@/request/axios";
  fetch({url: "v1/helloworld", method: "get"} ).then(res => {
    console.log(res)
  })
  let count = ref(0)
  function onClickBtn(){
    count.value++
  }


</script>

<style scoped>

</style>
