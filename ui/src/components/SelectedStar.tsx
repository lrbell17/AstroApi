import React from 'react';
import { Star } from '../types';
import StarSystemDiagram from './StarSystemDiagram';

type Props = {
  star: Star | null;
};

const SelectedStar: React.FC<Props> = ({ star }) => {
  if (!star) return null;

  return (
    <div>
      <h2 className="title">{star.name}</h2>
      <StarSystemDiagram star={star} />
    </div>
  );
};

export default SelectedStar;
