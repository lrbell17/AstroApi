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
    <div>
      <button onClick={handleLogin} disabled={loading}>
        {loading ? 'Logging in...' : 'Access star search'}
      </button>
      {error && <p className="text-red-500">{error}</p>}
    </div>
  );
};

export default Login;
