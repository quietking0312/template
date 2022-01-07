<template>
  <div style="padding-bottom: 10px">
    <el-button type="primary">
      新增
    </el-button>
  </div>
  <el-table :data="tableList" fit border>
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
