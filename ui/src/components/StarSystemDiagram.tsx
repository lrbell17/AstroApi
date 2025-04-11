import React from 'react';
import { Star } from '../types';
import '../styles/StarSystemDiagram.css';

type Props = {
  star: Star;
};

const StarSystemDiagram: React.FC<Props> = ({ star }) => {
  const fixedStarRadius = 12;
  const baseSvgSize = 600;

  // Get the maximum radius needed to dynamically scale diagram
  const maxOrbitAU = Math.max(
    ...star.planets.map(p => p.distance.value),
    star.habitable_zone_upper_bound?.value ?? 0
  );

  const margin = 50;
  const diagramRadius = baseSvgSize / 2 - margin;
  const AU_TO_PX = diagramRadius / maxOrbitAU;

  const center = baseSvgSize / 2;

  const getStarColor = (temp: number) => {
    if (temp > 10000) return '#a2c0ff'; // Blue
    if (temp > 7500) return '#fefeff';  // White
    if (temp > 6000) return '#fce9a5';  // Yellow
    if (temp > 4500) return '#f4b06a';  // Orange
    return '#ff5e5e';                   // Red
  };

  const habZoneInner = (star.habitable_zone_lower_bound?.value ?? 0) * AU_TO_PX + 2*fixedStarRadius;
  const habZoneOuter = (star.habitable_zone_upper_bound?.value ?? 0) * AU_TO_PX + 2*fixedStarRadius;

  return (
    <div className="star-diagram-container">
      <svg width={baseSvgSize} height={baseSvgSize}>
        <defs>
          <mask id="habitable-zone-mask">
            <rect width="100%" height="100%" fill="white" />
            <circle
              cx={center}
              cy={center}
              r={habZoneInner}
              fill="black"
            />
          </mask>
        </defs>

        {/* Habitable Zone */}
        {(habZoneInner > 0 && habZoneOuter > 0) && (
          <circle
            cx={center}
            cy={center}
            r={habZoneOuter}
            fill="rgba(0, 162, 255, 0.15)"
            mask="url(#habitable-zone-mask)"
          />
        )}

        {/* Star */}
        <circle
          cx={center}
          cy={center}
          r={fixedStarRadius}
          fill={getStarColor(star.temp?.value ?? 5800)}
        >
          <title>{star.name + '\nTemperature: ' + star.temp.value + ' K'}</title>
        </circle>

        {/* Planets and Orbits */}
        {star.planets.map((planet) => {
          const distance = planet.distance.value * AU_TO_PX + 2*fixedStarRadius;
          const orbitalPeriod = planet.orbital_period?.value ?? 365;
          const speedScale = 10; // Lower = faster
          const duration = orbitalPeriod / speedScale;

          return (
            <g key={planet.name}>
              {/* Orbit Path */}
              <circle
                cx={center}
                cy={center}
                r={distance}
                fill="none"
                stroke="#999"
                strokeDasharray="4 4"
              />

              {/* Rotating Planet */}
              <g
                className="planet-orbit"
                style={{ animationDuration: `${duration}s` }}
              >
                <circle
                  cx={center + distance}
                  cy={center}
                  r={6}
                  fill="#4fc3f7"
                >
                  <title>
                    {planet.name +
                      '\nDistance: ' + planet.distance.value + ' AU' +
                      '\nOrbital Period: ' + orbitalPeriod + ' days' +
                      '\nMass: ' + (planet.mass?.value ?? 'N/A') + ' ' + (planet.mass?.unit ?? '')}
                  </title>
                </circle>
              </g>
            </g>
          );
        })}
      </svg>
    </div>
  );
};

export default StarSystemDiagram;
