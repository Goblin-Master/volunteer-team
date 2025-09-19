<template>
  <div class="order-detail-page">
    <el-card shadow="never" v-loading="loading">
      <template #header>
        <div class="card-header">
          <el-button :icon="ArrowLeft" circle @click="goBack" />
          <span class="title">订单详情</span>
        </div>
      </template>

      <!-- 详情列表 -->
      <el-descriptions :column="2" border>
        <el-descriptions-item label="报修时间">
          {{ formatTime(detail.ctime) }}
        </el-descriptions-item>

        <el-descriptions-item label="学号">
          {{ detail.student_id || '-' }}
        </el-descriptions-item>

        <el-descriptions-item label="姓名">
          {{ detail.username || '-' }}
        </el-descriptions-item>

        <el-descriptions-item label="校区">
          {{ detail.campus_location || '-' }}
        </el-descriptions-item>

        <el-descriptions-item label="学院">
          {{ detail.department || '-' }}
        </el-descriptions-item>

        <el-descriptions-item label="手机号">
          {{ detail.phone_number || '-' }}
        </el-descriptions-item>

        <el-descriptions-item label="微信">
          {{ detail.wechat_id || '-' }}
        </el-descriptions-item>

        <el-descriptions-item label="地址">
          {{ detail.address || '-' }}
        </el-descriptions-item>

        <el-descriptions-item label="设备型号">
          {{ detail.device_model || '-' }}
        </el-descriptions-item>

        <el-descriptions-item label="操作系统">
          {{ detail.os_version || '-' }}
        </el-descriptions-item>

        <el-descriptions-item label="问题描述" :span="2">
          {{ detail.problem_description || '-' }}
        </el-descriptions-item>

        <el-descriptions-item label="备注" :span="2">
          {{ detail.notes || '-' }}
        </el-descriptions-item>
      </el-descriptions>
    </el-card>
  </div>
</template>

<script setup lang="ts" name="OrderDetail">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { ArrowLeft } from '@element-plus/icons-vue';
import { GetOrderDetail } from '@/api/order.ts';
import type { OrderDetailResp } from '@/types/order';

/* ---------- 数据 ---------- */
const route = useRoute()
const router = useRouter()

const loading = ref(true)
const detail = ref<OrderDetailResp>({} as OrderDetailResp)

/* ---------- 生命周期 ---------- */
onMounted(() => {
  const order_id = String(route.query.order_id)
  if (!order_id) {
    ElMessage.error('缺少订单 ID')
    return
  }
  fetchDetail(order_id)
})

/* ---------- 方法 ---------- */
async function fetchDetail(order_id:string) {
  try {
    loading.value = true
    const res = await GetOrderDetail(order_id)
    if (res.code === 0 && res.data) {
      detail.value = res.data
    } else {
      ElMessage.error(res.message || '获取详情失败')
    }
  } catch {
    ElMessage.error('网络错误或服务器无响应')
  } finally {
    loading.value = false
  }
}

function formatTime(ms: number) {
  const d = new Date(ms)
  const Y = d.getFullYear()
  const M = String(d.getMonth() + 1).padStart(2, '0')
  const D = String(d.getDate()).padStart(2, '0')
  const h = String(d.getHours()).padStart(2, '0')
  const m = String(d.getMinutes()).padStart(2, '0')
  return `${Y}-${M}-${D} ${h}:${m}`
}

function goBack() {
  router.back()
}
</script>

<style scoped>
.order-detail-page {
  max-width: 800px;
  margin: 40px auto;
  padding: 0 20px;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 12px;
}

.title {
  font-size: 18px;
  color: #303133;
}
</style>