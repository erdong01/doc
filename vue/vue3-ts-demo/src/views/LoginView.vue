<template>
    <div class="login-box">
        <h2>后台管理系统</h2>
        <el-form ref="ruleFormRef" :model="ruleForm" status-icon :rules="rules" label-width="80px"
            class="demo-ruleForm">

            <el-form-item label="账号：" prop="username">
                <el-input v-model="ruleForm.username" autocomplete="off" />
            </el-form-item>
            <el-form-item label="密码：" prop="password">
                <el-input v-model="ruleForm.password" type="password" autocomplete="off" />
            </el-form-item>
            <el-form-item>
                <el-button class="loginBtn" type="primary" @click="submitForm(ruleFormRef)">登录</el-button>
                <el-button class="loginBtn" @click="resetForm()">重置</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script lang="ts">
import type { FormInstance } from 'element-plus';
import { defineComponent, reactive, ref, toRefs } from 'vue';
import { useRouter } from 'vue-router';
import { login } from '../request/api';
import { LoginData } from '../type/login';
const formSize = ref('default')

export default defineComponent({
    setup() {
        const data = reactive(new LoginData())
        const rules = {
            username: [
                { required: true, message: '请输入账号', trigger: 'blur' },
                { min: 3, max: 10, message: '账号长度在 3 和 10之间', trigger: 'blur' },
            ],
            password: [
                { required: true, message: '请输入密码', trigger: 'blur' },
                { min: 3, max: 10, message: '密码长度在 3 和 10之间', trigger: 'blur' },
            ],
        }
        const router = useRouter()
        //登录
        const ruleFormRef = ref<FormInstance>()
        const submitForm = async (formEl: FormInstance | undefined) => {
            if (!formEl) return
            await formEl.validate((valid, fields) => {
                if (valid) {
                    console.log('submit!')
                    login(data.ruleForm).then((res)=>{
                        localStorage.setItem('token',res.data.token)
                        console.log(res)
                        router.push("/")
                    })
                } else {
                    console.log('error submit!', fields)
                }
            })
        }
        //重置
        const resetForm = () => {
            data.ruleForm.username = "";
            data.ruleForm.password = "";
        };


        return { ...toRefs(data), rules, ruleFormRef ,submitForm,resetForm}
    }
})
</script>

<style lang="scss" scoped>
.login-box {
    width: 100%;
    height: 100%;
    text-align: center;
}

.demo-ruleForm {
    width: 500px;
    margin: 200px auto;
    background-color: #fff;
    padding: 30px;
    border-radius: 40px;
}

.loginBtn {
    width: 48%;
}

h2 {
    margin-bottom: 10px;
}
</style>