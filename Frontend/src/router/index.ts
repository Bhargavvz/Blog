import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from '../stores/auth';
import Home from '../views/Home.vue';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/blog',
      name: 'blog',
      component: () => import('../views/Blog.vue'),
    },
    {
      path: '/blog/:slug',
      name: 'blog-post',
      component: () => import('../views/BlogPost.vue'),
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/About.vue'),
    },
    {
      path: '/contact',
      name: 'contact',
      component: () => import('../views/Contact.vue'),
    },
    {
      path: '/admin/login',
      name: 'Login',
      component: () => import('../views/Login.vue')
    },
    {
      path: '/admin',
      name: 'Admin',
      component: () => import('../views/Admin.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/admin/posts/new',
      name: 'admin-new-post',
      component: () => import('../views/AdminEditPost.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/admin/posts/:id/edit',
      name: 'admin-edit-post',
      component: () => import('../views/AdminEditPost.vue'),
      meta: { requiresAuth: true }
    }
  ],
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    }
    return { top: 0 };
  },
});

// Navigation guard
router.beforeEach((to, from, next) => {
  const auth = useAuthStore();
  
  // Check if route requires authentication
  if (to.meta.requiresAuth && !auth.isAuthenticated) {
    // Redirect to login page
    next({ path: '/admin/login' });
  } else if (to.path === '/admin/login' && auth.isAuthenticated) {
    // If already logged in and trying to access login page, redirect to admin
    next({ path: '/admin' });
  } else {
    next();
  }
});

export default router;