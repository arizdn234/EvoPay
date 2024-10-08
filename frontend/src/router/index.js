import { createRouter, createWebHistory } from 'vue-router'
import DocumentationView from '@/views/DocumentationView.vue'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import VerifyEmailView from '@/views/VerifyEmailView.vue'
import ResetPasswordView from '@/views/ResetPasswordView.vue'
import UserHomepage from '@/views/UserHomepage.vue'
import UserProfile from '@/views/UserProfile.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // {
    //   path: '/',
    //   name: 'home',
    //   component: HomeView
    // },
    // {
    //   path: '/about',
    //   name: 'about',
    //   component: () => import('../views/AboutView.vue')
    // },
    {
      path: '/',
      name: 'documentation',
      component: DocumentationView
    },
    {
      path: '/users/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/users/register',
      name: 'register',
      component: RegisterView
    },
    {
      path: '/users/verify-email',
      name: 'verify-email',
      component: VerifyEmailView
    },
    {
      path: '/users/reset-password',
      name: 'reset-password',
      component: ResetPasswordView
    },
    {
      path: '/users/homepage',
      name: 'user-homepage',
      component: UserHomepage
    },
    {
      path: '/users/me',
      name: 'user-profile',
      component: UserProfile
    }
  ]
})

export default router
