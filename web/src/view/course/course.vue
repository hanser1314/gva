<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div  v-auth="btnAuth.Courseoperate" class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        stripe =true
        border =true
        size = 'large'
        empty-text='No data'
        highlight-current-row
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="课程号" prop="cno" width="120" />
        <el-table-column align="left" label="课程名" prop="cname" width="120" />
        <el-table-column align="left" label="先行课编号" prop="cpno" width="120" />
        <el-table-column align="left" label="学分" prop="ccredit" width="120" />
        <el-table-column align="left" label="教工号" prop="ctno" width="120" />
        <el-table-column align="left" label="上课地点、时间" prop="cplace" width="240" />
        <el-table-column align="left" label="最大选课人数" prop="cmax" width="120" />
        <el-table-column align="left" label="已选人数" prop="cremain" width="120" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button v-auth="btnAuth.Courseoperate" type="primary" link icon="edit" class="table-button" @click="updateCourseFunc(scope.row)" >变更</el-button>
            <el-button v-auth="btnAuth.Courseoperate" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            <el-button color="#008000" plain :dark="isDark"  :disabled="isDisabled" @click="isSelect(scope.row)">选课</el-button>
            <el-button color="#FF0000" plain :dark="isDark"   @click="isDeselect(scope.row)">退选</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>

    </div>

<!-- <el-card>
  <div>
          {{ searchData }}
        </div>
</el-card> -->

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">
            <el-form-item label="课程号:"  prop="cno" >
              <el-input v-model="formData.cno" :clearable="true"  placeholder="请输入课程号" />
            </el-form-item>
            <el-form-item label="课程名:"  prop="cname" >
              <el-input v-model="formData.cname" :clearable="true"  placeholder="请输入课程名" />
            </el-form-item>
            <el-form-item label="先行课编号:"  prop="cpno" >
              <el-input v-model="formData.cpno" :clearable="true"  placeholder="请输入先行课编号" />
            </el-form-item>
            <el-form-item label="学分:"  prop="ccredit" >
              <el-input v-model.number="formData.ccredit" :clearable="true" placeholder="请输入学分" />
            </el-form-item>
            <el-form-item label="教工号:"  prop="ctno" >
              <el-input v-model="formData.ctno" :clearable="true"  placeholder="请输入教工号" />
            </el-form-item>
            <el-form-item label="上课地点、时间:"  prop="cplace" >
              <el-input v-model="formData.cplace" :clearable="true"  placeholder="请输入上课地点、时间" />
            </el-form-item>
            <el-form-item label="最大选课人数:"  prop="cmax" >
              <el-input v-model.number="formData.cmax" :clearable="true" placeholder="请输入最大选课人数" />
            </el-form-item>
            <el-form-item label="已选人数:"  prop="cremain" >
              <el-input v-model.number="formData.cremain" :clearable="true" placeholder="请输入已选人数" />
            </el-form-item>
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
                <el-descriptions-item label="课程号">
                        {{ formData.cno }}
                </el-descriptions-item>
                <el-descriptions-item label="课程名">
                        {{ formData.cname }}
                </el-descriptions-item>
                <el-descriptions-item label="先行课编号">
                        {{ formData.cpno }}
                </el-descriptions-item>
                <el-descriptions-item label="学分">
                        {{ formData.ccredit }}
                </el-descriptions-item>
                <el-descriptions-item label="教工号">
                        {{ formData.ctno }}
                </el-descriptions-item>
                <el-descriptions-item label="上课地点、时间">
                        {{ formData.cplace }}
                </el-descriptions-item>
                <el-descriptions-item label="最大选课人数">
                        {{ formData.cmax }}
                </el-descriptions-item>
                <el-descriptions-item label="已选人数">
                        {{ formData.cremain }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createCourse,
  deleteCourse,
  deleteCourseByIds,
  updateCourse,
  findCourse,
  getCourseList
} from '@/api/course'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

import { useBtnAuth } from '@/utils/btnAuth'
const btnAuth = useBtnAuth()


defineOptions({
    name: 'Course'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        cno: '',
        cname: '',
        cpno: '',
        ccredit: 0,
        ctno: '',
        cplace: '',
        cmax: 0,
        cremain: 0
        })


// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getCourseList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteCourseFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteCourseByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateCourseFunc = async(row) => {
    const res = await findCourse({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.recourse
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteCourseFunc = async(row) => {
    const res = await deleteCourse({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findCourse({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recourse
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          cno: '',
          cname: '',
          cpno: '',
          ccredit: 0,
          ctno: '',
          cplace: '',
          cmax: 0,
          cremain: 0,
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        cno: '',
        cname: '',
        cpno: '',
        ccredit: 0,
        ctno: '',
        cplace: '',
        cmax: 0,
        cremain: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createCourse(formData.value)
                  break
                case 'update':
                  res = await updateCourse(formData.value)
                  break
                default:
                  res = await createCourse(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

const SctformData = ref({
        sno: '',
        cno: '',
        cname: '',
        ccredit: 0,
        tno: '',
        cplace: '',
        grade: 0
})

// const UpdateSctformData = ref({
//         sno: '',
//         cno: '',
//         cname: '',
//         ccredit: 0,
//         tno: '',
//         cplace: '',
//         grade: 0
// })

// 我的函数
import { useUserStore } from '@/pinia/modules/user'
import { createSct, getSctListBySno, getSctListBySnoAndCno, deleteSct } from '@/api/sct'
// import { provide } from 'vue'
const userStore = useUserStore()
const stuSno = ref('')
stuSno.value = userStore.userInfo.userName

const searchData = ref([])

const isSelect = (row) => {
  ElMessageBox.confirm(
    '确认要选该课程吗?',
    'Warning',
    {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'warning',
    }
  )
    .then(() => {
      optCourse(row)
    })
    // .catch(() => {
    //   ElMessage({
    //     type: 'info',
    //     message: '取消选课',
    //   })
    // })
}

const isDeselect = (row) => {
  ElMessageBox.confirm(
    '确认要退选该课程吗?',
    'Warning',
    {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'warning',
    }
  )
    .then(() => {
      Deselect(row)
    })
    // .catch(() => {
    //   ElMessage({
    //     type: 'info',
    //     message: '取消选课',
    //   })
    // })
}



const Deselect = async(row) => {
  const table = await getSctListBySnoAndCno({ page: page.value, pageSize: pageSize.value, sno: stuSno.value, cno: row.cno })
  if(table.data.list.length === 0) {
      ElMessage({
    showClose: true,
    message: '你还没选择该课程！',
    type: 'error',
  })
    return
    }
    if (table.code === 0 && row.cremain > 0) {
      searchData.value = table.data.list
      const json = JSON.stringify(table.data.list)
      const arr = JSON.parse(json)
      var id = arr[0].ID
      deleteSct({ ID: id })

  const res = await findCourse({ ID: row.ID })
    formData.value = res.data.recourse
          formData.value.cremain -= 1
          updateCourse(formData.value)
          getTableData()
      ElMessage({
        showClose: true,
        message: '退选成功！',
        type: 'success',
      })
    }

    else {
      ElMessage({
    showClose: true,
    message: '当前状态无法退选！',
    type: 'error',
  })
  return
    }
}

getTableData()

const optCourse = async(row) => {
  // provide('courseID', row.ID)
  const res = await findCourse({ ID: row.ID })

  if (res.data.recourse.cremain >= res.data.recourse.cmax) {
    ElMessage.error('该课程选课人数已满！')
      return
  }

  if (res.data.recourse.cremain < res.data.recourse.cmax) {
    const table = await getSctListBySno({ page: page.value, pageSize: pageSize.value, sno: stuSno.value })
    if (table.code === 0) {
      const json = JSON.stringify(table.data.list)
      const arr = JSON.parse(json)
      for (let i = 0; i < arr.length; i++) {
             if (res.data.recourse.cplace === arr[i].cplace) {
              ElMessage.error('上课地点或时间重复！')
              return
             }
             if (res.data.recourse.cname === arr[i].cname) {
              ElMessage.error('你已选过该类课程！')
              return
             }
      }
}
    // const isRepeat
    // console.log('yes')
    SctformData.value.sno = stuSno
    SctformData.value.cno = res.data.recourse.cno
    SctformData.value.cname = res.data.recourse.cname
    SctformData.value.ccredit = res.data.recourse.ccredit
    SctformData.value.tno = res.data.recourse.ctno
    SctformData.value.cplace = res.data.recourse.cplace
    SctformData.value.grade = 0

   const flag = await createSct(SctformData.value)
    if (flag.code === 0) {
      ElMessage({
                type: 'success',
                message: '选课成功'
            })
          // isDisabled = true
          formData.value = res.data.recourse
          formData.value.cremain += 1
          updateCourse(formData.value)
          getTableData()
    }
  }
  // console.log(res.data.recourse.cname)
}
getTableData()
</script>

<style>

</style>
