import axios from 'axios';
import { useUserStore } from '@/stores/user';
import { API_BASE_URL, REQUEST_TIMEOUT } from '@/config/env';
import router from '@/router';
// 1. 创建 Axios 实例
const service = axios.create({
  baseURL: API_BASE_URL,
  timeout: REQUEST_TIMEOUT,
});

// 2. 添加请求拦截器 (Request Interceptor)
service.interceptors.request.use(
  (config) => {
    const userStore = useUserStore();
    const token = userStore.token;
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    console.error(error);
    return Promise.reject(error);
  },
);

// 3. 添加响应拦截器 (Response Interceptor)
service.interceptors.response.use(
  (response) => {
    const { data } = response;
    return data;
  },
  (error) => {
    if (error?.response?.status === 401) {
      const userStore = useUserStore();
      userStore.logout();
      const current = router.currentRoute.value.fullPath;
      router.replace({ name: 'Login', query: { redirect: current } });
    }
    console.error('Response Error:', error);
    return Promise.reject(error);
  },
);

export default service;
