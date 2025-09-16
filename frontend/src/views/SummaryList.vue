<template>
  <div class="summary-list-page-container">
    <div class="header-title">修机总结</div>

    <main class="main-content">
      <el-card class="summary-card" shadow="never">
        <!-- 总结列表 -->
        <div v-if="summaries.length" class="summary-list">
          <div
            v-for="item in summaries"
            :key="item.id"
            class="summary-item"
          >
            <div class="summary-left">
              <div class="summary-time">{{ formatTime(item.utime) }}</div>
              <div class="summary-problem-type">问题类型：{{ item.problem_type }}</div>
              <div class="summary-problem-desc">问题描述：{{ item.problem_description }}</div>
              <div class="summary-repair">修机总结：{{ item.repair_summary }}</div>
              <div class="summary-receiver">接单人员：{{ item.receiver_name }}</div>
            </div>

            <div class="summary-right">
              <el-button
                type="primary"
                size="small"
                plain
                @click="goOrderDetail(item.order_id)"
              >
                查看订单详情
              </el-button>

              <el-button
                type="warning"
                size="small"
                plain
                @click="goUpdateSummary(item.id)"
              >
                更改修机总结
              </el-button>
            </div>
          </div>
        </div>

        <!-- 空状态 -->
        <el-empty v-else description="暂无修机总结" />
      </el-card>
    </main>
  </div>
</template>

<script setup lang="ts" name="SummaryList">
import { ref, onMounted } from 'vue';
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';
import { GetSummaryList } from '@/api/summary.ts';
import type { SummaryItem, SummaryListResp } from '@/types/summary.ts';

/* ---------- 数据 ---------- */
const router = useRouter()
const summaries = ref<SummaryItem[]>([])

/* ---------- 生命周期 ---------- */
onMounted(() => {
  fetchSummaryList()
})

/* ---------- 方法 ---------- */
const fetchSummaryList = async () => {
  try {
    const res = await GetSummaryList()
    if (res.code === 0) {
      summaries.value = (res.data as SummaryListResp)?.summaries ?? []
    } else {
      ElMessage.error(res.message || '获取修机总结失败')
    }
  } catch {
    ElMessage.error('网络错误或服务器无响应')
  }
}

const formatTime = (unix_ms: number) => {
  const date = new Date(unix_ms)
  const y = date.getFullYear()
  const m = String(date.getMonth() + 1).padStart(2, '0')
  const d = String(date.getDate()).padStart(2, '0')
  const h = String(date.getHours()).padStart(2, '0')
  const min = String(date.getMinutes()).padStart(2, '0')
  return `${y}-${m}-${d} ${h}:${min}`
}

const goOrderDetail = (order_id: number) => {
  router.push({ name: 'OrderDetail', query: { order_id } })
}

const goUpdateSummary = (summary_id: number) => {
  router.push({ name: 'UpdateSummary', query: { summary_id } })
}
</script>

<style scoped>
/* 样式与之前完全一致，无需改动 */
.summary-list-page-container {
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

.summary-card {
  border-radius: 16px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.05);
  padding: 8px 0;
}

.summary-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.summary-item {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 16px 20px;
  background: #ffffff;
  border-radius: 12px;
  transition: box-shadow 0.2s;
}

.summary-item:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.summary-left {
  flex: 1;
}

.summary-time {
  font-size: 13px;
  color: #909399;
  margin-bottom: 8px;
}

.summary-problem-type,
.summary-problem-desc,
.summary-repair,
.summary-receiver {
  font-size: 14px;
  color: #303133;
  line-height: 1.6;
}

.summary-right {
  display: flex;
  gap: 8px;
  margin-left: 16px;
}

@media (max-width: 600px) {
  .summary-item {
    flex-direction: column;
  }
  .summary-right {
    margin-left: 0;
    margin-top: 12px;
    width: 100%;
    justify-content: flex-end;
  }
}
</style>