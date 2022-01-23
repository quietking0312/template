<template>
  <div style="padding-bottom: 10px">
    <el-button type="primary" @click="handleAddUser">
      新增
    </el-button>
  </div>
  <el-table :data="userTableList" fit border>
    <el-table-column label="姓名" prop="name"  align="center">
    </el-table-column>
    <el-table-column label="用户名" prop="username"  align="center">
    </el-table-column>
    <el-table-column label="邮箱" prop="email"  align="center">
    </el-table-column>
    <el-table-column label="创建时间"  align="center" width="180">
      <template #default="{row}">
        {{ formatTime(row.create_time * 1000, "yyyy-MM-dd HH:mm:ss") }}
      </template>
    </el-table-column>
    <el-table-column label="状态" align="center" width="100">
      <template #default="{row}">
        {{ Array2Object(Status, 'value')[row.state].label }}
      </template>
    </el-table-column>
    <el-table-column label="操作" align="center">
      <template #default="{ row }">
        <el-button @click="handleUpdateUser(row)" type="primary" size="small">编辑</el-button>
        <el-popconfirm v-if="row.state===State.on"
                       confirm-button-text="ok"
                       cancel-button-text="cancel"
                       title="是否重置该用户密码？"
                       @confirm="InitUserPass(row.uid)"
        >
          <template #reference>
            <el-button type="danger" size="small">密码初始化</el-button>
          </template>
        </el-popconfirm>
<!--        <el-button @click="handleSetUserPermission(row)" type="primary" size="small">授权</el-button>-->
        <el-button v-if="row.state===State.off" type="danger" size="small">删除</el-button>
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


  <el-dialog v-model="dialogVisible" :title="dialogTitleMap[dialogTitleKey]" width="800px">
    <el-form ref="formRef" :model="dialogForm" label-width="120px" :rules="formRules">
      <el-form-item label="姓名" prop="name">
        <el-input v-model.trim="dialogForm.name" :disabled="dialogTitleKey==='setPid'"></el-input>
      </el-form-item>
      <el-form-item label="用户名" prop="username">
        <el-input  v-model.trim="dialogForm.username" :disabled="dialogTitleKey==='setPid'"></el-input>
      </el-form-item>
      <el-form-item label="email">
        <el-input v-model.trim="dialogForm.email" :disabled="dialogTitleKey==='setPid'"></el-input>
      </el-form-item>
      <el-form-item v-if="dialogTitleKey === 'update'" label="状态">
        <el-select v-model.number="dialogForm.state">
          <el-option v-for="item in Status" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item  v-if="dialogTitleKey === 'setPid'">
        <el-tree ref="treeRef"
            :data="permissionTreeData"
            :default-checked-keys="defaultCheckedKeys"
            :check-strictly="true"
            :show-checkbox="CheckPermission([102003])"
            check-on-click-node
            :props="TreeProp"
            default-expand-all
            node-key="permission_id"></el-tree>
      </el-form-item>
      <el-form-item v-if="roleList">
        <el-transfer :data="roleList"
                     v-model="dialogForm.rids"
                     filterable
                     :props="{key:'rid', label: 'role_name'}"
                     :titles="['角色列表', '已选择']"></el-transfer>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="dialogVisible=false">{{ $t("common.cancel") }}</el-button>
      <el-button v-if="dialogTitleKey === 'setPid'? CheckPermission([102002]): true"  type="primary" @click="handleConfirm()">{{ $t("common.confirm") }}</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import {reactive, ref, unref} from "vue";
import {Array2Object, formatTime} from "@/utils";
import {
  deleteUserPassApi,
  getPermissionListApi, getRoleAllApi,
  getUserListApi,
  postUserApi,
  postUserPermissionApi,
  updateUserApi
} from "@/api/permission";
import {CheckPermission, PermissionListToTree} from "@/utils/permission";
import {Message} from "@/components/Message";
import config from "@/request/config";
import {ElTree} from "element-plus";

const formRef = ref<HTMLElement | null>(null)
const userTableList = ref([])
const dialogVisible = ref(false)


const dialogTitleMap = reactive({
  create: "新增用户",
  update: "修改用户",
  setPid: "授权"
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
  state: State.on,
  rids: [] as any[]
})

function resetDialogForm() {
  dialogForm.uid = 0
  dialogForm.name = ""
  dialogForm.username = ""
  dialogForm.email = ""
  dialogForm.state = State.on
  dialogForm.rids = []
}

const formRules = reactive({
  name: [{required: true, trigger: 'blur'}, {min: 2, max:8, trigger: 'blur'}],
  username: [{required: true, trigger:'blur'}, {min: 3, max:18, trigger: 'blur'}],
})

const roleList = ref<any[]>([])
getRoleAllApi().then(res => {
  const {code, data} = res as any
  if (code === 0) {
    roleList.value = data.data
  }
})

const pageLimit = reactive({
  page: 1,
  limit: 20
})

const tableTotal = ref(0)

getUserList()

// 树组件
const permissionTreeData = ref([])
const defaultCheckedKeys = ref([])
const TreeProp = {
  label: "title",
  children: "children",
  disabled: "false"
}

const treeRef = ref<InstanceType<typeof ElTree>>()


function getUserList() {
  getUserListApi({page: pageLimit.page, limit: pageLimit.limit}).then(res => {
    const { code, data } = res as any
    if (code === 0) {
      tableTotal.value = data.total
      userTableList.value = data.data
    }
  })
}

// dialog确认按钮响应事件
function handleConfirm() {
  if (dialogTitleKey.value === "create") {
    AddUser()
  } else if (dialogTitleKey.value == "update") {
    UpdateUser()
  } else if (dialogTitleKey.value === 'setPid' && CheckPermission([102002])) {
    SetUserPermission()
  }
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
        await postUserApi(dialogForm).then(res => {
          const {code} = res as any
          if (code === config.result_code) {
            getUserList()
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
function handleUpdateUser(row: any) {
  dialogTitleKey.value = "update"
  Object.keys(dialogForm).map(key => {
    dialogForm[key] = row[key]
  })
  dialogVisible.value =true
}

// 修改用户信息
function UpdateUser() {
  const formWarp = unref(formRef) as any
  if (!formWarp) return
  try {
    formWarp.validate(async (valid: boolean) => {
      if (valid) {
        await updateUserApi(dialogForm).then(res => {
          const {code} = res as any
          if (code === config.result_code) {
            getUserList()
            Message.success("操作成功")
          }
        })
      }
    })
  }catch (err) {
    console.log(err)
  }
}


if (CheckPermission([102001])){
  GetPermissionList()
}
// 授权按钮响应事件
function handleSetUserPermission(row: any) {
  dialogTitleKey.value = "setPid"
  Object.keys(dialogForm).map(key => {
    dialogForm[key] = row[key]
  })
  defaultCheckedKeys.value = row?.permission_ids
  dialogVisible.value = true
}

function GetPermissionList() {
  getPermissionListApi().then(res => {
    const {code, data} = (res as any)
    if (code == config.result_code) {
      permissionTreeData.value = PermissionListToTree(data.data)
    }
  })
}

function SetUserPermission() {
  const treeRefWarp = unref(treeRef)
  if (!treeRefWarp) return
  postUserPermissionApi({uid: dialogForm.uid, p_ids: treeRefWarp?.getCheckedKeys(false)}).then(res => {
    const {code} = res as any
    if (code === config.result_code) {
      getUserList()
      Message.success("操作成功")
    }
  })
}

function InitUserPass(uid: string) {
  deleteUserPassApi({uid: uid}).then(res => {
    const {code} =res as any
    if (code === 0) {
      Message.success("重置成功")
    }
  })
}

</script>

<style lang="less" scoped>
.pagination_wrap {
  padding: 10px;
  margin-top: 15px;
  background: #fff;
}
</style>
