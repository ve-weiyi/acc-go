<script lang="ts" setup>
import {reactive, ref, watch} from "vue"
import {createTableDataApi, deleteTableDataApi, updateTableDataApi, getTableDataApi} from "@/api/table"
import {type FormInstance, type FormRules, ElMessage, ElMessageBox} from "element-plus"
import {Search, Refresh, CirclePlus, Delete, Download, RefreshRight} from "@element-plus/icons-vue"
import {usePagination} from "@/hooks/usePagination"
import {parseInt} from "lodash-es";

const loading = ref<boolean>(false)
const {paginationData, handleCurrentChange, handleSizeChange} = usePagination()

//#region 增
const dialogVisible = ref<boolean>(false)
const formRef = ref<FormInstance | null>(null)
const formData = reactive({
  username: "",
  password: ""
})

const handleCreate = () => {
  formRef.value?.validate((valid: boolean) => {
    if (valid) {
      if (currentUpdateId.value === undefined) {
        createTableDataApi({}).then(() => {
          ElMessage.success("新增成功")
          dialogVisible.value = false
          getTableData()
        })
      } else {
        updateTableDataApi({}).then(() => {
          ElMessage.success("修改成功")
          dialogVisible.value = false
          getTableData()
        })
      }
    } else {
      return false
    }
  })
}
const resetForm = () => {
  currentUpdateId.value = undefined
  formData.username = ""
  formData.password = ""
}
//#endregion

//#region 删
const handleDelete = (row: any) => {
  console.log(row)
  ElMessageBox.confirm(`正在删除用户：${row.username}，确认删除？`, "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(() => {
    deleteTableDataApi({
        id: row.details.id
      }
    ).then(() => {
      ElMessage.success("删除成功")
      getTableData()
    })
  })
}
//#endregion

//#region 改
const currentUpdateId = ref<undefined | string>(undefined)
const handleUpdate = (row: any) => {
  currentUpdateId.value = row.details.id
  formData.username = row.details.username
  formData.password = row.details.password
  dialogVisible.value = true
}
//#endregion

//#region 查
const tableData = ref<any[]>([])
const searchFormRef = ref<FormInstance | null>(null)
const searchData = reactive({
  keywords: "",
  startTime: "",
  endTime: "",
})
const getTableData = () => {
  loading.value = true
  getTableDataApi(
    paginationData.currentPage - 1,
    {
      keywords: searchData.keywords || undefined,
      startTime: searchData.startTime || undefined,
      endTime: searchData.endTime || undefined,
    })
    .then((res: any) => {
      paginationData.total = res.data.total
      tableData.value = res.data.datas
      console.log(tableData.value)
    })
    .catch(() => {
      tableData.value = []
    })
    .finally(() => {
      loading.value = false
    })
}
const handleSearch = () => {
  if (paginationData.currentPage === 1) {
    getTableData()
  }
  paginationData.currentPage = 1
}
const resetSearch = () => {
  searchFormRef.value?.resetFields()
  searchData.keywords = ""
  if (paginationData.currentPage === 1) {
    getTableData()
  }
  paginationData.currentPage = 1
}
const handleRefresh = () => {
  paginationData.currentPage = 1
  getTableData()
}


const disabledDate = (time: Date) => {
  return time.getTime() < Number(searchData.startTime)
}
const options = [
  {
    value: 'Option1',
    label: 'Option1',
  },
  {
    value: 'Option2',
    label: 'Option2',
  },
  {
    value: 'Option3',
    label: 'Option3',
  },
]

const value = ref('')
//#endregion

/** 监听分页参数的变化 */
watch([() => paginationData.currentPage, () => paginationData.pageSize], getTableData, {immediate: true})
</script>

<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never" class="search-wrapper">
      <el-form ref="searchFormRef" :inline="true" :model="searchData">
        <el-form-item prop="title" label="标题">
          <el-input v-model="searchData.keywords" placeholder="请输入"/>
        </el-form-item>

        <el-form-item prop="status" label="状态">
          <el-select v-model="value" placeholder="Select">
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>


        <el-form-item prop="date" label="开始日期">
          <el-date-picker
            v-model="searchData.startTime"
            type="date"
            placeholder="选择开始日期"
          >
          </el-date-picker>
        </el-form-item>

        <el-form-item prop="date" label="结束日期">
          <el-date-picker
            v-model="searchData.endTime"
            type="date"
            placeholder="选择结束日期"
            :disabled-date="disabledDate"
          >
          </el-date-picker>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" :icon="Search" @click="handleSearch">查询</el-button>
          <el-button :icon="Refresh" @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>


    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <div>
          <el-button type="primary" :icon="CirclePlus" @click="dialogVisible = true">新增用户</el-button>
          <el-button type="danger" :icon="Delete">批量删除</el-button>
        </div>
        <div>
          <el-tooltip content="下载">
            <el-button type="primary" :icon="Download" circle/>
          </el-tooltip>
          <el-tooltip content="刷新表格">
            <el-button type="primary" :icon="RefreshRight" circle @click="handleRefresh"/>
          </el-tooltip>
        </div>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData">
          <el-table-column type="selection" width="50" align="center"/>
          <el-table-column prop="username" label="用户名" align="center"/>
          <el-table-column prop="details.title" label="标题" align="center"/>
          <el-table-column prop="details.content" label="内容" align="center"/>
          <el-table-column prop="details.done" label="完成情况" align="center">
            <template #default="scope">
              <el-tag v-if="scope.row.details.done==0" type="danger" effect="plain">未完成</el-tag>
              <el-tag v-else-if="scope.row.details.done==1" type="warning" effect="plain">处理中</el-tag>
              <el-tag v-else-if="scope.row.details.done==2" type="success" effect="plain">完成</el-tag>
              <!--              <el-tag v-else type="primary" effect="plain">其他</el-tag> -->
            </template>
          </el-table-column>
          <el-table-column prop="details.start_time" label="开始时间" align="center"/>
          <el-table-column prop="tags" label="标签" align="center">
            <template #default="scope">
              <el-tag
                v-for="tag in scope.row.tags"
                :key="tag.name"
                class="mx-1"
                type="success"
              >
                {{ tag.name }}
              </el-tag>
            </template>
          </el-table-column>

          <el-table-column fixed="right" label="操作" width="150" align="center">
            <template #default="scope">
              <el-button type="primary" text bg size="small" @click="handleUpdate(scope.row)">修改</el-button>
              <el-button type="danger" text bg size="small" @click="handleDelete(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div class="pager-wrapper">
        <el-pagination
          background
          :page-sizes="paginationData.pageSizes"
          :total="paginationData.total"
          :page-size="paginationData.pageSize"
          :currentPage="paginationData.currentPage"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
    <!-- 新增/修改 -->
    <el-dialog
      v-model="dialogVisible"
      :title="currentUpdateId === undefined ? '新增用户' : '修改用户'"
      @close="resetForm"
      width="30%"
    >
      <el-form :model="formData" ref="vForm" :rules="rules" label-position="left" label-width="80px"
               size="default" @submit.prevent>
        <div class="card-container">
          <el-card style="{width: 100% !important}" shadow="never">
            <template #header>
              <div class="clear-fix">
                <span>card</span>
                <i class="float-right el-icon-arrow-down"></i>
              </div>
            </template>
            <el-form-item label="input" prop="input35927" class="label-center-align">
              <el-input v-model="formData.input35927" type="text" clearable></el-input>
            </el-form-item>
            <el-form-item label="input" prop="input18856" class="label-center-align">
              <el-input v-model="formData.input18856" type="text" clearable></el-input>
            </el-form-item>
            <el-form-item label="input" prop="input35791" class="label-center-align">
              <el-input v-model="formData.input35791" type="text" clearable></el-input>
            </el-form-item>
            <el-form-item label="input" prop="input107657" class="label-center-align">
              <el-input v-model="formData.input107657" type="text" clearable></el-input>
            </el-form-item>
            <el-form-item label="input" prop="input90110" class="label-center-align">
              <el-input v-model="formData.input90110" type="text" clearable></el-input>
            </el-form-item>
            <div class="static-content-item">
              <el-divider direction="horizontal"></el-divider>
            </div>
          </el-card>
        </div>
      </el-form>

    </el-dialog>


  </div>
</template>

<style lang="scss" scoped>
.search-wrapper {
  margin-bottom: 20px;

  :deep(.el-card__body) {
    padding-bottom: 2px;
  }

}

.toolbar-wrapper {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

.table-wrapper {
  margin-bottom: 20px;
}

.pager-wrapper {
  display: flex;
  justify-content: flex-end;
}
</style>
