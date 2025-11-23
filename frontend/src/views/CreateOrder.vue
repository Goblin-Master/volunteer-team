<template>
  <PageContainer>
    <BaseCard class="form-card">
      <div class="page-header">
        <SectionTitle>提交报修</SectionTitle>
      </div>
        <el-form
          ref="orderFormRef"
          :model="orderForm"
          :rules="orderModelRules"
          label-position="top"
          size="large"
          class="modern-form"
        >
          <div class="form-section">
            <h3 class="section-header">个人信息</h3>
            <el-row :gutter="20">
              <el-col :span="12" :xs="24">
                <el-form-item label="姓名" prop="username">
                  <el-input v-model="orderForm.username" placeholder="您的姓名" />
                </el-form-item>
              </el-col>
              <el-col :span="12" :xs="24">
                <el-form-item label="学号" prop="studentID">
                  <el-input v-model="orderForm.studentID" placeholder="请输入学号" />
                </el-form-item>
              </el-col>
            </el-row>

            <el-row :gutter="20">
              <el-col :span="12" :xs="24">
                <el-form-item label="联系电话" prop="phoneNumber">
                  <el-input v-model="orderForm.phoneNumber" placeholder="常用手机号" />
                </el-form-item>
              </el-col>
              <el-col :span="12" :xs="24">
                <el-form-item label="微信号" prop="wechatID">
                  <el-input v-model="orderForm.wechatID" placeholder="便于联系" />
                </el-form-item>
              </el-col>
            </el-row>
          </div>

          <div class="form-section">
            <h3 class="section-header">位置与设备</h3>
            <el-row :gutter="20">
              <el-col :span="12" :xs="24">
                <el-form-item label="校区" prop="campusLocation">
                  <el-select v-model="orderForm.campusLocation" placeholder="选择校区" class="w-100">
                    <el-option label="松山湖校区" value="松山湖校区" />
                    <el-option label="莞城校区" value="莞城校区" />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="12" :xs="24">
                <el-form-item label="院系" prop="department">
                  <el-input v-model="orderForm.department" placeholder="所属院系" />
                </el-form-item>
              </el-col>
            </el-row>
            
            <el-form-item label="详细地址" prop="address">
               <el-input v-model="orderForm.address" placeholder="例如：学生公寓A栋302" />
            </el-form-item>

            <el-row :gutter="20">
              <el-col :span="12" :xs="24">
                <el-form-item label="电脑机型" prop="deviceModel">
                  <el-input v-model="orderForm.deviceModel" placeholder="如：联想小新 Pro 14" />
                </el-form-item>
              </el-col>
              <el-col :span="12" :xs="24">
                <el-form-item label="系统版本" prop="osVersion">
                  <el-input v-model="orderForm.osVersion" placeholder="如：Win 11 / macOS" />
                </el-form-item>
              </el-col>
            </el-row>
          </div>

          <div class="form-section no-border">
            <h3 class="section-header">故障详情</h3>
            <el-form-item label="问题描述" prop="problemDescription">
              <el-input
                v-model="orderForm.problemDescription"
                type="textarea"
                :rows="5"
                maxlength="500"
                show-word-limit
                placeholder="请详细描述故障现象、发生时间等..."
                resize="none"
              />
            </el-form-item>

            <el-form-item label="备注" prop="notes">
              <el-input v-model="orderForm.notes" placeholder="选填：其他需要注意的事项" />
            </el-form-item>
          </div>

          <div class="form-footer">
             <el-button type="primary" size="large" class="submit-btn" :loading="isLoading" @click="submitForm">
               提交报修单
             </el-button>
          </div>
        </el-form>
    </BaseCard>
  </PageContainer>
</template>

<script setup lang="ts" name="CreateOrder">
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage, ElNotification } from 'element-plus';
import type { FormInstance } from 'element-plus';
import { type CreateOrderItem, orderModelRules } from '@/types/order.ts';
import { CreateOrder } from '@/api/order.ts';
import PageContainer from '@/components/common/PageContainer.vue';
import BaseCard from '@/components/common/BaseCard.vue';
import SectionTitle from '@/components/common/SectionTitle.vue';

const router = useRouter();
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
        duration: 2000,
      });
      orderFormRef.value.resetFields();
      router.back();
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
.page-header {
  padding: 0 0 16px 0;
  display: flex;
  align-items: center;
  gap: 16px;
}
.form-card { padding: 24px; }

.section-header {
  font-size: 16px;
  font-weight: 600;
  color: #303133;
  margin: 0 0 20px 0;
  padding-left: 10px;
  border-left: 4px solid #e4e7ed;
}

.form-section {
  margin-bottom: 30px;
  border-bottom: 1px solid #f0f2f5;
  padding-bottom: 10px;
}
.form-section.no-border {
  border-bottom: none;
}

/* 定制 Input */
:deep(.el-input__wrapper), :deep(.el-textarea__inner) {
  box-shadow: none !important;
  background-color: #f5f7fa;
  border: 1px solid #e4e7ed;
  border-radius: 8px;
  transition: all 0.2s ease;
}
:deep(.el-input__wrapper:hover), :deep(.el-textarea__inner:hover) {
  background-color: #fff;
  border-color: #c0c4cc;
}
:deep(.el-input__wrapper.is-focus), :deep(.el-textarea__inner:focus) {
  background-color: #fff;
  border-color: #c0c4cc;
  box-shadow: 0 0 0 3px rgba(0,0,0,0.06) !important;
}

:deep(.el-form-item__label) {
  font-weight: 500;
  color: #606266;
}

.w-100 { width: 100%; }

.submit-btn {
  width: 100%;
  font-weight: 600;
  border-radius: 12px;
  margin-top: 10px;
  height: 48px;
  box-shadow: 0 8px 16px rgba(0,0,0,0.08);
}

@media (max-width: 600px) {
  .form-card {
    padding: 20px;
  }
}
</style>