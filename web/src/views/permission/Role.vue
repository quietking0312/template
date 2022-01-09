<template>
  <div style="padding-bottom: 10px">
    <el-button type="primary">
      新增
    </el-button>
  </div>
  <el-table :data="roleTableList" fit border>
    <el-table-column label="角色名">
    </el-table-column>
  </el-table>
</template>

<script setup lang="ts">
import {reactive, ref, unref} from "vue";
import {getPermissionListApi, getRoleListApi, postRoleApi, updateRoleApi} from "@/api/permission";
import config from "@/request/config";
import {Message} from "@/components/Message";
import {PermissionListToTree} from "@/utils/permission";
import {ElTree} from "element-plus/es";

const formRef = ref<HTMLElement | null>(null)
const roleTableList = ref([])
const dialogVisible = ref(false)

const dialogTitleMap = reactive({
  create: "新增角色",
  update: "修改角色"
})

let dialogTitleKey = ref("create")

const dialogForm = reactive({
  rid: 0,
  role_name: ""
})

function resetDialogForm() {
  dialogForm.rid = 0
  dialogForm.role_name = ""
}

const pageLimit = reactive({
  page: 1,
  limit: 20
})

const tableTotal = ref(1)

// 树组件
const permissionTreeData = ref([])

const TreeProp = {
  label: function (data:any, node:any) {
    return data.title
  }
}
const treeRef = ref<InstanceType<typeof ElTree>>()

// 初始化
getRoleList()

function getRoleList() {
  getRoleListApi({page: pageLimit.page, limit: pageLimit.limit}).then(res => {
    const { code, data } = res as any
    if (code === 0) {
      tableTotal.value = data.total
      roleTableList.value =data.data
    }
  })
}

function handleConfirm() {
  if (dialogTitleKey.value === 'create') {

  } else if (dialogTitleKey.value === 'update') {

  }
  dialogVisible.value = false
}

// 新增按钮响应事件
function handleAddRole() {
  dialogTitleKey.value = "create"
  resetDialogForm()
  dialogVisible.value = true
}

// 新增角色
function AddRole() {
  const formWarp = unref(formRef) as any
  if (!formWarp) return
  try {
    formWarp.validate(async (valid: boolean) => {
      if (valid) {
        await postRoleApi(dialogForm).then(res => {
          const {code} = res as any
          if (code === config.result_code) {
            getRoleList()
            Message.success("操作成功")
          }
        })
      }
    })
  } catch (err) {
    console.log(err)
  }
}

// 编辑按钮响应事件
function handleUpdateRole(row: any) {
  dialogTitleKey.value = "update"
  Object.keys(dialogForm).map(key => {
    dialogForm[key] = row[key]
  })
  dialogVisible.value = true
}
// 修改角色信息
function UpdateRole() {
  const formWarp = unref(formRef) as any
  if (!formWarp) return
  try {
    formWarp.validate(async (valid: boolean) => {
      if (valid) {
        await updateRoleApi(dialogForm).then(res => {
          const {code} = res as any
          if (code === config.result_code) {
            getRoleList()
            Message.success("操作成功")
          }
        })
      }
    })
  }catch (err) {
    console.log(err)
  }
}

function GetPermissionList() {
  getPermissionListApi().then(res => {
    const {code, data} = (res as any)
    if (code == config.result_code) {
      permissionTreeData.value = PermissionListToTree(data.data)
    }
  })
}

</script>

<style scoped>

</style>
