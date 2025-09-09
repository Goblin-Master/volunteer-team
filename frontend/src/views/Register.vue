<template>
    <div class="register-container">
        <el-card class="register-card" shadow="never">
            <template #header>
                <div class="card-header">
                    <h2>用户注册</h2>
                    <span>创建您的账户，开始使用我们的服务</span>
                </div>
            </template>

            <el-form ref="registerFormRef" :model="registerForm" :rules="registerRules" label-position="top"
                size="large" @submit.prevent="handleRegister">

                <!-- 邮箱 -->
                <el-form-item label="邮箱" prop="email">
                    <el-input v-model="registerForm.email" type="email" placeholder="请输入您的邮箱地址" clearable />
                </el-form-item>

                <!-- 邮箱验证码 -->
                <el-form-item label="邮箱验证码" prop="code">
                    <div class="verification-code-wrapper">
                        <el-input v-model="registerForm.code" placeholder="请输入邮箱验证码" clearable />
                        <el-button class="send-code-btn" type="primary" :disabled="isSendingCode || !isEmailValid"
                            @click="sendVerificationCode">
                            {{ codeButtonText }}
                        </el-button>
                    </div>
                </el-form-item>
                <!-- 用户昵称（新增） -->
                <el-form-item label="用户昵称" prop="username">
                    <el-input v-model="registerForm.username" placeholder="请输入用户昵称" clearable />
                </el-form-item>
                <!-- 用户名（统一为 account） -->
                <el-form-item label="账号" prop="account">
                    <el-input v-model="registerForm.account" placeholder="请输入账号" clearable />
                </el-form-item>

                <!-- 密码 -->
                <el-form-item label="密码" prop="password">
                    <el-input v-model="registerForm.password" type="password" placeholder="请输入密码" show-password
                        clearable />
                </el-form-item>

                <!-- 确认密码 -->
                <el-form-item label="确认密码" prop="confirm_password">
                    <el-input v-model="registerForm.confirm_password" type="password" placeholder="请再次输入密码"
                        show-password clearable />
                </el-form-item>

                <!-- 内部人员 -->
                <el-form-item label="内部人员">
                    <el-switch v-model="registerForm.is_internal" @change="handleInternalSwitchChange" />
                </el-form-item>

                <!-- 协议勾选 -->
                <el-form-item prop="agree_terms">
                    <el-checkbox v-model="registerForm.agree_terms">
                        我已阅读并同意
                        <el-link type="primary" :underline="false">用户协议</el-link>
                        和
                        <el-link type="primary" :underline="false">隐私政策</el-link>
                    </el-checkbox>
                </el-form-item>

                <!-- 提交 -->
                <el-form-item>
                    <el-button type="primary" class="register-btn" native-type="submit" :loading="loading">
                        注 册
                    </el-button>
                </el-form-item>
            </el-form>

            <div class="login-link">
                已有账号？
                <el-link type="primary" @click="goToLogin">立即登录</el-link>
            </div>
        </el-card>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { useRouter } from 'vue-router'
import type { FormInstance } from 'element-plus'
import { ElMessage } from 'element-plus'
import { Register, GetRegisterCode } from '@/api/register'
import type { RegisterForm, RegisterFormRules } from '@/types/register'

/* ---------- 表单数据 ---------- */
const registerForm = reactive<RegisterForm>({
    email: '',
    code: '',
    username: '',
    account: '',               // ✅ 统一为 account
    password: '',
    confirm_password: '',
    is_internal: false,
    agree_terms: false,
})

const registerFormRef = ref<FormInstance>()
const router = useRouter()
const loading = ref(false)

/* ---------- 验证码逻辑 ---------- */
const isSendingCode = ref(false)
const countdown = ref(60)
const codeButtonText = computed(() =>
    isSendingCode.value ? `${countdown.value}秒后重试` : '发送验证码'
)

const isEmailValid = computed(() =>
    /^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$/.test(registerForm.email)
)

const sendVerificationCode = async () => {
    if (!isEmailValid.value) {
        ElMessage.error('请输入有效的邮箱地址！')
        return
    }
    isSendingCode.value = true
    try {
        const res = await GetRegisterCode(registerForm.email)
        if (res.code === 0) {
            ElMessage.success('验证码已发送至您的邮箱！')
            startCountdown()
        } else {
            ElMessage.error(res.message || '获取验证码失败！')
            isSendingCode.value = false
        }
    } catch (e) {
        ElMessage.error('网络错误或服务器无响应')
        isSendingCode.value = false
    }
}

const startCountdown = () => {
    countdown.value = 60
    const timer = setInterval(() => {
        countdown.value--
        if (countdown.value <= 0) {
            clearInterval(timer)
            isSendingCode.value = false
        }
    }, 1000)
}

/* ---------- 校验规则 ---------- */
const validatePass = (rule: any, value: string, callback: any) => {
    if (!value) callback(new Error('请输入密码'))
    else {
        if (registerForm.confirm_password) {
            registerFormRef.value?.validateField('confirmPassword')
        }
        callback()
    }
}

const validatePass2 = (rule: any, value: string, callback: any) => {
    if (!value) callback(new Error('请再次输入密码'))
    else if (value !== registerForm.password) {
        callback(new Error('两次输入的密码不一致！'))
    } else callback()
}

const validateAgreement = (rule: any, value: boolean, callback: any) => {
    value ? callback() : callback(new Error('请先阅读并同意用户协议和隐私政策'))
}

const registerRules = reactive<RegisterFormRules>({
    username: [
        { required: true, message: '请输入用户昵称', trigger: 'blur' },
        { min: 1, max: 14, message: '昵称长度 1-14 个字符', trigger: 'blur' }
    ],
    email: [
        { required: true, message: '请输入邮箱地址', trigger: 'blur' },
        { type: 'email', message: '邮箱地址格式不正确', trigger: ['blur', 'change'] },
    ],
    code: [
        { required: true, message: '请输入邮箱验证码', trigger: 'blur' },
    ],
    account: [                      // ✅ 统一为 account
        { required: true, message: '请输入用户名', trigger: 'blur' },
        { min: 3, max: 15, message: '长度在 3 到 15 个字符', trigger: 'blur' },
    ],
    password: [
        { validator: validatePass, trigger: 'blur', required: true },
        { min: 6, message: '密码长度不能少于6位', trigger: 'blur' },
    ],
    confirm_password: [
        { validator: validatePass2, trigger: 'blur', required: true },
    ],
    agree_terms: [
        { validator: validateAgreement, trigger: 'change' },
    ],
})

/* ---------- 提交 ---------- */
const handleRegister = async () => {
    if (!registerFormRef.value) return
    try {
        await registerFormRef.value.validate()
        loading.value = true
        const resp = await Register(registerForm)
        if (resp.code === 0) {
            ElMessage.success('注册成功！')
            router.replace('/login')
        } else {
            ElMessage.error(resp.message || '注册失败')
        }
    } catch (err: any) {
        err === false
            ? ElMessage.error('请检查表单输入项！')
            : ElMessage.error(err.message || '网络异常')
    } finally {
        loading.value = false
    }
}

/* ---------- 内部人员提示 ---------- */
const handleInternalSwitchChange = (val: boolean | string | number) => {
    if (val === true) {
        ElMessage.info({ message: '内部账号注册后需要管理员审核。', duration: 4000 })
    }
}

/* ---------- 跳转登录 ---------- */
const goToLogin = () => router.push('/login')
</script>

<style scoped>
/* 样式与之前完全一致，无需改动 */
.register-container {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    padding: 20px;
    background-color: #f0f2f5;
}

.register-card {
    width: 100%;
    max-width: 380px;
    padding: 40px;
    background: #ffffff;
    border-radius: 16px;
    box-shadow: 0 10px 30px -10px rgba(0, 0, 0, 0.1);
}

.card-header {
    text-align: center;
    margin-bottom: 30px;
}

.card-header h2 {
    font-size: 24px;
    font-weight: 600;
    color: #303133;
    margin: 0 0 8px;
}

.card-header span {
    font-size: 14px;
    color: #909399;
}

.el-form-item {
    margin-bottom: 22px;
}

:deep(.el-form-item__label) {
    font-weight: 500;
    color: #606266;
    margin-bottom: 6px !important;
    line-height: 1.2;
}

:deep(.el-input__wrapper) {
    height: 48px;
    padding: 0 15px;
    border-radius: 8px;
    border: 1px solid #dcdfe6;
    box-shadow: none !important;
    transition: border-color 0.2s;
}

:deep(.el-input__wrapper:hover) {
    border-color: #c0c4cc;
}

:deep(.el-input__wrapper.is-focus) {
    border-color: var(--el-color-primary);
}

.verification-code-wrapper {
    display: flex;
    align-items: center;
    gap: 10px;
    width: 100%;
}

.verification-code-wrapper .el-input {
    flex-grow: 1;
}

.send-code-btn {
    flex-shrink: 0;
    height: 48px;
    border-radius: 8px;
    background-color: #e8f3ff;
    color: var(--el-color-primary);
    border: none;
    font-weight: 500;
}

.send-code-btn:hover {
    background-color: #d9e9ff;
}

.register-btn {
    width: 100%;
    height: 48px;
    font-size: 16px;
    border-radius: 8px;
    border: none;
    background: var(--el-color-primary);
}

.login-link {
    margin-top: 24px;
    text-align: center;
    font-size: 14px;
}

@media (max-width: 480px) {
    .register-card {
        padding: 24px;
    }
}
</style>