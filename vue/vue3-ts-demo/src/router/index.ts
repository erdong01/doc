import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      children: [
        {
          path: "goods",
          name: "goods",
          meta: { isShow: true, title: "商品列表" },
          component: () => import('../views/GoodsView.vue')
        },
        {
          path: "user",
          name: "user",
          meta: { isShow: true, title: "用户列表" },
          component: () => import('../views/UserView.vue')
        },
        {
          path: "role",
          name: "role",
          meta: { isShow: true, title: "角色列表" },
          component: () => import('../views/RoleView.vue')
        },
        {
          path: "/authority",
          name: "authority",
          meta: {
            isShow: false,
            title: "权限列表"
          },
          component: () => import('../views/AuthorityView.vue')
        }
      ]
    }, {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    }, {
      path: '/login',
      name: 'login',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/LoginView.vue')
    },
  ]
})

export default router
