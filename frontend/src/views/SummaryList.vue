<!-- SummaryList.vue -->
<template>
  <PageContainer>
    <SectionTitle>修机总结</SectionTitle>
    <BaseCard shadow="never">
      <div v-if="summaries.length" class="summary-list">
        <ItemCard
          v-for="item in summaries"
          :key="item.summaryID"
          :time="formatTime(item.updateTime || 0)"
          :lines="[
            { label: '问题类型', value: item.problemType },
            { label: '问题描述', value: item.problemDescription },
            { label: '修机总结', value: item.repairSummary },
            { label: '接单人员', value: item.receiverName }
          ]"
          :highlightIndex="2"
        >
          <BaseButton type="primary" size="small" plain @click="goOrderDetail(item.orderID)">查看详情</BaseButton>
          <BaseButton type="warning" size="small" plain @click="goUpdateSummary(item.summaryID || '')">修改总结</BaseButton>
        </ItemCard>
      </div>
      <el-empty v-else description="暂无修机总结" />
    </BaseCard>
  </PageContainer>
</template>

<script setup lang="ts" name="SummaryList">
import { ref, onMounted } from 'vue';
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';
import { GetSummaryList } from '@/api/summary';
import type { SummaryItemModel } from '@/types/summary';
import PageContainer from '@/components/common/PageContainer.vue';
import SectionTitle from '@/components/common/SectionTitle.vue';
import BaseCard from '@/components/common/BaseCard.vue';
import BaseButton from '@/components/common/BaseButton.vue';
import ItemCard from '@/components/common/ItemCard.vue';

/* ---------- 数据 ---------- */
const router = useRouter();
const summaries = ref<SummaryItemModel[]>([]);

/* ---------- 生命周期 ---------- */
onMounted(() => {
  fetchSummaryList();
});

/* ---------- 方法 ---------- */
const fetchSummaryList = async () => {
  try {
    const res = await GetSummaryList();
    if (res.code === 0) {
      summaries.value = res.data?.summaries ?? [];
    } else {
      ElMessage.error(res.message || '获取修机总结失败');
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

const goUpdateSummary = (summaryID: string) => {
  // 整行数据一起带走
  const target = summaries.value.find((v) => v.summaryID === summaryID);
  router.push({
    name: 'UpdateSummary',
    query: { summaryID: String(summaryID) },
    state: { preload: JSON.stringify(target) },
  });
};
</script>

<style scoped>
/* ---------- 唯一一套浅色样式 ---------- */
.summary-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}
:deep(.el-empty__description) { color: #909399; }
</style>
