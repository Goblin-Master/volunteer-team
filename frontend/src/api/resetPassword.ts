import req from './axios';
import type BaseResp from '@/types/base.ts';
import type { ForgetPasswordModel, ResetPasswordReq } from '@/types/resetPassword.ts';
import { toResetPasswordReq } from '@/types/resetPassword.ts';

export const ResetPassword = (data: ForgetPasswordModel | ResetPasswordReq): Promise<BaseResp<string>> => {
  const dto: ResetPasswordReq = 'newPassword' in data ? toResetPasswordReq(data as ForgetPasswordModel) : (data as ResetPasswordReq);
  return req({ url: '/api/user/resetPassword', method: 'post', data: dto });
};

export const GetResetPasswordCode = (email: string): Promise<BaseResp<string>> =>
  req({ url: '/api/user/code/reset', method: 'post', data: { email } });
