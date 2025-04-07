import React from 'react';
import { Star } from '../types';

type Props = {
  star: Star | null;
};

const SelectedStar: React.FC<Props> = ({ star }) => {
  if (!star) return null;

  return (
    <div className="mt-4 p-4 border rounded shadow bg-white max-w-md mx-auto">
      <h2 className="text-xl font-bold mb-2">{star.name}</h2>
      <p><strong>Mass:</strong> {star.mass?.value ?? 'N/A'}  {star.mass?.unit ?? ''}</p>
      <p><strong>Radius:</strong> {star.radius?.value ?? 'N/A'} {star.radius?.unit ?? ''}</p>
      <p><strong>Temperature:</strong> {star.temp?.value ?? 'N/A'} {star.temp?.unit ?? ''}</p>
      <p><strong>Luminosity:</strong> {star.luminosity?.value ?? 'N/A'} {star.luminosity?.unit ?? ''}</p>
      <p><strong>Habitable Zone:</strong> {star.habitable_zone_lower_bound?.value ?? 'N/A'} {star.habitable_zone_lower_bound?.unit ?? ''} - {star.habitable_zone_upper_bound?.value ?? 'N/A'} {star.habitable_zone_upper_bound?.unit ?? ''}</p>
    </div>
  );
};

export default SelectedStar;
