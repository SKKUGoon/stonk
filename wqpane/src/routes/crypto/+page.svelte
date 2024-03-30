<svelte:head>
  <script src="https://cdn.plot.ly/plotly-latest.min.js" type="text/javascript" />
</svelte:head>

<script lang="ts">
  import { onMount } from 'svelte';
  import { getOptionData } from './data';

  const asset1 = "BTCUSDT";  // Downward trend
//   const asset1 = "BTCUSDT";
  const asset2 = "ETHUSDT";
  const asset3 = "XRPUSDT";
  const asset4 = "BNBUSDT";
  const asset5 = "DOGEUSDT";

  onMount(async () => {

    const dataCall = await getOptionData(asset1, "call");
    const dataPut = await getOptionData(asset1, "put");

    const data = [...dataCall.data, ...dataPut.data];

    const strikes = data.map(p => parseFloat(p.strike));
    const timeToMat = data.map(p => (p.t2mat / 1000) / (365 * 24 * 60 * 60)); // Convert seconds to years
    const impliedV = data.map(p => parseFloat(p.iv));
  
    const plotData = [{
      x: strikes,
      y: timeToMat,
      z: impliedV,
      mode: 'markers',
      type: 'scatter3d',
      marker: {
        size: 5,
        color: impliedV, // Color by implied volatility
        colorscale: 'Viridis',
        opacity: 0.8
      }

    //   type: 'surface',
    //   contours: {
    //     z: {
    //         show:true,
    //         usecolormap: true,
    //         highlightcolor:"#42f462",
    //         project:{z: true}
    //     }
    //   }
    }];
    
    const layout = {
      title: `Option Implied Volatility Surface (${asset1})`,
      scene: {
        xaxis: {title: 'Strike Price'},
        yaxis: {title: 'Time to Maturity (Years)'},
        zaxis: {title: 'Implied Volatility'}
      },
      autosize: true
    };

	let plotDiv = document.getElementById(`plotDiv-${asset1}`);
	new Plotly.newPlot(plotDiv, plotData, layout); 
  });
</script>

<div id="plotly">
  <!-- Draw plotly chart inside `plotDiv` -->
  <div id={`plotDiv-${asset1}`} style="width: 600px; height: 600px;" />
</div>

<div class="section">
  <div class="section-title">Volatility Smile:</div>
  <div class="section-content">
    <p><strong>Shape:</strong> A U-shape or smile occurs when the implied volatility (IV) is higher for options that are in-the-money (ITM) and out-of-the-money (OTM) compared to at-the-money (ATM) options.</p>
    <p><strong>Interpretation:</strong> Suggests significant expected volatility with a probability of large price swings in either direction, possibly due to upcoming news, events, or general market uncertainty.</p>
    <p><strong>Derivable Information:</strong> Indicates a leptokurtic distribution of asset returns, suggesting the market is pricing in significant moves away from the current price.</p>
  </div>
</div>
  
<div class="section">
  <div class="section-title">Volatility Smirk/Skew:</div>
  <div class="section-content">
    <p><strong>Shape:</strong> Observed when OTM puts have higher IVs compared to ATM and OTM calls, causing the curve to skew to one side.</p>
    <p><strong>Interpretation:</strong> Indicates anticipation of potential downside, reflecting market fear or bearish sentiment, often seen in equity markets.</p>
    <p><strong>Derivable Information:</strong> Suggests higher probabilities of downward movements, influencing hedging strategies and indicating a demand for downside protection.</p>
  </div>
</div>
  
<div class="section">
  <div class="section-title">Flat Surface:</div>
  <div class="section-content">
    <p><strong>Shape:</strong> Implied volatilities are relatively uniform across strikes and maturities.</p>
    <p><strong>Interpretation:</strong> Suggests a consensus that the underlying asset will experience volatility consistent with historical levels without significant shifts.</p>
    <p><strong>Derivable Information:</strong> Indicates a market in equilibrium or with balanced expectations, where arbitrage opportunities may be limited.</p>
  </div>
</div>
  
<div class="section">
  <div class="section-title">Forward Skew:</div>
  <div class="section-content">
    <p><strong>Shape:</strong> Implied volatility increases for options with higher strike prices.</p>
    <p><strong>Interpretation:</strong> Less common, occurring in markets or assets expected to move upwards, such as commodities facing supply shortages.</p>
    <p><strong>Derivable Information:</strong> Reflects bullish sentiment or market conditions driving expectations of price increases.</p>
  </div>
</div>

<style>
  .section {
    margin-bottom: 20px;
  }

  .section-title {
    color: #333;
    font-size: 20px;
    margin-bottom: 10px;
  }

  .section-content {
    background-color: #f9f9f9;
    border-left: 5px solid #007bff;
    padding: 10px 20px;
    margin: 10px 0;
  }

  .section-content p {
    margin: 10px 0;
  }
  
  .section-content strong {
    color: #000;
  }
</style>