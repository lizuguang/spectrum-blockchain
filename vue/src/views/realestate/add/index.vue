<template>
  <div class="app-container">
    <el-form ref="ruleForm" v-loading="loading" :model="ruleForm" :rules="rules" label-width="100px">

      <el-form-item label="用　　户" prop="proprietor">
        <el-select v-model="ruleForm.proprietor" placeholder="请选择用户" @change="selectGet">
          <el-option
            v-for="item in accountList"
            :key="item.accountId"
            :label="item.userName"
            :value="item.accountId"
          >
            <span style="float: left">{{ item.userName }}</span>
            <span style="float: right; color: #8492a6; font-size: 13px">{{ item.accountId }}</span>
          </el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="频率范围" prop="minfrequency">
        <el-col :span="5">
          <el-form-item prop="minfrequency">
            <el-input-number v-model="ruleForm.minfrequency" :precision="2" :step="100" :min="0" placeholder="最小值"></el-input-number>
            <span style="font-weight:bold">&#12288 &#12288 ——</span>
          </el-form-item>
        </el-col>
        <el-col :span="5">
          <el-form-item prop="maxfrequency">
            <el-input-number v-model="ruleForm.maxfrequency" :precision="2" :step="100" :min="0" placeholder="最大值"></el-input-number> &#12288 Hz
          </el-form-item>
        </el-col>
      </el-form-item>

      <el-form-item label="授权时间" prop="startdate">
        <el-col :span="5">
          <el-form-item prop="startdate">
        <el-date-picker
          v-model="ruleForm.startdate"
          type="datetime"
          value-format="yyyy-MM-dd hh:mm:ss"
          @change="dateChange1"
          placeholder="开始日期" :picker-options="pickerOptions0"></el-date-picker><span style="font-weight:bold">&#12288 &#12288 至</span>
      </el-form-item>
        </el-col>
      <el-col :span="5">
        <el-form-item prop="enddate">
          <el-date-picker
            v-model="ruleForm.enddate"
            type="datetime"
            value-format="yyyy-MM-dd hh:mm:ss"
            @change="dateChange2"
            placeholder="结束日期" :picker-options="pickerOptions1">
          </el-date-picker>
        </el-form-item>
      </el-col>
      </el-form-item>


      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/account'
import { createRealEstate } from '@/api/realEstate'

export default {
  name: 'AddRealeState',
  data() {
    var checkArea = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error('必须大于0'))
      } else {
        callback()
      }
    }

    return {
      //比较两输入框的时间，开始日期必须早于结束日期
      pickerOptions0: {
        disabledDate: (time) => {
          if (this.ruleForm.enddate != "") {
            return time.getTime() <= Date.now() - 8.64e7 || time.getTime() <= this.ruleForm.enddate  - 8.64e7;
          } else {
            return time.getTime() <= Date.now()  - 8.64e7;
          }
        }

      },
      pickerOptions1: {
        disabledDate: (time) => {
          if(this.ruleForm.startdate == ''){
            return time.getTime() > this.ruleForm.startdate  || time.getTime() <  Date.now() ;
          } else {
            var date = new Date(this.ruleForm.startdate);
            var time1 = date.getTime();
            return time.getTime() > this.ruleForm.startdate   || time.getTime() <  time1 - 3600 * 1000 * 24 ;
          }
        }
      },
      ruleForm: {
        proprietor: '',
        minfrequency: 0,
        maxfrequency: 0,
        startdate: '',
        enddate: ''
      },
      accountList: [],
      rules: {
        proprietor: [
          { required: true, message: '请选择用户', trigger: 'change' }
        ],
        minfrequency: [
          { validator: checkArea, trigger: 'blur' }
        ],
        maxfrequency: [
          { validator: checkArea, trigger: 'blur' }
        ]
      },
      loading: false
    }
  },
  computed: {
    ...mapGetters([
      'accountId'
    ])
  },
  created() {
    queryAccountList().then(response => {
      if (response !== null) {
        // 过滤掉管理员
        this.accountList = response.filter(item =>
          item.userName !== '管理员'
        )
      }
    })
  },
  methods: {
    dateChange1(val){
      this.ruleForm.startdate = val;
    },
    dateChange2(val){
      this.ruleForm.enddate = val;
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即创建?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loading = true
            createRealEstate({
              accountId: this.accountId,
              proprietor: this.ruleForm.proprietor,
              minfrequency: this.ruleForm.minfrequency,
              maxfrequency: this.ruleForm.maxfrequency,
              startdate: this.ruleForm.startdate,
              enddate: this.ruleForm.enddate
            }).then(response => {
              this.loading = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '创建成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '创建失败!'
                })
              }
            }).catch(_ => {
              this.loading = false
            })
          }).catch(() => {
            this.loading = false
            this.$message({
              type: 'info',
              message: '已取消创建'
            })
          })
        } else {
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    selectGet(accountId) {
      this.ruleForm.proprietor = accountId
    }
  }
}
</script>

<style scoped>
</style>
