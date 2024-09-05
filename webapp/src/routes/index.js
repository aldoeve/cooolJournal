import { createRouter, createWebHistory } from 'vue-router'
import Login from '../components/Login.vue'
import SignUp from '../components/SignUp.vue';

const routes = [
  {
    path: '/',
    name: 'Login',
    component: Login,
  },
  {
    path: '/signup',
    name: 'SignUp',
    component: SignUp,
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
