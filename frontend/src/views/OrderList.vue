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
/* 容器 */
.order-list-page-container {
  min-height: 100vh;
  background: #f5f7fa;
  padding: 24px 16px;
  box-sizing: border-box;
}

.header-title {
  font-size: 22px;
  font-weight: 700;
  color: #1a1a1a;
  text-align: center;
  margin-bottom: 24px;
}

.main-content {
  max-width: 840px;
  margin: 0 auto;
}

/* 移除外层 Card 的阴影，让内部 item 承担视觉重量 */
.order-card {
  background: transparent;
  border: none;
  box-shadow: none;
}
:deep(.el-card__body) {
  padding: 0;
}

.order-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

/* 单个订单项 - 现代化改造 */
.order-item {
  background: #ffffff;
  border-radius: 16px;
  padding: 24px;
  border: 1px solid rgba(0,0,0,0.03);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.04);
  display: flex;
  flex-direction: column; /* 移动端优先 */
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.order-item::before {
  content: "";
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 4px;
  background: var(--el-color-primary);
  opacity: 0;
  transition: opacity 0.3s;
}

.order-item:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 20px rgba(0, 0, 0, 0.08);
}
.order-item:hover::before {
  opacity: 1;
}

/* 内部布局 desktop */
@media(min-width: 640px) {
  .order-item {
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
  }
  .order-left {
    flex: 1;
    padding-right: 20px;
  }
  .order-right {
    flex-shrink: 0;
    display: flex;
    gap: 12px;
  }
}

.order-time {
  font-size: 12px;
  color: #909399;
  background: #f4f4f5;
  padding: 4px 8px;
  border-radius: 4px;
  display: inline-block;
  margin-bottom: 12px;
}

.order-row {
  font-size: 15px;
  color: #303133;
  line-height: 1.6;
}
.order-row .label {
  color: #606266;
  font-weight: 500;
}

.order-right .el-button {
  border-radius: 8px;
  padding: 8px 16px;
}

/* 空状态 */
:deep(.el-empty__description) {
  color: #909399;
}
</style>