<!-- SummaryList.vue -->
<template>
  <div class="summary-list-page-container">
    <h1 class="header-title">修机总结</h1>

    <main class="main-content">
      <el-card class="summary-card" shadow="never">
        <!-- 列表 -->
        <div v-if="summaries.length" class="summary-list">
          <div
            v-for="item in summaries"
            :key="item.summary_id"
            class="summary-item"
          >
            <!-- 左侧信息 -->
            <div class="summary-left">
              <div class="summary-time">{{ formatTime(item.utime) }}</div>
              <div class="summary-row">
                <span class="label">问题类型：</span>
                <span>{{ item.problem_type }}</span>
              </div>
              <div class="summary-row">
                <span class="label">问题描述：</span>
                <span>{{ item.problem_description }}</span>
              </div>
              <div class="summary-row">
                <span class="label">修机总结：</span>
                <span>{{ item.repair_summary }}</span>
              </div>
              <div class="summary-row">
                <span class="label">接单人员：</span>
                <span>{{ item.receiver_name }}</span>
              </div>
            </div>

            <!-- 右侧操作 -->
            <div class="summary-right">
              <el-button
                type="primary"
                size="small"
                plain
                @click="goOrderDetail(item.order_id)"
              >
                查看详情
              </el-button>
              <el-button
                type="warning"
                size="small"
                plain
                @click="goUpdateSummary(item.summary_id)"
              >
                修改总结
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
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { GetSummaryList } from '@/api/summary'
import type { SummaryItem, SummaryListResp } from '@/types/summary'

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

const goOrderDetail = (order_id: string) => {
  router.push({ name: 'OrderDetail', query: { order_id } })
}

const goUpdateSummary = (summary_id: string) => {
   // 整行数据一起带走
  const target = summaries.value.find(v => v.summary_id === summary_id)
     router.push({
        name: 'UpdateSummary',
        query: { summary_id: String(summary_id) },
        state: { preload: JSON.stringify(target) }
   })
}
</script>

<style scoped>
/* ---------- 唯一一套浅色样式 ---------- */
.summary-list-page-container {
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

.summary-card {
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, .05);
  background: #fff;
  padding: 12px 0;
}

.summary-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.summary-item {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 20px;
  background: #ffffff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, .05);
  transition: transform .2s, box-shadow .2s;
}
.summary-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(0, 0, 0, .08);
}

.summary-left {
  flex: 1;
}

.summary-time {
  font-size: 13px;
  color: #909399;
  margin-bottom: 12px;
}

.summary-row {
  font-size: 14px;
  color: #303133;
  line-height: 1.7;
}
.summary-row .label {
  font-weight: 500;
  margin-right: 4px;
}

.summary-right {
  display: flex;
  gap: 8px;
  margin-left: 16px;
}
.summary-right .el-button {
  min-width: 88px;
}

/* ---------- 空状态 ---------- */
:deep(.el-empty__description) {
  color: #909399;
}

/* ---------- 响应式 ---------- */
@media (max-width: 640px) {
  .summary-item {
    flex-direction: column;
  }
  .summary-right {
    width: 100%;
    margin: 12px 0 0 0;
    justify-content: flex-end;
  }
}
</style>