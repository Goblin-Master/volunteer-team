import req from './axios';
import type BaseResp from '@/types/base.ts';
import type { LoginResp, LoginModel } from '@/types/login.ts';
import { toLoginModel } from '@/types/login.ts';
import { mapOk } from '@/api/helpers';

export const Login = async (data: any): Promise<BaseResp<LoginModel | null>> => {
  const payload = 'loginType' in data ? { ...data, login_type: data.loginType } : data;
  const raw = (await req({ url: '/api/user/login', method: 'post', data: payload })) as unknown as BaseResp<LoginResp | null>;
  return mapOk(raw, toLoginModel);
};

export const GetLoginCode = (email: string): Promise<BaseResp<string>> =>
  req({ url: '/api/user/code/login', method: 'post', data: { email } });
