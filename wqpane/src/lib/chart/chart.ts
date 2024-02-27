/* eslint-disable @typescript-eslint/no-explicit-any */

export type chartJSDataset = {
  labels: string[];
  datasets: {
    label?: string;
    data: any[];
  }[];
}