# API Template

A TypeScript-based REST API template for mobile app development using Express.js.

## Tech Stack

- **Node.js** v22.14.0+
- **Express.js** - Web framework
- **TypeScript** - Type-safe JavaScript
- **CORS** - Cross-Origin Resource Sharing for mobile apps
- **Helmet** - Security middleware
- **Morgan** - HTTP request logger

## Quick Start

1. **Install dependencies:**
   ```bash
   npm install
   ```

2. **Setup environment variables:**
   ```bash
   cp .env.example .env
   ```
   Edit `.env` file with your configuration.

3. **Development mode:**
   ```bash
   npm run dev
   ```

4. **Build for production:**
   ```bash
   npm run build
   npm start
   ```

## Available Scripts

- `npm run dev` - Start development server with hot reload
- `npm run build` - Build TypeScript to JavaScript
- `npm start` - Start production server
- `npm test` - Run tests
- `npm run lint` - Run ESLint
- `npm run lint:fix` - Fix ESLint issues

## API Endpoints

### Health Check
- `GET /health` - Basic health check
- `GET /health/detailed` - Detailed system information

### API Routes
- `GET /api/v1/users` - Get all users (sample)
- `GET /api/v1/users/:id` - Get user by ID (sample)
- `POST /api/v1/users` - Create new user (sample)
- `GET /api/v1/app-config` - Get mobile app configuration

### Root
- `GET /` - API information and available endpoints

## Project Structure

```
src/
├── index.ts          # Main server file
├── routes/           # Route handlers
│   ├── health.ts     # Health check routes
│   └── api.ts        # Main API routes
├── middleware/       # Custom middleware (add as needed)
├── controllers/      # Route controllers (add as needed)
└── types/           # TypeScript type definitions (add as needed)
```

## Environment Variables

Create a `.env` file based on `.env.example`:

```env
PORT=3000
NODE_ENV=development
API_BASE_URL=http://localhost:3000
```

## Connecting to Mobile App

This API is designed to work seamlessly with the React Native/Expo mobile app in this template:

1. The API includes CORS middleware to handle cross-origin requests
2. Uses JSON responses that are mobile-app friendly
3. Includes app configuration endpoint for mobile app settings
4. Structured error responses for consistent mobile app error handling

## Next Steps

This is a basic template. To extend it for your mobile app needs:

1. **Add Authentication:**
   - Install JWT middleware
   - Create login/register endpoints
   - Add protected routes

2. **Add Database:**
   - Choose your database (PostgreSQL, MongoDB, SQLite, etc.)
   - Add ORM/ODM (Prisma, TypeORM, Mongoose, etc.)
   - Create data models

3. **Add More Features:**
   - File upload handling
   - Push notifications
   - Real-time features (Socket.io)
   - API documentation (Swagger)
   - Rate limiting
   - Input validation

4. **Production Deployment:**
   - Add Docker configuration
   - Set up CI/CD
   - Configure production environment
   - Add monitoring and logging
