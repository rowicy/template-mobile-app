# API

This is a template for building APIs using Node.js, Express, and TypeScript.

## Technology Stack

- [Node.js](https://nodejs.org/) v22.14.0
- [Express](https://expressjs.com/) - Web framework
- [TypeScript](https://www.typescriptlang.org/) - Type-safe JavaScript
- [Helmet](https://helmetjs.github.io/) - Security middleware
- [CORS](https://github.com/expressjs/cors) - Cross-origin resource sharing
- [Morgan](https://github.com/expressjs/morgan) - HTTP request logger

## Getting Started

### Prerequisites

- Node.js v22.14.0 or later
- npm or yarn

### Installation

```bash
npm install
```

### Development

Start the development server with hot reload:

```bash
npm run dev
```

The API will be available at `http://localhost:3000`

### Building

Build the TypeScript code:

```bash
npm run build
```

### Production

Run the production build:

```bash
npm start
```

## API Endpoints

### Health Check
- `GET /health` - Health check endpoint

### Main Routes
- `GET /` - API information
- `GET /api/users` - Get all users (sample)
- `POST /api/users` - Create a new user (sample)

## Development Commands

- `npm run dev` - Start development server with hot reload
- `npm run build` - Build TypeScript to JavaScript
- `npm start` - Start production server
- `npm run lint` - Run ESLint
- `npm run format` - Format code with Prettier
- `npm run format:check` - Check code formatting

## Project Structure

```
api/
├── src/
│   └── index.ts          # Main server file
├── dist/                 # Build output (generated)
├── package.json          # Dependencies and scripts
├── tsconfig.json         # TypeScript configuration
├── .eslintrc.js          # ESLint configuration
├── .prettierrc           # Prettier configuration
├── .gitignore            # Git ignore rules
└── README.md             # This file
```

## Environment Variables

Create a `.env` file in the root directory:

```
PORT=3000
NODE_ENV=development
```

## Adding New Routes

1. Create route files in the `src/` directory
2. Import and use them in `src/index.ts`
3. Follow the existing pattern for error handling and middleware
