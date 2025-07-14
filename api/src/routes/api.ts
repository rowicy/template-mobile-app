import express from 'express';

const router = express.Router();

// Example users endpoint
router.get('/users', (req, res) => {
  // This is a sample endpoint - replace with your actual logic
  const sampleUsers = [
    { id: 1, name: 'John Doe', email: 'john@example.com' },
    { id: 2, name: 'Jane Smith', email: 'jane@example.com' }
  ];

  res.json({
    success: true,
    data: sampleUsers,
    message: 'Users retrieved successfully'
  });
});

// Example user by ID endpoint
router.get('/users/:id', (req, res) => {
  const userId = parseInt(req.params.id);
  
  // This is a sample endpoint - replace with your actual logic
  const sampleUser = { id: userId, name: 'Sample User', email: 'user@example.com' };

  res.json({
    success: true,
    data: sampleUser,
    message: 'User retrieved successfully'
  });
});

// Example POST endpoint
router.post('/users', (req, res) => {
  const { name, email } = req.body;

  // This is a sample endpoint - replace with your actual logic
  const newUser = {
    id: Date.now(), // Simple ID generation for demo
    name,
    email,
    createdAt: new Date().toISOString()
  };

  res.status(201).json({
    success: true,
    data: newUser,
    message: 'User created successfully'
  });
});

// Example mobile app specific endpoint
router.get('/app-config', (req, res) => {
  res.json({
    success: true,
    data: {
      appName: 'Mobile App Template',
      version: '1.0.0',
      features: {
        authentication: true,
        pushNotifications: false,
        analytics: true
      },
      api: {
        baseUrl: process.env.API_BASE_URL || 'http://localhost:3000',
        version: 'v1'
      }
    },
    message: 'App configuration retrieved successfully'
  });
});

export { router as apiRouter };