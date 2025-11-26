<template>
  <PageContainer>
    <SectionTitle>个人中心</SectionTitle>
      <div v-if="!userStore.isUserLogin" class="auth-card-wrapper">
         <el-empty description="您尚未登录">
            <el-button type="primary" size="large" round @click="router.push('/login')">
              立即登录
            </el-button>
         </el-empty>
      </div>

      <template v-else>
        <BaseCard class="profile-card">
          <div class="profile-bg"></div>
          <div class="profile-content">
             <div class="avatar-section" @click="showImagePreview">
               <el-avatar 
                :size="100" 
                :src="userStore.fullAvatarUrl" 
                class="user-avatar"
                shape="circle"
               />
               <div class="avatar-upload-trigger">
                  <el-icon><Camera /></el-icon>
               </div>
             </div>
             <div class="info-section">
                <div class="name-row">
                  <h2 class="username">{{ userStore.username }}</h2>
                  <el-tag 
                    effect="dark" 
                    round 
                    size="small"
                    :type="userStore.userRole === Role.COMMON_USER ? 'success' : 'warning'"
                  >
                    {{ userStore.userRole === Role.COMMON_USER ? '普通用户' : '内部人员' }}
                  </el-tag>
                </div>
                <p class="time-display">{{ currentTime }}</p>
                
                <div class="action-row">
                   <el-button size="small" round @click="openAvatarDialog">更换头像</el-button>
                   <el-button size="small" round>编辑资料</el-button>
                   <el-button size="small" type="danger" plain round @click="handleLogout">退出登录</el-button>
                </div>
             </div>
          </div>
        </BaseCard>

        

        <SectionTitle>快捷服务</SectionTitle>
        <div class="features-grid">
          
          <template v-if="userStore.userRole === Role.COMMON_USER">
            <div class="feature-item" @click="goToCreateOrderPage">
               <div class="icon-box blue">
                 <el-icon><Setting /></el-icon>
               </div>
               <div class="text-box">
                 <h3>我要报修</h3>
                 <p>设备故障？点此快速提交申请</p>
               </div>
               <el-icon class="arrow"><ArrowRight /></el-icon>
            </div>
            
            <div class="feature-item" @click="goToOrderListPage('mine')">
               <div class="icon-box red">
                 <el-icon><List /></el-icon>
               </div>
               <div class="text-box">
                 <h3>我的订单</h3>
                 <p>查看维修进度和历史记录</p>
               </div>
               <el-icon class="arrow"><ArrowRight /></el-icon>
            </div>
            
            <div class="feature-item" @click="showIntroduction">
               <div class="icon-box green">
                 <el-icon><User /></el-icon>
               </div>
               <div class="text-box">
                 <h3>关于师友</h3>
                 <p>了解更多服务计划详情</p>
               </div>
               <el-icon class="arrow"><ArrowRight /></el-icon>
            </div>
          </template>

          <template v-else-if="userStore.userRole === Role.INTERNAL_USER">
            <div class="feature-item" @click="goToCreateOrderPage">
               <div class="icon-box blue">
                 <el-icon><Setting /></el-icon>
               </div>
               <div class="text-box">
                 <h3>我要报修</h3>
                 <p>设备故障？点此快速提交申请</p>
               </div>
               <el-icon class="arrow"><ArrowRight /></el-icon>
            </div>

            <div class="feature-item" @click="goToOrderListPage('pending')">
               <div class="icon-box red">
                 <el-icon><List /></el-icon>
               </div>
               <div class="text-box">
                 <h3>未处理订单</h3>
                 <p>查看并处理用户的报修请求</p>
               </div>
               <el-icon class="arrow"><ArrowRight /></el-icon>
            </div>

            <div class="feature-item" @click="goToOrderListPage('mine')">
               <div class="icon-box red">
                 <el-icon><List /></el-icon>
               </div>
               <div class="text-box">
                 <h3>我的订单</h3>
                 <p>查看维修进度和历史记录</p>
               </div>
               <el-icon class="arrow"><ArrowRight /></el-icon>
            </div>

            <div class="feature-item" @click="goToSummaryListPage">
               <div class="icon-box orange">
                 <el-icon><Document /></el-icon>
               </div>
               <div class="text-box">
                 <h3>修机总结</h3>
                 <p>维修经验库与知识沉淀</p>
               </div>
               <el-icon class="arrow"><ArrowRight /></el-icon>
            </div>

            <div class="feature-item" @click="showIntroduction">
               <div class="icon-box green">
                 <el-icon><User /></el-icon>
               </div>
               <div class="text-box">
                 <h3>关于师友</h3>
                 <p>了解更多服务计划详情</p>
               </div>
               <el-icon class="arrow"><ArrowRight /></el-icon>
            </div>
          </template>
        </div>
        
      </template>
  </PageContainer>

    <el-dialog v-model="imagePreviewVisible" width="90%" center align-center class="glass-dialog">
      <div style="text-align: center">
        <el-image :src="userStore.fullAvatarUrl" fit="contain" style="max-width: 100%; max-height: 80vh"></el-image>
      </div>
    </el-dialog>

    <el-dialog v-model="dialogVisible" title="更换头像" width="320px" center class="rounded-dialog">
      <div class="upload-area">
        <el-button type="primary" size="large" round class="w-100" :icon="FolderOpened" @click="selectFromAlbum">
           选择图片
        </el-button>
      </div>
    </el-dialog>

    <el-dialog v-model="introductionDialogVisible" title="关于师友" width="600px" align-center>
      <div class="intro-wrapper">
        <el-image :src="introductionImage" fit="contain" />
      </div>
    </el-dialog>

    <input ref="fileInputRef" type="file" accept="image/*" style="display: none" @change="handleFileSelected" />
</template>

<script setup lang="ts" name="Home">
import { Role } from '@/enums/role.ts';
import { ref, onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { useUserStore } from '@/stores/user';
import { updateAvatar } from '@/api/updateAvatar.ts';
import {
  Camera, Edit, Setting, User, List, Document, FolderOpened, ArrowRight
} from '@element-plus/icons-vue'; // 引入 ArrowRight
import introductionImage from '@/assets/introduce.jpg';
import PageContainer from '@/components/common/PageContainer.vue';
import SectionTitle from '@/components/common/SectionTitle.vue';
import BaseCard from '@/components/common/BaseCard.vue';

// ...逻辑部分完全保持原样，不做任何修改...
const router = useRouter();
const userStore = useUserStore();
const imagePreviewVisible = ref(false);
const dialogVisible = ref(false);
const introductionDialogVisible = ref(false);
const fileInputRef = ref<HTMLInputElement | null>(null);
const currentTime = ref('');
let timer: number | undefined;

//

const updateTime = () => {
  const now = new Date();
  const year = now.getFullYear();
  const month = (now.getMonth() + 1).toString().padStart(2, '0');
  const day = now.getDate().toString().padStart(2, '0');
  const hour = now.getHours().toString().padStart(2, '0');
  const minute = now.getMinutes().toString().padStart(2, '0');
  const second = now.getSeconds().toString().padStart(2, '0');
  currentTime.value = `${year}-${month}-${day} ${hour}:${minute}:${second}`;
};

const handleLogout = () => {
  userStore.logout();
  ElMessage.success('已退出登录');
  router.replace('/login');
};

const openAvatarDialog = () => dialogVisible.value = true;

const selectFromAlbum = () => {
  if (fileInputRef.value) {
    fileInputRef.value.removeAttribute('capture');
    fileInputRef.value.click();
  }
};

const handleFileSelected = async (event: Event) => {
  const target = event.target as HTMLInputElement;
  const file = target.files?.[0];
  if (!file) return;
  if (!file.type.startsWith('image/')) {
    ElMessage.error('请选择图片文件!');
    return;
  }
  if (file.size / 1024 / 1024 > 10) {
    ElMessage.error('图片大小不能超过 10MB!');
    return;
  }
  try {
    const resp = await updateAvatar(file);
    if (resp.code === 0 && resp.data) {
      const imageUrl = resp.data;
      userStore.updateAvatar(imageUrl);
      ElMessage.success('图片上传成功！');
    } else {
      ElMessage.error(resp.message || '图片上传失败');
    }
  } catch (error) {
    ElMessage.error('网络请求失败，请稍后重试。');
  }
};

const showIntroduction = () => introductionDialogVisible.value = true;
const showImagePreview = () => imagePreviewVisible.value = true;
const goToCreateOrderPage = () => router.push({ name: 'CreateOrder' });
const goToOrderListPage = (view: 'mine' | 'pending' = 'mine') => router.push({ name: 'OrderList', query: { view } });
const goToSummaryListPage = () => router.push({ name: 'SummaryList' });

onMounted(() => {
  updateTime();
  timer = setInterval(updateTime, 1000);
});

onUnmounted(() => clearInterval(timer));

//
</script>

<style scoped>

/* 个人卡片 */
.profile-card { overflow: hidden; margin-bottom: 24px; position: relative; }

.profile-bg {
  height: 140px;
  background:
    linear-gradient(135deg, color-mix(in srgb, var(--el-color-primary), #ffffff 92%) 0%, #f6f8fb 55%, #eef1f6 100%),
    radial-gradient(600px 380px at 85% 10%, color-mix(in srgb, var(--el-color-primary), #ffffff 95%) 0%, transparent 70%);
  background-blend-mode: soft-light;
}

.profile-content {
  padding: 0 30px 30px;
  margin-top: -50px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.avatar-section {
  position: relative;
  cursor: pointer;
  padding: 4px;
  background: #fff;
  border-radius: 50%;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
}
.user-avatar {
  border: 2px solid #fff;
  display: block;
}
.avatar-upload-trigger {
  position: absolute;
  bottom: 5px;
  right: 5px;
  background: var(--el-color-primary);
  color: #fff;
  width: 28px;
  height: 28px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  border: 2px solid #fff;
}

.info-section {
  margin-top: 16px;
  width: 100%;
}

.name-row {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 4px;
}

.username {
  font-size: 24px;
  font-weight: 700;
  color: var(--color-text);
  letter-spacing: -0.2px;
  margin: 0;
}

.time-display {
  font-size: 14px;
  color: #909399;
  margin-bottom: 20px;
}

.action-row {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
}

/* 快捷服务网格 */
.section-title {
  font-size: 18px;
  font-weight: 600;
  color: #333;
  margin-bottom: 16px;
  padding-left: 8px;
}

.features-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 16px;
}
@media (min-width: 600px) {
  .features-grid {
    grid-template-columns: 1fr 1fr;
  }
}

.feature-item {
  background: #fff;
  border-radius: 16px;
  padding: 24px;
  display: flex;
  align-items: center;
  gap: 20px;
  cursor: pointer;
  transition: all 0.3s ease;
  border: 1px solid transparent;
  box-shadow: 0 4px 12px rgba(0,0,0,0.02);
}

.feature-item:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 24px rgba(0,0,0,0.06);
  border-color: #ebeef5;
}

.icon-box {
  width: 56px;
  height: 56px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 28px;
  flex-shrink: 0;
}
.icon-box.blue { background: #ecf5ff; color: #409eff; }
.icon-box.red { background: #fef0f0; color: #f56c6c; }
.icon-box.green { background: #f0f9eb; color: #67c23a; }
.icon-box.orange { background: #fdf6ec; color: #e6a23c; }

.text-box {
  flex: 1;
}
.text-box h3 {
  margin: 0 0 6px;
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}
.text-box p {
  margin: 0;
  font-size: 13px;
  color: #909399;
}

.arrow {
  color: #dcdfe6;
  font-size: 18px;
}
.feature-item:hover .arrow {
  color: var(--el-color-primary);
}

/* 弹窗美化 */
.w-100 {
  width: 100%;
}
.upload-area {
  padding: 20px 0;
  text-align: center;
}
</style>