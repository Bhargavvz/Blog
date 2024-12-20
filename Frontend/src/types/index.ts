export interface Post {
  id: string;
  title: string;
  slug: string;
  content: string;
  published: boolean;
  createdAt: string;
  updatedAt: string;
  // Add other fields as needed
}

export interface User {
  id: string;
  email: string;
  role: 'admin' | 'user';
}