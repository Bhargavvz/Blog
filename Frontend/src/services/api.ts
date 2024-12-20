import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
});

// Add auth token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export const login = async (email: string, password: string) => {
  const response = await api.post('/login', { email, password });
  return response.data;
};

export const createPost = async (formData: FormData) => {
  // Debug: Log form data contents
  for (const pair of formData.entries()) {
    console.log(`${pair[0]}: ${pair[1]}`);
  }

  const response = await api.post('/posts', formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
  return response.data;
};

export const updatePost = async (id: string, formData: FormData) => {
  const response = await api.put(`/posts/${id}`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
  return response.data;
};

export const getPosts = async () => {
  const response = await api.get('/posts');
  return response.data;
};

export const getPost = async (id: string) => {
  const response = await api.get(`/posts/${id}`);
  return response.data;
};

export const deletePost = async (id: string) => {
  const response = await api.delete(`/posts/${id}`);
  return response.data;
};

export const sendContactMessage = async (data: {
  name: string;
  email: string;
  subject: string;
  message: string;
}) => {
  try {
    console.log('Sending contact request with data:', data)
    const response = await api.post('/contact', data, {
      headers: {
        'Content-Type': 'application/json',
      },
    })
    console.log('Contact response:', response.data)
    return response.data
  } catch (error) {
    console.error('Contact form error:', error)
    throw error
  }
}

export const logout = () => {
  // Clear the stored token
  localStorage.removeItem('token');
  // You might also want to clear any other user-related data
  localStorage.removeItem('user');
};

export default api; 