# 🚀 Modern Blog Platform

A modern blog platform built with Vue 3, TypeScript, and Go, featuring a clean design and robust content management system.

## 📚 Table of Contents
- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Getting Started](#-getting-started)
- [Project Structure](#-project-structure)
- [Configuration](#-configuration)

## ✨ Features

### Frontend
- 📱 Responsive design with Tailwind CSS
- 🌙 Dark/Light mode support
- 📝 Blog post management
- 📧 Contact form with email notifications
- 🔐 Secure admin authentication

### Backend
- 🔒 JWT authentication
- 📨 SMTP email integration
- 🗄️ MongoDB database
- 🖼️ Image upload support
- 🔄 RESTful API

## 🛠️ Tech Stack

### Frontend
- [Vue 3](https://vuejs.org/) - Progressive JavaScript Framework
- [TypeScript](https://www.typescriptlang.org/) - JavaScript with syntax for types
- [Vite](https://vitejs.dev/) - Next Generation Frontend Tooling
- [Pinia](https://pinia.vuejs.org/) - State Management
- [Vue Router](https://router.vuejs.org/) - Official Vue.js router
- [TailwindCSS](https://tailwindcss.com/) - Utility-first CSS framework

### Backend
- [Go](https://golang.org/) - Backend programming language
- [Gin](https://gin-gonic.com/) - Web framework for Go
- [MongoDB](https://www.mongodb.com/) - NoSQL database
- [JWT-Go](https://github.com/golang-jwt/jwt) - JSON Web Token implementation

## 🚀 Getting Started

### Prerequisites
- Node.js (v16.0 or higher)
- Go (v1.16 or higher)
- MongoDB
- Git

### Frontend Setup

1. Clone the repository
```bash
git clone <repository-url>
cd Frontend
```

2. Install dependencies
```bash
npm install
```

3. Create .env file
```env
VITE_API_URL=http://localhost:8080/api
```

4. Start development server
```bash
npm run dev
```

### Backend Setup

1. Navigate to backend directory
```bash
cd Backend
```

2. Create .env file
```env
PORT=8080
MONGODB_URI=mongodb://localhost:27017/blog
JWT_SECRET=your_jwt_secret
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_FROM=your-email@gmail.com
SMTP_PASSWORD=your-app-password
ADMIN_EMAIL=admin-email@gmail.com
```

3. Run the server
```bash
go run main.go
```

## 🏗️ Project Structure

```
Frontend/
├── src/
│   ├── views/              # Page components
│   ├── components/         # Reusable components
│   ├── stores/            # Pinia stores
│   ├── services/          # API services
│   ├── router/            # Vue Router config
│   └── types/             # TypeScript types
└── ...

Backend/
├── handlers/              # Request handlers
├── middleware/           # Custom middleware
├── models/              # Data models
├── db/                  # Database configuration
└── main.go             # Entry point
```

## ⚙️ Configuration

### Environment Variables

#### Frontend (.env)
```env
VITE_API_URL=http://localhost:8080/api
```

#### Backend (.env)
```env
PORT=8080
MONGODB_URI=mongodb://localhost:27017/blog
JWT_SECRET=your_jwt_secret
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_FROM=your-email@gmail.com
SMTP_PASSWORD=your-app-password
ADMIN_EMAIL=admin-email@gmail.com
```

## 📝 API Endpoints

### Authentication
- `POST /api/login` - Admin login

### Posts
- `GET /api/posts` - Get all posts
- `GET /api/posts/:id` - Get single post
- `POST /api/posts` - Create new post
- `PUT /api/posts/:id` - Update post
- `DELETE /api/posts/:id` - Delete post

### Contact
- `POST /api/contact` - Submit contact form

## 👥 Development Team

- Developer - [Your Name](https://github.com/yourusername)

## 📄 License

This project is licensed under the MIT License.

---

[⬆ back to top](#-table-of-contents)
