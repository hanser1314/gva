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

        <el-form-item>
          <el-input
           placeholder="姓名"
           v-model="searchInfo.tname"
           type="text"
           clearable
            >
                <template #prefix>
                  <el-icon class="el-input__icon"><search /></el-icon>
                </template>
              </el-input>
     
        </el-form-item>

        <el-form-item>
          <el-input
           placeholder="教工号"
           v-model="searchInfo.tno"
           type="text"
           clearable
            >
                <template #prefix>
                  <el-icon class="el-input__icon"><search /></el-icon>
                </template>
              </el-input>
     
        </el-form-item>
       

        <el-form-item>
          <el-button type="primary" icon="search" @click="handleQuery">查询2</el-button>
        </el-form-item>



      </el-form>

    </div>


    <div class="gva-table-box">
        <div class="gva-btn-list">
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

               <el-button type="success" icon="Printer" @click="onPrint">打印</el-button>

        </div>

        <el-table
        stripe =true
        border =true
        size = large
        empty-text
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        id="box"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column sortable align="left" label="教工号" prop="tno" width="120" />
        <el-table-column align="left" label="姓名" prop="tname" width="120" />
        <el-table-column align="left" label="性别" prop="tsex" width="120" />
        <el-table-column align="left" label="年龄" prop="tage" width="120" />
        <el-table-column align="left" label="学历" prop="teb" width="120" />
        <el-table-column align="left" label="职称" prop="tpt" width="120" />
        <el-table-column align="left" label="主讲课程一" prop="cno1" width="120" />
        <el-table-column align="left" label="主讲课程二" prop="cno2" width="120" />
        <el-table-column align="left" label="主讲课程三" prop="cno3" width="120" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateTeacherFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="教工号:"  prop="tno" >
              <el-input v-model="formData.tno" :clearable="true"  placeholder="请输入教工号" />
            </el-form-item>
            <el-form-item label="姓名:"  prop="tname" >
              <el-input v-model="formData.tname" :clearable="true"  placeholder="请输入姓名" />
            </el-form-item>
            <el-form-item label="性别:" prop="tsex">
              <el-select v-model="formData.tsex" placeholder="请选择性别">
              <el-option label="男" value="男"></el-option> <el-option label="女" value="女"></el-option>
              </el-select>
              </el-form-item>
            <el-form-item label="年龄:"  prop="tage" >
              <el-input v-model.number="formData.tage" :clearable="true" placeholder="请输入年龄" />
            </el-form-item>
             <el-form-item label="学历:" prop="teb">
                  <el-select v-model="formData.teb" placeholder="请选择学历">
                  <el-option label="博士" value="博士"></el-option>
                  <el-option label="硕士" value="硕士"></el-option>
                  <el-option label="学士" value="学士"></el-option>
                  </el-select>
                  </el-form-item>
                  <el-form-item label="职称:" prop="tpt">
                    <el-select v-model="formData.tpt" placeholder="请选择职称">
                              <el-option label="助教" value="助教"></el-option>
                              <el-option label="讲师" value="讲师"></el-option>
                              <el-option label="副教授" value="副教授"></el-option>
                              <el-option label="教授" value="教授"></el-option>
                              </el-select>
                    </el-form-item>
            <el-form-item label="主讲课程一:"  prop="cno1" >
              <el-input v-model="formData.cno1" :clearable="true"  placeholder="请输入主讲课程一" />
            </el-form-item>
            <el-form-item label="主讲课程二:"  prop="cno2" >
              <el-input v-model="formData.cno2" :clearable="true"  placeholder="请输入主讲课程二" />
            </el-form-item>
            <el-form-item label="主讲课程三:"  prop="cno3" >
              <el-input v-model="formData.cno3" :clearable="true"  placeholder="请输入主讲课程三" />
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
        <el-descriptions column="1" border=true>
                <el-descriptions-item label="教工号">
                        {{ formData.tno }}
                </el-descriptions-item>
                <el-descriptions-item label="姓名">
                        {{ formData.tname }}
                </el-descriptions-item>
                <el-descriptions-item label="性别">
                        {{ formData.tsex }}
                </el-descriptions-item>
                <el-descriptions-item label="年龄">
                        {{ formData.tage }}
                </el-descriptions-item>
                <el-descriptions-item label="学历">
                        {{ formData.teb }}
                </el-descriptions-item>
                <el-descriptions-item label="职称">
                        {{ formData.tpt }}
                </el-descriptions-item>
                <el-descriptions-item label="主讲课程一">
                        {{ formData.cno1 }}
                </el-descriptions-item>
                <el-descriptions-item label="主讲课程二">
                        {{ formData.cno2 }}
                </el-descriptions-item>
                <el-descriptions-item label="主讲课程三">
                        {{ formData.cno3 }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createTeacher,
  deleteTeacher,
  deleteTeacherByIds,
  updateTeacher,
  findTeacher,
  getTeacherList,
  getTeacherByNameOrTno
} from '@/api/teacher'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import printJS from 'print-js'

defineOptions({
    name: 'Teacher'
})

//打印按钮
const onPrint = () => {
  //获取元素
  // const node = multipleTable.value;
  printJS({
    printable: "box", // 标签元素id
    type: "html",
    targetStyles: ["*"], //添加样式
    maxWidth: '1000px'
  })
}

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        tno: '',
        tname: '',
        tsex: '',
        tage: 0,
        teb: '',
        tpt: '',
        cno1: '',
        cno2: '',
        cno3: '',
        })

// 验证规则
const rule = reactive({
               tno : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               tname : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               tsex : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               tage : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               teb : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               cno1 : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
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
// 排序
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
  getTableData()
}

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

const handleQuery = () => {
          page.value = 1
          pageSize.value = 10
          getTableDataByNameOrTno()
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
  const table = await getTeacherList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}


const getTableDataByNameOrTno = async() => {
    const table = await getTeacherByNameOrTno({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteTeacherFunc(row)
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
      const res = await deleteTeacherByIds({ ids })
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
const updateTeacherFunc = async(row) => {
    const res = await findTeacher({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reteacher
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteTeacherFunc = async (row) => {
    const res = await deleteTeacher({ ID: row.ID })
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
  const res = await findTeacher({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reteacher
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          tno: '',
          tname: '',
          tsex: '',
          tage: 0,
          teb: '',
          tpt: '',
          cno1: '',
          cno2: '',
          cno3: '',
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
        tno: '',
        tname: '',
        tsex: '',
        tage: 0,
        teb: '',
        tpt: '',
        cno1: '',
        cno2: '',
        cno3: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createTeacher(formData.value)
                  break
                case 'update':
                  res = await updateTeacher(formData.value)
                  break
                default:
                  res = await createTeacher(formData.value)
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
