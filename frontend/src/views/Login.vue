<template>
    <div class="login-container">
        <main class="main-content">
            <div class="login-box" :class="{ 'shake-animation': isShaking }">
                <div class="logo-section">
                    <img src="@/assets/logo.png" alt="Logo" />
                    <h3>欢迎回来</h3>
                    <p class="subtitle">请登录您的账户</p>
                </div>

                <el-tabs v-model="activeTab" class="login-tabs" stretch>
                    <el-tab-pane name="account">
                        <template #label>
                            <span class="tab-label">
                                <el-icon>
                                    <User />
                                </el-icon>
                                <span>账号登录</span>
                            </span>
                        </template>
                        <el-form ref="accountFormRef" :model="accountForm" :rules="accountRules" size="large"
                            @keyup.enter="handleLogin(accountFormRef)">
                            <el-form-item prop="account">
                                <el-input v-model="accountForm.account" placeholder="账号">
                                    <template #prefix>
                                        <el-icon>
                                            <User />
                                        </el-icon>
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="password">
                                <el-input v-model="accountForm.password" type="password" placeholder="密码" show-password>
                                    <template #prefix>
                                        <el-icon>
                                            <Lock />
                                        </el-icon>
                                    </template>
                                </el-input>
                            </el-form-item>
                        </el-form>
                    </el-tab-pane>

                    <el-tab-pane name="email">
                        <template #label>
                            <span class="tab-label">
                                <el-icon>
                                    <Message />
                                </el-icon>
                                <span>邮箱登录</span>
                            </span>
                        </template>
                        <el-form ref="emailFormRef" :model="emailForm" :rules="emailRules" size="large"
                            @keyup.enter="handleLogin(emailFormRef)">
                            <el-form-item prop="email">
                                <el-input v-model="emailForm.email" placeholder="邮箱">
                                    <template #prefix>
                                        <el-icon>
                                            <Message />
                                        </el-icon>
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="code">
                                <div class="verification-code-wrapper">
                                    <el-input v-model="emailForm.code" placeholder="验证码">
                                        <template #prefix>
                                            <el-icon>
                                                <Message />
                                            </el-icon> </template>
                                    </el-input>
                                    <el-button class="send-code-btn" :disabled="isSendingCode || !isEmailValid"
                                        @click="sendVerificationCode">
                                        {{ codeButtonText }}
                                    </el-button>
                                </div>
                            </el-form-item>
                        </el-form>
                    </el-tab-pane>
                </el-tabs>

                <div class="additional-links">
                    <el-link type="primary" :underline="false" @click="goToRegister">快速注册</el-link>
                    <el-link type="primary" :underline="false" @click="handleForgotPassword">忘记密码?</el-link>
                </div>

                <el-button type="primary" size="large" class="login-btn" :loading="isLoading"
                    @click="handleLogin(activeTab === 'account' ? accountFormRef : emailFormRef)">
                    登 录
                </el-button>
            </div>
        </main>

        <footer class="login-footer">
            ...
        </footer>
    </div>
</template>

<script setup lang="ts" name="Login">
import { ref, reactive, computed } from 'vue';
import { User, Lock, Message } from '@element-plus/icons-vue'; // [修改点 2] 移除 Iphone
import { ElMessage } from 'element-plus';
import { Login, GetCode } from '@/api/login.ts';
import type { LoginResp } from '@/types/login.ts';
import type BaseResp from '@/types/base.ts';
import type { FormInstance } from 'element-plus';

const activeTab = ref('account');
const isLoading = ref(false);
const isShaking = ref(false);

// Account Login Form
const accountFormRef = ref<FormInstance>();
const accountForm = reactive({
    account: '',
    password: '',
});
const accountRules = reactive({
    account: [{ required: true, message: '请输入您的账号', trigger: 'blur' }],
    password: [{ required: true, message: '请输入您的密码', trigger: 'blur' }],
});

// [修改点 3] Mobile Form -> Email Form
const emailFormRef = ref<FormInstance>();
const emailForm = reactive({
    email: '',
    code: '',
});
const emailRules = reactive({
    email: [
        { required: true, message: '请输入您的邮箱', trigger: 'blur' },
        { type: 'email', message: '请输入正确的邮箱格式', trigger: ['blur', 'change'] }
    ],
    code: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
});

// [修改点 4] 添加发送验证码逻辑
const isSendingCode = ref(false);
const countdown = ref(60);
const codeButtonText = computed(() => {
    return isSendingCode.value ? `${countdown.value}秒后重试` : '发送验证码';
});
const isEmailValid = computed(() =>
    /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(emailForm.email)
);
const sendVerificationCode = async () => {
    if (!isEmailValid.value) {
        ElMessage.error('请输入有效的邮箱地址！');
        return;
    }
    isSendingCode.value = true; // 1. 开始加载

    try {
        const res = await GetCode(emailForm.email);
        if (res.code === 0) {
            ElMessage.success('验证码已发送至您的邮箱!');
            startCountdown(); // 2. 成功才倒计时
        } else {
            ElMessage.error(res.message || '获取验证码失败!');
            isSendingCode.value = false; // 3. 失败立即恢复按钮
        }
    } catch (error) {
        console.error('GetCode request failed:', error);
        ElMessage.error('网络错误或服务器无响应');
        isSendingCode.value = false;
    }
};

const startCountdown = () => {
    countdown.value = 60;
    const timer = setInterval(() => {
        countdown.value--;
        if (countdown.value <= 0) {
            clearInterval(timer);
            isSendingCode.value = false;
        }
    }, 1000);
};


// [修改点 5] 更新登录处理器
const handleLogin = (formEl?: FormInstance) => {
    if (!formEl || isLoading.value) return;

    formEl.validate(async (valid: boolean) => {
        if (valid) {
            isLoading.value = true;
            try {
                let data: any;
                if (activeTab.value === 'account') {
                    data = {
                        account: accountForm.account,
                        password: accountForm.password,
                        login_type: 'account' // 建议增加登录类型字段
                    }
                } else { // activeTab.value === 'email'
                    data = {
                        email: emailForm.email,
                        code: emailForm.code,
                        login_type: 'email' // 建议增加登录类型字段
                    }
                }

                const res: BaseResp<LoginResp> = await Login(data);

                if (res.code === 0) {
                    ElMessage.success('登录成功!');
                    localStorage.setItem('token', res.data.token);
                    // TODO: Redirect to dashboard
                } else {
                    ElMessage.error(res.message || '登录失败');
                }
            } catch (error) {
                console.error('Login request failed:', error);
                ElMessage.error('网络错误或服务器无响应');
            } finally {
                isLoading.value = false;
            }
        } else {
            console.log('error submit!');
            isShaking.value = true;
            setTimeout(() => {
                isShaking.value = false;
            }, 600);
        }
    });
};

import { useRouter } from 'vue-router';
const router = useRouter();

const goToRegister = () => {
    router.push('/register');
};
const handleForgotPassword = () => {
    ElMessage.info('忘记密码功能正在开发中...');
};
</script>

<style scoped>
/* --- 整体布局 --- */
.login-container {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    background-color: #f0f2f5;
    /* 使用柔和的浅灰色背景 */
}

.main-content {
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 20px;
}

/* --- 登录卡片 --- */
.login-box {
    width: 100%;
    max-width: 380px;
    padding: 40px;
    background: #ffffff;
    border-radius: 16px;
    /* 更大的圆角 */
    box-shadow: 0 10px 30px -10px rgba(0, 0, 0, 0.1);
}

/* --- 头部 Logo --- */
.logo-section {
    text-align: center;
    margin-bottom: 30px;
}

.logo-section img {
    width: 60px;
}

.logo-section h3 {
    font-size: 24px;
    font-weight: 600;
    color: #303133;
    margin: 16px 0 8px;
}

.logo-section .subtitle {
    font-size: 14px;
    color: #909399;
    margin: 0;
}

/* --- Tabs 导航 --- */
.login-tabs {
    margin-bottom: 20px;
}

.tab-label {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 16px;
    font-weight: 500;
}

:deep(.el-tabs__header) {
    margin-bottom: 25px;
    /* 拉开和表单的距离 */
}

:deep(.el-tabs__nav-wrap::after) {
    height: 1px;
    /* 底部分割线变细 */
}

:deep(.el-tabs__active-bar) {
    height: 3px;
    /* 激活指示器加粗 */
    border-radius: 2px;
}

/* --- 表单与输入框核心对齐样式 --- */
.el-form-item {
    margin-bottom: 22px;
    /* 统一表单项间距 */
}

/* 使用 :deep() 穿透组件，修改内部样式 */
:deep(.el-input__wrapper) {
    height: 48px;
    /* 统一高度 */
    padding: 0 15px;
    border-radius: 8px;
    border: 1px solid #dcdfe6;
    box-shadow: none !important;
    /* 移除默认的 focus shadow */
    transition: border-color 0.2s;
}

:deep(.el-input__wrapper:hover) {
    border-color: #c0c4cc;
}

:deep(.el-input__wrapper.is-focus) {
    border-color: var(--el-color-primary);
}

/* [核心] 图标对齐样式 */
:deep(.el-input__prefix) {
    margin-right: 10px;
    color: #909399;
}

:deep(.el-input__prefix .el-icon) {
    font-size: 18px;
}

:deep(.el-input__inner) {
    font-size: 14px;
}

/* --- 验证码输入框与按钮对齐 --- */
.verification-code-wrapper {
    display: flex;
    align-items: center;
    /* 垂直居中对齐 */
    gap: 10px;
    width: 100%;
    /* <-- 就是这一行! */
}

.verification-code-wrapper .el-input {
    flex-grow: 1;
}

.send-code-btn {
    flex-shrink: 0;
    /* 防止按钮被压缩 */
    height: 48px;
    /* 与输入框高度一致 */
    border-radius: 8px;
    background-color: #e8f3ff;
    /* 图片中的浅蓝色 */
    color: var(--el-color-primary);
    border: none;
    font-weight: 500;
}

.send-code-btn:hover {
    background-color: #d9e9ff;
}

/* --- 底部链接与主按钮 --- */
.additional-links {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 25px;
    font-size: 14px;
}

.login-btn {
    width: 100%;
    height: 48px;
    font-size: 16px;
    border-radius: 8px;
    border: none;
    background: var(--el-color-primary);
    /* 使用 Element Plus 的主题色 */
}

/* --- 响应式 --- */
@media (max-width: 480px) {
    .login-box {
        padding: 24px;
    }
}

/* --- Animations --- */
.shake-animation {
    animation: shake 0.6s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
}

@keyframes shake {

    10%,
    90% {
        transform: translate3d(-1px, 0, 0);
    }

    20%,
    80% {
        transform: translate3d(2px, 0, 0);
    }

    30%,
    50%,
    70% {
        transform: translate3d(-4px, 0, 0);
    }

    40%,
    60% {
        transform: translate3d(4px, 0, 0);
    }
}
</style>