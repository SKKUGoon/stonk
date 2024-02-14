<script lang="ts">
  import { Chart } from "chart.js/auto";
  import { onMount } from "svelte";

  
  // Define a variable to bind to the canvas element
  let canvasElement: HTMLCanvasElement;
  let chart: Chart | null = null;

  let chartValues = {
    data: [20, 10, 5, 2, 20, 30, 45],
    label: [
      "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul"
    ],
  };

  onMount(() => {
    if (canvasElement) {
      chart = new Chart(
        canvasElement,
        {
          type: 'line',
          data: {
            labels: chartValues.label,
            datasets: [
              {
                label: 'test by year',
                data: chartValues.data,
                backgroundColor: 'rgb(255, 99, 132)',
                borderColor: 'rgb(255, 99, 132)',
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
  