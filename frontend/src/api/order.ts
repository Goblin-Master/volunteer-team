import req from './axios';
import type BaseResp from '@/types/base';
import type { OrderItem, OrderDetailResp, CreateOrderReq, OrderItemModel, OrderDetailModel, CreateOrderItem } from '@/types/order.ts';
import { toOrderItemModel, toOrderDetailModel, toCreateOrderReq } from '@/types/order.ts';

export const GetOrderDetail = async (orderID: string): Promise<BaseResp<OrderDetailModel>> => {
  const raw = (await req({ url: '/api/order/detail', method: 'get', params: { order_id: orderID } })) as unknown as BaseResp<OrderDetailResp>;
  const dto = toOrderDetailModel(raw.data as OrderDetailResp);
  return { code: raw.code, message: raw.message, data: dto } as BaseResp<OrderDetailModel>;
};

export const GetOrderList = async (): Promise<BaseResp<{ orders: OrderItemModel[] }>> => {
  const raw = (await req({ url: '/api/order/list', method: 'get' })) as unknown as BaseResp<{ orders: OrderItem[] }>;
  const dto = (raw.data?.orders ?? []).map((x: OrderItem) => toOrderItemModel(x));
  return { code: raw.code, message: raw.message, data: { orders: dto } } as BaseResp<{ orders: OrderItemModel[] }>;
};

export const CreateOrder = (payload: CreateOrderItem): Promise<BaseResp<string>> => {
  const dto: CreateOrderReq = toCreateOrderReq(payload);
  return req({ url: '/api/order/create', method: 'post', data: dto });
};

// removed D/DTO aliases: API 内部直接做转换，外部只使用 Model

export const FinishOrder = (orderID: string): Promise<BaseResp<string>> =>
  req({ url: '/api/order/finish', method: 'put', params: { order_id: orderID } });
