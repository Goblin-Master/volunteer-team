import req from './axios';
import type BaseResp from '@/types/base';
import type { OrderListResp } from '@/types/order.ts';
import type { OrderDetailResp } from '@/types/order.ts';
import type { CreateOrderReq } from '@/types/order.ts';

export const GetOrderDetail = (
  order_id: string,
): Promise<BaseResp<OrderDetailResp>> =>
  req({
    url: '/api/order/detail',
    method: 'get',
    params: { order_id },
  });

export const GetOrderList = (): Promise<BaseResp<OrderListResp>> =>
  req({ url: '/api/order/list', method: 'get' });

export const CreateOrder = (data: CreateOrderReq): Promise<BaseResp<string>> =>
  req({ url: '/api/order/create', method: 'post', data: data });
