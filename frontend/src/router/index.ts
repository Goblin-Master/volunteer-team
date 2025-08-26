import { createRouter, createWebHistory } from 'vue-router'
import Login from "@/views/Login.vue"
import Register from "@/views/Register.vue"
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/register',
      name: 'Register',
      component: Register,
    },
    {
      path: '/',
      redirect: '/login',
    }
  ],
})

export default router
