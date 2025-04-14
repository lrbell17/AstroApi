  import React, { useRef, useEffect, useState } from 'react';

  interface SvgTooltipProps {
    x: number;
    y: number;
    title: string;
    content: Record<string, string>;
    svgWidth: number;
    svgHeight: number;
  }

  const SvgTooltip: React.FC<SvgTooltipProps> = ({ x, y, title, svgWidth, svgHeight, content }) => {
    const tooltipRef = useRef<SVGGElement | null>(null); 
    const [adjustedPos, setAdjustedPos] = useState({ x, y });

    useEffect(() => {
      const { x: newX, y: newY } = adjustTooltipPosition(x, y, svgWidth, svgHeight, tooltipRef);
      setAdjustedPos({ x: newX, y: newY });
    }, [x, y, svgWidth, svgHeight, content]);

    const { width: windowWidth, height: windowHeight } = calculateTooltipDimensions(content)

    return (
      <g ref={tooltipRef} transform={`translate(${adjustedPos.x}, ${adjustedPos.y})`} className="tooltip-group">
        <rect 
          x={-100} 
          y={-30} 
          width={windowWidth} 
          height={windowHeight} 
          rx={8} 
          fill="#222" 
          stroke="#999" 
          strokeWidth={0.5} />
        <text x="0" y="-35" textAnchor="middle" fill="white" fontWeight="bold" fontSize="12">{title}</text>
        {typeof content === 'object' && content !== null && !Array.isArray(content) ? (
          Object.entries(content).map(([key, value], index) => (
            <text key={key} x="-95" y={-20 + index * 15} textAnchor="start" fill="white" fontSize="10">
              <tspan fontWeight="bold">{key}: </tspan>
              <tspan>{String(value)}</tspan>
            </text>
          ))
        ) : (
          <text x="0" y="-15" textAnchor="middle" fill="white" fontSize="10">
            {String(content)}
          </text>
        )}
      </g>
    );
  };

  // Get tooltip dimentions based on content
  const calculateTooltipDimensions = (content: Record<string, string>) => {
    const entries = Object.entries(content);

    const contentHeight = entries.length * 12 + 10 // Approximate line height per entry
    const charWidth = 6; // Approximate character width in pixels

    let maxLength = 0;
    
    entries.forEach(([key, value]) => {
      const entryLength = `${key}: ${value}`.length * charWidth;
      maxLength = Math.max(maxLength, entryLength);
    });
    
    const contentWidth = maxLength + 10;
    return { width: contentWidth, height: contentHeight };
  };

  // adjust tooltip position to prevent it going out of the SVG boundary
  const adjustTooltipPosition = (x: number, y: number, svgWidth: number, svgHeight: number, tooltipRef: React.RefObject<SVGElement | null>) => {
    const tooltipElement = tooltipRef.current;
    if (!tooltipElement) return { x, y };

    const bbox = (tooltipElement as SVGGElement).getBBox();
    const tooltipWidth = bbox.width;
    const tooltipHeight = bbox.height;

    let adjustedX = x;
    let adjustedY = y;

    if (x + tooltipWidth / 2 > svgWidth) {
      adjustedX = svgWidth - tooltipWidth / 2;
    }
    if (x - tooltipWidth / 2 < 0) {
      adjustedX = tooltipWidth / 2;
    }
    if (y + tooltipHeight / 2 > svgHeight) {
      adjustedY = svgHeight - tooltipHeight / 2;
    }
    if (y - tooltipHeight / 2 < 0) {
      adjustedY = tooltipHeight / 2;
    }

    return { x: adjustedX, y: adjustedY };
  };


  export default SvgTooltip
