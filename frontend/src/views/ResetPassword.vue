<template>
  <div class="forget-password-container">
    <main class="main-content">
      <el-card class="password-card" shadow="never">
        <div class="card-header">
          <h2>忘记密码</h2>
          <p>通过邮箱找回您的密码</p>
        </div>

        <el-form
          ref="form_ref"
          :model="form_data"
          :rules="form_rules"
          size="large"
          @keyup.enter="handleResetPassword(form_ref)"
        >
          <el-form-item prop="email_address">
            <el-input
              v-model="form_data.email_address"
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
                v-model="form_data.code"
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
                :disabled="is_sending_code || !is_email_valid"
                @click="sendVerificationCode"
              >
                {{ code_button_text }}
              </el-button>
            </div>
          </el-form-item>

          <el-form-item prop="new_password">
            <el-input
              v-model="form_data.new_password"
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

          <el-form-item prop="confirm_new_password">
            <el-input
              v-model="form_data.confirm_new_password"
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
              :loading="is_loading"
              @click="handleResetPassword(form_ref)"
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
import { ref, reactive, computed } from "vue";
import { useRouter } from "vue-router";
import { Message, Lock } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import type { FormInstance, FormRules } from "element-plus";
import { GetResetPasswordCode, ResetPassword } from "@/api/resetPassword.ts"; // 确保路径正确

const router = useRouter();
const form_ref = ref<FormInstance>();
const is_loading = ref(false);

const form_data = reactive({
  email_address: "",
  code: "",
  new_password: "",
  confirm_new_password: "",
});

// 验证规则
const validatePass = (rule: any, value: string, callback: any) => {
  if (!value) {
    callback(new Error("请输入新密码"));
  } else if (value.length < 6) {
    callback(new Error("密码长度不能少于6位"));
  } else {
    if (form_data.confirm_new_password) {
      form_ref.value?.validateField("confirm_new_password");
    }
    callback();
  }
};

const validateConfirmPass = (rule: any, value: string, callback: any) => {
  if (!value) {
    callback(new Error("请再次输入密码"));
  } else if (value !== form_data.new_password) {
    callback(new Error("两次输入的密码不一致！"));
  } else {
    callback();
  }
};

const form_rules = reactive<FormRules>({
  email_address: [
    { required: true, message: "请输入邮箱地址", trigger: "blur" },
    {
      type: "email",
      message: "邮箱地址格式不正确",
      trigger: ["blur", "change"],
    },
  ],
  code: [
    { required: true, message: "请输入验证码", trigger: "blur" },
    { len: 6, message: "验证码长度为6位", trigger: "blur" },
  ],
  new_password: [{ required: true, validator: validatePass, trigger: "blur" }],
  confirm_new_password: [
    { required: true, validator: validateConfirmPass, trigger: "blur" },
  ],
});

// 验证码发送逻辑
const is_sending_code = ref(false);
const countdown_time = ref(60);
const code_button_text = computed(() =>
  is_sending_code.value ? `${countdown_time.value}秒后重试` : "发送验证码"
);

const is_email_valid = computed(() =>
  /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(
    form_data.email_address
  )
);

const sendVerificationCode = async () => {
  if (!is_email_valid.value) {
    ElMessage.error("请输入有效的邮箱地址！");
    return;
  }
  is_sending_code.value = true;
  try {
    const res = await GetResetPasswordCode(form_data.email_address);
    if (res.code === 0) {
      ElMessage.success("验证码已发送至您的邮箱!");
      startCountdown();
    } else {
      ElMessage.error(res.message || "获取验证码失败!");
      is_sending_code.value = false;
    }
  } catch (error) {
    console.error("GetCode request failed:", error);
    ElMessage.error("网络错误或服务器无响应");
    is_sending_code.value = false;
  }
};

const startCountdown = () => {
  countdown_time.value = 60;
  const timer = setInterval(() => {
    countdown_time.value--;
    if (countdown_time.value <= 0) {
      clearInterval(timer);
      is_sending_code.value = false;
    }
  }, 1000);
};

// 重置密码逻辑
const handleResetPassword = (form_el?: FormInstance) => {
  if (!form_el || is_loading.value) return;

  form_el.validate(async (valid: boolean) => {
    if (valid) {
      is_loading.value = true;
      try {
        const data = {
          email: form_data.email_address,
          code: form_data.code,
          new_password: form_data.new_password,
        };
        const res = await ResetPassword(data);
        if (res.code === 0) {
          ElMessage.success("密码重置成功，请使用新密码登录！");
          router.replace("/login");
        } else {
          ElMessage.error(res.message || "重置密码失败");
        }
      } catch (error) {
        console.error("ResetPassword request failed:", error);
        ElMessage.error("网络错误或服务器无响应");
      } finally {
        is_loading.value = false;
      }
    } else {
      ElMessage.error("请检查表单输入项！");
    }
  });
};

// 返回登录页面
const goToLogin = () => {
  router.push("/login");
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
