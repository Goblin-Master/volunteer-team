import req from './axios'
import type BaseResp from '@/types/base.ts'
import type { LoginResp } from '@/types/login.ts'

export const Login = (data: any): Promise<BaseResp<LoginResp>> =>
  req({ url: '/api/user/login', method: 'post', data })

export const GetLoginCode = (email: string): Promise<BaseResp<string>> =>
  req({
    url: '/api/user/code/login',
    method: 'post',
    data: {email} // ✅ 作为请求体字段
  })