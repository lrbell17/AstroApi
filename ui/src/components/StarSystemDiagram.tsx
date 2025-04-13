import React, { useState } from 'react';
import { Star } from '../types';
import '../styles/StarSystemDiagram.css';

// Custom tooltip component
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
          {title}
        </text>
        <text
          x="0"
          y="-15"
          textAnchor="middle"
          fill="white"
          fontSize="10"
        >
          {content}
        </text>
      </g>
    );
  }
  
  if (typeof content === 'object' && content !== null) {
    const lines = Object.entries(content);
    const lineHeight = 15;
    const boxHeight = 30 + (lines.length * lineHeight);
    
    return (
      <g transform={`translate(${x}, ${y})`}>
        <rect
          x="-100"
          y="-55"
          width="200"
          height={boxHeight}
          rx="5"
          ry="5"
          fill="rgba(0, 0, 0, 0.8)"
          filter="url(#drop-shadow)"
        />
        <text
          x="0"
          y="-35"
          textAnchor="middle"
          fill="white"
          fontWeight="bold"
          fontSize="12"
        >
          {title}
        </text>
        {lines.map(([key, value], index) => (
          <text
            key={key}
            x="-95" // Left-aligned text
            y={-20 + (index * lineHeight)}
            textAnchor="start"
            fill="white"
            fontSize="10"
          >
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
  
  // Tooltip state
  const [tooltipType, setTooltipType] = useState<TooltipType>(null);
  const [tooltipPosition, setTooltipPosition] = useState({ x: 0, y: 0 });
  const [tooltipData, setTooltipData] = useState<{ title: string, content: any }>({ title: '', content: '' });
  const [activeElement, setActiveElement] = useState<string | null>(null);

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

  const handleMouseMove = (e: React.MouseEvent<SVGSVGElement>) => {
    const svgRect = e.currentTarget.getBoundingClientRect();
    const x = e.clientX - svgRect.left;
    const y = e.clientY - svgRect.top;
    
    const dx = x - center;
    const dy = y - center;
    const distance = Math.sqrt(dx * dx + dy * dy);
    
    // Only update tooltip if no other tooltip is active
    if (activeElement) {
      return;
    }
    
    if (distance <= fixedStarRadius) {
      // Mouse over star
      setTooltipType('star');
      setTooltipPosition({ x, y: y - 5 });
      setTooltipData(getTooltipContent('star'));
    } else if (distance >= habZoneInner && distance <= habZoneOuter) {
      // Mouse over habitable zone
      setTooltipType('habitable');
      setTooltipPosition({ x, y: y - 5 });
      setTooltipData(getTooltipContent('habitable'));
    } else {
      // Check if mouse is over any planet orbit
      let isOverOrbit = false;
      
      for (let i = 0; i < star.planets.length; i++) {
        const planet = star.planets[i];
        const orbitRadius = planet.distance.value * AU_TO_PX + 2*fixedStarRadius;
        const orbitDistance = Math.abs(distance - orbitRadius);
        
        // Consider mouse over orbit if within 3 pixels of orbit path
        if (orbitDistance < 3) {
          isOverOrbit = true;
          setTooltipType('planet');
          setTooltipPosition({ x, y: y - 5 });
          setActiveElement(`planet-${i}`);
          setTooltipData(getTooltipContent('planet', i));
          break;
        }
      }
      
      if (!isOverOrbit) {
        setTooltipType(null);
        setActiveElement(null);
      }
    }
  };

  // Handle mouse interaction with elements
  const handleElementMouseEnter = (elementType: TooltipType, id: string, index: number, e: React.MouseEvent<SVGElement>) => {
    const svgRect = (e.currentTarget.ownerSVGElement as SVGSVGElement).getBoundingClientRect();
    const x = e.clientX - svgRect.left;
    const y = e.clientY - svgRect.top;
    
    setActiveElement(id);
    setTooltipType(elementType);
    setTooltipPosition({ x, y: y - 5 });
    setTooltipData(getTooltipContent(elementType, index));
  };
  
  const handleElementMouseMove = (e: React.MouseEvent<SVGElement>) => {
    if (activeElement) {
      const svgRect = (e.currentTarget.ownerSVGElement as SVGSVGElement).getBoundingClientRect();
      const x = e.clientX - svgRect.left;
      const y = e.clientY - svgRect.top;
      setTooltipPosition({ x, y: y - 5 });
    }
  };
  
  const handleElementMouseLeave = () => {
    setActiveElement(null);
    setTooltipType(null);
  };

  return (
    <div className="star-diagram-container">
      <svg 
        width={baseSvgSize} 
        height={baseSvgSize}
        onMouseMove={handleMouseMove}
        onMouseLeave={() => setTooltipType(null)}
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
          onMouseEnter={(e) => handleElementMouseEnter('star', 'star', -1, e)}
          onMouseMove={handleElementMouseMove}
          onMouseLeave={handleElementMouseLeave}
        />

        {/* Planets and Orbits */}
        {star.planets.map((planet, index) => {
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
                strokeWidth="2"
                onMouseEnter={(e) => handleElementMouseEnter('planet', `planet-${index}`, index, e)}
                onMouseMove={handleElementMouseMove}
                onMouseLeave={handleElementMouseLeave}
                style={{ cursor: 'pointer' }}
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
                  onMouseEnter={(e) => handleElementMouseEnter('planet', `planet-${index}`, index, e)}
                  onMouseMove={handleElementMouseMove}
                  onMouseLeave={handleElementMouseLeave}
                  style={{ cursor: 'pointer' }}
                />
              </g>
            </g>
          );
        })}
        
        {/* SVG-based Tooltip - rendered using the component */}
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