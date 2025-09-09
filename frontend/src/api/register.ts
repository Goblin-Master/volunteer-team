import req from './axios';
import type BaseResp from '@/types/base.ts';
import type { RegisterReq } from '@/types/register.ts';

export const Register = (data: RegisterReq): Promise<BaseResp<string>> =>
  req({ url: '/api/user/register', method: 'post', data })

export const GetRegisterCode = (email: string): Promise<BaseResp<string>> =>
  req({
    url: '/api/user/code/register',
    method: 'post',
    data: { email } // ✅ 作为请求体字段
  })