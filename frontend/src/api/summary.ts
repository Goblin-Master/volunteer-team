import req from './axios';
import type BaseResp from '@/types/base';
import type { SummaryListResp, SummaryPayload,SummaryDetailResp,UpdateSummaryReq } from "@/types/summary.ts";

export const CreateSummary = (data:SummaryPayload): Promise<BaseResp<string>> =>
  req({
    url: '/api/summary/create',
    method: 'post',
    data:data,
  })

export const GetSummaryList = (): Promise<BaseResp<SummaryListResp>> =>
  req({ url: '/api/summary/list', method: 'get' })

export const GetSummaryDetail = (summary_id:string): Promise<BaseResp<SummaryDetailResp>> =>
  req({
    url: '/api/summary/detail',
    method: 'get',
    params: { summary_id }
  })

export const UpdateSummary = (data: UpdateSummaryReq): Promise<BaseResp<SummaryDetailResp>> =>
  req(
    {
      url: '/api/summary/update',
      method: 'put',
      data: data
    }
  )