import React, { useState } from 'react';
import { Star } from '../types';
import '../styles/StarSystemDiagram.css';

type Props = {
  star: Star;
};

const StarSystemDiagram: React.FC<Props> = ({ star }) => {
  const fixedStarRadius = 12;
  const baseSvgSize = 600;

  const [tooltipVisible, setTooltipVisible] = useState(false);
  const [tooltipPosition, setTooltipPosition] = useState({ x: 0, y: 0 });

  // Get the maximum radius needed to dynamically scale diagram
  const maxOrbitAU = Math.max(
    ...star.planets.map(p => p.distance.value),
    star.habitable_zone_upper_bound?.value ?? 0
  );

  const margin = 50;
  const diagramRadius = baseSvgSize / 2 - margin;
  const AU_TO_PX = diagramRadius / maxOrbitAU;
  const center = baseSvgSize / 2;

  const formatValue = (value: number | undefined): string => {
    if (value === undefined || value === 0) return "N/A";
    return value.toFixed(2);
  };


  const getStarColor = (temp: number) => {
    if (temp > 10000) return '#a2c0ff'; // Blue
    if (temp > 7500) return '#fefeff';  // White
    if (temp > 6000) return '#fce9a5';  // Yellow
    if (temp > 4500) return '#f4b06a';  // Orange
    return '#ff5e5e';                   // Red
  };

  const habZoneInner = (star.habitable_zone_lower_bound?.value ?? 0) * AU_TO_PX + 2*fixedStarRadius;
  const habZoneOuter = (star.habitable_zone_upper_bound?.value ?? 0) * AU_TO_PX + 2*fixedStarRadius;

  const handleMouseMove = (e: React.MouseEvent<SVGSVGElement>) => {
    const svgRect = e.currentTarget.getBoundingClientRect();
    const x = e.clientX - svgRect.left;
    const y = e.clientY - svgRect.top;
    
    const dx = x - center;
    const dy = y - center;
    const distance = Math.sqrt(dx * dx + dy * dy);
    
    if (distance >= habZoneInner && distance <= habZoneOuter) {
      setTooltipVisible(true);
      setTooltipPosition({ x, y: y - 5 }); 
    } else {
      setTooltipVisible(false);
    }
  };

  return (
    <div className="star-diagram-container">
      <svg 
        width={baseSvgSize} 
        height={baseSvgSize}
        onMouseMove={handleMouseMove}
        onMouseLeave={() => setTooltipVisible(false)}
      >
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
          
          {/* Filter for tooltip shadow */}
          <filter id="drop-shadow" x="-20%" y="-20%" width="140%" height="140%">
            <feDropShadow dx="0" dy="1" stdDeviation="2" floodOpacity="0.3" />
          </filter>
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
        
        {/* SVG-based Tooltip */}
        {tooltipVisible && (
          <g transform={`translate(${tooltipPosition.x}, ${tooltipPosition.y})`}>
            <rect
              x="-100"
              y="-45"
              width="200"
              height="40"
              rx="5"
              ry="5"
              fill="rgba(0, 0, 0, 0.8)"
              filter="url(#drop-shadow)"
            />
            <text
              x="0"
              y="-30"
              textAnchor="middle"
              fill="white"
              fontWeight="bold"
              fontSize="12"
            >
              Habitable Zone
            </text>
            <text
              x="0"
              y="-15"
              textAnchor="middle"
              fill="white"
              fontSize="10"
            >
              {formatValue(star.habitable_zone_lower_bound?.value)} {star.habitable_zone_lower_bound?.unit ?? 'AU'} - {formatValue(star.habitable_zone_upper_bound?.value)} {star.habitable_zone_upper_bound?.unit ?? 'AU'}
            </text>
          </g>
        )}
      </svg>
    </div>
  );
};

export default StarSystemDiagram;