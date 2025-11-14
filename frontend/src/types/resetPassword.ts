// 校验规则（Element-Plus 风格）
import type { FormRules } from 'element-plus';
export type ForgetPasswordFormRules = FormRules<ForgetPasswordModel>;

// 发送给后端的重置密码请求体
// 只包含后端需要的字段，与表单数据解耦
export interface ResetPasswordReq {
  email: string;
  code: string;
  new_password: string;
}

export interface ForgetPasswordModel {
  email: string;
  code: string;
  newPassword: string;
  confirmNewPassword?: string;
}


export const toResetPasswordReq = (x: ForgetPasswordModel): ResetPasswordReq => ({
  email: x.email,
  code: x.code,
  new_password: x.newPassword,
});
