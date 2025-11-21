<template>
  <div class="reset-container">
    <div class="bg-circle circle-1"></div>
    <div class="bg-circle circle-2"></div>

    <main class="main-content">
      <div class="glass-box">
        <div class="header-section">
          <h2>重置密码</h2>
          <p>设置一个新的安全密码以保护您的账户</p>
        </div>

        <el-form
          ref="formRef"
          :model="formData"
          :rules="formRules"
          size="large"
          class="modern-form"
          @keyup.enter="handleResetPassword(formRef)"
        >
          <el-form-item prop="email">
            <el-input
              v-model="formData.email"
              placeholder="注册邮箱"
              class="custom-input"
            >
              <template #prefix><el-icon><Message /></el-icon></template>
            </el-input>
          </el-form-item>

          <el-form-item prop="code">
            <div class="code-flex">
              <el-input
                v-model="formData.code"
                placeholder="验证码"
                class="custom-input code-input"
                :maxlength="6"
              >
                <template #prefix><el-icon><Key /></el-icon></template>
              </el-input>
              <el-button
                class="send-code-btn"
                plain
                :disabled="isSendingCode || !isEmailValid"
                @click="sendVerificationCode"
              >
                {{ codeButtonText }}
              </el-button>
            </div>
          </el-form-item>

          <el-form-item prop="newPassword">
            <el-input
              v-model="formData.newPassword"
              type="password"
              placeholder="新密码 (至少6位)"
              show-password
              class="custom-input"
            >
              <template #prefix><el-icon><Lock /></el-icon></template>
            </el-input>
          </el-form-item>

          <el-form-item prop="confirmNewPassword">
            <el-input
              v-model="formData.confirmNewPassword"
              type="password"
              placeholder="确认新密码"
              show-password
              class="custom-input"
            >
              <template #prefix><el-icon><Lock /></el-icon></template>
            </el-input>
          </el-form-item>

          <el-button
            type="primary"
            class="submit-btn"
            :loading="isLoading"
            @click="handleResetPassword(formRef)"
          >
            确认重置
          </el-button>
        </el-form>

        <div class="footer-link">
          <el-link type="info" :underline="false" @click="goToLogin" class="back-link">
            <el-icon class="icon-arrow"><ArrowLeft /></el-icon> 返回登录
          </el-link>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts" name="ResetPassword">
import { ref, reactive, computed } from 'vue';
import { useRouter } from 'vue-router';
import { Message, Lock, Key, ArrowLeft } from '@element-plus/icons-vue'; // 新增 Key, ArrowLeft
import { ElMessage } from 'element-plus';
import type { FormInstance, FormRules } from 'element-plus';
import { GetResetPasswordCode, ResetPassword } from '@/api/resetPassword.ts';

/* ------------------ 逻辑保持原样 ------------------ */
const router = useRouter();
const formRef = ref<FormInstance>();
const isLoading = ref(false);

const formData = reactive({
  email: '',
  code: '',
  newPassword: '',
  confirmNewPassword: '',
});

const validatePass = (rule: any, value: string, callback: any) => {
  if (!value) {
    callback(new Error('请输入新密码'));
  } else if (value.length < 6) {
    callback(new Error('密码长度不能少于6位'));
  } else {
    if (formData.confirmNewPassword) {
      formRef.value?.validateField('confirmNewPassword');
    }
    callback();
  }
};

const validateConfirmPass = (rule: any, value: string, callback: any) => {
  if (!value) {
    callback(new Error('请再次输入密码'));
  } else if (value !== formData.newPassword) {
    callback(new Error('两次输入的密码不一致！'));
  } else {
    callback();
  }
};

const formRules = reactive<FormRules>({
  email: [
    { required: true, message: '请输入邮箱地址', trigger: 'blur' },
    { type: 'email', message: '邮箱地址格式不正确', trigger: ['blur', 'change'] },
  ],
  code: [
    { required: true, message: '请输入验证码', trigger: 'blur' },
    { len: 6, message: '验证码长度为6位', trigger: 'blur' },
  ],
  newPassword: [{ required: true, validator: validatePass, trigger: 'blur' }],
  confirmNewPassword: [{ required: true, validator: validateConfirmPass, trigger: 'blur' }],
});

const isSendingCode = ref(false);
const countdownTime = ref(60);
const codeButtonText = computed(() =>
  isSendingCode.value ? `${countdownTime.value}s` : '获取验证码',
);

const isEmailValid = computed(() =>
  /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(formData.email),
);

const sendVerificationCode = async () => {
  if (!isEmailValid.value) {
    ElMessage.error('请输入有效的邮箱地址！');
    return;
  }
  isSendingCode.value = true;
  try {
    const res = await GetResetPasswordCode(formData.email);
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
  countdownTime.value = 60;
  const timer = setInterval(() => {
    countdownTime.value--;
    if (countdownTime.value <= 0) {
      clearInterval(timer);
      isSendingCode.value = false;
    }
  }, 1000);
};

const handleResetPassword = (formEl?: FormInstance) => {
  if (!formEl || isLoading.value) return;
  formEl.validate(async (valid: boolean) => {
    if (valid) {
      isLoading.value = true;
      try {
        const res = await ResetPassword({
          email: formData.email,
          code: formData.code,
          newPassword: formData.newPassword,
        });
        if (res.code === 0) {
          ElMessage.success('密码重置成功，请使用新密码登录！');
          router.replace('/login');
        } else {
          ElMessage.error(res.message || '重置密码失败');
        }
      } catch (error) {
        console.error('ResetPassword request failed:', error);
        ElMessage.error('网络错误或服务器无响应');
      } finally {
        isLoading.value = false;
      }
    } else {
      ElMessage.error('请检查表单输入项！');
    }
  });
};

const goToLogin = () => {
  router.push('/login');
};
</script>

<style scoped>
.reset-container {
  position: relative;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(135deg, #f5f7fa 0%, #e4e9f2 100%);
  overflow: hidden;
  padding: 20px;
}

/* 背景装饰 */
.bg-circle {
  position: absolute;
  border-radius: 50%;
  filter: blur(60px);
  z-index: 0;
  opacity: 0.5;
}
.circle-1 { width: 300px; height: 300px; background: #a0cfff; top: -50px; left: -50px; }
.circle-2 { width: 400px; height: 400px; background: #d9ecff; bottom: -100px; right: -100px; }

.main-content {
  width: 100%;
  max-width: 440px;
  z-index: 1;
}

.glass-box {
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.6);
  border-radius: 24px;
  padding: 40px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.08);
}

.header-section {
  text-align: center;
  margin-bottom: 32px;
}
.header-section h2 {
  margin: 0 0 8px;
  font-size: 24px;
  color: #303133;
}
.header-section p {
  margin: 0;
  font-size: 14px;
  color: #909399;
}

.el-form-item { margin-bottom: 24px; }

/* 输入框样式统一 */
:deep(.el-input__wrapper) {
  background-color: #f5f7fa;
  box-shadow: none !important;
  border: 1px solid transparent;
  padding: 0 15px;
  height: 48px;
  border-radius: 12px;
  transition: all 0.3s;
}
:deep(.el-input__wrapper:hover) { background-color: #eef1f6; }
:deep(.el-input__wrapper.is-focus) {
  background-color: #fff;
  border-color: var(--el-color-primary);
  box-shadow: 0 0 0 4px rgba(64, 158, 255, 0.1) !important;
}

.code-flex { display: flex; gap: 12px; }
.send-code-btn {
  height: 48px;
  border-radius: 12px;
  background: #ecf5ff;
  border: none;
  color: var(--el-color-primary);
  font-weight: 600;
  padding: 0 20px;
}

.submit-btn {
  width: 100%;
  height: 48px;
  font-size: 16px;
  font-weight: 600;
  border-radius: 12px;
  box-shadow: 0 8px 16px rgba(64, 158, 255, 0.25);
  margin-top: 10px;
  transition: all 0.3s;
}
.submit-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 12px 24px rgba(64, 158, 255, 0.35);
}

.footer-link {
  margin-top: 24px;
  text-align: center;
}
.back-link {
  font-size: 14px;
  display: inline-flex;
  align-items: center;
  gap: 4px;
}
.icon-arrow { font-size: 14px; margin-right: 2px; }

@media(max-width: 480px) {
  .glass-box { padding: 30px 20px; }
}
</style>