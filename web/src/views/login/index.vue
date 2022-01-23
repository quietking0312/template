<template>
  <div class="login-wrap" @keydown.enter="login">
    <div class="login-con">
      <el-card class="box-card">
        <template #header>
          <span class="login--header">{{ $t('login.titleLogin') }}</span>
          <LangSelect v-if="showVersion" class="lang--header" />
        </template>
        <el-form ref="loginForm" :model="form" class="login-form">
          <el-form-item prop="username">
            <el-input v-model="form.username" :placeholder="$t('login.placeholderUsername')" class="form--input">
              <template #prefix>
                <span class="svg-container">
                  <svg-icon icon-class="user" />
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input v-model="form.password" type="text"
                      show-password :minlength="3" :maxlength="18"
                      :placeholder="$t('login.placeholderPassword')" class="form--input">
              <template #prefix>
                <span class="svg-container">
                  <svg-icon icon-class="password" />
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button :loading="loading" type="primary" class="login--button" @click="login">{{ $t('login.btnLogin') }}</el-button>
          </el-form-item>
        </el-form>
        <div class="login-tips">
          <span>{{ $t('login.tipsVersion') }}： {{ version }}</span>
          <span style="float: right">
            <span v-if="registerOk">
              <a href="javascript:void(0)" @click="HandleRegister">注册账号</a> |
            </span>
            <a href="javascript:void(0)" @click="HandleResetPass">修改密码</a>
          </span>
        </div>
      </el-card>
    </div>
    <el-dialog :title="dialogTitleMap[dialogTitleKey]" v-model="dialogVisible" width="600px">
      <el-form ref="dialogForm" v-model="dialogFormData" label-width="80px">
        <el-form-item v-if="dialogTitleKey === 'register'" label="姓名">
          <el-input v-model="dialogFormData.name"></el-input>
        </el-form-item>
        <el-form-item v-if="dialogTitleKey === 'register'" label="email">
          <el-input v-model="dialogFormData.email"></el-input>
        </el-form-item>
        <el-form-item label="账号">
          <el-input v-model.trim="dialogFormData.username"></el-input>
        </el-form-item>
        <el-form-item v-if="dialogTitleKey==='resetPass'" label="旧密码">
          <el-input v-model.trim="dialogFormData.oldPassword" show-password :minlength="3" :maxlength="18" ></el-input>
        </el-form-item>
        <el-form-item :label="dialogTitleKey === 'resetPass'? '新密码': '密码'">
          <el-input v-model.trim="dialogFormData.password" show-password :minlength="3" :maxlength="18" ></el-input>
        </el-form-item>
        <el-form-item label="确认密码">
          <el-input v-model.trim="dialogFormData.newPassword" show-password :minlength="3" :maxlength="18" ></el-input>
        </el-form-item>

      </el-form>
      <template #footer>
        <el-button @click="dialogVisible=false">取消</el-button>
        <el-button type="primary" @click="dialogTitleKey === 'resetPass'?ResetPass():register()">确认</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import {ref, watch, reactive, unref, computed} from "vue";
import {useRouter} from "vue-router";
import wsCache, {cacheKey} from "@/cache";
import {appInfoApi, loginApi, registerApi, resetPassApi} from "@/api/login";
import LangSelect from "@/components/LangSelect/index.vue";
import {appStore} from "@/store/modules/app";
import {respType} from "@/request/request";
import config from "@/request/config";
import {Message} from "@/components/Message";
const registerOk = ref<boolean>(false)
// app版本
let version = ref<string>("")
appInfoApi().then(res => {
  const {code, data} = res as any
  if (code == config.result_code) {
    registerOk.value = data?.register
    version.value = data.version
    wsCache.set(cacheKey.conf, data)
  }
})

const showVersion = computed(() => appStore.showLanguage)

// d登录
interface FormModule {
  username: string
  password: string
}

const { push, currentRoute } = useRouter()
const loginForm = ref<HTMLElement | null>(null)
const loading = ref<boolean>(false)
const redirect = ref<string>('')
watch(() => {return currentRoute.value},
    (route) =>{redirect.value=(route.query && route.query.redirect as string)},
    {immediate: true})
const form = reactive<FormModule>({
  username: '',
  password: ''
})
async function login(): Promise<void> {
  const formWrap = unref(loginForm) as any
  if (!formWrap) return
  loading.value = true
  try {
    formWrap.validate(async (valid: boolean) => {
      if (valid) {
        loginApi(form).then(res => {
          if (res) {
            wsCache.set(cacheKey.userInfo, (res as unknown as respType).data.token)
            push({path: redirect.value || '/'})
          }
        })

      } else {
        console.log("error submit!!")
        return false
      }
    })
  }catch (err) {
    console.log(err)
  } finally {
    loading.value = false
  }
}

// ======= 修改密码 =====
const dialogTitleMap = reactive({
  resetPass: "重置密码",
  register: "注册"
})
const dialogTitleKey = ref("resetPass")
const dialogVisible = ref<boolean>(false)
const dialogForm = ref<HTMLElement | null>(null)
const dialogFormData = reactive({
  name: "",
  email: "",
  username: "",
  oldPassword: "",
  password: "",
  newPassword: ""
})

function resetDialogForm() {
  dialogFormData.name = ""
  dialogFormData.email = ""
  dialogFormData.username = ""
  dialogFormData.oldPassword = ""
  dialogFormData.password = ""
  dialogFormData.newPassword = ""
}

function HandleResetPass() {
  dialogTitleKey.value = "resetPass"
  resetDialogForm()
  dialogVisible.value = true
}

function ResetPass() {
  const formWrap = unref(dialogForm) as any
  if (!formWrap) return
  try {
    let data = {
      username: dialogFormData.username,
      oldPassword: dialogFormData.oldPassword,
      password: dialogFormData.password,
    }
    resetPassApi(data).then(res => {
      const {code} = res as any
      if (code === 0) {
        Message.success("修改成功")
        dialogVisible.value = false
      }
    })
  } catch (err) {
    console.log(err)
  }
}


function HandleRegister() {
  dialogTitleKey.value = "register"
  resetDialogForm()
  dialogVisible.value = true
}

function register() {
  const formWrap = unref(dialogForm) as any
  if (!formWrap) return
  try {
    let data = {
      name: dialogFormData.name,
      email: dialogFormData.email,
      username: dialogFormData.username,
      password: dialogFormData.password,
    }
    registerApi(data).then(res => {
      const {code, data} = res as any
      if (code === 0) {
        registerOk.value = data?.register
        Message.success("注册成功")
        dialogVisible.value = false
      }
    })
  } catch (err) {
    console.log(err)
  }
}

</script>

<style lang="less" scoped>
.login-wrap {
  width: 100%;
  height: 100%;
  //background-image: url('~@/assets/img/login-bg.jpg');
  background-size: cover;
  background-position: center;
  position: relative;
  .box-card {
    width: 400px;
    .login--header {
      font-size: 24px;
      font-weight: 600;
    }
    @{deep}(.lang--header) {
      right: 35px;
      position: absolute;
    }
    .svg-container {
      color: #889aa4;
      vertical-align: middle;
      width: 30px;
      display: inline-block;
    }
    .form--input {
      width: 100%;
      @{deep}(.el-input__inner) {
        padding-left: 40px;
      }
    }
    .login--button {
      width: 100%;
    }
  }
  .login-con {
    position: absolute;
    right: 160px;
    top: 50%;
    transform: translateY(-60%);
  }
  .login-tips {
    font-size: 12px;
    line-height: 15px;
  }
}
</style>
