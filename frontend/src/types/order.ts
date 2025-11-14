import type { FormRules } from 'element-plus';

export interface OrderItem {
  order_id: string;
  ctime: number;
  problem_description: string;
}
export interface OrderListResp {
  orders: OrderItem[];
}

export interface OrderDetailResp {
  username: string;
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
  order_id: string;
  ctime: number;
}

export interface CreateOrderReq {
  username: string;
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

// 业务模型规则（CamelCase）
export const orderModelRules: FormRules<CreateOrderItem> = {
  username: [
    { required: true, message: '请输入姓名', trigger: 'blur' },
    { min: 2, max: 10, message: '长度在 2 到 10 个字符', trigger: 'blur' },
  ],
  studentID: [{ required: true, message: '请输入学号', trigger: 'blur' }],
  campusLocation: [{ required: true, message: '请选择所在校区', trigger: 'change' }],
  department: [{ required: true, message: '请输入所在院系', trigger: 'blur' }],
  phoneNumber: [{ required: true, message: '请输入手机号', trigger: 'blur' }],
  wechatID: [{ required: false, message: '请输入微信号', trigger: 'blur' }],
  address: [{ required: true, message: '请输入详细地址', trigger: 'blur' }],
  deviceModel: [{ required: true, message: '请输入电脑机型', trigger: 'blur' }],
  osVersion: [{ required: true, message: '请输入操作系统版本', trigger: 'blur' }],
  problemDescription: [
    { required: true, message: '请输入问题描述', trigger: 'blur' },
    { min: 2, max: 500, message: '长度在 2 到 500 个字符', trigger: 'blur' },
  ],
  notes: [{ required: false }],
};


export interface OrderItemModel {
  orderID: string;
  createTime: number;
  problemDescription: string;
}

export interface OrderDetailModel {
  orderID: string;
  createTime: number;
  username: string;
  studentID: string;
  campusLocation: string;
  department: string;
  phoneNumber: string;
  wechatID: string;
  address: string;
  deviceModel: string;
  osVersion: string;
  problemDescription: string;
  notes: string;
}

export interface CreateOrderItem {
  username: string;
  studentID: string;
  campusLocation: string;
  department: string;
  phoneNumber: string;
  wechatID: string;
  address: string;
  deviceModel: string;
  osVersion: string;
  problemDescription: string;
  notes: string;
}

export const orderModelRulesEx: FormRules<CreateOrderItem> = orderModelRules;

export const toOrderItemModel = (x: OrderItem): OrderItemModel => ({
  orderID: x.order_id,
  createTime: x.ctime,
  problemDescription: x.problem_description,
});

export const toOrderDetailModel = (x: OrderDetailResp): OrderDetailModel => ({
  orderID: x.order_id,
  createTime: x.ctime,
  username: x.username,
  studentID: x.student_id,
  campusLocation: x.campus_location,
  department: x.department,
  phoneNumber: x.phone_number,
  wechatID: x.wechat_id,
  address: x.address,
  deviceModel: x.device_model,
  osVersion: x.os_version,
  problemDescription: x.problem_description,
  notes: x.notes,
});

export const toCreateOrderReq = (x: CreateOrderItem): CreateOrderReq => ({
  username: x.username,
  student_id: x.studentID,
  campus_location: x.campusLocation,
  department: x.department,
  phone_number: x.phoneNumber,
  wechat_id: x.wechatID,
  address: x.address,
  device_model: x.deviceModel,
  os_version: x.osVersion,
  problem_description: x.problemDescription,
  notes: x.notes,
});
