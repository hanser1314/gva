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
           v-model="searchInfo.sname"
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
           placeholder="年龄"
           v-model="searchInfo.sage"
           type="text"
           clearable
            >
                <template #prefix>
                  <el-icon class="el-input__icon"><search /></el-icon>
                </template>
              </el-input>

        </el-form-item>


        <el-form-item>
          <el-button type="primary" icon="search" @click="handleQueryName">查询2</el-button>
        </el-form-item>

       <el-form-item>
        <span>
          根据学院进行筛选
          <el-tooltip content="从以下学院中选中">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
        <el-select v-model="searchInfo.sdept" class="m-2" placeholder="Select" size="large"
        @visible-change="handleQueryDept"
        >
    <el-option
      v-for="item in depart"
      :key="item.value"
      :label="item.label"
      :value="item.value"
    />
        </el-select>
       </el-form-item>
      </el-form>
    </div>

      <!-- <div class="my-search">
        <el-form :model="searchInfo">

        <el-form-item>
          <el-input class="query-input"
           placeholder="Type something"
           v-model="searchInfo.sname"
           type="text"
           clearable
            >
                <template #prefix>
                  <el-icon class="el-input__icon"><search /></el-icon>
                </template>
              </el-input>

        </el-form-item>
          </el-form>

        <el-form-item>
          <el-button type="primary" icon="search" @click="handleQueryName">查询2</el-button>
        </el-form-item>

      </div> -->

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


            <!-- <el-button type="primary"  @click="onPrint">
              <el-icon><Printer /></el-icon>
              打印</el-button> -->

              <el-button type="success" icon="Printer" @click="onPrint">打印</el-button>

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
        id="box"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="学号" prop="sno" width="140" />
        <el-table-column align="left" label="姓名" prop="sname" width="120" />
        <el-table-column align="left" label="性别" prop="ssex" width="120" />
        <el-table-column align="left" label="年龄" prop="sage" width="120" />
        <el-table-column align="left" label="出生日期" prop="sborn" width="240" />
        <el-table-column align="left" label="系别" prop="sdept" width="120" />
        <el-table-column align="left" label="操作">

            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
               <el-icon con style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateStudentFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            <!-- <el-button type="primary" link icon="delete" @click="onPrint">打印</el-button> -->
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
              <el-input v-model="formData.sno" :clearable="true"  placeholder="请输入学号" />
            </el-form-item>
            <el-form-item label="姓名:"  prop="sname" >
              <el-input v-model="formData.sname" :clearable="true"  placeholder="请输入姓名" />
            </el-form-item>
             <el-form-item label="性别:" prop="ssex">
              <el-select v-model="formData.ssex" placeholder="请选择性别">
              <el-option label="男" value="男"></el-option> <el-option label="女" value="女"></el-option>
              </el-select>
              </el-form-item>
            <el-form-item label="年龄:"  prop="sage" >
              <el-input v-model.number="formData.sage" :clearable="true" placeholder="请输入年龄" />
            </el-form-item>
            <el-form-item label="出生日期:"  prop="sborn" >
              <el-date-picker
                  v-model="formData.sborn"
                  type="date"
                  placeholder="请选择出生日期"
                  :disabled-date="disabledDate"
                  :shortcuts="shortcuts"
                  :size="large"
                />
            </el-form-item>
            <el-form-item label="系别:"  prop="sdept" >
              <el-input v-model="formData.sdept" :clearable="true"  placeholder="请输入系别" />
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
        <el-descriptions column="1"  border = true>
                <el-descriptions-item label="学号">
                        {{ formData.sno }}
                </el-descriptions-item>
                <el-descriptions-item label="姓名">
                        {{ formData.sname }}
                </el-descriptions-item>
                <el-descriptions-item label="性别">
                        {{ formData.ssex }}
                </el-descriptions-item>
                <el-descriptions-item label="年龄">
                        {{ formData.sage }}
                </el-descriptions-item>
                  <el-descriptions-item label="出生日期">
                    {{ formData.sborn}}
                  </el-descriptions-item>
                <el-descriptions-item label="系别">
                        {{ formData.sdept }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>

    <br>
    <br>
    <div class="gva-search-box" ref="main" style="width: 100%; height: 400px"></div>

  </div>

</template>

<script setup lang="ts">
import {
  createStudent,
  deleteStudent,
  deleteStudentByIds,
  updateStudent,
  findStudent,
  getStudentList,
  getStudentByName,
  getStudentByDept
} from '@/api/student'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive,onMounted } from 'vue'
import printJS from 'print-js'
import * as echarts from 'echarts'
// 我的函数
// import stuchart from '@/view/mine/studentchart.vue'

const shortcuts = [
  {
    text: 'Today',
    value: new Date(),
  },
  {
    text: 'Yesterday',
    value: () => {
      const date = new Date()
      date.setTime(date.getTime() - 3600 * 1000 * 24)
      return date
    },
  },
  {
    text: 'A week ago',
    value: () => {
      const date = new Date()
      date.setTime(date.getTime() - 3600 * 1000 * 24 * 7)
      return date
    },
  },
]

const disabledDate = (time: Date) => {
  return time.getTime() > Date.now()
}

const main = ref() // 使用ref创建虚拟DOM引用，使用时用main.value

var jsj = ref(0)
var wx = ref(0)
// var ty = ref(0)
// var yy = ref(0)
// var sw = ref(0)
// var sx = ref(0)

// const getStuNum = async() => {
//   const table = await getStudentByDept({ page: page.value, pageSize: pageSize.value, sdept: '计算机学院' })
//   const json = JSON.stringify(table.data.list)
//   const arr = JSON.parse(json)
//   console.log(arr.length)
//   var jsj= ref(0)
//    jsj.value = arr.length
//    return jsj.value
// }

onMounted(
  () => {
    init()
  }
)


// function init() {
const init = async() => {

    var table = await getStudentByDept({ page: page.value, pageSize: pageSize.value, sdept: '计算机学院' })
    jsj.value = table.data.list.length

    table = await getStudentByDept({ page: page.value, pageSize: pageSize.value, sdept: '文学院' })
    wx.value = table.data.list.length
  // 基于准备好的dom，初始化echarts实例


  var myChart = echarts.init(main.value)

  // 指定图表的配置项和数据
  var option = {
    title: {
      text: '各级学院人数'
    },
    tooltip: {},
    legend: {
      data: ['人数']
    },
    xAxis: {
      data: ['计算机学院', '文学院', '体育学院', '音乐学院', '生物科学学院']
      // data: jsj.value
    },
    yAxis: {},
    series: [
      {
        name: '人数',
        type: 'bar',
        data: [jsj.value,wx.value,34,45,26]
      }
    ]
  }
  // 使用刚指定的配置项和数据显示图表。
  myChart.setOption(option)
}

const depart = [
  {
    value: '',
    label: '全部(All)',
  },
  {
    value: '计算机学院',
    label: '计算机学院',
  },
  {
    value: '文学院',
    label: '文学院',
  },
  {
    value: '体育学院',
    label: '体育学院',
  },
  {
    value: '音乐学院',
    label: '音乐学院',
  },
  {
    value: '生物科学学院',
    label: '生物科学学院',
  },
]

defineOptions({
    name: 'Student'
})

const multipleTable = ref(null)
// const loading = ref(true)

// 打印按钮
const onPrint = () => {
  // 获取元素
  // const node = multipleTable.value;
  printJS({
    printable: 'box', // 标签元素id
    type: 'html',
    targetStyles: ['*'], // 添加样式
    maxWidth: '1000px'
    // scanStyles: false
  })
}

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        sno: '',
        sname: '',
        ssex: '',
        sage: 0,
        sborn: new Date(2000,1,1),
        sdept: '',
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

// const tableDataCopy = Object.assign(tableData)

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

const handleQueryName = () => {
    // if (val.length > 0) {
    //   tableData.value = tableData.value.filter(item => (item.value).toLowerCase().match(val.toLowerCase()))
    // } else {
    //   tableData.value = tableDataCopy.value
    // }
    // elSearchFormRef.value?.validate(async(valid) => {
    //   if(!valid) return
          page.value = 1
          pageSize.value = 10
          getTableDataByName()
    // })
}

const handleQueryDept = () => {
    // if (val.length > 0) {
    //   tableData.value = tableData.value.filter(item => (item.value).toLowerCase().match(val.toLowerCase()))
    // } else {
    //   tableData.value = tableDataCopy.value
    // }
    // elSearchFormRef.value?.validate(async(valid) => {
    //   if(!valid) return
          page.value = 1
          pageSize.value = 10
          getTableDataByDept()
    // })
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
  const table = await getStudentList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    // loading.value = false
  }
}

const getTableDataByName = async() => {
    const table = await getStudentByName({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
}

const getTableDataByDept = async() => {
    const table = await getStudentByDept({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteStudentFunc(row)
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
      const res = await deleteStudentByIds({ ids })
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
const updateStudentFunc = async(row) => {
    const res = await findStudent({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.restudent
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteStudentFunc = async (row) => {
    const res = await deleteStudent({ ID: row.ID })
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
  const res = await findStudent({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.restudent
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          sno: '',
          sname: '',
          ssex: '',
          sage: 0,
          sborn: new Date(2000,1,1),
          sdept: '',
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
        sname: '',
        ssex: '',
        sage: 0,
        sborn: new Date(2000,1,1),
        sdept: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createStudent(formData.value)
                  break
                case 'update':
                  res = await updateStudent(formData.value)
                  break
                default:
                  res = await createStudent(formData.value)
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
