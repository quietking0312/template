<template>
  <div class="login-wrap" @keydown.enter="login">
    <div class="login-con">
      <el-card class="box-card">
        <template #header>
          <span class="login--header">{{ $t('login.titleLogin') }}</span>
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
            <el-input v-model="form.password"
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
      </el-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, reactive, unref} from "vue";
import {useRouter} from "vue-router";
import wsCache, {cacheKey} from "@/cache";
import {loginApi} from "@/api/login";
import {PROTO_MESSAGE} from "@/proto/message";
import {ArrayBufferToStr, BytesToUint16, StrToArrayBuffer, BlobToArrayBuffer} from "@/utils/code";

interface FormModule {
  username: string
  password: string
}

const { push, addRoute, currentRoute } = useRouter()
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
        const msg = PROTO_MESSAGE.Login.create()
        msg.username = form.username
        msg.password = form.password
        const req = PROTO_MESSAGE.Login.encode(msg).finish()
        await StrToArrayBuffer(req, (s) => {
          loginApi(s).then(res => {
            BlobToArrayBuffer(res, (rs) => {
              const resp = PROTO_MESSAGE.Login.decode(rs)
              console.log(resp)
            })
            wsCache.set(cacheKey.userInfo, form)
            push({path: redirect.value || '/'})
          })
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
}
</style>
