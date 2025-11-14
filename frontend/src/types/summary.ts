export interface SummaryItem {
  summary_id?: string;
  order_id: string;
  utime?: number;
  problem_type: string;
  problem_description: string;
  repair_summary: string;
  receiver_name: string;
}

/* ---------- 校验规则（element-plus 格式） ---------- */
import type { FormRules } from 'element-plus';

// 业务模型规则（CamelCase）
export const summaryModelRules: FormRules = {
  problemType: [{ required: true, message: '请选择问题类型', trigger: 'change' }],
  problemDescription: [
    { required: true, message: '请输入问题描述', trigger: 'blur' },
    { min: 2, message: '至少 2 个字', trigger: 'blur' },
  ],
  repairSummary: [
    { required: true, message: '请填写修机总结', trigger: 'blur' },
    { min: 2, message: '至少 2 个字', trigger: 'blur' },
  ],
  receiverName: [
    { required: true, message: '请输入接单人员姓名', trigger: 'blur' },
    { pattern: /^.{2,20}$/, message: '长度 2-20 位', trigger: 'blur' },
  ],
};

export interface SummaryListResp {
  summaries: SummaryItem[];
}

export type SummaryDetailResp = SummaryItem;

export interface UpdateSummaryReq {
  summary_id: string;
  problem_type: string;
  problem_description: string;
  repair_summary: string;
  receiver_name: string;
}


export interface SummaryItemModel {
  summaryID: string;
  orderID: string;
  updateTime: number;
  problemType: string;
  problemDescription: string;
  repairSummary: string;
  receiverName: string;
}

export interface SummaryDetailModel {
  summaryID: string;
  orderID: string;
  updateTime: number;
  problemType: string;
  problemDescription: string;
  repairSummary: string;
  receiverName: string;
}

export const summaryModelRulesEx: FormRules<SummaryItemModel> = summaryModelRules as FormRules<SummaryItemModel>;

// 业务模型统一使用 Item 后缀

export interface UpdateSummaryModel {
  summaryID: string;
  problemType: string;
  problemDescription: string;
  repairSummary: string;
  receiverName: string;
}

export const toSummaryItemModel = (x: SummaryItem): SummaryItemModel => ({
  summaryID: x.summary_id ?? '',
  orderID: x.order_id,
  updateTime: x.utime ?? 0,
  problemType: x.problem_type,
  problemDescription: x.problem_description,
  repairSummary: x.repair_summary,
  receiverName: x.receiver_name,
});

export const toSummaryDetailModel = (x: SummaryDetailResp): SummaryDetailModel => ({
  summaryID: x.summary_id ?? '',
  orderID: x.order_id,
  updateTime: x.utime ?? 0,
  problemType: x.problem_type,
  problemDescription: x.problem_description,
  repairSummary: x.repair_summary,
  receiverName: x.receiver_name,
});

export const toSummaryItem = (x: SummaryItemModel): SummaryItem => ({
  summary_id: x.summaryID,
  order_id: x.orderID,
  utime: x.updateTime,
  problem_type: x.problemType,
  problem_description: x.problemDescription,
  repair_summary: x.repairSummary,
  receiver_name: x.receiverName,
});

export const toUpdateSummaryReq = (x: UpdateSummaryModel): UpdateSummaryReq => ({
  summary_id: x.summaryID,
  problem_type: x.problemType,
  problem_description: x.problemDescription,
  repair_summary: x.repairSummary,
  receiver_name: x.receiverName,
});
