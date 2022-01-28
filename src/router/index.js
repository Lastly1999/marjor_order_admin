import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from "../views/Login"
import List from "../views/List"

Vue.use(VueRouter)

const routes = [
  {
    path: "/login",
    component: Login
  },
  {
    path: "/list",
    component: List
  },
  {
    path: '/',
    redirect: '/login'
  },
]

const router = new VueRouter({
  mode: 'hash',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, _from, next) => {
  const loginAuth = localStorage.getItem("userToken")
  if (to.path === "/login") {
    next()
  } else {
    if (!loginAuth) {
      next("/login")
    }
    // 放行
    next()
  }
})

export default router
