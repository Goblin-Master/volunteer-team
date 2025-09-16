import req from './axios';
import type BaseResp from '@/types/base';
import type {SummaryPayload} from '@/types/summary.ts';
export const CreateSummary = (data:SummaryPayload): Promise<BaseResp<string>> =>
  req({
    url: '/api/summary/create',
    method: 'post',
    data:data,
  })