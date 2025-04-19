### 🪐 Welcome to Axo ✨

Axo is a Restful API for Go, built on top of stdlib and gorm.
It is designed to be simple, fast, and easy to use.
For more information, please visit: https://haume.me/axo

License: MIT
Copyright (c) 2025 Haume
It's not neccesary but i'll be greatful if you give me a star on GitHub and mention me in your project.


## Roadmap
- Auth System **[Priority]**
    - [x] Implement Refresh Token and Access Token
    - [x] Store Tokens in cookies for quick access
    - [x] Implement Mail Verification and 2FA (later)
    - [x] Implement Password Reset or Recovery
    - [x] Implement Account Creation
    - [x] Implement Account Login
    - [x] Implement Account Logout
    - [x] Implement Account Activation
    - [x] Implement Account Deactivation
- Roles & Permissions **[Priority]**
    - [x] Implement Roles and Permissions
    - [x] Implement Method:Route based permissions
    - [x] Implement Default Permissions (default,admin)
    - [x] Implement Permission Handlers
    - [ ] Implement Permission Middleware
    - [ ] Implement Role Hierarchy
    - [ ] Implement Role-based Access Control (RBAC)
- Database Integration *[GORM]*
    - [x] PostgreSQL
    - [ ] MySQL
    - [ ] SQLite
- Image Optimization
    - [x] Realtime Image Resizing
    - [x] Realtime Image Compression
    - [x] Realtime Format Conversion
    - [ ] Build-time Image Optimization *(maybe-later)*
- Logging
    - [x] Archive logs based on date
    - [x] Log to console
- Frontend
    - [x] Serve Vite single page applications. [*/axo/frontends/ServeSpa.go*]
    - [x] Serving static web pages. [*/axo/frontends/ServeStatic.go*]
    - [ ] Serve SSR applications with nodejs. [*/axo/frontends/ServeSSR.go*]
- Payment Systems
    - [ ] Payten(Turkey)
    - [ ] Iyzico
    - [ ] PayTR
    - [ ] WePay
    - [ ] Stripe
- Other
    - [x] Onboarding for env variables
    - [x] Static file server
    - [x] Mail System
    - [ ] Auto SSL with Let's Encrypt
    - [ ] Dockerize the project
- Extensions
    - [ ] E-Commerce module
    - [ ] Real-time app module (chat, notifications, etc.)
