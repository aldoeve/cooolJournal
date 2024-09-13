import { createRouter, createWebHistory } from 'vue-router'
import Login from '../components/Login.vue'
import SignUp from '../components/SignUp.vue'
import Profile from '../components/Profile.vue';
import Create_Username from '@/components/Create_Username.vue';
import Create_Avatar from '@/components/Create_Avatar.vue';

const routes = [
  {
    path: '/profile',
    name: 'Profile',
    component: Profile,
  },
  {
    path: '/',
    name: 'Login',
    component: Login,
  },
  {
    path: '/signup',
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
  }

];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
