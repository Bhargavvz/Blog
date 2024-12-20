import { ref, computed } from 'vue';
import type { User } from '../types';

const currentUser = ref<User | null>(null);

export function useAuth() {
  const isAdmin = computed(() => currentUser.value?.role === 'admin');
  const isAuthenticated = computed(() => !!currentUser.value);

  const login = async (email: string, password: string) => {
    // Implement authentication logic here
    currentUser.value = {
      id: '1',
      email,
      role: 'admin'
    };
  };

  const logout = () => {
    currentUser.value = null;
  };

  return {
    currentUser,
    isAdmin,
    isAuthenticated,
    login,
    logout
  };
}