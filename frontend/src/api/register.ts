import req from './axios';
import type BaseResp from '@/types/base.ts';
import type { RegisterModel, RegisterReq } from '@/types/register.ts';
import { toRegisterReq } from '@/types/register.ts';

export const Register = (data: RegisterModel): Promise<BaseResp<string>> => {
  const dto: RegisterReq = toRegisterReq(data);
  return req({ url: '/api/user/register', method: 'post', data: dto });
};

export const GetRegisterCode = (email: string): Promise<BaseResp<string>> =>
  req({ url: '/api/user/code/register', method: 'post', data: { email } });
