import req from './axios';
import type BaseResp from '@/types/base';
import type {
  SummaryItem,
  SummaryDetailResp,
  SummaryItemModel,
  SummaryDetailModel,
  UpdateSummaryModel,
} from '@/types/summary.ts';
import { toSummaryItemModel, toSummaryDetailModel, toSummaryItem, toUpdateSummaryReq } from '@/types/summary.ts';

export const CreateSummary = (payload: SummaryItemModel): Promise<BaseResp<string>> => {
  const dto: SummaryItem = toSummaryItem(payload);
  return req({ url: '/api/summary/create', method: 'post', data: dto });
};

export const GetSummaryList = async (): Promise<BaseResp<{ summaries: SummaryItemModel[] }>> => {
  const raw = (await req({ url: '/api/summary/list', method: 'get' })) as unknown as BaseResp<{ summaries: SummaryItem[] }>;
  const dto = (raw.data?.summaries ?? []).map((x: SummaryItem) => toSummaryItemModel(x));
  return { code: raw.code, message: raw.message, data: { summaries: dto } } as BaseResp<{ summaries: SummaryItemModel[] }>;
};

export const GetSummaryDetail = async (summaryID: string): Promise<BaseResp<SummaryDetailModel>> => {
  const raw = (await req({ url: '/api/summary/detail', method: 'get', params: { summary_id: summaryID } })) as unknown as BaseResp<SummaryDetailResp>;
  const dto = toSummaryDetailModel(raw.data as SummaryDetailResp);
  return { code: raw.code, message: raw.message, data:dto } as BaseResp<SummaryDetailModel>;
};

export const UpdateSummary = (reqModel: UpdateSummaryModel): Promise<BaseResp<SummaryDetailResp>> => {
  const dto = toUpdateSummaryReq(reqModel);
  return req({ url: '/api/summary/update', method: 'put', data: dto });
};

