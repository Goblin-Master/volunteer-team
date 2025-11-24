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
import { mapOk } from '@/api/helpers';

export const CreateSummary = (payload: SummaryItemModel): Promise<BaseResp<string>> => {
  const dto: SummaryItem = toSummaryItem(payload);
  return req({ url: '/api/summary/create', method: 'post', data: dto });
};

export const GetSummaryList = async (): Promise<BaseResp<{ summaries: SummaryItemModel[] } | null>> => {
  const raw = (await req({ url: '/api/summary/list', method: 'get' })) as unknown as BaseResp<{ summaries: SummaryItem[] } | null>;
  return mapOk(raw, (d) => ({ summaries: (d.summaries ?? []).map((x: SummaryItem) => toSummaryItemModel(x)) })) as BaseResp<{ summaries: SummaryItemModel[] } | null>;
};

export const GetSummaryDetail = async (summaryID: string): Promise<BaseResp<SummaryDetailModel | null>> => {
  const raw = (await req({ url: '/api/summary/detail', method: 'get', params: { summary_id: summaryID } })) as unknown as BaseResp<SummaryDetailResp | null>;
  return mapOk(raw, toSummaryDetailModel) as BaseResp<SummaryDetailModel | null>;
};

export const UpdateSummary = (reqModel: UpdateSummaryModel): Promise<BaseResp<SummaryDetailResp>> => {
  const dto = toUpdateSummaryReq(reqModel);
  return req({ url: '/api/summary/update', method: 'put', data: dto });
};

