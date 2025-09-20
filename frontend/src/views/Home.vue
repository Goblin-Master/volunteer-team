<template>
  <div class="home-page-container">
    <div class="header-title">管理您的个人信息</div>

    <main class="main-content">
      <div v-if="userStore.is_user_login" class="profile-card">
        <div class="profile-header">
          <div class="avatar-wrapper" @click="showImagePreview">
            <el-image
              class="avatar-image"
              :src="userStore.full_avatar_url"
              fit="cover"
            />
          </div>
        </div>
        <div class="profile-body">
          <h2 class="user-name">{{ userStore.user_name }}</h2>
          <p class="current-time">{{ currentTime }}</p>
          <div class="action-buttons">
            <el-button
              type="primary"
              :icon="Camera"
              round
              @click="openAvatarDialog"
            >
              更换头像
            </el-button>
            <el-button :icon="Edit" round>编辑资料</el-button>
          </div>
        </div>
      </div>

      <div v-else class="not-logged-in">
        <p>
          您尚未登录，请先
          <router-link to="/login">登录</router-link>
        </p>
      </div>

      <template v-if="userStore.is_user_login">
        <div class="user-role-title">
          {{
            userStore.user_role === Role.COMMON_USER ? '普通用户' : '内部人员'
          }}
        </div>

        <div class="features-grid">
          <template v-if="userStore.user_role === Role.COMMON_USER">
            <div class="feature-card" @click="goToCreateOrderPage">
              <div
                class="feature-icon-wrapper"
                style="background-color: #eef5ff"
              >
                <el-icon :size="32" color="#409eff"><Setting /></el-icon>
              </div>
              <h3 class="feature-title">我要报修</h3>
              <p class="feature-description">快速提交设备维修申请</p>
            </div>
            <div class="feature-card" @click="goToOrderListPage">
              <!-- ← 新增点击 -->
              <div
                class="feature-icon-wrapper"
                style="background-color: #fef0f0"
              >
                <el-icon :size="32" color="#f56c6c"><List /></el-icon>
              </div>
              <h3 class="feature-title">我的订单</h3>
              <p class="feature-description">查看并管理我的报修订单</p>
            </div>
            <div class="feature-card" @click="showIntroduction">
              <div
                class="feature-icon-wrapper"
                style="background-color: #e9f8f1"
              >
                <el-icon :size="32" color="#67c23a"><User /></el-icon>
              </div>
              <h3 class="feature-title">关于师友</h3>
              <p class="feature-description">了解师友计划详情</p>
            </div>
          </template>

          <template v-else-if="userStore.user_role === Role.INTERNAL_USER">
            <div class="feature-card" @click="goToOrderListPage">
              <!-- ← 新增点击 -->
              <div
                class="feature-icon-wrapper"
                style="background-color: #fef0f0"
              >
                <el-icon :size="32" color="#f56c6c"><List /></el-icon>
              </div>
              <h3 class="feature-title">未处理订单</h3>
              <p class="feature-description">查看并处理维修订单</p>
            </div>
            <div class="feature-card" @click="goToSummaryListPage">
              <div
                class="feature-icon-wrapper"
                style="background-color: #fdf6ec"
              >
                <el-icon :size="32" color="#e6a23c"><Document /></el-icon>
              </div>
              <h3 class="feature-title">修机总结</h3>
              <p class="feature-description">撰写和查阅维修经验总结</p>
            </div>
          </template>
        </div>

        <el-button class="logout-btn" type="danger" round @click="handleLogout">
          退出登录
        </el-button>
      </template>
    </main>

    <el-dialog v-model="image_preview_visible" width="90%" center>
      <div style="text-align: center">
        <el-image
          :src="userStore.full_avatar_url"
          fit="contain"
          style="max-width: 100%; max-height: 80vh"
        ></el-image>
      </div>
    </el-dialog>

    <el-dialog v-model="dialog_visible" title="更换头像" width="350px" center>
      <div class="upload-options">
        <el-button
          type="primary"
          plain
          size="large"
          :icon="FolderOpened"
          @click="selectFromAlbum"
          >从图库选择</el-button
        >
      </div>
    </el-dialog>

    <el-dialog
      v-model="introduction_dialog_visible"
      title="关于师友"
      width="600px"
      custom-class="no-padding-dialog"
    >
      <div class="introduction-image-container">
        <el-image
          :src="introduction_image"
          fit="contain"
          style="width: 100%; height: 100%"
        />
      </div>
    </el-dialog>

    <input
      ref="file_input_ref"
      type="file"
      accept="image/*"
      style="display: none"
      @change="handleFileSelected"
    />
  </div>
</template>

<script setup lang="ts" name="Home">
import { Role } from '@/enums/role.ts';
import { ref, onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { useUserStore } from '@/stores/user';
import { updateAvatar } from '@/api/updateAvatar.ts';
import {
  Camera,
  Edit,
  Setting,
  User,
  List,
  Document,
  FolderOpened,
} from '@element-plus/icons-vue';

import introduction_image from '@/assets/introduce.jpg';

const router = useRouter();
const userStore = useUserStore();

const image_preview_visible = ref(false);
const dialog_visible = ref(false);
const introduction_dialog_visible = ref(false);
const file_input_ref = ref<HTMLInputElement | null>(null);
const currentTime = ref('');
let timer: number | undefined;

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

const openAvatarDialog = () => {
  dialog_visible.value = true;
};

const selectFromAlbum = () => {
  if (file_input_ref.value) {
    file_input_ref.value.removeAttribute('capture');
    file_input_ref.value.click();
  }
};

const handleFileSelected = async (event: Event) => {
  const target = event.target as HTMLInputElement;
  const file = target.files?.[0];
  if (!file) {
    return;
  }
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
    console.error('上传失败:', error);
  }
};

const showIntroduction = () => {
  introduction_dialog_visible.value = true;
};

const showImagePreview = () => {
  image_preview_visible.value = true;
};

// ✅ 新增的导航函数
const goToCreateOrderPage = () => {
  router.push({ name: 'CreateOrder' });
};

const goToOrderListPage = () => {
  router.push({ name: 'OrderList' });
};

const goToSummaryListPage = () => {
  router.push({ name: 'SummaryList' });
};

onMounted(() => {
  updateTime();
  timer = setInterval(updateTime, 1000);
});

onUnmounted(() => {
  clearInterval(timer);
});
</script>

<style scoped>
/* 样式保持不变 */
.home-page-container {
  width: 100%;
  min-height: 100vh;
  background-color: #f5f7fa;
  padding: 40px 20px;
  box-sizing: border-box;
  font-family:
    'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB',
    'Microsoft YaHei', Arial, sans-serif;
}

.header-title {
  font-size: 18px;
  color: #606266;
  margin-bottom: 24px;
  text-align: center;
}

.main-content {
  max-width: 960px;
  margin: 0 auto;
}

.profile-card {
  background-color: #ffffff;
  border-radius: 16px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.05);
  overflow: hidden;
  margin-bottom: 30px;
}

.profile-header {
  height: 120px;
  background: linear-gradient(to right, #e0c3fc, #8ec5fc);
  display: flex;
  justify-content: center;
  position: relative;
}

.avatar-wrapper {
  position: absolute;
  top: 50%;
  transform: translateY(-20%);
  width: 100px;
  height: 100px;
  border-radius: 50%;
  background-color: white;
  padding: 5px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  cursor: pointer;
}

.avatar-image {
  width: 100%;
  height: 100%;
  border-radius: 50%;
}

.profile-body {
  padding: 80px 20px 30px;
  text-align: center;
}

.user-name {
  font-size: 24px;
  font-weight: 500;
  color: #303133;
  margin: 0 0 8px;
}

.current-time {
  font-size: 14px;
  color: #909399;
  margin: 0 0 24px;
}

.action-buttons .el-button {
  margin: 0 8px;
}

.user-role-title {
  color: #909399;
  font-size: 14px;
  margin-bottom: 16px;
  display: flex;
  align-items: center;
  gap: 12px;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.feature-card {
  background-color: #ffffff;
  padding: 24px;
  border-radius: 12px;
  text-align: center;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.04);
  transition:
    transform 0.3s,
    box-shadow 0.3s;
  cursor: pointer;
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.08);
}

.feature-icon-wrapper {
  width: 64px;
  height: 64px;
  border-radius: 50%;
  margin: 0 auto 16px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.feature-title {
  font-size: 18px;
  font-weight: 500;
  color: #303133;
  margin: 0 0 8px;
}

.feature-description {
  font-size: 14px;
  color: #909399;
  margin: 0;
}

.upload-options {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.upload-options .el-button {
  width: 100%;
  margin: 0;
}

.logout-btn {
  display: block;
  margin: 0 auto;
  width: 200px;
}

.not-logged-in {
  max-width: 600px;
  margin: 50px auto;
  padding: 20px;
  border: 1px solid #ddd;
  border-radius: 8px;
  text-align: center;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  color: #888;
}

.not-logged-in a {
  color: #409eff;
  text-decoration: none;
}

.introduction-image-container {
  display: flex;
  justify-content: center;
  align-items: center;
  max-height: 80vh;
}

:deep(.no-padding-dialog .el-dialog__body) {
  padding: 0px !important;
}
</style>
