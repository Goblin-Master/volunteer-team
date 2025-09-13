import { createRouter, createWebHistory } from "vue-router";
import Login from "@/views/Login.vue";
import Register from "@/views/Register.vue";
import ResetPassword from "@/views/ResetPassword.vue";
import Home from "@/views/Home.vue";
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/login",
      name: "Login",
      component: Login,
    },
    {
      path: "/register",
      name: "Register",
      component: Register,
    },
    {
      path: "/resetPassword",
      name: "ResetPassword",
      component: ResetPassword,
    },
    {
      path: "/home",
      name: "Home",
      component: Home,
      meta: { requiresAuth: true }, // 添加 requiresAuth 字段，表示需要登录
    },
    {
      path: "/",
      redirect: "/home",
    },
  ],
});

export default router;

import { useUserStore } from "@/stores/user";

router.beforeEach((to, from, next) => {
  const userStore = useUserStore();
  const isLogin = userStore.is_user_login;

  // 定义需要重定向的公开页面
  const publicPages = ["Login", "Register", "ResetPassword"];
  const isPublicPage = publicPages.includes(to.name as string);

  // 情况一：用户已登录，但访问的是登录/注册/重置密码页
  if (isLogin && isPublicPage) {
    // 强制跳转到主页
    next({ name: "Home" });
  }
  // 情况二：需要登录的页面
  else if (to.meta.requiresAuth && !isLogin) {
    // 未登录，但访问的是需要鉴权的页面，重定向到登录页
    next({
      name: "Login",
      query: { redirect: to.fullPath },
    });
  }
  // 情况三：其他情况，正常放行
  else {
    next();
  }
});
