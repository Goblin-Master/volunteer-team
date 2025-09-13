<template>
  <div class="home-page-container">
    <div class="header-title">管理您的个人信息</div>

    <main class="main-content">
      <div v-if="userStore.is_user_login" class="profile-card">
        <div class="profile-header">
          <div class="avatar-wrapper">
            <el-image
              class="avatar-image"
              :src="userStore.full_avatar_url"
              :preview-src-list="[userStore.full_avatar_url]"
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
        <p>您尚未登录，请先
          <router-link to="/login">登录</router-link>
        </p>
      </div>

      <template v-if="userStore.is_user_login">
        <div class="user-role-title">
          {{ userStore.user_role === 2 ? '普通用户' : '内部人员' }}
        </div>

        <div class="features-grid">
          <template v-if="userStore.user_role === 2">
            <div class="feature-card">
              <div class="feature-icon-wrapper" style="background-color: #eef5ff;">
                <el-icon :size="32" color="#409eff"><Setting /></el-icon>
              </div>
              <h3 class="feature-title">我要报修</h3>
              <p class="feature-description">快速提交设备维修申请</p>
            </div>
            <div class="feature-card" @click="showIntroduction">
              <div class="feature-icon-wrapper" style="background-color: #e9f8f1;">
                <el-icon :size="32" color="#67c23a"><User /></el-icon>
              </div>
              <h3 class="feature-title">关于师友</h3>
              <p class="feature-description">了解师友计划详情</p>
            </div>
          </template>

          <template v-else-if="userStore.user_role === 1">
            <div class="feature-card">
              <div class="feature-icon-wrapper" style="background-color: #fef0f0;">
                <el-icon :size="32" color="#f56c6c"><List /></el-icon>
              </div>
              <h3 class="feature-title">未处理订单</h3>
              <p class="feature-description">查看并处理待维修订单</p>
            </div>
            <div class="feature-card">
              <div class="feature-icon-wrapper" style="background-color: #fdf6ec;">
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

    <el-dialog v-model="dialog_visible" title="更换头像" width="350px" center>
      <div class="upload-options">
        <el-button type="primary" plain size="large" :icon="FolderOpened" @click="selectFromAlbum">从图库选择</el-button>
      </div>
    </el-dialog>

    <el-dialog
      v-model="introduction_dialog_visible"
      title="关于师友"
      width="600px"
      custom-class="no-padding-dialog"
    >
      <div class="introduction-image-container">
        <el-image :src="introduction_image" fit="contain" style="width: 100%; height: 100%;"/>
      </div>
    </el-dialog>

    <input ref="file_input_ref" type="file" accept="image/*" style="display: none" @change="handleFileSelected" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import { useRouter } from "vue-router";
import { ElMessage } from 'element-plus';
import { useUserStore } from "@/stores/user";
import { updateAvatar } from "@/api/updateAvatar.ts";
import {
  Camera,
  Edit,
  Setting,
  User,
  List,
  Document,
  FolderOpened
} from '@element-plus/icons-vue';

// Import the image directly. Webpack/Vite will handle the path.
import introduction_image from '@/assets/introduce.jpg';

// --- Vue Router 和 Pinia Store ---
const router = useRouter();
const userStore = useUserStore();

// --- 响应式数据 ---
const dialog_visible = ref(false);
const introduction_dialog_visible = ref(false);
const file_input_ref = ref<HTMLInputElement | null>(null);
const currentTime = ref('');
let timer: number | undefined;

// --- 方法 ---
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

  // 1. 进行前端验证
  if (!file.type.startsWith('image/')) {
    ElMessage.error('请选择图片文件!');
    return;
  }
  if (file.size / 1024 / 1024 > 10) {
    ElMessage.error('图片大小不能超过 10MB!');
    return;
  }

  // 2. 调用封装好的 API 接口
  try {
    const resp = await updateAvatar(file);

    // 假设后端 code=0 表示成功
    if (resp.code === 0 && resp.data) {
      const imageUrl = resp.data; // 这里拿到图片 URL

      // 在这里处理后续逻辑，例如更新用户头像
      userStore.updateAvatar(imageUrl);
      ElMessage.success('图片上传成功！');
    } else {
      ElMessage.error(resp.message || '图片上传失败');
    }
  } catch (error) {
    // 处理网络错误或后端返回的非 2xx 状态码
    ElMessage.error('网络请求失败，请稍后重试。');
    console.error('上传失败:', error);
  }
};

const showIntroduction = () => {
  introduction_dialog_visible.value = true;
};

// --- 生命周期钩子 ---
onMounted(() => {
  updateTime(); // 首次加载时立即更新时间
  timer = setInterval(updateTime, 1000); // 每秒更新一次
});

onUnmounted(() => {
  clearInterval(timer); // 组件卸载时清除定时器，防止内存泄漏
});
</script>

<style scoped>
/* (Keep all existing styles as they are) */
.home-page-container {
  width: 100%;
  min-height: 100vh;
  background-color: #f5f7fa;
  padding: 40px 20px;
  box-sizing: border-box;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", Arial, sans-serif;
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
  transition: transform 0.3s, box-shadow 0.3s;
  /* Add pointer cursor to indicate it's clickable */
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

/* 调整关于师友图片显示的样式 */
.introduction-image-container {
  /* 确保容器内图片按比例缩小，并居中显示 */
  display: flex;
  justify-content: center;
  align-items: center;
  /* 限制图片容器的最大高度，避免在小屏幕上过高 */
  max-height: 80vh;
}

/* 使用深度选择器和自定义类名，精确地去除指定对话框的内边距 */
:deep(.no-padding-dialog .el-dialog__body) {
  padding: 0px !important;
}
</style>