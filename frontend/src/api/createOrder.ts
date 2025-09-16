import req from './axios';
import type BaseResp from '@/types/base.ts';
import type { CreateOrderReq } from '@/types/createOrder.ts';

export const CreateOrder = (data: CreateOrderReq): Promise<BaseResp<string>> =>
  req({ url: '/api/order/create', method: 'post', data:data })