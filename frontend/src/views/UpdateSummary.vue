<template>
  <div class="detail-page">
    <div class="detail-card">
      <div class="card-title">更新修机总结</div>

      <el-form
        ref="form_ref"
        :model="form"
        :rules="rules"
        label-width="90px"
        size="default"
      >
        <el-form-item label="订单号">{{ form.order_id }}</el-form-item>

        <!-- 问题类型 -->
        <el-form-item label="问题类型" prop="problem_type" required>
          <el-select
            v-model="form.problem_type"
            placeholder="请选择"
            class="w_full"
            :class="{ old: isCache }"
            @change="onceChange"
          >
            <el-option label="硬件故障" value="硬件故障" />
            <el-option label="软件问题" value="软件问题" />
            <el-option label="网络异常" value="网络异常" />
            <el-option label="其他" value="其他" />
          </el-select>
        </el-form-item>

        <!-- 问题描述 -->
        <el-form-item label="问题描述" prop="problem_description" required>
          <el-input
            v-model="form.problem_description"
            type="textarea"
            :rows="3"
            maxlength="200"
            show-word-limit
            placeholder="简要描述问题"
            :class="{ old: isCache }"
            @input="onceChange"
          />
        </el-form-item>

        <!-- 修机总结 -->
        <el-form-item label="修机总结" prop="repair_summary" required>
          <el-input
            v-model="form.repair_summary"
            type="textarea"
            :rows="4"
            maxlength="500"
            show-word-limit
            placeholder="维修过程、解决方案"
            :class="{ old: isCache }"
            @input="onceChange"
          />
        </el-form-item>

        <!-- 接单人员 -->
        <el-form-item label="接单人员" prop="receiver_name" required>
          <el-input
            v-model="form.receiver_name"
            placeholder="请输入姓名,人名之间用 , 隔开"
            :class="{ old: isCache }"
            @input="onceChange"
          />
        </el-form-item>

        <!-- 按钮 -->
        <el-form-item>
          <el-button
            type="primary"
            @click="handle_submit_summary"
            :loading="loading"
          >
            提交
          </el-button>
          <el-button @click="on_back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup lang="ts" name="UpdateSummary">
import { ref, reactive, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import type { FormInstance } from 'element-plus';
import {
  summary_rules,
  type SummaryPayload,
  type SummaryItem,
  type UpdateSummaryReq,
} from '@/types/summary';
import { GetSummaryDetail, UpdateSummary } from '@/api/summary';

/* ---------------- 基础准备 ---------------- */
const route = useRoute();
const router = useRouter();
const summary_id = String(route.query.summary_id);
const form_ref = ref<FormInstance>();
const loading = ref(false);

/* 表单 */
const form = reactive<SummaryPayload>({
  order_id: '',
  problem_type: '',
  problem_description: '',
  repair_summary: '',
  receiver_name: '',
});
const rules = summary_rules;

/* 旧数据标记 */
const isCache = ref(true);
const onceChange = () => (isCache.value = false);

/* ---------------- 生命周期 ---------------- */
onMounted(async () => {
  if (!summary_id) {
    ElMessage.error('缺少总结参数');
    return router.back();
  }

  const cache = history.state?.preload as string | undefined;
  if (cache) {
    try {
      const data = JSON.parse(cache) as SummaryItem;
      Object.assign(form, data);
      isCache.value = true; // 保持灰
      return;
    } catch {}
  }

  await load_old_summary();
  isCache.value = false; // 接口完成 → 正常颜色
});

/* ---------------- 方法 ---------------- */
async function load_old_summary() {
  try {
    const res = await GetSummaryDetail(summary_id);
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

async function handle_submit_summary() {
  if (!form_ref.value) return;
  const valid = await form_ref.value.validate().catch(() => false);
  if (!valid) return ElMessage.error('请检查表单输入项！');

  loading.value = true;
  try {
    const payload: UpdateSummaryReq = {
      summary_id: summary_id,
      ...form,
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

function on_back() {
  router.back();
}
</script>

<style scoped>
/* -------------- 布局 -------------- */
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
.w_full {
  width: 100%;
}

/* -------------- 旧数据统一淡灰色 -------------- */
/* 统一覆盖 Element 内部所有文本节点 */
.old *,
.old .el-input__inner,
.old .el-textarea__inner,
.old .el-select__placeholder,
.old .el-select__selected-item {
  color: #c0c4cc !important; /* 与你截图完全一致 */
  transition: color 0.2s;
}
</style>
