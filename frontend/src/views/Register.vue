<template>
  <div class="register-container">
    <div class="bg-circle circle-1"></div>
    <div class="bg-circle circle-2"></div>

    <main class="main-content">
      <div class="register-box">
        <div class="header-section">
          <h2>创建账户</h2>
          <p>加入我们，开始您的体验</p>
        </div>

        <el-form
          ref="registerFormRef"
          :model="registerForm"
          :rules="registerRules"
          size="large"
          class="modern-form"
          @submit.prevent="handleRegister"
        >
          <el-form-item prop="email">
            <el-input
              v-model="registerForm.email"
              placeholder="邮箱地址"
              class="custom-input"
            >
              <template #prefix><el-icon><Message /></el-icon></template>
            </el-input>
          </el-form-item>

          <el-form-item prop="code">
            <div class="code-flex">
              <el-input
                v-model="registerForm.code"
                placeholder="验证码"
                class="custom-input code-input"
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

          <el-form-item prop="username">
            <el-input
              v-model="registerForm.username"
              placeholder="用户昵称"
              class="custom-input"
            >
              <template #prefix><el-icon><User /></el-icon></template>
            </el-input>
          </el-form-item>

          <el-form-item prop="account">
            <el-input
              v-model="registerForm.account"
              placeholder="设置账号 (3-15位)"
              class="custom-input"
            >
              <template #prefix><el-icon><Postcard /></el-icon></template>
            </el-input>
          </el-form-item>

          <el-row :gutter="16">
            <el-col :span="12">
              <el-form-item prop="password">
                <el-input
                  v-model="registerForm.password"
                  type="password"
                  placeholder="设置密码"
                  show-password
                  class="custom-input"
                >
                  <template #prefix><el-icon><Lock /></el-icon></template>
                </el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item prop="confirmPassword">
                <el-input
                  v-model="registerForm.confirmPassword"
                  type="password"
                  placeholder="确认密码"
                  show-password
                  class="custom-input"
                >
                   <template #prefix><el-icon><Lock /></el-icon></template>
                </el-input>
              </el-form-item>
            </el-col>
          </el-row>

          <div class="switch-row">
            <span class="switch-label">我是内部工作人员</span>
            <el-switch
              v-model="registerForm.isInternal"
              @change="handleInternalSwitchChange"
              style="--el-switch-on-color: var(--el-color-primary)"
            />
          </div>

          <el-form-item prop="agreeTerms" class="terms-item">
            <el-checkbox v-model="registerForm.agreeTerms">
              阅读并同意 <el-link type="primary" :underline="false">用户协议</el-link> 与 <el-link type="primary" :underline="false">隐私政策</el-link>
            </el-checkbox>
          </el-form-item>

          <el-button
            type="primary"
            class="submit-btn"
            :loading="loading"
            @click="handleRegister"
          >
            立即注册
          </el-button>
        </el-form>

        <div class="footer-link">
          <span>已有账号？</span>
          <el-link type="primary" :underline="false" @click="goToLogin" class="link-text">直接登录</el-link>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
/* 逻辑保持完全一致，仅补充图标引入 */
import { ref, reactive, computed } from 'vue';
import { useRouter } from 'vue-router';
import type { FormInstance } from 'element-plus';
import { ElMessage } from 'element-plus';
import { Message, Lock, User, Key, Postcard } from '@element-plus/icons-vue'; // 新增 Key, Postcard
import { Register, GetRegisterCode } from '@/api/register';
import type { RegisterModel, RegisterFormRules } from '@/types/register';

/* ---------- (以下逻辑代码未变，复制即可) ---------- */
const registerForm = reactive<RegisterModel>({
  email: '',
  code: '',
  username: '',
  account: '',
  password: '',
  confirmPassword: '',
  isInternal: false,
  agreeTerms: false,
});

const registerFormRef = ref<FormInstance>();
const router = useRouter();
const loading = ref(false);

const isSendingCode = ref(false);
const countdown = ref(60);
const codeButtonText = computed(() =>
  isSendingCode.value ? `${countdown.value}s` : '获取验证码',
);

const isEmailValid = computed(() =>
  /^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$/.test(registerForm.email),
);

const sendVerificationCode = async () => {
  if (!isEmailValid.value) {
    ElMessage.error('请输入有效的邮箱地址！');
    return;
  }
  isSendingCode.value = true;
  try {
    const res = await GetRegisterCode(registerForm.email);
    if (res.code === 0) {
      ElMessage.success('验证码已发送至您的邮箱！');
      startCountdown();
    } else {
      ElMessage.error(res.message || '获取验证码失败！');
      isSendingCode.value = false;
    }
  } catch (e) {
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

const validatePass = (rule: any, value: string, callback: any) => {
  if (!value) callback(new Error('请输入密码'));
  else {
    if (registerForm.confirmPassword) {
      registerFormRef.value?.validateField('confirmPassword');
    }
    callback();
  }
};

const validatePass2 = (rule: any, value: string, callback: any) => {
  if (!value) callback(new Error('请再次输入密码'));
  else if (value !== registerForm.password) {
    callback(new Error('两次输入的密码不一致！'));
  } else callback();
};

const validateAgreement = (rule: any, value: boolean, callback: any) => {
  value ? callback() : callback(new Error('请同意协议'));
};

const registerRules = reactive<RegisterFormRules>({
  username: [
    { required: true, message: '请输入用户昵称', trigger: 'blur' },
    { min: 1, max: 14, message: '昵称长度 1-14 个字符', trigger: 'blur' },
  ],
  email: [
    { required: true, message: '请输入邮箱地址', trigger: 'blur' },
    {
      type: 'email',
      message: '邮箱地址格式不正确',
      trigger: ['blur', 'change'],
    },
  ],
  code: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
  account: [
    { required: true, message: '请输入账号', trigger: 'blur' },
    { min: 3, max: 15, message: '长度在 3 到 15 个字符', trigger: 'blur' },
  ],
  password: [
    { validator: validatePass, trigger: 'blur', required: true },
    { min: 6, message: '不少于6位', trigger: 'blur' },
  ],
  confirmPassword: [
    { validator: validatePass2, trigger: 'blur', required: true },
  ],
  agreeTerms: [{ validator: validateAgreement, trigger: 'change' }],
});

const handleRegister = async () => {
  if (!registerFormRef.value) return;
  const valid = await registerFormRef.value.validate().catch(() => false);
  if (!valid) return;

  loading.value = true;
  try {
    const resp = await Register(registerForm);
    if (resp.code === 0) {
      ElMessage.success('注册成功！');
      router.replace('/login');
    } else {
      ElMessage.error(resp.message || '注册失败');
    }
  } catch (err: any) {
    ElMessage.error(err.message || '网络异常');
  } finally {
    loading.value = false;
  }
};

const handleInternalSwitchChange = (val: boolean | string | number) => {
  if (val === true) {
    ElMessage.info({
      message: '内部账号注册后需要管理员审核。',
      duration: 4000,
    });
  }
};

const goToLogin = () => router.push('/login');
</script>

<style scoped>
.register-container {
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
  max-width: 460px;
  z-index: 1;
}

.register-box {
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.6);
  border-radius: 24px;
  padding: 40px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.08);
}

.header-section {
  text-align: center;
  margin-bottom: 30px;
}
.header-section h2 {
  margin: 0;
  font-size: 24px;
  color: #303133;
}
.header-section p {
  margin: 8px 0 0;
  color: #909399;
  font-size: 14px;
}

/* 输入框样式复用 */
.el-form-item { margin-bottom: 20px; }

:deep(.el-input__wrapper) {
  background-color: #f5f7fa;
  box-shadow: none !important;
  border: 1px solid transparent;
  padding: 0 15px;
  height: 44px;
  border-radius: 12px;
  transition: all 0.3s;
}
:deep(.el-input__wrapper:hover) { background-color: #eef1f6; }
:deep(.el-input__wrapper.is-focus) {
  background-color: #fff;
  border-color: var(--el-color-primary);
  box-shadow: 0 0 0 4px rgba(64, 158, 255, 0.1) !important;
}

/* 验证码行 */
.code-flex { display: flex; gap: 12px; }
.send-code-btn {
  height: 44px;
  border-radius: 12px;
  background: #ecf5ff;
  border: none;
  color: var(--el-color-primary);
  font-weight: 600;
  padding: 0 20px;
}

/* 内部人员开关 */
.switch-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding: 0 4px;
}
.switch-label { font-size: 14px; color: #606266; }

/* 提交按钮 */
.submit-btn {
  width: 100%;
  height: 48px;
  font-size: 16px;
  font-weight: 600;
  border-radius: 12px;
  box-shadow: 0 8px 16px rgba(64, 158, 255, 0.25);
  margin-top: 8px;
}
.submit-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 12px 20px rgba(64, 158, 255, 0.35);
}

.footer-link {
  margin-top: 24px;
  text-align: center;
  font-size: 14px;
  color: #909399;
}
.link-text { font-weight: 600; margin-left: 4px; }

@media(max-width: 480px) {
  .register-box { padding: 30px 20px; }
}
</style>