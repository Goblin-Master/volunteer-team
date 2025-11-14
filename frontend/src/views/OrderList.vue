<!-- OrderList.vue -->
<template>
  <div class="order-list-page-container">
    <h1 class="header-title">报修订单</h1>

    <main class="main-content">
      <el-card class="order-card" shadow="never">
        <!-- 订单列表 -->
        <div v-if="orderList.length" class="order-list">
          <div
            v-for="item in orderList"
            :key="item.orderID"
            class="order-item"
          >
            <!-- 左侧信息 -->
            <div class="order-left">
              <div class="order-time">{{ formatTime(item.createTime) }}</div>
              <div class="order-row">
                <span class="label">问题描述：</span>
                <span>{{ item.problemDescription }}</span>
              </div>
            </div>

            <!-- 右侧操作 -->
            <div class="order-right">
              <el-button
                type="primary"
                size="small"
                plain
                @click="goOrderDetail(item.orderID)"
              >
                查看详情
              </el-button>

              <el-button
                v-if="userRole === Role.INTERNAL_USER"
                type="warning"
                size="small"
                plain
                @click="goWriteSummary(item.orderID)"
              >
                写修机总结
              </el-button>
            </div>
          </div>
        </div>

        <!-- 空状态 -->
        <el-empty v-else description="暂无订单" />
      </el-card>
    </main>
  </div>
</template>

<script setup lang="ts" name="OrderList">
import { Role } from '@/enums/role.ts';
import { ref, onMounted } from 'vue';
import { ElMessage } from 'element-plus';
import { useUserStore } from '@/stores/user.ts';
import { GetOrderList } from '@/api/order.ts';
import { useRouter } from 'vue-router';
import type { OrderItemModel } from '@/types/order.ts';

/* ---------- 数据 ---------- */
const userStore = useUserStore();
const router = useRouter();

const orderList = ref<OrderItemModel[]>([]);
const userRole = userStore.userRole;

/* ---------- 生命周期 ---------- */
onMounted(() => {
  fetchOrderList();
});

/* ---------- 方法 ---------- */
const fetchOrderList = async () => {
  try {
    const res = await GetOrderList();
    if (res.code === 0) {
      orderList.value = res.data?.orders ?? [];
    } else {
      ElMessage.error(res.message || '获取订单失败');
    }
  } catch {
    ElMessage.error('网络错误或服务器无响应');
  }
};

const formatTime = (unixMs: number) => {
  const date = new Date(unixMs);
  const y = date.getFullYear();
  const m = String(date.getMonth() + 1).padStart(2, '0');
  const d = String(date.getDate()).padStart(2, '0');
  const h = String(date.getHours()).padStart(2, '0');
  const min = String(date.getMinutes()).padStart(2, '0');
  return `${y}-${m}-${d} ${h}:${min}`;
};

const goOrderDetail = (orderID: string) => {
  router.push({ name: 'OrderDetail', query: { orderID } });
};

const goWriteSummary = (orderID: string) => {
  router.push({ name: 'CreateSummary', query: { orderID } });
};
</script>

<style scoped>
/* ---------- 与修机总结完全同款样式 ---------- */
.order-list-page-container {
  min-height: 100vh;
  background: #f5f7fa;
  padding: 40px 16px;
  box-sizing: border-box;
}

.header-title {
  font-size: 20px;
  font-weight: 500;
  color: #303133;
  text-align: center;
  margin-bottom: 32px;
}

.main-content {
  max-width: 840px;
  margin: 0 auto;
}

.order-card {
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  background: #fff;
  padding: 12px 0;
}

.order-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.order-item {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 20px;
  background: #ffffff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  transition:
    transform 0.2s,
    box-shadow 0.2s;
}
.order-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.08);
}

.order-left {
  flex: 1;
}

.order-time {
  font-size: 13px;
  color: #909399;
  margin-bottom: 12px;
}

.order-row {
  font-size: 14px;
  color: #303133;
  line-height: 1.7;
}
.order-row .label {
  font-weight: 500;
  margin-right: 4px;
}

.order-right {
  display: flex;
  gap: 8px;
  margin-left: 16px;
}
.order-right .el-button {
  min-width: 88px;
}

/* ---------- 空状态 ---------- */
:deep(.el-empty__description) {
  color: #909399;
}

/* ---------- 响应式 ---------- */
@media (max-width: 640px) {
  .order-item {
    flex-direction: column;
  }
  .order-right {
    width: 100%;
    margin: 12px 0 0 0;
    justify-content: flex-end;
  }
}
</style>
