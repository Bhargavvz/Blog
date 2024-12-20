import { defineStore } from 'pinia'

interface User {
  id: string
  email: string
  isAdmin: boolean
}

interface AuthState {
  isAuthenticated: boolean
  user: User | null
  token: string | null
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    isAuthenticated: false,
    user: null,
    token: null,
  }),

  actions: {
    login(userData: { token: string; user: User }) {
      this.isAuthenticated = true;
      this.user = userData.user;
      this.token = userData.token;
      localStorage.setItem('token', userData.token);
      localStorage.setItem('user', JSON.stringify(userData.user));
    },

    logout() {
      this.isAuthenticated = false;
      this.user = null;
      this.token = null;
      localStorage.removeItem('token');
      localStorage.removeItem('user');
    },

    initializeFromStorage() {
      const token = localStorage.getItem('token');
      const userStr = localStorage.getItem('user');

      if (token && userStr) {
        try {
          const user = JSON.parse(userStr);
          this.isAuthenticated = true;
          this.token = token;
          this.user = user;
        } catch (error) {
          console.error('Error parsing user data:', error);
          // Clear invalid data
          this.logout();
        }
      } else {
        // No stored data or incomplete data
        this.logout();
      }
    }
  },

  getters: {
    isAdmin: (state): boolean => state.user?.isAdmin ?? false,
    currentUser: (state): User | null => state.user
  }
}); 