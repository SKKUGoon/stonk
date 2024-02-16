<script lang="ts">
  import { Chart } from "chart.js/auto";
  import { onMount } from "svelte";
	import type { periodProfit } from "./type";
	import { generateHistory } from "./chartData";

  
  export let pp: { [ key: string ]: periodProfit };
  export let accntBal: number;

  // Define a variable to bind to the canvas element
  let canvasElement: HTMLCanvasElement;
  let chart: Chart | null = null;

  let chartValues = {
    data: [] as number[],
    label: [] as string[],
  };

  // Set date
  let now = new Date();
  const year = now.getFullYear().toString();
  const month = (now.getMonth() + 1).toString().padStart(2, '0');
  const day = now.getDate().toString().padStart(2, '0');
  let nowString = `${year}${month}${day}`;

  $: if (pp && "payload0" in pp && typeof (accntBal) !== undefined) {
    let accntTS: periodProfit = pp['payload0'];

    let hasToday: boolean = false;
    for (let trade of accntTS.Output1) {
      // Realized profit amount
      chartValues.data.push(+trade.ovrs_rlzt_pfls_amt);
      chartValues.label.push(trade.trad_day);

      if (trade.trad_day === nowString) {
        hasToday = true;
      }
    }

    // Reverse
    chartValues.data = generateHistory(chartValues.data.reverse(), accntBal);
    chartValues.label = chartValues.label.reverse();

    if (!hasToday) {
      chartValues.data.push(accntBal);
      chartValues.label.push(nowString);
    }

    if (chart) {
      chart.data.labels = chartValues.label;
      chart.data.datasets.forEach((dataset) => {
        if (dataset.label === 'Account Balance') {
          dataset.data = chartValues.data;
          console.log(chartValues.data);
        }
      });
      chart.update();
    }
  } 


  onMount(async () => {
    if (canvasElement) {
      chart = new Chart(
        canvasElement,
        {
          type: 'line',
          options: {
            responsive: true,
          },
          data: {
            labels: [],
            datasets: [
              {
                label: 'Account Balance',
                data: [],
                backgroundColor: 'rgb(255, 99, 132)',
                borderColor: 'rgb(255, 99, 132)',
                lineTension: 0.4,  // Works fine
              },
            ],
          },
        }
      );
    };
  });
</script>

<div style="width: 100%">
  <canvas bind:this={canvasElement} />
</div>
  