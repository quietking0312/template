<template>
  <div>
    <div>hello world</div>
    <svg-icon iconClass="dashboard" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { PROTO_MESSAGE } from "@/proto/message"

export default defineComponent({
  name: "Dashboard",
  setup() {
    const  a = PROTO_MESSAGE.Test.create()
    a.data = [{time: 1}]
    console.log(a.toJSON())
    return {}
  }
})
</script>

<style scoped>

</style>
