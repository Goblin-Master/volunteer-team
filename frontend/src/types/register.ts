// 校验规则（Element-Plus 风格，可换其他库）
import type { FormRules } from 'element-plus';
export type RegisterFormRules = FormRules<RegisterModel>;

// 真正发给后端的 DTO（可以比表单少字段）
export interface RegisterReq {
  email: string;
  code: string;
  username: string;
  account: string;
  password: string;
  is_internal: boolean;
}

export interface RegisterModel {
  email: string;
  code: string;
  username: string;
  account: string;
  password: string;
  confirmPassword?: string;
  isInternal: boolean;
  agreeTerms?: boolean;
}

export const toRegisterReq = (x: RegisterModel): RegisterReq => ({
  email: x.email,
  code: x.code,
  username: x.username,
  account: x.account,
  password: x.password,
  is_internal: x.isInternal,
});


