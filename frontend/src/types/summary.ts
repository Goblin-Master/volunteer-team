// 提交 payload 类型
export interface SummaryPayload {
  order_id: string
  problem_type: string
  problem_description: string
  repair_summary: string
  receiver_name: string
}

/* ---------- 校验规则（element-plus 格式） ---------- */
import type { FormRules } from 'element-plus'

export const summary_rules: FormRules = {
  problem_type: [
    { required: true, message: '请选择问题类型', trigger: 'change' }
  ],
  problem_description: [
    { required: true, message: '请输入问题描述', trigger: 'blur' },
    { min: 2, message: '至少 2 个字', trigger: 'blur' }
  ],
  repair_summary: [
    { required: true, message: '请填写修机总结', trigger: 'blur' },
    { min: 2, message: '至少 2 个字', trigger: 'blur' }
  ],
  receiver_name: [
    { required: true, message: '请输入接单人员姓名', trigger: 'blur' },
    { pattern: /^.{2,20}$/,
      message: '长度 2-20 位',
      trigger: 'blur' }
  ]
}

export interface SummaryItem {
  summary_id: string
  order_id: string
  utime: number
  problem_type: string
  problem_description: string
  repair_summary: string
  receiver_name: string
}

export interface SummaryListResp {
  summaries: SummaryItem[]
}

export interface SummaryDetailResp {
  summary_id: string
  order_id: string
  utime: number
  problem_type: string
  problem_description: string
  repair_summary: string
  receiver_name: string
}

export interface UpdateSummaryReq {
  summary_id: string
  problem_type: string
  problem_description: string
  repair_summary: string
  receiver_name: string
}