<template>
  <div v-if="shouldShow" class="notice-container">
    <div class="notice-circle" @click="toggleModal">
      <el-icon class="circle-icon"><BellFilled /></el-icon>
    </div>

    <transition name="fade">
      <div
        v-if="isModalOpen"
        class="notice-overlay"
        @click.self="toggleModal"
      >
        <div class="notice-card animate-pop">
          <div class="card-header">
            <h3>电脑报修须知</h3>
            <button class="close-btn" @click="toggleModal">
              <el-icon><Close /></el-icon>
            </button>
          </div>
          
          <div class="card-body custom-scrollbar">
            <div class="notice-section">
              <h4><span class="emoji">🛠️</span> 报修指引</h4>
              <ol>
                <li>
                  <strong>网络问题：</strong> 请直接联系学校网络中心。
                  <div class="sub-info">松山湖: 22861378 | 其它校区: 22680756</div>
                </li>
                <li>
                  <strong>硬件限制：</strong> 师友团队暂不支持硬件故障维修（如主板损坏），请联系厂家处理。
                </li>
                <li>
                  <strong>描述规范：</strong> 报修时请尽可能详细描述故障现象。
                </li>
                <li>
                  <strong>响应时间：</strong> 一般为工作日，1-2个工作日内会与您联系。
                </li>
                <li>
                  <strong>紧急联系：</strong>
                  <div class="contact-grid">
                     <div class="contact-item">陆晓彤: 14778189977</div>
                     <div class="contact-item">刘奕豪: 13714721480</div>
                  </div>
                </li>
              </ol>
            </div>

            <div class="notice-section warning">
              <h4><span class="emoji">💡</span> 温馨提示</h4>
              <ul>
                <li>莞城校区报修系统尚未开放。</li>
                <li>电话仅用于紧急情况，课后时间提供咨询。</li>
                <li>请务必确认联系方式正确。</li>
                <li>师友是校内唯一的免费义务维修组织，您的理解是我们最大的动力。</li>
                <li>师友提供的硅脂性能普通，如有高要求请自备。</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useRoute } from 'vue-router';
import { BellFilled, Close } from '@element-plus/icons-vue'; // 引入图标

// State
const isModalOpen = ref(false);

// Router
const route = useRoute();
const loginPaths = ['/login', '/register', '/resetPassword'];

// Computed
const shouldShow = computed(() => {
  return !loginPaths.includes(route.path);
});

// Methods
const toggleModal = () => {
  isModalOpen.value = !isModalOpen.value;
};
</script>

<style scoped>
.notice-container {
  position: fixed;
  bottom: 30px;
  right: 30px;
  z-index: 2000;
}

/* 悬浮球优化 */
.notice-circle {
  width: 56px;
  height: 56px;
  background: linear-gradient(135deg, #409eff, #337ecc);
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  box-shadow: 0 8px 20px rgba(64, 158, 255, 0.4);
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}
.notice-circle:hover {
  transform: scale(1.1) translateY(-2px);
  box-shadow: 0 12px 25px rgba(64, 158, 255, 0.5);
}
.circle-icon {
  color: white;
  font-size: 24px;
}

/* 遮罩与卡片 */
.notice-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(0, 0, 0, 0.4);
  backdrop-filter: blur(4px);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 2001;
  padding: 20px;
}

.notice-card {
  background: #fff;
  width: 100%;
  max-width: 600px;
  max-height: 85vh;
  border-radius: 20px;
  display: flex;
  flex-direction: column;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
  overflow: hidden;
}

.card-header {
  padding: 20px 24px;
  border-bottom: 1px solid #f0f2f5;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #fff;
}
.card-header h3 {
  margin: 0;
  font-size: 18px;
  color: #303133;
}

.close-btn {
  background: transparent;
  border: none;
  font-size: 20px;
  color: #909399;
  cursor: pointer;
  padding: 4px;
  border-radius: 50%;
  transition: background 0.2s;
  display: flex;
}
.close-btn:hover { background: #f2f3f5; color: #606266; }

.card-body {
  padding: 24px;
  overflow-y: auto;
  font-size: 14px;
  line-height: 1.6;
  color: #606266;
}

/* 内容样式优化 */
.notice-section { margin-bottom: 24px; }
.notice-section h4 {
  margin: 0 0 12px;
  color: #303133;
  font-size: 16px;
  display: flex;
  align-items: center;
  gap: 8px;
}
.emoji { font-size: 18px; }

ol, ul { padding-left: 20px; margin: 0; }
li { margin-bottom: 8px; }
li:last-child { margin-bottom: 0; }

.sub-info {
  font-size: 13px;
  color: #909399;
  margin-top: 4px;
  background: #f9f9f9;
  padding: 4px 8px;
  border-radius: 4px;
  display: inline-block;
}

.contact-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px;
  margin-top: 6px;
}
.contact-item {
  background: #ecf5ff;
  color: #409eff;
  padding: 6px 10px;
  border-radius: 6px;
  font-size: 13px;
  font-weight: 500;
}

.notice-section.warning {
  background: #fdf6ec;
  padding: 16px;
  border-radius: 12px;
  border: 1px solid #faecd8;
}
.notice-section.warning h4 { color: #e6a23c; }
.notice-section.warning li::marker { color: #e6a23c; }

/* 滚动条美化 */
.custom-scrollbar::-webkit-scrollbar { width: 6px; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #dcdfe6; border-radius: 3px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }

/* 动画 */
.fade-enter-active, .fade-leave-active { transition: opacity 0.2s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }

.animate-pop { animation: popIn 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275); }
@keyframes popIn {
  from { opacity: 0; transform: scale(0.95); }
  to { opacity: 1; transform: scale(1); }
}
</style>