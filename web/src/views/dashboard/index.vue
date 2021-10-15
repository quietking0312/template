<template>
  <div>
    <div>hello world</div>
    <svg-icon iconClass="dashboard" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { PROTO_MESSAGE } from "@/proto/message"
import { fetch } from "@/request/axios";

export default defineComponent({
  name: "Dashboard",
  setup() {
    const  a = PROTO_MESSAGE.Test.create()
    a.data = [{time: 1}]
    console.log(a.toJSON())
    fetch({url: "v1/helloworld", method: "get"} ).then(res => {
      console.log(res)
    })
    return {}
  }
})
</script>

<style scoped>

</style>
