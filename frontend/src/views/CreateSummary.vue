<template>
  <PageContainer>
    <BaseCard class="form-card">
      <div class="header-bar">
        <SectionTitle>撰写修机总结</SectionTitle>
      </div>

      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-position="top"
        size="large"
        class="modern-form"
      >
        <div class="order-tag">
          <span class="label">关联订单：</span>
          <span class="value">{{ form.orderID }}</span>
        </div>

        <el-row :gutter="20">
          <el-col :span="12" :xs="24">
            <el-form-item label="问题类型" prop="problemType">
              <el-select
                v-model="form.problemType"
                placeholder="请选择类型"
                class="w-full"
              >
                <el-option label="硬件故障" value="硬件故障" />
                <el-option label="软件问题" value="软件问题" />
                <el-option label="网络异常" value="网络异常" />
                <el-option label="其他" value="其他" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12" :xs="24">
            <el-form-item label="接单人员" prop="receiverName">
              <ReceiversInput v-model="form.receiverName" placeholder="输入姓名，回车添加" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="问题描述" prop="problemDescription">
          <el-input
            v-model="form.problemDescription"
            type="textarea"
            :rows="3"
            maxlength="200"
            show-word-limit
            placeholder="简要描述用户遇到的问题..."
            resize="none"
          />
        </el-form-item>

        <el-form-item label="维修方案与总结" prop="repairSummary">
          <el-input
            v-model="form.repairSummary"
            type="textarea"
            :rows="5"
            maxlength="500"
            show-word-limit
            placeholder="详细记录排查过程和最终解决方案..."
            resize="none"
          />
        </el-form-item>

        <el-form-item class="form-footer">
          <el-button
            type="primary"
            class="submit-btn"
            @click="handleSubmitSummary"
            :loading="loading"
          >
            提交总结
          </el-button>
        </el-form-item>
      </el-form>
    </BaseCard>
  </PageContainer>
</template>

<script setup lang="ts" name="CreateSummary">
import { ref, reactive, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import type { FormInstance } from 'element-plus';
import { summaryModelRules, type SummaryItemModel } from '@/types/summary';
import { CreateSummary } from '@/api/summary';
import ReceiversInput from '@/components/common/ReceiversInput.vue';
import PageContainer from '@/components/common/PageContainer.vue';
import BaseCard from '@/components/common/BaseCard.vue';
import SectionTitle from '@/components/common/SectionTitle.vue';

/* ---------- 逻辑部分完全保持 ---------- */
const route = useRoute();
const router = useRouter();
const orderID = String(route.query.orderID);
const formRef = ref<FormInstance>();
const loading = ref(false);

const form = reactive<SummaryItemModel>({
  summaryID: '',
  orderID: orderID,
  updateTime: 0,
  problemType: '',
  problemDescription: '',
  repairSummary: '',
  receiverName: '',
});

const rules = summaryModelRules;

onMounted(() => {
  if (!orderID) {
    ElMessage.error('缺少订单参数');
    router.back();
  }
});

const handleSubmitSummary = async () => {
  if (!formRef.value) return;

  const valid = await formRef.value.validate().catch(() => false);
  if (!valid) {
    ElMessage.error('请检查表单输入项！');
    return;
  }

  loading.value = true;
  try {
    const resp = await CreateSummary(form);
    if (resp.code === 0) {
      ElMessage.success('已记录');
      router.back();
    } else {
      ElMessage.error(resp.message || '提交失败');
    }
  } catch (err: any) {
    ElMessage.error(err.message || '网络异常');
    console.error(err);
  } finally {
    loading.value = false;
  }
};

</script>

<style scoped>
.form-card { padding: 24px; }

.header-bar {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid #f0f2f5;
}
.back-btn {
  background: #f4f4f5;
  border: none;
  color: #606266;
}

.order-tag {
  background: #f6f7fb;
  border: 1px solid #eff2f6;
  padding: 8px 16px;
  border-radius: 8px;
  display: inline-flex;
  align-items: center;
  margin-bottom: 24px;
  font-size: 14px;
}
.order-tag .label { color: #909399; margin-right: 4px; }
.order-tag .value { color: #303133; font-weight: 500; font-family: monospace; }

/* 深度定制表单控件 */
:deep(.el-input__wrapper), :deep(.el-textarea__inner) {
  box-shadow: none !important;
  background-color: #f5f7fa;
  border: 1px solid #e4e7ed;
  border-radius: 12px;
  transition: all 0.2s ease;
  padding: 10px 15px;
}
:deep(.el-textarea__inner) { padding: 12px; }

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
  margin-bottom: 6px;
}

.w-full { width: 100%; }

.submit-btn {
  width: 100%;
  height: 48px;
  font-size: 16px;
  border-radius: 12px;
  margin-top: 10px;
  box-shadow: 0 8px 20px rgba(64, 158, 255, 0.2);
}
</style>