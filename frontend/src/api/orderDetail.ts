import type { OrderDetailResp } from '@/types/order'
import req from './axios'
import type BaseResp from '@/types/base'

export const GetOrderDetail = (id: number): Promise<BaseResp<OrderDetailResp>> =>
  req({
    url: '/api/order/detail',
    method: 'get',
    params: { id }
  })