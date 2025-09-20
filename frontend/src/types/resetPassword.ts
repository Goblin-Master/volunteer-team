// 忘记密码表单数据
// 这里的字段名与前端表单的 v-model 绑定
export interface ForgetPasswordForm {
  email_address: string;
  code: string;
  new_password: string;
  confirm_new_password: string;
}

// 校验规则（Element-Plus 风格）
import type { FormRules } from 'element-plus';
export type ForgetPasswordFormRules = FormRules<ForgetPasswordForm>;

// 发送给后端的重置密码请求体
// 只包含后端需要的字段，与表单数据解耦
export interface ResetPasswordReq {
  email: string;
  code: string;
  new_password: string;
}
