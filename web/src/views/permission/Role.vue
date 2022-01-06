<template>
  <el-table :data="tableList">
    <el-table-column label="角色名">

    </el-table-column>
  </el-table>
</template>

<script setup lang="ts">
import {reactive} from "vue";

const tableList = reactive([])

</script>

<style scoped>

</style>
