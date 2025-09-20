import type { FormRules } from 'element-plus';

export interface OrderItem {
  order_id: string;
  ctime: number;
  problem_description: string;
}
export interface OrderListResp {
  orders: OrderItem[];
}
/**
 * 报修订单详情接口返回数据结构
 * 路由：GET /api/order/detail?id=xxx
 */
export interface OrderDetailResp {
  /* 用户基本信息 */
  username: string; // 姓名
  student_id: string; // 学号
  campus_location: string; // 校区
  department: string; // 学院

  /* 联系方式 */
  phone_number: string; // 手机号
  wechat_id: string; // 微信 ID

  /* 地址 */
  address: string; // 详细地址

  /* 设备信息 */
  device_model: string; // 设备型号
  os_version: string; // 操作系统版本

  /* 报修内容 */
  problem_description: string; // 问题描述
  notes: string; // 备注

  /* 系统字段 */
  order_id: string; // 订单id
  ctime: number; // 创建时间（毫秒时间戳）
}
// 定义发送给后端的请求数据格式
// 所有变量都使用 a_b_c 命名规则
export interface CreateOrderReq {
  username: string; // ✅ 将 student_name 修改为 username
  student_id: string;
  campus_location: string;
  department: string;
  phone_number: string;
  wechat_id: string;
  address: string;
  device_model: string;
  os_version: string;
  problem_description: string;
  notes: string;
}

// 表单的前端验证规则
export const order_rules: FormRules = {
  username: [
    // ✅ 将 student_name 修改为 username
    { required: true, message: '请输入姓名', trigger: 'blur' },
    { min: 2, max: 10, message: '长度在 2 到 10 个字符', trigger: 'blur' },
  ],
  student_id: [
    { required: true, message: '请输入学号', trigger: 'blur' },
    { pattern: /^[0-9]+$/, message: '学号只能包含数字', trigger: 'blur' },
  ],
  campus_location: [
    { required: true, message: '请选择所在校区', trigger: 'change' },
  ],
  department: [{ required: true, message: '请输入所在院系', trigger: 'blur' }],
  phone_number: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    {
      pattern: /^1[3-9]\d{9}$/,
      message: '请输入正确的手机号',
      trigger: 'blur',
    },
  ],
  wechat_id: [{ required: false, message: '请输入微信号', trigger: 'blur' }],
  address: [{ required: true, message: '请输入详细地址', trigger: 'blur' }],
  device_model: [
    { required: true, message: '请输入电脑机型', trigger: 'blur' },
  ],
  os_version: [
    { required: true, message: '请输入操作系统版本', trigger: 'blur' },
  ],
  problem_description: [
    { required: true, message: '请输入问题描述', trigger: 'blur' },
    { min: 2, max: 500, message: '长度在 2 到 500 个字符', trigger: 'blur' },
  ],
};
