import { createRouter, createWebHashHistory } from 'vue-router'
import About from '../views/About.vue'
import News from '../views/News.vue'
import NotFound from '../views/NotFound.vue'
import Page from '../views/Page.vue'
import Parent from '../views/Parent.vue'
import ShopFooter from '../views/ShopFooter.vue'
import ShopMain from '../views/ShopMain.vue'
import ShopTop from '../views/ShopTop.vue'
import StyleOne from '../views/StyleOne.vue'
import StyleTow from '../views/StyleTow.vue'
import User from '../views/User.vue'
// 2. 定义一些路由
// 每个路由都需要映射到一个组件。
// 我们后面再讨论嵌套路由。
const routes = [
  { path: '/', component:()=>import('../views/Home.vue') },
  { path: '/about', component: About ,beforeEach:(to,from)=>{
    console.log(to,from);
  }},
  { path: '/user/:id', component: User },
  { path: '/:path(.*)', component: NotFound },
  { path: '/news/:id+', component: News, name: "news" },
  {
    path: '/parent', component: Parent,
    children: [
      { path: 'styleone', component: StyleOne },
      { path: 'styletow', component: StyleTow },]
  },
  { path: '/page', component: Page },
  {
    path: '/shop', component: {
      default: ShopMain,
      ShopTop: ShopTop,
      ShopFooter: ShopFooter,
    },
    porps:{default:true,ShopFooter:true,ShopTop:false}
  }
]

// 3. 创建路由实例并传递 `routes` 配置
// 你可以在这里输入更多的配置，但我们在这里
// 暂时保持简单
const router = createRouter({
  // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
  history: createWebHashHistory(),
  routes, // `routes: routes` 的缩写
})
router.beforeEach((to,from,next)=>{
  console.log(to);
  console.log(from);
  next()
})
export default router
