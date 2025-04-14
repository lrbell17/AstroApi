  import React, { useRef, useEffect, useState } from 'react';

  interface SvgTooltipProps {
    x: number;
    y: number;
    title: string;
    content: React.ReactNode;
    svgWidth: number;
    svgHeight: number;
  }

  const SvgTooltip: React.FC<SvgTooltipProps> = ({ x, y, svgWidth, svgHeight, content }) => {
    const tooltipRef = useRef<SVGGElement | null>(null); 
    const [adjustedPos, setAdjustedPos] = useState({ x, y });

    useEffect(() => {
      const { x: newX, y: newY } = adjustTooltipPosition(x, y, svgWidth, svgHeight, tooltipRef);
      setAdjustedPos({ x: newX, y: newY });
    }, [x, y, svgWidth, svgHeight, content]);

    return (
      <g ref={tooltipRef} transform={`translate(${adjustedPos.x}, ${adjustedPos.y})`} className="tooltip-group">
        <rect x={-100} y={-30} width={200} height={50} rx={8} fill="#222" stroke="#999" strokeWidth={0.5} />
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
