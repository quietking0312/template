<template>
  <div style="padding-bottom: 10px">
    <el-button type="primary" @click="dialogVisible=true">
      <el-icon style="">
        <circle-plus style="vertical-align: middle" />
      </el-icon>
      <span style="vertical-align: middle">新增</span>
    </el-button>
  </div>
  <el-table :data="userTabelList" fit border>
    <el-table-column label="姓名">
    </el-table-column>
    <el-table-column label="用户名">
    </el-table-column>
    <el-table-column label="邮箱">
    </el-table-column>
    <el-table-column label="创建时间">
    </el-table-column>
    <el-table-column label="状态">
    </el-table-column>
    <el-table-column label="操作">
      <el-button>编辑</el-button>
      <el-button>删除</el-button>
    </el-table-column>
  </el-table>

  <div class="pagination_wrap">
    <el-pagination
        v-model="pageLimit.page"
        style="text-align: right"
        :page-size="pageLimit.limit"
        :total="tableTotal"
        :page-sizes="[10, 20, 30, 40, 50, 1000]"
        layout="total, prev, pager, next"></el-pagination>
  </div>


  <el-dialog v-model="dialogVisible" :title="dialogTitleMap[dialogTitleKey]">
    <el-form ref="formRef" :model="dialogForm" label-width="120px">
      <el-form-item label="姓名">
        <el-input v-model="dialogForm.name"></el-input>
      </el-form-item>
      <el-form-item label="用户名">
        <el-input v-model="dialogForm.username"></el-input>
      </el-form-item>
      <el-form-item label="email">
        <el-input v-model="dialogForm.email"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="dialogVisible=false">Cancel</el-button>
      <el-button type="primary" @click="dialogVisible=false">confirm</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import {reactive, ref} from "vue";
import {CirclePlus} from "@element-plus/icons-vue";
import {getUserListApi} from "@/api/permission";
const userTabelList = ref([])
const dialogVisible = ref(false)
const dialogTitleMap = {
  "create": "create",
  "update": "update"
}
const dialogTitleKey = "create"

const dialogForm = reactive({
  name: "",
  username: "",
  email: ""
})

const pageLimit = reactive({
  page: 1,
  limit: 20
})

const tableTotal = ref(100)

function getUserList() {
  getUserListApi({}).then(res => {
    console.log(res)
  })
}
getUserList()


</script>

<style lang="less" scoped>
.pagination_wrap {
  padding: 10px;
  margin-top: 15px;
  background: #fff;
}
</style>
