import type { VolatilitySurfaceAxis, VolatilitySurfaceRequest } from "./type";

const binanceOptionVolSuf = "http://0.0.0.0:10501/api/v1/bn/volatility";

export const getOptionData = async (
  asset: string, 
  callput: "call" | "put"
): Promise<VolatilitySurfaceAxis> => {
  const opt = await fetch(binanceOptionVolSuf, {
    method: "POST",
    body: JSON.stringify({
      underlying: asset,
      callput: callput,
    } as VolatilitySurfaceRequest),
  });

  return opt.json();
}