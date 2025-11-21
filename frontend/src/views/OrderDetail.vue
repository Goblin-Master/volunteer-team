<template>
  <div class="detail-page">
    <div class="glass-card">
      <div class="header-row">
        <el-button icon="ArrowLeft" circle class="back-btn" @click="goBack" />
        <h1>报修详情单</h1>
      </div>

      <div class="content-body" v-loading="loading">
        <div class="info-grid">
          <div class="info-item full-width">
             <span class="label">问题描述</span>
             <div class="value highlight">{{ detail.problemDescription || '-' }}</div>
          </div>

          <div class="info-item">
            <span class="label">报修时间</span>
            <div class="value">{{ formatTime(detail.createTime) }}</div>
          </div>
          <div class="info-item">
            <span class="label">设备型号</span>
            <div class="value">{{ detail.deviceModel || '-' }}</div>
          </div>

          <div class="info-item">
            <span class="label">申请人</span>
            <div class="value">{{ detail.username || '-' }}</div>
          </div>
          <div class="info-item">
            <span class="label">学号</span>
            <div class="value">{{ detail.studentID || '-' }}</div>
          </div>

          <div class="info-item">
            <span class="label">联系电话</span>
            <div class="value">{{ detail.phoneNumber || '-' }}</div>
          </div>
          <div class="info-item">
            <span class="label">微信号</span>
            <div class="value">{{ detail.wechatID || '-' }}</div>
          </div>

          <div class="info-item full-width">
            <span class="label">详细地址</span>
            <div class="value">{{ detail.address || '-' }} ({{ detail.campusLocation }})</div>
          </div>

          <div class="info-item full-width" v-if="detail.notes">
             <span class="label">备注信息</span>
             <div class="value text-gray">{{ detail.notes }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts" name="OrderDetail">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { ArrowLeft } from '@element-plus/icons-vue';
import { GetOrderDetail } from '@/api/order.ts';
import type { OrderDetailModel } from '@/types/order';

/* (逻辑保持不变) */
const route = useRoute();
const router = useRouter();
const loading = ref(true);
const detail = ref<OrderDetailModel>({} as OrderDetailModel);

onMounted(() => {
  const orderID = String(route.query.orderID);
  if (!orderID) return ElMessage.error('缺少订单 ID');
  fetchDetail(orderID);
});

async function fetchDetail(orderID: string) {
  try {
    loading.value = true;
    const res = await GetOrderDetail(orderID);
    if (res.code === 0 && res.data) detail.value = res.data;
    else ElMessage.error(res.message || '获取详情失败');
  } catch {
    ElMessage.error('网络错误');
  } finally {
    loading.value = false;
  }
}

function formatTime(ms: number) {
  if(!ms) return '-';
  const d = new Date(ms);
  return `${d.getFullYear()}-${String(d.getMonth()+1).padStart(2,'0')}-${String(d.getDate()).padStart(2,'0')} ${String(d.getHours()).padStart(2,'0')}:${String(d.getMinutes()).padStart(2,'0')}`;
}

function goBack() { router.back(); }
</script>

<style scoped>
.detail-page {
  min-height: 100vh;
  background: #f5f7fa;
  padding: 30px 20px;
  display: flex;
  justify-content: center;
}

.glass-card {
  width: 100%;
  max-width: 700px;
  background: #fff;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0,0,0,0.05);
  padding: 30px;
  height: fit-content;
}

.header-row {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 30px;
  padding-bottom: 20px;
  border-bottom: 1px solid #f0f2f5;
}
.header-row h1 {
  margin: 0;
  font-size: 20px;
  color: #303133;
}
.back-btn {
  background: #f4f4f5;
  border: none;
  color: #606266;
}

/* Grid 布局模拟 Descriptions */
.info-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
}

.info-item {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.info-item.full-width {
  grid-column: span 2;
}

.label {
  font-size: 13px;
  color: #909399;
  font-weight: 500;
}

.value {
  font-size: 15px;
  color: #303133;
  background: #f9fafc;
  padding: 10px 12px;
  border-radius: 8px;
  border: 1px solid #eff2f6;
}

.value.highlight {
  background: #ecf5ff;
  color: #409eff;
  border-color: #d9ecff;
  line-height: 1.6;
}

.text-gray { color: #606266; }

@media (max-width: 600px) {
  .info-grid { grid-template-columns: 1fr; }
  .info-item.full-width { grid-column: span 1; }
}
</style>