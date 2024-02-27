export const cumulativeSum = (arr: number[]): number[] => {
  if (arr.length <= 0) {
    return [];
  }

  const cumul = [arr[0]];
  for (let i = 1; i < arr.length; i++) {
    cumul.push(cumul[i-1] + arr[i]);
  }

  return cumul;
}

export const generateHistory = (arr: number[], final: number) => {
  // arr contains information on each trade's realized pnl
  // final is the final account balance
  // Create account balance history
  
  if (arr.length === 0) {
    return [];
  }

  // arr is ordered by time
  const history: number[] = [];
  let mov = final;

  for (let i = arr.length - 1; i >= 0; i--) {
    history.push(mov - arr[i]);
    mov -= arr[i];
  }

  return history.reverse();
}