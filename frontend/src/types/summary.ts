// 提交 payload 类型
export interface SummaryPayload {
  order_id: number;
  problem_type: string;
  problem_description: string;
  repair_summary: string;
  receiver_name: string;
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

// summary.ts
export interface SummaryItem {
  id: number
  order_id: number
  utime: number
  problem_type: string
  problem_description: string
  repair_summary: string
  receiver_name: string
}

export interface SummaryListResp {
  summaries: SummaryItem[]
}
/* ---------- 本地模拟提交 ---------- */
export async function submitSummaryMock(payload: SummaryPayload): Promise<void> {
  // 随机 0.5-1s 延迟，模拟网络
  await new Promise(r => setTimeout(r, 500 + Math.random() * 500))
  if (Math.random() > 0.9) throw new Error('模拟服务器异常')
  console.log('【模拟提交成功】', payload)
}

export interface SummaryDetailResp {
  id: number
  order_id: number
  utime: number
  problem_type: string
  problem_description: string
  repair_summary: string
  receiver_name: string
}

export interface UpdateSummaryReq {
  id: number // 总结的订单
  problem_type: string
  problem_description: string
  repair_summary: string
  receiver_name: string
}