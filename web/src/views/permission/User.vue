<template>
  <div style="padding-bottom: 10px">
    <el-button type="primary" @click="handleAddUser">
      <span style="vertical-align: middle">新增</span>
    </el-button>
  </div>
  <el-table :data="userTabelList" fit border>
    <el-table-column label="姓名" prop="name">
    </el-table-column>
    <el-table-column label="用户名" prop="username">
    </el-table-column>
    <el-table-column label="邮箱" prop="email">
    </el-table-column>
    <el-table-column label="创建时间">
      <template #default="{row}">
        {{ formatTime(row.create_time * 1000, "yyyy-MM-dd HH:mm:ss") }}
      </template>
    </el-table-column>
    <el-table-column label="状态">
      <template #default="{row}">
        {{ row.state }}
      </template>
    </el-table-column>
    <el-table-column label="操作">
      <template #default="{ row }">
        <el-button @click="handleUpdateUser(row)" type="primary" size="small">编辑</el-button>
        <el-button v-if="row.state===State.on" type="danger" size="small">删除</el-button>
      </template>
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
        <el-input  v-model="dialogForm.username"></el-input>
      </el-form-item>
      <el-form-item label="email">
        <el-input v-model="dialogForm.email"></el-input>
      </el-form-item>
      <el-form-item v-if="dialogTitleKey === 'update'" label="状态">
        <el-select v-model="dialogForm.state">
          <el-option v-for="item in Status" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
<!--        <el-input v-model="dialogForm.state"></el-input>-->
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="dialogVisible=false">Cancel</el-button>
      <el-button type="primary" @click="handleConfirm()">confirm</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import {reactive, ref, unref} from "vue";
import { formatTime } from "@/utils";
import {getUserListApi, postUserApi} from "@/api/permission";

const formRef = ref<HTMLElement | null>(null)
const userTabelList = ref([])
const dialogVisible = ref(false)


const dialogTitleMap = reactive({
  "create": "create",
  "update": "update"
})

let dialogTitleKey = ref("create")

enum State {
  on = 1, // 可用
  off = 2 // 不可用
}

const Status = [
  {value: State.on, label: "on"},
  {value: State.off, label: "off"}
]

const dialogForm = reactive({
  uid: 0,
  name: "",
  username: "",
  email: "",
  state: State.on
})

function resetDialogForm() {
  dialogForm.uid = 0
  dialogForm.name = ""
  dialogForm.username = ""
  dialogForm.email = ""
  dialogForm.state = State.on
}

const pageLimit = reactive({
  page: 1,
  limit: 20
})

const tableTotal = ref(100)

function getUserList() {
  getUserListApi({page: pageLimit.page, limit: pageLimit.limit}).then(res => {
    const { code, data } = res as any
    if (code === 0) {
      tableTotal.value = data.total
      userTabelList.value = data.data
    }
  })
}
getUserList()

// dialog确认按钮响应事件
function handleConfirm() {
  if (dialogTitleKey.value === "create") {
    AddUser()
  } else if (dialogTitleKey.value == "update") {
  }
  getUserList()
  dialogVisible.value = false
}

// 新增按钮响应事件
function handleAddUser() {
  dialogTitleKey.value = "create"
  resetDialogForm()
  dialogVisible.value = true
}

// 新增用户
function AddUser() {
  const formWarp = unref(formRef) as any
  if (!formWarp) return
  try {
    formWarp.validate(async (valid: boolean) => {
      if (valid) {
        postUserApi(dialogForm).then(res => {
          console.log(res)
        })
      }
    })
  } catch (err) {
    console.log(err)
  }
}

// 编辑按钮响应事件
function handleUpdateUser(row: any) {

  dialogTitleKey.value = "update"
  Object.keys(dialogForm).map(key => {
    dialogForm[key] = row[key]
  })
  dialogVisible.value =true
}

</script>

<style lang="less" scoped>
.pagination_wrap {
  padding: 10px;
  margin-top: 15px;
  background: #fff;
}
</style>
