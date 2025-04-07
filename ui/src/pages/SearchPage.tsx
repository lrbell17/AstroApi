import React, { useState } from 'react';
import StarSearch from '../components/StarSearch';
import SelectedStar from '../components/SelectedStar';
import { Star } from '../types';

const SearchPage: React.FC = () => {
  const [selectedStar, setSelectedStar] = useState<Star | null>(null);

  return (
    <div>
      <StarSearch onSelect={setSelectedStar} />
      <SelectedStar star={selectedStar} />
    </div>
  );
};

export default SearchPage;