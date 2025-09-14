import type { FormRules } from 'element-plus';

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
const order_rules: FormRules = {
  username: [ // ✅ 将 student_name 修改为 username
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
  department: [
    { required: true, message: '请输入所在院系', trigger: 'blur' },
  ],
  phone_number: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号', trigger: 'blur' },
  ],
  wechat_id: [
    { required: false, message: '请输入微信号', trigger: 'blur' },
  ],
  address: [
    { required: true, message: '请输入详细地址', trigger: 'blur' },
  ],
  device_model: [
    { required: true, message: '请输入电脑机型', trigger: 'blur' },
  ],
  os_version: [
    { required: true, message: '请输入操作系统版本', trigger: 'blur' },
  ],
  problem_description: [
    { required: true, message: '请输入问题描述', trigger: 'blur' },
    { min: 10, max: 500, message: '长度在 10 到 500 个字符', trigger: 'blur' },
  ],
};

export { order_rules };