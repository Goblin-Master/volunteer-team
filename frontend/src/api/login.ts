import req from './axios';
import type BaseResp from '@/types/base.ts';
import type { LoginResp, LoginModel } from '@/types/login.ts';
import { toLoginModel } from '@/types/login.ts';

export const Login = async (data: any): Promise<BaseResp<LoginModel>> => {
  const payload = 'loginType' in data ? { ...data, login_type: data.loginType } : data;
  const raw = (await req({ url: '/api/user/login', method: 'post', data: payload })) as unknown as BaseResp<LoginResp>;
  const model = toLoginModel(raw.data as LoginResp);
  return { code: raw.code, message: raw.message, data: model } as BaseResp<LoginModel>;
};

export const GetLoginCode = (email: string): Promise<BaseResp<string>> =>
  req({ url: '/api/user/code/login', method: 'post', data: { email } });
