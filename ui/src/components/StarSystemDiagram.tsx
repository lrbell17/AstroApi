import React, { useRef, useState } from 'react';
import { Star } from '../types';
import SvgTooltip from './Tooltip';
import '../styles/StarSystemDiagram.css';

interface Props {
  star: Star;
}

const StarSystemDiagram: React.FC<Props> = ({ star }) => {
  const fixedStarRadius = 12;
  const baseSvgSize = 600;
  const center = baseSvgSize / 2;

  const [tooltipData, setTooltipData] = useState<{ title: string, content: any } | null>(null);
  const [tooltipPosition, setTooltipPosition] = useState({ x: 0, y: 0 });
  const tooltipMoveTimeout = useRef<number | null>(null);

  const maxOrbitAU = Math.max(
    ...star.planets.map(p => p.distance.value),
    star.habitable_zone_upper_bound?.value ?? 0
  );

  const margin = 50;
  const diagramRadius = baseSvgSize / 2 - margin;
  const AU_TO_PX = diagramRadius / maxOrbitAU;

  const habZoneInner = (star.habitable_zone_lower_bound?.value ?? 0) * AU_TO_PX + 2 * fixedStarRadius;
  const habZoneOuter = (star.habitable_zone_upper_bound?.value ?? 0) * AU_TO_PX + 2 * fixedStarRadius;

  const formatValue = (value: number | undefined): string => value && value !== 0 ? value.toFixed(2) : "N/A";

  const getStarColor = (temp: number) => {
    if (temp > 10000) return '#a2c0ff';
    if (temp > 7500) return '#fefeff';
    if (temp > 6000) return '#fce9a5';
    if (temp > 4500) return '#f4b06a';
    return '#ff5e5e';
  };

  const getTooltipContent = (type: 'star' | 'habitable' | 'planet', index: number = 0) => {
    switch (type) {
      case 'star':
        return {
          title: "Star " + star.name,
          content: {
            "Mass": `${formatValue(star.mass?.value)} ${star.mass?.unit ?? ''}`,
            "Radius": `${formatValue(star.radius?.value)} ${star.radius?.unit ?? ''}`,
            "Temperature": `${formatValue(star.temp?.value)} ${star.temp?.unit}`,
            "Luminosity": `${formatValue(star.luminosity?.value)} ${star.luminosity?.unit ?? ''}`
          }
        };
      case 'habitable':
        return {
          title: 'Habitable Zone',
          content: {
            "Inner Boundary": `${formatValue(star.habitable_zone_lower_bound?.value)} ${star.habitable_zone_lower_bound?.unit ?? ''}`,
            "Outer Boundary": `${formatValue(star.habitable_zone_upper_bound?.value)} ${star.habitable_zone_upper_bound?.unit ?? ''}`
          }
        };
      case 'planet':
        const planet = star.planets[index];
        return {
          title: "Planet " + planet.name,
          content: {
            "Mass": `${formatValue(planet.mass?.value)} ${planet.mass?.unit ?? ''}`,
            "Radius": `${formatValue(planet.radius?.value)} ${planet.radius?.unit ?? ''}`,
            "Distance": `${formatValue(planet.distance?.value)} ${planet.distance?.unit ?? ''}`,
            "Orbital Period": `${formatValue(planet.orbital_period?.value)} ${planet.orbital_period?.unit ?? ''}`
          }
        };
    }
  };

  const updateTooltip = (type: 'star' | 'habitable' | 'planet' | null, index: number, x: number, y: number) => {
    if (tooltipMoveTimeout.current) clearTimeout(tooltipMoveTimeout.current);
    tooltipMoveTimeout.current = window.setTimeout(() => {
      if (type === null) {
        setTooltipData(null);
        return;
      }
      setTooltipPosition({ x, y });
      setTooltipData(getTooltipContent(type, index));
    }, 10);
  };

  const handleMouseMove = (e: React.MouseEvent<SVGSVGElement>) => {
    const svgRect = e.currentTarget.getBoundingClientRect();
    const x = e.clientX - svgRect.left;
    const y = e.clientY - svgRect.top;

    const dx = x - center;
    const dy = y - center;
    const distance = Math.sqrt(dx * dx + dy * dy);

    for (let i = 0; i < star.planets.length; i++) {
      const planetRadius = star.planets[i].distance.value * AU_TO_PX + 2 * fixedStarRadius;
      if (Math.abs(distance - planetRadius) < 6) {
        updateTooltip('planet', i, x, y - 5);
        return;
      }
    }

    if (distance >= habZoneInner && distance <= habZoneOuter) {
      updateTooltip('habitable', -1, x, y - 5);
    } else if (distance <= fixedStarRadius) {
      updateTooltip('star', -1, x, y - 5);
    } else {
      updateTooltip(null, -1, x, y);
    }
  };

  return (
    <div className="star-diagram-container">
      <svg
        width={baseSvgSize}
        height={baseSvgSize}
        onMouseMove={handleMouseMove}
        onMouseLeave={() => setTooltipData(null)}
      >
        <defs>
          <mask id="habitable-zone-mask">
            <rect width="100%" height="100%" fill="white" />
            <circle cx={center} cy={center} r={habZoneInner} fill="black" />
          </mask>
          <filter id="drop-shadow" x="-20%" y="-20%" width="140%" height="140%">
            <feDropShadow dx="0" dy="1" stdDeviation="2" floodOpacity="0.3" />
          </filter>
        </defs>

        {(habZoneInner > 0 && habZoneOuter > 0) && (
          <circle cx={center} cy={center} r={habZoneOuter} fill="rgba(0, 162, 255, 0.15)" mask="url(#habitable-zone-mask)" />
        )}

        <circle cx={center} cy={center} r={fixedStarRadius} fill={getStarColor(star.temp?.value ?? 5800)} />

        {star.planets.map((planet) => {
          const distance = planet.distance.value * AU_TO_PX + 2 * fixedStarRadius;
          const orbitalPeriod = planet.orbital_period?.value ?? 365;
          const speedScale = 10;
          const duration = orbitalPeriod / speedScale;

          return (
            <g key={planet.name} >
              <circle
                cx={center}
                cy={center}
                r={distance}
                fill="none"
                stroke="#999"
                strokeDasharray="4 4"
                strokeWidth="2"
              />
              <g className="planet-orbit" style={{ animationDuration: `${duration}s` }}>
                <circle
                    cx={center + distance}
                    cy={center}
                    r={6}
                    fill="#4fc3f7"
                />
              </g>
            </g>
          );
        })}

        {tooltipData && (
          <SvgTooltip
            x={tooltipPosition.x}
            y={tooltipPosition.y}
            title={tooltipData.title}
            content={tooltipData.content}
            svgWidth={baseSvgSize}
            svgHeight={baseSvgSize}
          />
        )}
      </svg>
    </div>
  );
};

export default StarSystemDiagram;
