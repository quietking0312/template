<template>
  <div style="padding-bottom: 10px">
    <el-button v-permission="[101002]" type="primary" @click="handleAddRole()">
      新增
    </el-button>
  </div>
  <el-table :data="roleTableList" fit border>
    <el-table-column label="角色名" prop="role_name">
    </el-table-column>
    <el-table-column label="操作">
      <template #default="{row}">
        <el-button v-permission="[101003]" type="primary" @click="handleUpdateRole(row)" size="small">编辑</el-button>
        <el-button v-permission="[102001]" type="primary" @click="handleSetRolePermission(row)" size="small">授权</el-button>
<!--        <el-button v-permission=[101004]>删除</el-button>-->
      </template>
    </el-table-column>
  </el-table>
  <el-dialog :title="dialogTitleMap[dialogTitleKey]" v-model="dialogVisible">
    <el-form ref="formRef" :model="dialogForm" :rules="formRules" label-width="100px">
      <el-form-item label="角色名" prop="role_name">
        <el-input :disabled="dialogTitleKey === 'setPid'" v-model.trim="dialogForm.role_name" maxlength="5"></el-input>
      </el-form-item>
      <el-form-item v-if="dialogTitleKey === 'setPid'">
        <el-tree
            ref="treeRef"
            :default-checked-keys="defaultCheckedKeys"
            :data="permissionTreeData"
            :check-strictly="true"
            :show-checkbox="CheckPermission([102003])"
            check-on-click-node
            :props="TreeProp"
            node-key="permission_id"
            default-expand-all
        >
        </el-tree>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="dialogVisible=false">{{ $t("common.cancel") }}</el-button>
      <el-button v-if="dialogTitleKey === 'setPid'? CheckPermission([102003]): true" type="primary" @click="handleConfirm">{{ $t("common.confirm") }}</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import {nextTick, reactive, ref, unref} from "vue";
import {
  getPermissionListApi,
  getRoleListApi,
  postRoleApi, postRolePermissionApi,
  updateRoleApi
} from "@/api/permission";
import config from "@/request/config";
import {Message} from "@/components/Message";
import {CheckPermission, PermissionListToTree} from "@/utils/permission";
import {ElTree} from "element-plus/es";
import type { ElForm} from "element-plus/es";

type FormInstance = InstanceType<typeof ElForm>
const formRef = ref<FormInstance | null>(null)
const roleTableList = ref([])
const dialogVisible = ref(false)

const dialogTitleMap = reactive({
  create: "新增角色",
  update: "修改角色",
  setPid: "授权"
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

const formRules = reactive({
  role_name: [{required: true,  trigger: 'blur'}, {min: 2, max:5, trigger: 'blur'}]
})

const pageLimit = reactive({
  page: 1,
  limit: 20
})

const tableTotal = ref(1)

// 树组件
const permissionTreeData = ref([])
const defaultCheckedKeys = ref([])
const TreeProp = {
  label: "title",
  children: 'children',
  disabled: 'false'
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
    AddRole()
  } else if (dialogTitleKey.value === 'update') {
    UpdateRole()
  } else if (dialogTitleKey.value === 'setPid' && CheckPermission([102003])) {
    SetRolePermission()
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


// 授权按钮响应事件
function handleSetRolePermission(row: any) {
  dialogTitleKey.value = "setPid"
  Object.keys(dialogForm).map(key => {
    dialogForm[key] = row[key]
  })
  // defaultCheckedKeys.value = row?.permission_ids? row.permission_ids: []
  dialogVisible.value = true
  nextTick(() => {
    formRef.value?.clearValidate()
    treeRef.value?.setCheckedKeys(row?.permission_ids? row.permission_ids: [], false)
  })
}

if (CheckPermission([102001])){
  GetPermissionList()
}
function GetPermissionList() {
  getPermissionListApi().then(res => {
    const {code, data} = (res as any)
    if (code == config.result_code) {
      permissionTreeData.value = PermissionListToTree(data.data)
    }
  })
}

function SetRolePermission() {
  const treeRefWarp = unref(treeRef)
  if (!treeRefWarp) return
  postRolePermissionApi({rid: dialogForm.rid, p_ids: treeRefWarp?.getCheckedKeys(false)}).then(res => {
    const {code} = res as any
    if (code === config.result_code) {
      getRoleList()
      Message.success("操作成功")
    }
  })
}
</script>

<style scoped>

</style>
