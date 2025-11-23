<template>
  <PageContainer>
    <BaseCard class="form-card">
      <div class="header-bar">
        <SectionTitle>更新修机总结</SectionTitle>
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

        <el-form-item label="问题类型" prop="problemType" required>
          <el-select
            v-model="form.problemType"
            placeholder="请选择"
            class="w-full"
            :class="{ old: isCache }"
            @change="onceChange"
          >
            <el-option label="硬件故障" value="硬件故障" />
            <el-option label="软件问题" value="软件问题" />
            <el-option label="网络异常" value="网络异常" />
            <el-option label="其他" value="其他" />
          </el-select>
        </el-form-item>

        <el-form-item label="问题描述" prop="problemDescription" required>
          <el-input
            v-model="form.problemDescription"
            type="textarea"
            :rows="3"
            maxlength="200"
            show-word-limit
            placeholder="简要描述问题"
            :class="{ old: isCache }"
            resize="none"
            @input="onceChange"
          />
        </el-form-item>

        <el-form-item label="修机总结" prop="repairSummary" required>
          <el-input
            v-model="form.repairSummary"
            type="textarea"
            :rows="5"
            maxlength="500"
            show-word-limit
            placeholder="维修过程、解决方案"
            :class="{ old: isCache }"
            resize="none"
            @input="onceChange"
          />
        </el-form-item>

        <el-form-item label="接单人员" prop="receiverName" required>
          <ReceiversInput
            v-model="form.receiverName"
            :class="{ old: isCache }"
            placeholder="输入姓名，回车添加"
          />
        </el-form-item>

        <el-form-item class="form-footer">
          <el-button
            type="primary"
            class="submit-btn"
            @click="handleSubmitSummary"
            :loading="loading"
          >
            保存更改
          </el-button>
        </el-form-item>
      </el-form>
    </BaseCard>
  </PageContainer>
</template>

<script setup lang="ts" name="UpdateSummary">
import { ref, reactive, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import type { FormInstance } from 'element-plus';
import { summaryModelRules, type SummaryItemModel, type UpdateSummaryModel } from '@/types/summary';
import { GetSummaryDetail, UpdateSummary } from '@/api/summary';
import ReceiversInput from '@/components/common/ReceiversInput.vue';
import PageContainer from '@/components/common/PageContainer.vue';
import BaseCard from '@/components/common/BaseCard.vue';
import SectionTitle from '@/components/common/SectionTitle.vue';

/* ---------- 逻辑保持原样 ---------- */
const route = useRoute();
const router = useRouter();
const summaryID = String(route.query.summaryID);
const formRef = ref<FormInstance>();
const loading = ref(false);

const form = reactive<SummaryItemModel>({
  summaryID: '',
  orderID: '',
  updateTime: 0,
  problemType: '',
  problemDescription: '',
  repairSummary: '',
  receiverName: '',
});
const rules = summaryModelRules;

const isCache = ref(true);
const onceChange = () => (isCache.value = false);

onMounted(async () => {
  if (!summaryID) {
    ElMessage.error('缺少总结参数');
    return router.back();
  }

  const cache = history.state?.preload as string | undefined;
  if (cache) {
    try {
      const data = JSON.parse(cache) as SummaryItemModel;
      Object.assign(form, data);
      isCache.value = true;
      return;
    } catch {}
  }

  await loadOldSummary();
  isCache.value = false;
});

async function loadOldSummary() {
  try {
    const res = await GetSummaryDetail(summaryID);
    if (res.code === 0) {
      Object.assign(form, res.data);
    } else {
      ElMessage.error(res.message || '获取详情失败');
      router.back();
    }
  } catch {
    ElMessage.error('网络异常');
    router.back();
  }
}

async function handleSubmitSummary() {
  if (!formRef.value) return;
  const valid = await formRef.value.validate().catch(() => false);
  if (!valid) return ElMessage.error('请检查表单输入项！');

  loading.value = true;
  try {
    const payload: UpdateSummaryModel = {
      summaryID: summaryID,
      problemType: form.problemType,
      problemDescription: form.problemDescription,
      repairSummary: form.repairSummary,
      receiverName: form.receiverName,
    };
    const resp = await UpdateSummary(payload);
    if (resp.code === 0) {
      ElMessage.success('更新成功');
      router.back();
    } else {
      ElMessage.error(resp.message || '更新失败');
    }
  } catch (e: any) {
    ElMessage.error(e.message || '网络异常');
  } finally {
    loading.value = false;
  }
}

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
.back-btn { background: #f4f4f5; border: none; color: #606266; }

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

.w-full { width: 100%; }

/* 核心：现代输入框样式 */
:deep(.el-input__wrapper), :deep(.el-textarea__inner) {
  box-shadow: none !important;
  background-color: #f5f7fa;
  border: 1px solid #e4e7ed;
  border-radius: 12px;
  transition: all 0.2s ease;
  padding: 10px 15px;
  color: #303133;
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

.submit-btn {
  width: 100%;
  height: 48px;
  font-size: 16px;
  border-radius: 12px;
  margin-top: 10px;
  box-shadow: 0 8px 20px rgba(64, 158, 255, 0.2);
}

/* -------------- 旧数据置灰逻辑优化 -------------- */
/* 使用 :deep 穿透到 Element 组件内部的 input */
:deep(.old .el-input__inner),
:deep(.old .el-textarea__inner),
:deep(.old .el-select__selected-item) {
  color: #909399 !important; /* 浅灰色 */
}
</style>