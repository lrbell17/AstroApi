import React, { useState } from 'react';

const baseUrl = import.meta.env.VITE_ASTRO_API_URL;

type Props = {
  onLogin: () => void;
};

const Login: React.FC<Props> = ({ onLogin }) => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleLogin = async () => {
    setLoading(true);
    setError(null);

    // Dummy credentials
    const formData = new FormData();
    formData.append('username', 'admin');
    formData.append('password', 'admin');

    try {
      const response = await fetch(`${baseUrl}/api/login`, {
        method: 'POST',
        body: formData,
        credentials: 'include', // Allows the backend to set the cookie
      });

      if (!response.ok) {
        throw new Error('Login failed');
      }

      alert('Login successful!');
      onLogin();
    } catch (error) {
      setError('Failed to log in');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="login-container">
      <div className="login-card">
        <h1 className="login-title">Explore Distant Stars and Exoplanets</h1>
        <button
          onClick={handleLogin}
          disabled={loading}
          className="login-button"
        >
          {loading ? 'Establishing connection...' : 'Begin Exploration'}
        </button>
        {error && <p className="login-error">{error}</p>}
      </div>
    </div>
  );
};

export default Login;
