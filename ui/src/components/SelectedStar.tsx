import React from 'react';
import { Star } from '../types';
import StarSystemDiagram from './StarSystemDiagram';

type Props = {
  star: Star | null;
};

const SelectedStar: React.FC<Props> = ({ star }) => {
  if (!star) return null;

  return (
    <div className="mt-4 p-4 border rounded shadow bg-white max-w-md mx-auto">
      <h2 className="text-xl font-bold mb-2">{star.name}</h2>
      <StarSystemDiagram star={star} />
    </div>
  );
};

export default SelectedStar;
