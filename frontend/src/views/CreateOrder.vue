<template>
  <div class="order-page-container">
    <div class="header-title">提交报修订单</div>

    <main class="main-content">
      <el-card class="order-card">
        <el-form
          ref="order_form_ref"
          :model="order_form"
          :rules="order_rules"
          label-position="top"
          class="order-form"
        >
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="姓名" prop="username">
                <el-input v-model="order_form.username" placeholder="请输入姓名" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="学号" prop="student_id">
                <el-input v-model="order_form.student_id" placeholder="请输入学号" />
              </el-form-item>
            </el-col>
          </el-row>

          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="所在校区" prop="campus_location">
                <el-select
                  v-model="order_form.campus_location"
                  placeholder="请选择校区"
                  style="width: 100%;"
                >
                  <el-option label="松山湖校区" value="松山湖校区" />
                  <el-option label="莞城校区" value="莞城校区" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="院系" prop="department">
                <el-input v-model="order_form.department" placeholder="请输入院系" />
              </el-form-item>
            </el-col>
          </el-row>

          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="手机号" prop="phone_number">
                <el-input v-model="order_form.phone_number" placeholder="请输入手机号" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="微信号" prop="wechat_id">
                <el-input v-model="order_form.wechat_id" placeholder="请输入微信号" />
              </el-form-item>
            </el-col>
          </el-row>

          <el-form-item label="详细地址" prop="address">
            <el-input v-model="order_form.address" placeholder="请输入详细地址" />
          </el-form-item>

          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="电脑机型" prop="device_model">
                <el-input v-model="order_form.device_model" placeholder="如：联想小新Pro 14" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="操作系统版本" prop="os_version">
                <el-input v-model="order_form.os_version" placeholder="如：Windows 11, macOS Ventura" />
              </el-form-item>
            </el-col>
          </el-row>

          <el-form-item label="问题描述" prop="problem_description">
            <el-input
              v-model="order_form.problem_description"
              type="textarea"
              :rows="4"
              maxlength="500"
              show-word-limit
              placeholder="请详细描述您遇到的问题..."
            />
          </el-form-item>

          <el-form-item label="备注（选填）" prop="notes">
            <el-input
              v-model="order_form.notes"
              type="textarea"
              :rows="2"
              placeholder="其他需要说明的事项"
            />
          </el-form-item>

          <el-form-item>
            <el-button
              type="primary"
              class="submit-btn"
              :loading="isLoading"
              @click="submitForm"
            >
              提交订单
            </el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </main>
  </div>
</template>

<script setup lang="ts" name="CreateOrder">
import { reactive, ref } from 'vue';
import { ElMessage, ElNotification } from 'element-plus';
import type { FormInstance } from 'element-plus';
import { order_rules } from '@/types/createOrder';
import type { CreateOrderReq } from '@/types/createOrder';
import { CreateOrder } from '@/api/createOrder';

/* ---------- 表单实例 & 数据 ---------- */
const order_form_ref = ref<FormInstance | null>(null);
const isLoading = ref(false);

const order_form: CreateOrderReq = reactive({
  username: '',
  student_id: '',
  campus_location: '',
  department: '',
  phone_number: '',
  wechat_id: '',
  address: '',
  device_model: '',
  os_version: '',
  problem_description: '',
  notes: '',
});

/* ---------- 提交 ---------- */
const submitForm = async () => {
  if (!order_form_ref.value) return;
  isLoading.value = true;

  try {
    // 1. 校验
    await order_form_ref.value.validate();

    // 2. 请求
    const res = await CreateOrder(order_form);

    // 3. 业务处理
    if (res.code === 0) {
      ElNotification({
        title: '提交成功',
        message: '您的报修订单已提交，我们会尽快处理！',
        type: 'success',
      })
      order_form_ref.value.resetFields();
    } else {
      ElMessage.error(res.message || '提交失败');
    }
  } catch (e: any) {
    // 网络错误
    if (e?.response || e?.request) {
      ElMessage.error('网络错误或服务器无响应');
    }
    // 表单校验失败已由 Element 标红，不再 toast
  } finally {
    isLoading.value = false;
  }
};
</script>

<style scoped>
.order-page-container {
  width: 100%;
  min-height: 100vh;
  background-color: #f5f7fa;
  padding: 40px 20px;
  box-sizing: border-box;
}

.header-title {
  font-size: 18px;
  color: #606266;
  margin-bottom: 24px;
  text-align: center;
}

.main-content {
  max-width: 800px;
  margin: 0 auto;
}

.order-card {
  border-radius: 16px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.05);
}

.order-form {
  padding: 20px;
}

.el-form-item {
  margin-bottom: 24px;
}

.submit-btn {
  width: 100%;
  padding: 12px 0;
  font-size: 16px;
}
</style>