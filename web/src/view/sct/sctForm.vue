<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="学号:" prop="sno">
          <el-input v-model="formData.sno" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="课程号:" prop="cno">
          <el-input v-model="formData.cno" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="课程名:" prop="cname">
          <el-input v-model="formData.cname" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="学分:" prop="ccredit">
          <el-input v-model.number="formData.ccredit" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="教工号:" prop="tno">
          <el-input v-model="formData.tno" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="上课地点:" prop="cplace">
          <el-input v-model="formData.cplace" :clearable="true" placeholder="请输入" />
       </el-form-item>
       <el-form-item label="成绩:" prop="grade">
          <el-input v-model="formData.grade" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createSct,
  updateSct,
  findSct
} from '@/api/sct'

defineOptions({
    name: 'SctForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')

const formData = ref({
            sno: '',
            cno: '',
            cname: '',
            ccredit: 0,
            tno: '',
            cplace: '',
            grade: 0,
 })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSct({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.resct
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
