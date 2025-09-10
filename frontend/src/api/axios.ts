import axios from "axios";
import { useUserStore } from "@/stores/user";
// 1. 创建 Axios 实例
const service = axios.create({
  baseURL: "http://127.0.0.1:9000", // .env 文件中定义的基础URL
  timeout: 10000, // 请求超时时间
});

// 2. 添加请求拦截器 (Request Interceptor)
service.interceptors.request.use(
  (config) => {
    // 在发送请求前做些什么
    // 比如，统一为请求添加 token
    const userStore = useUserStore();
    const token = userStore.token;
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    // 对请求错误做些什么
    console.error(error); // for debug
    return Promise.reject(error);
  }
);

// 3. 添加响应拦截器 (Response Interceptor)
service.interceptors.response.use(
  (response) => {
    // 对响应数据做点什么
    // 例如，只返回 response.data 部分
    const res = response.data;
    return res;
  },
  (error) => {
    // 对响应错误做点什么
    // 例如，处理 401 Unauthorized 错误，跳转到登录页
    console.error("Response Error:" + error); // for debug
    return Promise.reject(error);
  }
);

export default service; // 导出配置好的实例
