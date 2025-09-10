// src/stores/user.ts
import { defineStore } from "pinia";
import type { LoginResp } from "@/types/login.ts";
const baseURL: string = "http://127.0.0.1:9000";

// 定义用户状态的类型
interface UserState {
  token: string | null;
  username: string | null;
  avatar: string | null;
}

export const useUserStore = defineStore("user", {
  state: (): UserState => ({
    token: localStorage.getItem("user-token"), // 从 localStorage 读取初始 token
    username: localStorage.getItem("user-username"),
    avatar: localStorage.getItem("user-avatar"),
  }),

  // 添加 getters
  getters: {
    // 检查用户是否已登录
    // 这个 getter 依赖于 token，当 token 变化时，它会自动重新计算
    isUserLogin(state): boolean {
      return !!state.token;
    },
    // 获取完整的头像 URL，如果后端只返回相对路径
    // 假设后端返回 "/avatars/user123.jpg"
    fullAvatarUrl(state): string | null {
      if (state.avatar) {
        return `${{ baseURL }}${state.avatar}`; // 组合成完整的 URL
      }
      return null;
    },
  },

  actions: {
    setUserInfo(userData: LoginResp) {
      // 1. 更新 Pinia store
      this.token = userData.token;
      this.username = userData.username;
      this.avatar = userData.avatar;

      // 2. 同时将数据持久化到 localStorage
      localStorage.setItem("user-token", userData.token);
      localStorage.setItem("user-username", userData.username);
      localStorage.setItem("user-avatar", userData.avatar);
    },

    logout() {
      // 1. 清空 Pinia store
      this.token = null;
      this.username = null;
      this.avatar = null;

      // 2. 清空 localStorage
      localStorage.removeItem("user-token");
      localStorage.removeItem("user-username");
      localStorage.removeItem("user-avatar");
    },
  },
});
