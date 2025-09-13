// src/stores/user.ts
import { defineStore } from "pinia";
import { ref, computed } from "vue";
import type { LoginResp } from "@/types/login.ts";
const base_url: string = "http://127.0.0.1:9000/uploads/";

// 这是一个简化的 token 解析函数，实际应用中建议使用 jwt-decode 等库
const parse_token = (token: string) => {
  try {
    const payload_base64 = token.split('.')[1];
    const payload = JSON.parse(atob(payload_base64));
    return {
      role: payload.role_id || 2, // 默认 role 为 2 (普通用户)
    };
  } catch (error) {
    console.error('Failed to parse token:', error);
    return null;
  }
};

export const useUserStore = defineStore("user", () => {
  // state
  const token = ref(localStorage.getItem("user-token") || null);
  const user_name = ref(localStorage.getItem("user-username") || null);
  const avatar = ref(localStorage.getItem("user-avatar") || null);

  // getters
  // 检查用户是否已登录
  const is_user_login = computed(() => !!token.value);

  // 核心：使用 computed 来实时解析 token 并返回 role
  const user_role = computed(() => {
    if (token.value) {
      const parsed = parse_token(token.value);
      return parsed ? parsed.role : null;
    }
    return null;
  });

  // 获取完整的头像 URL
  const full_avatar_url = computed(() => {
    if (avatar.value) {
      return `${base_url}${avatar.value}`;
    }
    // 提供一个默认头像
    return 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg';
  });

  // actions
  const setUserInfo = (userData: LoginResp) => {
    // 1. 更新 Pinia store
    token.value = userData.token;
    user_name.value = userData.username;
    avatar.value = userData.avatar;

    // 2. 同时将数据持久化到 localStorage
    localStorage.setItem("user-token", userData.token);
    localStorage.setItem("user-username", userData.username);
    localStorage.setItem("user-avatar", userData.avatar);
  };

  const logout = () => {
    // 1. 清空 Pinia store
    token.value = null;
    user_name.value = null;
    avatar.value = null;

    // 2. 清空 localStorage
    localStorage.removeItem("user-token");
    localStorage.removeItem("user-username");
    localStorage.removeItem("user-avatar");
  };

  const updateAvatar = (new_avatar_url: string) => {
    avatar.value = new_avatar_url;
    localStorage.setItem("user-avatar", new_avatar_url);
  }

  return {
    token,
    is_user_login,
    user_name,
    avatar,
    user_role,
    full_avatar_url,
    setUserInfo,
    logout,
    updateAvatar
  };
});