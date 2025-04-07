import React, { useEffect, useRef } from 'react';

const StarryBackground: React.FC = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  
  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    
    const ctx = canvas.getContext('2d');
    if (!ctx) return;
    
    // Set canvas to full screen
    const resizeCanvas = () => {
      canvas.width = window.innerWidth;
      canvas.height = window.innerHeight;
    };
    
    resizeCanvas();
    window.addEventListener('resize', resizeCanvas);
    
    // Create stars
    const stars: {
      x: number, 
      y: number, 
      size: number, 
      brightness: number, 
      twinkleSpeed: number,
      twinkleDirection: number,
      maxBrightness: number,
      minBrightness: number
    }[] = [];
    
    const createStars = () => {
      const starCount = Math.floor(canvas.width * canvas.height / 2500);
      stars.length = 0; 
      
      for (let i = 0; i < starCount; i++) {
        const maxBrightness = 0.7 + Math.random() * 0.3; 
        stars.push({
          x: Math.random() * canvas.width,
          y: Math.random() * canvas.height,
          size: Math.random() * 2.5 + 0.5, 
          brightness: Math.random() * maxBrightness,
          twinkleSpeed: 0.01 + Math.random() * 0.003, 
          twinkleDirection: Math.random() > 0.5 ? 1 : -1, 
          maxBrightness: maxBrightness,
          minBrightness: Math.random() * 0.2
        });
      }
    };
    
    createStars();
    window.addEventListener('resize', createStars);
    
    // Animation
    let animationFrameId: number;
    let lastTime = 0;
    
    const render = (time: number) => {
      const deltaTime = time - lastTime || 0;
      lastTime = time;
      
      ctx.clearRect(0, 0, canvas.width, canvas.height);
      ctx.fillStyle = "#000810"; // Very dark blue background
      ctx.fillRect(0, 0, canvas.width, canvas.height);
      
      // Draw stars
      stars.forEach(star => {
        // Update twinkle - more pronounced effect
        star.brightness += star.twinkleDirection * star.twinkleSpeed * (deltaTime / 16);
        
        // Change direction when reaching brightness limits
        if (star.brightness <= star.minBrightness) {
          star.brightness = star.minBrightness;
          star.twinkleDirection = 1;
        } else if (star.brightness >= star.maxBrightness) {
          star.brightness = star.maxBrightness;
          star.twinkleDirection = -1;
        }
        
        // Occasionally change direction randomly for more dynamic twinkling
        if (Math.random() < 0.005) {
          star.twinkleDirection *= -1;
        }
        
        // Draw star with glow effect for brighter stars
        const opacity = star.brightness;
        
        // Draw glow for larger/brighter stars
        if (star.size > 1.5) {
          const gradient = ctx.createRadialGradient(
            star.x, star.y, 0,
            star.x, star.y, star.size * 3
          );
          gradient.addColorStop(0, `rgba(255, 255, 255, ${opacity * 0.7})`);
          gradient.addColorStop(1, 'rgba(255, 255, 255, 0)');
          
          ctx.fillStyle = gradient;
          ctx.beginPath();
          ctx.arc(star.x, star.y, star.size * 3, 0, Math.PI * 2);
          ctx.fill();
        }
        
        // Draw star core
        ctx.fillStyle = `rgba(255, 255, 255, ${opacity})`;
        ctx.beginPath();
        ctx.arc(star.x, star.y, star.size, 0, Math.PI * 2);
        ctx.fill();
      });
      
      animationFrameId = window.requestAnimationFrame(render);
    };
    
    animationFrameId = window.requestAnimationFrame(render);
    
    // Cleanup
    return () => {
      window.cancelAnimationFrame(animationFrameId);
      window.removeEventListener('resize', resizeCanvas);
      window.removeEventListener('resize', createStars);
    };
  }, []);
  
  return (
    <canvas
      ref={canvasRef}
      style={{
        position: 'fixed',
        top: 0,
        left: 0,
        width: '100%',
        height: '100%',
        zIndex: -1,
        pointerEvents: 'none'  // Let clicks pass through
      }}
    />
  );
};

export default StarryBackground;