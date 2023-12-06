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
          <el-table-column align="left" label="学号" prop="sno" width="120" />
          <el-table-column align="left" label="课程号" prop="cno" width="120" />
          <el-table-column align="left" label="课程名" prop="cname" width="120" />
          <el-table-column align="left" label="学分" prop="ccredit" width="120" />
          <el-table-column align="left" label="教工号" prop="tno" width="120" />
          <el-table-column align="left" label="上课地点、时间" prop="cplace" width="240" />
          <el-table-column align="left" label="成绩" prop="grade" width="120" />
          <el-table-column align="left" label="操作">
              <template #default="scope">
              <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                  <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                  查看详情
              </el-button>
              <el-button type="primary" link icon="edit" class="table-button" @click="updateSctFunc(scope.row)">评分</el-button>
              <!-- <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button> -->
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
      <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
        <el-scrollbar height="500px">
            <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">
              <el-form-item label="学号:"  prop="sno" >
                <el-input v-model="formData.sno" :clearable="true"  placeholder="请输入学号" disabled/>
              </el-form-item>
              <el-form-item label="课程号:"  prop="cno" >
                <el-input v-model="formData.cno" :clearable="true"  placeholder="请输入课程名" disabled/>
              </el-form-item>
              <el-form-item label="课程名:"  prop="cname" >
                <el-input v-model="formData.cname" :clearable="true"  placeholder="请输入课程名" disabled/>
              </el-form-item>
              <el-form-item label="学分:"  prop="ccredit" >
                <el-input v-model.number="formData.ccredit" :clearable="true" placeholder="请输入学分" disabled/>
              </el-form-item>
              <el-form-item label="教工号:"  prop="tno" >
                <el-input v-model="formData.tno" :clearable="true"  placeholder="请输入教工号" disabled/>
              </el-form-item>
              <el-form-item label="上课地点、时间:"  prop="cplace" >
                <el-input v-model="formData.cplace" :clearable="true"  placeholder="请输入上课地点、时间" disabled/>
              </el-form-item>
              <el-form-item label="成绩:"  prop="grade" >
                <el-input v-model.number="formData.grade" :clearable="true"  placeholder="请输入成绩" />
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
                  <el-descriptions-item label="学号">
                          {{ formData.sno }}
                  </el-descriptions-item>
                  <el-descriptions-item label="课程号">
                          {{ formData.cno }}
                  </el-descriptions-item>
                  <el-descriptions-item label="课程名">
                          {{ formData.cname }}
                  </el-descriptions-item>
                  <el-descriptions-item label="学分">
                          {{ formData.ccredit }}
                  </el-descriptions-item>
                  <el-descriptions-item label="教工号">
                          {{ formData.tno }}
                  </el-descriptions-item>
                  <el-descriptions-item label="上课地点、时间">
                          {{ formData.cplace }}
                  </el-descriptions-item>
                  <el-descriptions-item label="成绩">
                          {{ formData.grade }}
                  </el-descriptions-item>
          </el-descriptions>
        </el-scrollbar>
      </el-dialog>
    </div>
  </template>

  <script setup>
  import {
    createSct,
    deleteSct,
    deleteSctByIds,
    updateSct,
    findSct,
    getSctList,
    getSctListByTno
  } from '@/api/sct'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 我的函数
import { findCourse, updateCourse } from '@/api/course'

import { useUserStore } from '@/pinia/modules/user'

import { useBtnAuth } from '@/utils/btnAuth'
const btnAuth = useBtnAuth()

const username = ref('')
const userStore = useUserStore()
username.value = userStore.userInfo.userName

  // var courseID = inject('courseID')
//   const test = (row) =>{
//     alert(row.ID)
//   }

  defineOptions({
      name: 'Sct'
  })

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
          sno: '',
          cno: '',
          cname: '',
          ccredit: 0,
          tno: '',
          cplace: '',
          grade: 0,
          })

  const CourseFormData = ref({
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
    const table = await getSctListByTno({ page: page.value, pageSize: pageSize.value, tno: username.value })
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
              deleteSctFunc(row)
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
        const res = await deleteSctByIds({ ids })
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
  const updateSctFunc = async(row) => {
      const res = await findSct({ ID: row.ID })
      type.value = 'update'
      if (res.code === 0) {
          formData.value = res.data.resct
          dialogFormVisible.value = true
      }
  }


  // 删除行
  const deleteSctFunc = async(row) => {
      // const info = await findCourse({ ID: courseID })
      const res = await deleteSct({ ID: row.ID })
      if (res.code === 0) {
          ElMessage({
                  type: 'success',
                  message: '删除成功'
              })
              if (tableData.value.length === 1 && page.value > 1) {
              page.value--
          }
          // CourseFormData.value = info.data.recourse
          // CourseFormData.value.cremain -= 1
          // updateCourse(CourseFormData.value)
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
    const res = await findSct({ ID: row.ID })
    if (res.code === 0) {
      formData.value = res.data.resct
      openDetailShow()
    }
  }


  // 关闭详情弹窗
  const closeDetailShow = () => {
    detailShow.value = false
    formData.value = {
            sno: '',
            cno: '',
            cname: '',
            ccredit: 0,
            tno: '',
            cplace: '',
            grade: 0,
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
          sno: '',
          cno: '',
          cname: '',
          ccredit: 0,
          tno: '',
          cplace: '',
          grade: 0
          }
  }
  // 弹窗确定
  const enterDialog = async () => {
       elFormRef.value?.validate( async (valid) => {
               if (!valid) return
                let res
                switch (type.value) {
                  case 'create':
                    res = await createSct(formData.value)
                    break
                  case 'update':
                    res = await updateSct(formData.value)
                    break
                  default:
                    res = await createSct(formData.value)
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

  </script>

  <style>

  </style>
