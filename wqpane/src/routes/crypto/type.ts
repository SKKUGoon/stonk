export type VolatilitySurfaceAxis = {
  rangeX: [number, number];
  rangeY: [number, number];
  rangeZ: [number, number];

  data: VolatilitySurfaceDataPoint[];
}

export type VolatilitySurfaceDataPoint = {
  strike: string; // number-able
  t2mat: number;
  iv: string;     // number-able
}

export type VolatilitySurfaceRequest = {
  underlying: string; // BTCUSDT etc.
  callput: "call" | "put";
}