import React, { useEffect, useState } from 'react';
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
  useLocation,
} from "react-router-dom";
import SearchPage from './pages/SearchPage';
import Login from './pages/Login';
import StarryBackground from './components/StarryBackground';

// Protected route component
const ProtectedRoute: React.FC<{
  isAuthenticated: boolean;
  authChecking: boolean;
  children: React.ReactNode;
}> = ({ isAuthenticated, authChecking, children }) => {
  const location = useLocation();
  
  if (authChecking) {
    return <div className="flex justify-center items-center h-screen">Loading...</div>;
  }
  
  if (!isAuthenticated) {
    return <Navigate to="/" state={{ from: location }} replace />;
  }
  
  return <>{children}</>;
};

const App: React.FC = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [authChecking, setAuthChecking] = useState(true);

  // Check for valid session on mount
  useEffect(() => {
    const checkAuth = async () => {
      try {
        const response = await fetch(`${import.meta.env.VITE_ASTRO_API_URL}/api/session`, {
          credentials: "include",
        });
        setIsAuthenticated(response.ok);
      } catch (error) {
        console.error("Authentication check failed:", error);
        setIsAuthenticated(false);
      } finally {
        setAuthChecking(false);
      }
    };
    
    checkAuth();
  }, []);

  return (
    <>
      <StarryBackground />
      <Router>
        <Routes>
          <Route
            path="/"
            element={
              authChecking ? (
                <div className="flex justify-center items-center h-screen">Loading...</div>
              ) : isAuthenticated ? (
                <Navigate to="/search" replace />
              ) : (
                <Login onLogin={() => setIsAuthenticated(true)} />
              )
            }
          />
          <Route
            path="/search"
            element={
              <ProtectedRoute isAuthenticated={isAuthenticated} authChecking={authChecking}>
                <SearchPage />
              </ProtectedRoute>
            }
          />
          <Route
            path="*"
            element={
              <div className="flex flex-col justify-center items-center h-screen text-white">
                <h1 className="text-4xl mb-4">Page Not Found</h1>
                <p className="mb-6">The page you're looking for doesn't exist.</p>
                <button 
                  onClick={() => window.location.href = '/'}
                  className="px-4 py-2 bg-blue-600 hover:bg-blue-700 rounded"
                >
                  Return Home
                </button>
              </div>
            }
          />
        </Routes>
      </Router>
    </>
  );
};

export default App;