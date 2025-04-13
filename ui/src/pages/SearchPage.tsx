import React, { useState } from 'react';
import StarSearch from '../components/StarSearch';
import SelectedStar from '../components/SelectedStar';
import { Star } from '../types';
import '../styles/StarSearch.css';

const SearchPage: React.FC = () => {
  const [selectedStar, setSelectedStar] = useState<Star | null>(null);

  return (
    <div className="container">
      <div className="card">
        <StarSearch onSelect={setSelectedStar} />
        <SelectedStar star={selectedStar} />
      </div>
    </div>
  );
};

export default SearchPage;