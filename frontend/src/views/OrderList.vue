<template>
  <div class="order-list-page-container">
    <div class="header-title">我的报修订单</div>

    <main class="main-content">
      <el-card class="order-card" shadow="never">
        <!-- 订单列表 -->
        <div v-if="order_list.length" class="order-list">
          <div
            v-for="item in order_list"
            :key="item.id"
            class="order-item"
          >
            <div class="order-left">
              <div class="order-time">{{ format_time(item.ctime) }}</div>
              <div class="order-problem">{{ item.problem_description }}</div>
            </div>

            <div class="order-right">
              <el-button
                type="primary"
                size="small"
                plain
                @click="go_order_detail(item.id)"
              >
                查看详情
              </el-button>

              <el-button
                v-if="user_role === 1"
                type="warning"
                size="small"
                plain
                @click="go_write_summary(item.id)"
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
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/stores/user'
import { GetOrderList } from '@/api/orderList'
import { useRouter } from 'vue-router'
import type { OrderItem } from '@/types/order'

/* ---------- 数据 ---------- */
const user_store = useUserStore()
const router = useRouter()

const order_list = ref<OrderItem[]>([])
const user_role = user_store.user_role

/* ---------- 生命周期 ---------- */
onMounted(() => {
  fetch_order_list()
})

/* ---------- 方法 ---------- */
const fetch_order_list = async () => {
  try {
    const res = await GetOrderList()
    if (res.code === 0) {
      order_list.value = res.data?.orders ?? []
    } else {
      ElMessage.error(res.message || '获取订单失败')
    }
  } catch {
    ElMessage.error('网络错误或服务器无响应')
  }
}

const format_time = (unixMs: number) => {
  const date = new Date(unixMs)
  const Y = date.getFullYear()
  const M = String(date.getMonth() + 1).padStart(2, '0')
  const D = String(date.getDate()).padStart(2, '0')
  const h = String(date.getHours()).padStart(2, '0')
  const m = String(date.getMinutes()).padStart(2, '0')
  return `${Y}-${M}-${D} ${h}:${m}`
}

const go_order_detail = (order_id: number) => {
  router.push({ name: 'OrderDetail', params: { order_id } })
}

const go_write_summary = (order_id: number) => {
  router.push({ name: 'WriteSummary', params: { order_id } })
}
</script>

<style scoped>
.order-list-page-container {
  width: 100%;
  min-height: 100vh;
  background-color: #f5f7fa;
  padding: 40px 20px;
  box-sizing: border-box;
}

.header-title {
  font-size: 18px;
  color: #606266;
  text-align: center;
  margin-bottom: 24px;
}

.main-content {
  max-width: 800px;
  margin: 0 auto;
}

.order-card {
  border-radius: 16px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.05);
  padding: 8px 0;
}

.order-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.order-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  background: #ffffff;
  border-radius: 12px;
  transition: box-shadow 0.2s;
}

.order-item:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.order-left {
  flex: 1;
}

.order-time {
  font-size: 13px;
  color: #909399;
  margin-bottom: 6px;
}

.order-problem {
  font-size: 15px;
  color: #303133;
  line-height: 1.5;
}

.order-right {
  display: flex;
  gap: 8px;
  margin-left: 16px;
}

@media (max-width: 600px) {
  .order-item {
    flex-direction: column;
    align-items: flex-start;
  }
  .order-right {
    margin-left: 0;
    margin-top: 12px;
    width: 100%;
    justify-content: flex-end;
  }
}
</style>