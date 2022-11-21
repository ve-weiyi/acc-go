<template>
  <el-config-provider :locale="locale" namespace="el" size="small">
    <router-view :key="key" />
  </el-config-provider>
</template>
<script setup>
import { setToken } from '@/utils/auth'
// import { useAppStore } from '@/store/app'

import zhCn from 'element-plus/lib/locale/lang/zh-cn'
import {computed} from "vue";
let locale = $ref(zhCn) //中文

// import en from 'element-plus/lib/locale/lang/en'
// let locale = $ref(en) //english

// const appStore = useAppStore()
// const settings = computed(() => {
//   return appStore.settings
// })

const key = computed(() =>
  "/views/Home.vue"
)
// onBeforeMount(() => {
//   //set tmp token when setting isNeedLogin false
//   if (!settings.value.isNeedLogin) setToken(settings.value.tmpToken)
// })
</script>
