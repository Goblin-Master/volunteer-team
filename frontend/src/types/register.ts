// 注册表单数据
export interface RegisterForm {
  email: string;
  code: string;
  username: string;
  account: string;
  password: string;
  confirm_password: string;
  is_internal: boolean;   // 是否内部员工
  agree_terms: boolean;   // 是否同意协议
}

// 校验规则（Element-Plus 风格，可换其他库）
import type { FormRules } from 'element-plus';
export type RegisterFormRules = FormRules<RegisterForm>;

// 真正发给后端的 DTO（可以比表单少字段）
export interface RegisterReq {
  email: string;
  code: string;
  username: string;
  account: string;
  password: string;
  is_internal: boolean;
}