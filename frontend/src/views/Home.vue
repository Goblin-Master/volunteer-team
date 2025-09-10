<template>
  <div class="home-page">
    <h2>用户主页</h2>

    <div v-if="userStore.isUserLogin" class="user-info">
      <p><strong>用户名:</strong> {{ userStore.username }}</p>

      <div v-if="userStore.fullAvatarUrl" class="avatar-container">
        <img :src="userStore.fullAvatarUrl" alt="用户头像" class="avatar" />
      </div>

      <button @click="handleLogout" class="logout-btn">退出登录</button>
    </div>

    <div v-else class="not-logged-in">
      <p>您尚未登录，请先<router-link to="/login">登录</router-link>。</p>
    </div>
  </div>
</template>

<script setup lang="ts" name=" Home">
import { useRouter } from "vue-router";
import { useUserStore } from "@/stores/user";

const router = useRouter();
const userStore = useUserStore();

const handleLogout = () => {
  userStore.logout();
  router.replace("/login");
};
</script>

<style scoped>
/* 样式与之前相同，确保组件美观 */
.home-page {
  max-width: 600px;
  margin: 50px auto;
  padding: 20px;
  border: 1px solid #ddd;
  border-radius: 8px;
  text-align: center;
  font-family: Arial, sans-serif;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.user-info {
  margin-top: 20px;
}

.avatar-container {
  margin: 20px 0;
}

.avatar {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #3498db;
}

.logout-btn {
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
  background-color: #e74c3c;
  color: white;
  border: none;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.logout-btn:hover {
  background-color: #c0392b;
}

.not-logged-in {
  margin-top: 20px;
  color: #888;
}

.not-logged-in a {
  color: #3498db;
  text-decoration: none;
}
</style>
