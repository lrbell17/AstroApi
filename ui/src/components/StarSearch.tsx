import React, { useState, useEffect } from "react";
import { Star } from "../types";
import '../styles/StarSearch.css';

const baseUrl = import.meta.env.VITE_ASTRO_API_URL;

type Props = {
  onSelect: (star: Star) => void;
};

const StarSearch: React.FC<Props> = ({ onSelect }) => {
  const [query, setQuery] = useState("");
  const [suggestions, setSuggestions] = useState<Star[]>([]);

  useEffect(() => {
    if (query.length === 0) {
      setSuggestions([]);
      return;
    }

    const timeout = setTimeout(() => {
      fetch(`${baseUrl}/api/stars?search=${encodeURIComponent(query)}&limit=10`, {
        credentials: 'include'
      })
        .then((res) => {
          if (!res.ok) {
            throw new Error("Network response was not ok");
          }
          return res.json();  
        })
        .then((data) => {
          setSuggestions(data.stars || []);
        })
        .catch((err) => console.error("Error fetching stars:", err));
    }, 300);

    return () => clearTimeout(timeout);
  }, [query]);

  return (
    <div className="relative w-full max-w-md mx-auto">
      <input
        type="text"
        className="star-search-input"
        placeholder="Search for a star..."
        value={query}
        onChange={(e) => setQuery(e.target.value)}
      />
      {suggestions.length > 0 && (
        <ul className="absolute z-10 w-full bg-white border rounded shadow mt-1 max-h-60 overflow-y-auto">
          {suggestions.map((star) => (
            <li
              key={star.id}
              className="px-4 py-2 hover:bg-gray-100 cursor-pointer"
              onClick={() => {
                onSelect(star);
                setQuery("");
                setSuggestions([]);
              }}
            >
              {star.name}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default StarSearch;