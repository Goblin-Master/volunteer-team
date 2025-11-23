<!-- ResetPassword.vue -->
<template>
  <div class="reset-container">
    <BaseCard class="reset-box" shadow="always">
      <div class="header-section">
        <img src="@/assets/logo.png" alt="Logo" />
        <h2>重置密码</h2>
        <p class="subtitle">设置一个新的安全密码以保护您的账户</p>
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
              maxlength="6"
              class="custom-input code-input"
            >
              <template #prefix><el-icon><Key /></el-icon></template>
            </el-input>
            <el-button
              class="code-btn"
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
            placeholder="新密码 (≥6位)"
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
          class="reset-btn"
          :loading="isLoading"
          @click="handleResetPassword(formRef)"
        >
          确认重置
        </el-button>
      </el-form>

      <div class="footer-links">
        <el-link type="info" :underline="false" @click="goToLogin">
          返回登录
        </el-link>
      </div>
    </BaseCard>
  </div>
</template>

<script setup lang="ts" name="ResetPassword">
import { ref, reactive, computed } from 'vue';
import { useRouter } from 'vue-router';
import { Message, Lock, Key } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';
import type { FormInstance, FormRules } from 'element-plus';
import { GetResetPasswordCode, ResetPassword } from '@/api/resetPassword';

/* ------------------ 响应数据 ------------------ */
const router = useRouter();
const formRef = ref<FormInstance>();
const isLoading = ref(false);

const formData = reactive({
  email: '',
  code: '',
  newPassword: '',
  confirmNewPassword: '',
});

/* ------------------ 校验规则 ------------------ */
const validatePass = (rule: any, value: string, callback: any) => {
  if (!value) callback(new Error('请输入新密码'));
  else if (value.length < 6) callback(new Error('密码长度不能少于6位'));
  else {
    if (formData.confirmNewPassword) formRef.value?.validateField('confirmNewPassword');
    callback();
  }
};
const validateConfirmPass = (rule: any, value: string, callback: any) => {
  if (!value) callback(new Error('请再次输入密码'));
  else if (value !== formData.newPassword) callback(new Error('两次输入的密码不一致！'));
  else callback();
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

/* ------------------ 验证码倒计时 ------------------ */
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

/* ------------------ 提交重置 ------------------ */
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

const goToLogin = () => router.push('/login');
</script>

<style scoped>
/* ---------- 容器 ---------- */
.reset-container{
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: 24px;
  overflow: hidden;
  background-attachment: fixed;
  background:
    radial-gradient(2200px circle at -15% 40%, rgba(64,158,255,.18) 0%, rgba(64,158,255,0) 65%),
    radial-gradient(2200px circle at 115% 60%, rgba(72,217,151,.16) 0%, rgba(72,217,151,0) 65%),
    linear-gradient(135deg, #eef3fb 0%, #f2f6ff 50%, #eef2f9 100%);
}
.reset-container::before{
  content:"";
  position:absolute;
  inset:-12% -12% -12% -12%;
  background: radial-gradient(1200px circle at 50% -10%, rgba(255,214,102,.12) 0%, rgba(255,214,102,0) 70%);
  filter: blur(18px);
  pointer-events:none;
}
.reset-container::after{
  content:"";
  position:absolute;
  inset:0;
  background: radial-gradient(80% 120% at 50% 50%, rgba(255,255,255,.08) 0%, rgba(255,255,255,0) 60%);
  pointer-events:none;
}

/* ---------- 卡片 ---------- */
.reset-box{
  width: 100%;
  max-width: 420px;
  padding: 40px 36px;
  border-radius: 12px;
}

/* ---------- 头部 ---------- */
.header-section{
  text-align: center;
  margin-bottom: 32px;
}
.header-section img{
  width: 64px;
  margin-bottom: 12px;
}
.header-section h2{
  font-size: 26px;
  font-weight: 700;
  color: #1a1a1a;
  margin: 0 0 6px;
  letter-spacing: -0.5px;
}
.subtitle{
  font-size: 14px;
  color: #909399;
}

/* ---------- 表单 ---------- */
.el-form-item{
  margin-bottom: 24px;
}

/* 输入框深度样式 – 直接复用 Login.vue */
:deep(.el-input__wrapper){
  background-color: #f5f7fa;
  box-shadow: none !important;
  border: 1px solid #e4e7ed;
  padding: 8px 15px;
  transition: all .2s ease;
}
:deep(.el-input__wrapper:hover){
  background-color: #eef1f6;
}
:deep(.el-input__wrapper.is-focus){
  background-color: #fff;
  border-color: #c0c4cc;
  box-shadow: 0 0 0 3px rgba(0,0,0,.06) !important;
}
:deep(.el-input__inner){
  height: 32px;
  font-weight: 500;
}

/* ---------- 验证码 ---------- */
.code-flex{
  display: flex;
  gap: 12px;
}
.code-btn{
  height: auto;
  padding: 0 20px;
  border-radius: 12px;
  background: linear-gradient(180deg,#ecf5ff 0%,#e7f0ff 100%);
  border: 1px solid #d9ecff;
  color: var(--el-color-primary);
  font-weight: 600;
  box-shadow: 0 2px 6px rgba(64,158,255,.12);
}
.code-btn:hover{
  background: var(--el-color-primary);
  color: #fff;
}

/* ---------- 按钮 ---------- */
.reset-btn{
  width: 100%;
  height: 52px;
  font-size: 18px;
  font-weight: 600;
  letter-spacing: 1px;
  margin-top: 10px;
  box-shadow: 0 10px 20px rgba(64,158,255,.3);
  transition: all .3s;
}
.reset-btn:hover{
  transform: translateY(-2px);
  box-shadow: 0 15px 25px rgba(64,158,255,.4);
}

/* ---------- 底部 ---------- */
.footer-links{
  margin-top: 24px;
  text-align: center;
  font-size: 14px;
}

/* ---------- 响应 ---------- */
@media (max-width: 480px){
  .reset-box{
    padding: 32px 24px;
  }
}
</style>