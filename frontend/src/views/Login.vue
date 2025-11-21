<template>
  <div class="login-container">
    <div class="bg-circle circle-1"></div>
    <div class="bg-circle circle-2"></div>

    <main class="main-content">
      <div class="login-box" :class="{ 'shake-animation': isShaking }">
        <div class="logo-section">
          <img src="@/assets/logo.png" alt="Logo" />
          <h3>欢迎回来</h3>
          <p class="subtitle">登录您的账户以继续</p>
        </div>

        <el-tabs v-model="activeTab" class="custom-tabs" stretch>
          <el-tab-pane name="account">
            <template #label>
              <span class="tab-inner">账号登录</span>
            </template>
            <el-form
              ref="accountFormRef"
              :model="accountForm"
              :rules="accountRules"
              size="large"
              @keyup.enter="handleLogin(accountFormRef)"
            >
              <el-form-item prop="account">
                <el-input 
                  v-model="accountForm.account" 
                  placeholder="请输入账号"
                  class="custom-input"
                >
                  <template #prefix><el-icon><User /></el-icon></template>
                </el-input>
              </el-form-item>
              <el-form-item prop="password">
                <el-input
                  v-model="accountForm.password"
                  type="password"
                  placeholder="请输入密码"
                  show-password
                  class="custom-input"
                >
                  <template #prefix><el-icon><Lock /></el-icon></template>
                </el-input>
              </el-form-item>
            </el-form>
          </el-tab-pane>

          <el-tab-pane name="email">
            <template #label>
              <span class="tab-inner">邮箱登录</span>
            </template>
            <el-form
              ref="emailFormRef"
              :model="emailForm"
              :rules="emailRules"
              size="large"
              @keyup.enter="handleLogin(emailFormRef)"
            >
              <el-form-item prop="email">
                <el-input 
                  v-model="emailForm.email" 
                  placeholder="请输入邮箱"
                  class="custom-input"
                >
                  <template #prefix><el-icon><Message /></el-icon></template>
                </el-input>
              </el-form-item>
              <el-form-item prop="code">
                <div class="code-flex">
                  <el-input 
                    v-model="emailForm.code" 
                    placeholder="验证码" 
                    class="custom-input code-input"
                  >
                    <template #prefix><el-icon><Key /></el-icon></template>
                  </el-input>
                  <el-button
                    class="code-btn"
                    :disabled="isSendingCode || !isEmailValid"
                    @click="sendVerificationCode"
                    plain
                  >
                    {{ codeButtonText }}
                  </el-button>
                </div>
              </el-form-item>
            </el-form>
          </el-tab-pane>
        </el-tabs>

        <el-button
          type="primary"
          size="large"
          class="login-btn"
          :loading="isLoading"
          @click="handleLogin(activeTab === 'account' ? accountFormRef : emailFormRef)"
        >
          立即登录
        </el-button>

        <div class="footer-links">
          <span class="text-gray">没有账号?</span>
          <el-link type="primary" :underline="false" @click="goToRegister">注册账号</el-link>
          <div class="divider"></div>
          <el-link type="info" :underline="false" @click="handleResetPassword">忘记密码</el-link>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts" name="Login">
/* 逻辑保持完全不变，仅为了兼容 icon 引入需确认 */
import { ref, reactive, computed } from 'vue';
import { User, Lock, Message, Key } from '@element-plus/icons-vue'; // 增加 Key icon
import { ElMessage } from 'element-plus';
import { Login, GetLoginCode } from '@/api/login.ts';
import type { LoginModel } from '@/types/login.ts';
import type BaseResp from '@/types/base.ts';
import type { FormInstance } from 'element-plus';
import { useUserStore } from '@/stores/user';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();

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

// Email Form
const emailFormRef = ref<FormInstance>();
const emailForm = reactive({
  email: '',
  code: '',
});
const emailRules = reactive({
  email: [
    { required: true, message: '请输入您的邮箱', trigger: 'blur' },
    {
      type: 'email',
      message: '请输入正确的邮箱格式',
      trigger: ['blur', 'change'],
    },
  ],
  code: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
});

const isSendingCode = ref(false);
const countdown = ref(60);
const codeButtonText = computed(() => {
  return isSendingCode.value ? `${countdown.value}s` : '获取验证码';
});
const isEmailValid = computed(() =>
  /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(emailForm.email),
);

const sendVerificationCode = async () => {
  if (!isEmailValid.value) {
    ElMessage.error('请输入有效的邮箱地址！');
    return;
  }
  isSendingCode.value = true;

  try {
    const res = await GetLoginCode(emailForm.email);
    if (res.code === 0) {
      ElMessage.success('验证码已发送至您的邮箱!');
      startCountdown();
    } else {
      ElMessage.error(res.message || '获取验证码失败!');
      isSendingCode.value = false;
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
            loginType: 'account',
          };
        } else {
          data = {
            email: emailForm.email,
            code: emailForm.code,
            loginType: 'email',
          };
        }

        const res: BaseResp<LoginModel> = await Login(data);

        if (res.code === 0) {
          ElMessage.success('登录成功!');
          useUserStore().setUserInfo(res.data);
          const redirectPath = (route.query.redirect as string) || '/';
          router.replace(redirectPath);
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
      isShaking.value = true;
      setTimeout(() => {
        isShaking.value = false;
      }, 600);
    }
  });
};

const goToRegister = () => router.push('/register');
const handleResetPassword = () => router.push('/resetPassword');
</script>

<style scoped>
.login-container {
  position: relative;
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #e4e9f2 100%);
  overflow: hidden;
}

/* 背景装饰圆 */
.bg-circle {
  position: absolute;
  border-radius: 50%;
  filter: blur(60px);
  z-index: 0;
  opacity: 0.6;
}
.circle-1 {
  width: 300px;
  height: 300px;
  background: #a0cfff;
  top: -50px;
  left: -50px;
}
.circle-2 {
  width: 400px;
  height: 400px;
  background: #d9ecff;
  bottom: -100px;
  right: -100px;
}

.main-content {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
  z-index: 1;
}

.login-box {
  width: 100%;
  max-width: 420px;
  padding: 48px 40px;
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.6);
  border-radius: 24px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.08);
  transition: transform 0.3s;
}

.logo-section {
  text-align: center;
  margin-bottom: 40px;
}
.logo-section img {
  width: 64px;
  margin-bottom: 16px;
}
.logo-section h3 {
  font-size: 26px;
  font-weight: 700;
  color: #1a1a1a;
  margin: 0 0 8px;
  letter-spacing: -0.5px;
}
.logo-section .subtitle {
  font-size: 14px;
  color: #909399;
}

/* Tabs 优化 */
:deep(.el-tabs__nav-wrap::after) {
  height: 2px;
  background-color: #f0f2f5;
}
:deep(.el-tabs__active-bar) {
  height: 3px;
  border-radius: 3px;
  background-color: var(--el-color-primary);
}
:deep(.el-tabs__item) {
  font-size: 16px;
  color: #606266;
  padding-bottom: 12px;
}
:deep(.el-tabs__item.is-active) {
  font-weight: 600;
  color: var(--el-color-primary);
}

/* 输入框深度定制 */
.el-form-item {
  margin-bottom: 24px;
}

:deep(.el-input__wrapper) {
  background-color: #f5f7fa; /* 灰色背景 */
  box-shadow: none !important;
  border: 1px solid transparent;
  padding: 8px 15px;
  transition: all 0.3s ease;
}

:deep(.el-input__wrapper:hover) {
  background-color: #eef1f6;
}

:deep(.el-input__wrapper.is-focus) {
  background-color: #fff;
  border-color: var(--el-color-primary);
  box-shadow: 0 0 0 4px rgba(64, 158, 255, 0.1) !important;
}

:deep(.el-input__inner) {
  height: 32px;
  font-weight: 500;
}

/* 验证码区域 */
.code-flex {
  display: flex;
  gap: 12px;
}
.code-btn {
  height: auto;
  padding: 0 20px;
  border-radius: 12px;
  background: #ecf5ff;
  border-color: #d9ecff;
  color: var(--el-color-primary);
  font-weight: 600;
}
.code-btn:hover {
  background: var(--el-color-primary);
  color: #fff;
}

/* 登录按钮 */
.login-btn {
  width: 100%;
  height: 52px;
  font-size: 18px;
  font-weight: 600;
  letter-spacing: 1px;
  margin-top: 10px;
  box-shadow: 0 10px 20px rgba(64, 158, 255, 0.3);
  transition: all 0.3s;
}
.login-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 15px 25px rgba(64, 158, 255, 0.4);
}

/* 底部链接 */
.footer-links {
  margin-top: 30px;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 14px;
  gap: 8px;
}
.text-gray {
  color: #909399;
}
.divider {
  width: 1px;
  height: 12px;
  background: #dcdfe6;
  margin: 0 8px;
}

/* 动画 */
.shake-animation {
  animation: shake 0.5s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
}
@keyframes shake {
  10%, 90% { transform: translate3d(-1px, 0, 0); }
  20%, 80% { transform: translate3d(2px, 0, 0); }
  30%, 50%, 70% { transform: translate3d(-4px, 0, 0); }
  40%, 60% { transform: translate3d(4px, 0, 0); }
}

/* 响应式 */
@media (max-width: 480px) {
  .login-box {
    padding: 32px 24px;
  }
}
</style>