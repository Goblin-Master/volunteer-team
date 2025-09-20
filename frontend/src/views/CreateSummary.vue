<template>
  <div class="detail-page">
    <div class="detail-card">
      <div class="card-title">修机总结</div>

      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-width="90px"
        size="default"
      >
        <el-form-item label="订单号">{{ order_id }}</el-form-item>

        <el-form-item label="问题类型" prop="problem_type" required>
          <el-select v-model="form.problem_type" placeholder="请选择" class="w-full">
            <el-option label="硬件故障" value="硬件故障" />
            <el-option label="软件问题" value="软件问题" />
            <el-option label="网络异常" value="网络异常" />
            <el-option label="其他" value="其他" />
          </el-select>
        </el-form-item>

        <el-form-item label="问题描述" prop="problem_description" required>
          <el-input
            v-model="form.problem_description"
            type="textarea"
            :rows="3"
            maxlength="200"
            show-word-limit
            placeholder="简要描述问题"
          />
        </el-form-item>

        <el-form-item label="修机总结" prop="repair_summary" required>
          <el-input
            v-model="form.repair_summary"
            type="textarea"
            :rows="4"
            maxlength="500"
            show-word-limit
            placeholder="维修过程、解决方案"
          />
        </el-form-item>

        <el-form-item label="接单人员" prop="receiver_name" required>
          <el-input v-model="form.receiver_name" placeholder="请输入姓名,人名之间用 , 隔开" />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="handleSubmitSummary" :loading="loading">提交</el-button>
          <el-button @click="onBack">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup lang="ts" name='CreateSummary'>
import { ref, reactive, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import type { FormInstance } from 'element-plus';
import { summary_rules, type SummaryPayload } from '@/types/summary';
import { CreateSummary } from '@/api/summary';

const route = useRoute();
const router = useRouter();
const order_id = String(route.query.order_id);
const formRef = ref<FormInstance>();
const loading = ref(false);

const form = reactive<SummaryPayload>({
  order_id,
  problem_type: '',
  problem_description: '',
  repair_summary: '',
  receiver_name: ''
});

const rules = summary_rules;

onMounted(() => {
  if (!order_id) {
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

function onBack() {
  router.back();
}
</script>

<style scoped>
.detail-page {
  max-width: 800px;
  margin: 0 auto;
  padding: 24px 16px;
  box-sizing: border-box;
}
.detail-card {
  background: #ffffff;
  border: 1px solid #ebeef5;
  border-radius: 4px;
}
.card-title {
  padding: 12px 24px;
  font-size: 16px;
  font-weight: 500;
  color: #303133;
  border-bottom: 1px solid #ebeef5;
}
.w-full {
  width: 100%;
}
</style>