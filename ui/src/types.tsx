export type Star = {
    id: number;
    name: string;
    mass: MeasuredValue;
    radius: MeasuredValue;
    temp: MeasuredValue;
    luminosity: MeasuredValue;
    habitable_zone_lower_bound: MeasuredValue;
    habitable_zone_upper_bound: MeasuredValue;
    planets: Planet[];
};

export type Planet = {
  name: string;
  mass: MeasuredValue;
  radius: MeasuredValue;
  distance: MeasuredValue;
}

export type MeasuredValue = {
  value: number;
  unit: string;
};
