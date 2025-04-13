import React, { useRef, useState } from 'react';
import { Star } from '../types';
import '../styles/StarSystemDiagram.css';

type TooltipType = 'star' | 'planet' | 'habitable' | null;

interface SvgTooltipProps {
  x: number;
  y: number;
  title: string;
  content: React.ReactNode;
}

const SvgTooltip: React.FC<SvgTooltipProps> = ({ x, y, title, content }) => {
  if (typeof content === 'string') {
    return (
      <g transform={`translate(${x}, ${y})`}>
        <rect x="-100" y="-45" width="200" height="40" rx="5" ry="5" fill="rgba(0, 0, 0, 0.8)" filter="url(#drop-shadow)" />
        <text x="0" y="-30" textAnchor="middle" fill="white" fontWeight="bold" fontSize="12">{title}</text>
        <text x="0" y="-15" textAnchor="middle" fill="white" fontSize="10">{content}</text>
      </g>
    );
  }

  if (typeof content === 'object' && content !== null) {
    const lines = Object.entries(content);
    const lineHeight = 15;
    const boxHeight = 30 + (lines.length * lineHeight);

    return (
      <g transform={`translate(${x}, ${y})`}>
        <rect x="-100" y="-55" width="200" height={boxHeight} rx="5" ry="5" fill="rgba(0, 0, 0, 0.8)" filter="url(#drop-shadow)" />
        <text x="0" y="-35" textAnchor="middle" fill="white" fontWeight="bold" fontSize="12">{title}</text>
        {lines.map(([key, value], index) => (
          <text key={key} x="-95" y={-20 + (index * lineHeight)} textAnchor="start" fill="white" fontSize="10">
            <tspan fontWeight="bold">{key}: </tspan>
            <tspan>{value}</tspan>
          </text>
        ))}
      </g>
    );
  }

  return null;
};

type Props = {
  star: Star;
};

const StarSystemDiagram: React.FC<Props> = ({ star }) => {
  const fixedStarRadius = 12;
  const baseSvgSize = 600;
  const center = baseSvgSize / 2;

  const [tooltipType, setTooltipType] = useState<TooltipType>(null);
  const [tooltipPosition, setTooltipPosition] = useState({ x: 0, y: 0 });
  const [tooltipData, setTooltipData] = useState<{ title: string, content: any }>({ title: '', content: '' });
  const [activeElement, setActiveElement] = useState<string | null>(null);
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

  const getTooltipContent = (type: TooltipType, index: number = 0) => {
    switch (type) {
      case 'star':
        return {
          title: star.name,
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
        if (index >= 0 && index < star.planets.length) {
          const planet = star.planets[index];
          return {
            title: planet.name,
            content: {
              "Mass": `${formatValue(planet.mass?.value)} ${planet.mass?.unit ?? ''}`,
              "Radius": `${formatValue(planet.radius?.value)} ${planet.radius?.unit ?? ''}`,
              "Distance": `${formatValue(planet.distance?.value)} ${planet.distance?.unit ?? ''}`,
              "Orbital Period": `${formatValue(planet.orbital_period?.value)} ${planet.orbital_period?.unit ?? ''}`
            }
          };
        }
        return { title: '', content: '' };
      default:
        return { title: '', content: '' };
    }
  };

  const updateTooltipPosition = (x: number, y: number) => {
    if (tooltipMoveTimeout.current) clearTimeout(tooltipMoveTimeout.current);
    tooltipMoveTimeout.current = window.setTimeout(() => {
      setTooltipPosition({ x, y });
    }, 10);
  };

  const handleMouseMove = (e: React.MouseEvent<SVGSVGElement>) => {
    if (activeElement) return;

    const svgRect = e.currentTarget.getBoundingClientRect();
    const x = e.clientX - svgRect.left;
    const y = e.clientY - svgRect.top;

    const dx = x - center;
    const dy = y - center;
    const distance = Math.sqrt(dx * dx + dy * dy);

    for (let i = 0; i < star.planets.length; i++) {
      const planet = star.planets[i];
      const orbitRadius = planet.distance.value * AU_TO_PX + 2 * fixedStarRadius;
      if (Math.abs(distance - orbitRadius) < 3) {
        setTooltipType('planet');
        setTooltipData(getTooltipContent('planet', i));
        updateTooltipPosition(x, y - 5);
        return;
      }
    }

    if (distance >= habZoneInner && distance <= habZoneOuter) {
      setTooltipType('habitable');
      setTooltipData(getTooltipContent('habitable'));
      updateTooltipPosition(x, y - 5);
    } else if (distance <= fixedStarRadius) {
      setTooltipType('star');
      setTooltipData(getTooltipContent('star'));
      updateTooltipPosition(x, y - 5);
    } else {
      setTooltipType(null);
      setTooltipData({ title: '', content: '' });
    }
  };

  const handleElementMouseEnter = (type: TooltipType, id: string, index: number, e: React.MouseEvent<SVGElement>) => {
    const svgRect = (e.currentTarget.ownerSVGElement as SVGSVGElement).getBoundingClientRect();
    const x = e.clientX - svgRect.left;
    const y = e.clientY - svgRect.top;

    setActiveElement(id);
    setTooltipType(type);
    setTooltipData(getTooltipContent(type, index));
    updateTooltipPosition(x, y - 5);
  };

  const handleElementMouseMove = (e: React.MouseEvent<SVGElement>) => {
    const svgRect = (e.currentTarget.ownerSVGElement as SVGSVGElement).getBoundingClientRect();
    const x = e.clientX - svgRect.left;
    const y = e.clientY - svgRect.top;
    updateTooltipPosition(x, y - 5);
  };

  const handleElementMouseLeave = () => {
    setActiveElement(null);
    setTooltipType(null);
    setTooltipData({ title: '', content: '' });
  };

  return (
    <div className="star-diagram-container">
      <svg
        width={baseSvgSize}
        height={baseSvgSize}
        onMouseMove={handleMouseMove}
        onMouseLeave={() => {
          setTooltipType(null);
          setTooltipData({ title: '', content: '' });
          setActiveElement(null);
        }}
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

        <circle
          cx={center}
          cy={center}
          r={fixedStarRadius}
          fill={getStarColor(star.temp?.value ?? 5800)}
          onMouseEnter={(e) => handleElementMouseEnter('star', 'star', -1, e)}
          onMouseMove={handleElementMouseMove}
          onMouseLeave={handleElementMouseLeave}
        />

        {star.planets.map((planet, index) => {
          const distance = planet.distance.value * AU_TO_PX + 2 * fixedStarRadius;
          const orbitalPeriod = planet.orbital_period?.value ?? 365;
          const speedScale = 10;
          const duration = orbitalPeriod / speedScale;

          return (
            <g key={planet.name}>
              <circle
                cx={center}
                cy={center}
                r={distance}
                fill="none"
                stroke="#999"
                strokeDasharray="4 4"
                strokeWidth="2"
                onMouseEnter={(e) => handleElementMouseEnter('planet', `planet-${index}`, index, e)}
                onMouseMove={handleElementMouseMove}
                onMouseLeave={handleElementMouseLeave}
                style={{ cursor: 'pointer' }}
              />
              <g className="planet-orbit" style={{ animationDuration: `${duration}s` }}>
                <circle
                  cx={center + distance}
                  cy={center}
                  r={6}
                  fill="#4fc3f7"
                  onMouseEnter={(e) => handleElementMouseEnter('planet', `planet-${index}`, index, e)}
                  onMouseMove={handleElementMouseMove}
                  onMouseLeave={handleElementMouseLeave}
                  style={{ cursor: 'pointer' }}
                />
              </g>
            </g>
          );
        })}

        {tooltipType && (
          <SvgTooltip
            x={tooltipPosition.x}
            y={tooltipPosition.y}
            title={tooltipData.title}
            content={tooltipData.content}
          />
        )}
      </svg>
    </div>
  );
};

export default StarSystemDiagram;
