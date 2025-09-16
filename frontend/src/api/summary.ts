import req from './axios';
import type BaseResp from '@/types/base';
import type { SummaryListResp, SummaryPayload } from "@/types/summary.ts";

export const CreateSummary = (data:SummaryPayload): Promise<BaseResp<string>> =>
  req({
    url: '/api/summary/create',
    method: 'post',
    data:data,
  })

export const GetSummaryList = (): Promise<BaseResp<SummaryListResp>> =>
  req({ url: '/api/summary/list', method: 'get' })
