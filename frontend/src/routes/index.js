import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import SignUp from '../views/SignUp.vue'
import Profile from '../views/Profile.vue';
import Create_Username from '@/views/Create_Username.vue';
import Create_Avatar from '@/views/Create_Avatar.vue';
import Home from '@/views/Home.vue';


const routes = [
  {
    path: '/profile',
    name: 'Profile',
    component: Profile,
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/create',
    name: 'SignUp',
    component: SignUp,
  },
  {
    path: '/create/username',
    name: 'Create_Username',
    component: Create_Username,
  },
  {
    path: '/create/avatar',
    name: 'Create_Avatar',
    component: Create_Avatar,
  },
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/',
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
