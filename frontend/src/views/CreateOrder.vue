<template>
  <div class="order-page-container">
    <div class="header-title">提交报修订单</div>

    <main class="main-content">
      <el-card class="order-card">
        <el-form
          ref="orderFormRef"
          :model="orderForm"
          :rules="orderModelRules"
          label-position="top"
          class="order-form"
        >
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="姓名" prop="username">
                <el-input v-model="orderForm.username" placeholder="请输入姓名" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="学号" prop="studentID">
                <el-input v-model="orderForm.studentID" placeholder="请输入学号" />
              </el-form-item>
            </el-col>
          </el-row>

          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="所在校区" prop="campusLocation">
                <el-select
                  v-model="orderForm.campusLocation"
                  placeholder="请选择校区"
                  style="width: 100%"
                >
                  <el-option label="松山湖校区" value="松山湖校区" />
                  <el-option label="莞城校区" value="莞城校区" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="院系" prop="department">
                <el-input v-model="orderForm.department" placeholder="请输入院系" />
              </el-form-item>
            </el-col>
          </el-row>

          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="手机号" prop="phoneNumber">
                <el-input v-model="orderForm.phoneNumber" placeholder="请输入手机号" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="微信号" prop="wechatID">
                <el-input v-model="orderForm.wechatID" placeholder="请输入微信号" />
              </el-form-item>
            </el-col>
          </el-row>

          <el-form-item label="详细地址" prop="address">
            <el-input v-model="orderForm.address" placeholder="请输入详细地址" />
          </el-form-item>

          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="电脑机型" prop="deviceModel">
                <el-input v-model="orderForm.deviceModel" placeholder="如：联想小新Pro 14" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="操作系统版本" prop="osVersion">
                <el-input v-model="orderForm.osVersion" placeholder="如：Windows 11, macOS Ventura" />
              </el-form-item>
            </el-col>
          </el-row>

          <el-form-item label="问题描述" prop="problemDescription">
            <el-input
              v-model="orderForm.problemDescription"
              type="textarea"
              :rows="4"
              maxlength="500"
              show-word-limit
              placeholder="请详细描述您遇到的问题..."
            />
          </el-form-item>

          <el-form-item label="备注（选填）" prop="notes">
            <el-input
              v-model="orderForm.notes"
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
import { useRouter } from 'vue-router'; // ① 引入
import { ElMessage, ElNotification } from 'element-plus';
import type { FormInstance } from 'element-plus';
import { type CreateOrderItem, orderModelRules } from '@/types/order.ts';
import { CreateOrder } from '@/api/order.ts';

/* ---------- 路由 ---------- */
const router = useRouter(); // ② 实例化

/* ---------- 表单实例 & 数据 ---------- */
const orderFormRef = ref<FormInstance | null>(null);
const isLoading = ref(false);

const orderForm: CreateOrderItem = reactive({
  username: '',
  studentID: '',
  campusLocation: '',
  department: '',
  phoneNumber: '',
  wechatID: '',
  address: '',
  deviceModel: '',
  osVersion: '',
  problemDescription: '',
  notes: '',
});

/* ---------- 提交 ---------- */
const submitForm = async () => {
  if (!orderFormRef.value) return;
  isLoading.value = true;

  try {
    await orderFormRef.value.validate();
    const res = await CreateOrder(orderForm);

    if (res.code === 0) {
      ElNotification({
        title: '提交成功',
        message: '您的报修订单已提交，我们会尽快处理！',
        type: 'success',
        duration: 2000, // 2 秒后自动关闭
      });
      orderFormRef.value.resetFields();

      router.back(); // ③ 返回上一页
    } else {
      ElMessage.error(res.message || '提交失败');
    }
  } catch (e: any) {
    if (e?.response || e?.request) {
      ElMessage.error('网络错误或服务器无响应');
    }
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
