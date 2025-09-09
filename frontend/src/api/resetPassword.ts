import req from './axios';
import type BaseResp from '@/types/base.ts';
import type { ResetPasswordReq } from '@/types/resetPassword.ts';

export const ResetPassword = (data: ResetPasswordReq): Promise<BaseResp<string>> =>
  req({ url: '/api/user/resetPassword', method: 'post', data })

export const GetResetPasswordCode = (email: string): Promise<BaseResp<string>> =>
  req({
    url: '/api/user/code/reset', method: 'post',
    data: { email } // ✅ 作为请求体字段
  })