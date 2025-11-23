<!-- OrderList.vue -->
<template>
  <PageContainer>
    <SectionTitle>{{ title }}</SectionTitle>
    <BaseCard shadow="never">
      <div v-if="orderList.length" class="order-list">
        <ItemCard
          v-for="item in orderList"
          :key="item.orderID"
          :time="formatTime(item.createTime)"
          :lines="[{ label: '问题描述', value: item.problemDescription }]"
          :highlightIndex="0"
        >
          <BaseButton type="primary" size="small" plain @click="goOrderDetail(item.orderID)">查看详情</BaseButton>
          <BaseButton v-if="userRole === Role.INTERNAL_USER" type="warning" size="small" plain @click="goWriteSummary(item.orderID)">写修机总结</BaseButton>
        </ItemCard>
      </div>
      <el-empty v-else :description="view === 'pending' ? '暂无未处理订单' : '暂无我的订单'" />
    </BaseCard>
  </PageContainer>
</template>

<script setup lang="ts" name="OrderList">
import { Role } from '@/enums/role.ts';
import { ref, onMounted, computed } from 'vue';
import { ElMessage } from 'element-plus';
import { useUserStore } from '@/stores/user.ts';
import { GetOrderList, GetOrderDetail } from '@/api/order.ts';
import { useRouter, useRoute } from 'vue-router';
import type { OrderItemModel } from '@/types/order.ts';
import PageContainer from '@/components/common/PageContainer.vue';
import SectionTitle from '@/components/common/SectionTitle.vue';
import BaseCard from '@/components/common/BaseCard.vue';
import BaseButton from '@/components/common/BaseButton.vue';
import ItemCard from '@/components/common/ItemCard.vue';

/* ---------- 数据 ---------- */
const userStore = useUserStore();
const router = useRouter();

const orderList = ref<OrderItemModel[]>([]);
const userRole = userStore.userRole;
const route = useRoute();
const view = computed(() => (route.query.view === 'pending' ? 'pending' : 'mine'));
const title = computed(() => (view.value === 'pending' ? '未处理订单' : '报修订单'));

/* ---------- 生命周期 ---------- */
onMounted(() => {
  fetchOrderList();
});

/* ---------- 方法 ---------- */
const fetchOrderList = async () => {
  try {
    const res = await GetOrderList();
    if (res.code === 0) {
      const all = res.data?.orders ?? [];
      if (view.value === 'mine' && userStore.username) {
        const details = await Promise.allSettled(all.map((x) => GetOrderDetail(x.orderID)));
        const mineIds = new Set(
          details
            .filter((d) => d.status === 'fulfilled' && d.value.code === 0 && d.value.data?.username === userStore.username)
            .map((d: any) => d.value.data.orderID),
        );
        orderList.value = all.filter((x) => mineIds.has(x.orderID));
      } else {
        orderList.value = all;
      }
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
.order-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}
:deep(.el-empty__description) { color: #909399; }
</style>