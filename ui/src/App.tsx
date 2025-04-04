import { useState } from "react";
import StarSearch from "./components/StarSearch";
import { Star } from "./types";

function App() {
  const [selectedStar, setSelectedStar] = useState<Star | null>(null);

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4" >Exoplanet Explorer</h1>

      <StarSearch onSelect={(star: Star) => setSelectedStar(star)} />

      {selectedStar && (
        <div className="mt-6 p-4 border rounded shadow bg-white">
          <h2 className="text-xl font-semibold">Selected Star</h2>
          <p><strong>Name:</strong> {selectedStar.name}</p>
          <p><strong>ID:</strong> {selectedStar.id}</p>
        </div>
      )}
    </div>
  );
}

export default App;