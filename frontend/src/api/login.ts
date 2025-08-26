import req from './axios'
import type BaseResp from '@/types/base.ts'
import type { LoginResp } from '@/types/login.ts'

export const Login = (data: any): Promise<BaseResp<LoginResp>> =>
  req({ url: '/api/login', method: 'post', data })