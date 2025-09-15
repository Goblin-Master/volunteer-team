export interface OrderItem {
  id: number
  ctime: number
  problem_description: string
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
  username: string // 姓名
  student_id: string // 学号
  campus_location: string // 校区
  department: string // 学院

  /* 联系方式 */
  phone_number: string // 手机号
  wachat_id: string // 微信 ID

  /* 地址 */
  address: string // 详细地址

  /* 设备信息 */
  device_model: string // 设备型号
  os_version: string // 操作系统版本

  /* 报修内容 */
  problem_description: string // 问题描述
  notes: string // 备注

  /* 系统字段 */
  id: number // 订单主键
  ctime: number // 创建时间（毫秒时间戳）
}