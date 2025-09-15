import req from './axios';
import type BaseResp from '@/types/base.ts';
import type { OrderListResp } from '@/types/order.ts';

export const GetOrderList = (): Promise<BaseResp<OrderListResp>> =>
  req({ url: '/api/order/list', method: 'get' })