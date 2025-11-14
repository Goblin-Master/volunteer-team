<template>
  <div class="forget-password-container">
    <main class="main-content">
      <el-card class="password-card" shadow="never">
        <div class="card-header">
          <h2>忘记密码</h2>
          <p>通过邮箱找回您的密码</p>
        </div>

        <el-form
          ref="formRef"
          :model="formData"
          :rules="formRules"
          size="large"
          @keyup.enter="handleResetPassword(formRef)"
        >
          <el-form-item prop="email">
            <el-input
              v-model="formData.email"
              placeholder="请输入您注册时使用的邮箱地址"
              clearable
            >
              <template #prefix>
                <el-icon>
                  <Message />
                </el-icon>
              </template>
            </el-input>
          </el-form-item>

          <el-form-item prop="code">
            <div class="verification-code-wrapper">
              <el-input
                v-model="formData.code"
                placeholder="请输入收到的验证码"
                clearable
                :maxlength="6"
              >
                <template #prefix>
                  <el-icon>
                    <Message />
                  </el-icon>
                </template>
              </el-input>
              <el-button
                class="send-code-btn"
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
              placeholder="请输入新密码"
              show-password
            >
              <template #prefix>
                <el-icon>
                  <Lock />
                </el-icon>
              </template>
            </el-input>
          </el-form-item>

          <el-form-item prop="confirmNewPassword">
            <el-input
              v-model="formData.confirmNewPassword"
              type="password"
              placeholder="请确认新密码"
              show-password
            >
              <template #prefix>
                <el-icon>
                  <Lock />
                </el-icon>
              </template>
            </el-input>
          </el-form-item>

          <el-form-item>
            <el-button
              type="primary"
              class="reset-password-btn"
              :loading="isLoading"
              @click="handleResetPassword(formRef)"
            >
              重置密码
            </el-button>
          </el-form-item>
        </el-form>

        <div class="additional-links">
          <el-link type="primary" :underline="false" @click="goToLogin"
            >记起密码了？返回登录</el-link
          >
        </div>
      </el-card>
    </main>
  </div>
</template>

<script setup lang="ts" name="ResetPassword">
import { ref, reactive, computed } from 'vue';
import { useRouter } from 'vue-router';
import { Message, Lock } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';
import type { FormInstance, FormRules } from 'element-plus';
import { GetResetPasswordCode, ResetPassword } from '@/api/resetPassword.ts'; // 确保路径正确

const router = useRouter();
const formRef = ref<FormInstance>();
const isLoading = ref(false);

const formData = reactive({
  email: '',
  code: '',
  newPassword: '',
  confirmNewPassword: '',
});

// 验证规则
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
    {
      type: 'email',
      message: '邮箱地址格式不正确',
      trigger: ['blur', 'change'],
    },
  ],
  code: [
    { required: true, message: '请输入验证码', trigger: 'blur' },
    { len: 6, message: '验证码长度为6位', trigger: 'blur' },
  ],
  newPassword: [{ required: true, validator: validatePass, trigger: 'blur' }],
  confirmNewPassword: [
    { required: true, validator: validateConfirmPass, trigger: 'blur' },
  ],
});

// 验证码发送逻辑
const isSendingCode = ref(false);
const countdownTime = ref(60);
const codeButtonText = computed(() =>
  isSendingCode.value ? `${countdownTime.value}秒后重试` : '发送验证码',
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

// 重置密码逻辑
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

// 返回登录页面
const goToLogin = () => {
  router.push('/login');
};
</script>

<style scoped>
/* 样式与登录和注册页面保持一致 */
.forget-password-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  background-color: #f0f2f5;
}

.main-content {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
}

.password-card {
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

.card-header p {
  font-size: 14px;
  color: #909399;
  margin: 0;
}

.el-form-item {
  margin-bottom: 22px;
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

:deep(.el-input__prefix) {
  margin-right: 10px;
  color: #909399;
}

:deep(.el-input__prefix .el-icon) {
  font-size: 18px;
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

.reset-password-btn {
  width: 100%;
  height: 48px;
  font-size: 16px;
  border-radius: 8px;
  border: none;
  background: var(--el-color-primary);
}

.additional-links {
  margin-top: 24px;
  text-align: center;
  font-size: 14px;
}

@media (max-width: 480px) {
  .password-card {
    padding: 24px;
  }
}
</style>
