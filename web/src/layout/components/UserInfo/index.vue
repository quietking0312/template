<template>
  <el-dropdown class="avatar-container" trigger="click">
    <div id="user-container">
      <div class="avatar-wrapper">
        <img :src="getUserImg()" class="user-avatar">
        <span class="name-item">{{ uName }}</span>
      </div>
    </div>
    <template #dropdown>
      <el-dropdown-item key="1">
        <span style="display: block;" @click="toHome">{{ generateTitle('dashboard') }}</span>
      </el-dropdown-item>
      <el-dropdown-item key="2">
        <span style="display: block;" @click="loginOut">{{ $t('userInfo.btnLogout') }}</span>
      </el-dropdown-item>
    </template>
  </el-dropdown>
</template>

<script lang="ts">
import {computed, defineComponent} from "vue";
import {useRouter} from "vue-router";
import wsCache, {cacheKey} from "@/cache";
import {resetRouter} from "@/router";
import {useTagsViewStore} from "@/store/modules/tagsView";
import { generateTitle } from "@/utils/i18n";
import {usePermissionStore} from "@/store/modules/permission";
import { useUserInfoStore } from "@/store/modules/userInfo";

export default defineComponent({
  name: "UserInfo",
  setup() {
    const tagsViewStore = useTagsViewStore()
    const permissionStore = usePermissionStore()
    const userInfoStore = useUserInfoStore()
    const { replace, push } = useRouter()
    async function loginOut(): Promise<void> {
      wsCache.delete(cacheKey.userInfo)
      await resetRouter()
      await permissionStore.SetIsAddRouters(false)
      await tagsViewStore.delAllViews()
      await replace('/login')
    }
    function toHome() {
      push('/')
    }
    function getUserImg() {
      return new URL('../../assets/avatar.png', import.meta.url).href
    }
    const uName = computed(() => userInfoStore.name)
    return {
      loginOut,
      toHome,
      getUserImg,
      generateTitle,
      uName,
    }
  }
})
</script>

<style lang="less" scoped>
.avatar-container {
  margin-right: 30px;
  padding: 0 10px;
  .avatar-wrapper {
    display: flex;
    align-items: center;
    height: 100%;
    cursor: pointer;
    .user-avatar {
      width: 30px;
      height: 30px;
      border-radius: 10px;
    }
    .name-item {
      font-size: 14px;
      font-weight: 600;
      display: inline-block;
      margin-left: 5px;
    }
  }
}
</style>
