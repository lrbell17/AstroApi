import React, { useState } from 'react';
import StarSearch from './components/StarSearch';
import Login from './components/Login';
import { Star } from './types';

const App: React.FC = () => {
  const [selectedStar, setSelectedStar] = useState<Star | null>(null);

  const handleStarSelect = (star: Star) => {
    setSelectedStar(star);
  };

  return (
    <div>
      <h1>Star Search</h1>
      <Login />
      <StarSearch onSelect={handleStarSelect} />
      {selectedStar && <div>Selected Star: {selectedStar.name}</div>}
    </div>
  );
};

export default App;
