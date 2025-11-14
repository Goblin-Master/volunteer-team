// src/stores/user.ts
import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import type { LoginModel } from '@/types/login.ts';
const base_url: string = 'http://127.0.0.1:9000/uploads/';

// 这是一个简化的 token 解析函数，实际应用中建议使用 jwt-decode 等库
const parseToken = (token: string) => {
  try {
    const payloadBase64 = token.split('.')[1];
    const payload = JSON.parse(atob(payloadBase64));
    return {
      role: payload.role || 1,
    };
  } catch (error) {
    console.error('Failed to parse token:', error);
    return null;
  }
};

export const useUserStore = defineStore('user', () => {
  // state
  const token = ref(localStorage.getItem('user-token') || null);
  const username = ref(localStorage.getItem('user-username') || null);
  const avatar = ref(localStorage.getItem('user-avatar') || null);

  // getters
  // 检查用户是否已登录
  const isUserLogin = computed(() => !!token.value);

  // 核心：使用 computed 来实时解析 token 并返回 role
  const userRole = computed(() => {
    if (token.value) {
      const parsed = parseToken(token.value);
      return parsed ? parsed.role : null;
    }
    return null;
  });

  // 获取完整的头像 URL
  const fullAvatarUrl = computed(() => {
    if (avatar.value) {
      return `${base_url}${avatar.value}`;
    }
    // 提供一个默认头像
    return 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg';
  });

  // actions
  const setUserInfo = (userData: LoginModel) => {
    // 1. 更新 Pinia store
    token.value = userData.token;
    username.value = userData.username;
    avatar.value = userData.avatar;

    // 2. 同时将数据持久化到 localStorage
    localStorage.setItem('user-token', userData.token);
    localStorage.setItem('user-username', userData.username);
    localStorage.setItem('user-avatar', userData.avatar);
  };

  const logout = () => {
    // 1. 清空 Pinia store
    token.value = null;
    username.value = null;
    avatar.value = null;

    // 2. 清空 localStorage
    localStorage.removeItem('user-token');
    localStorage.removeItem('user-username');
    localStorage.removeItem('user-avatar');
  };

  const updateAvatar = (newAvatarUrl: string) => {
    avatar.value = newAvatarUrl;
    localStorage.setItem('user-avatar', newAvatarUrl);
  };

  return {
    token,
    isUserLogin,
    username,
    avatar,
    userRole,
    fullAvatarUrl,
    setUserInfo,
    logout,
    updateAvatar,
  };
});
