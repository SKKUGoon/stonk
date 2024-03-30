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

<p>
    1. Volatility Smile:
    Shape: A volatility smile occurs when the implied volatility (IV) is higher for options that are in-the-money (ITM) and out-of-the-money (OTM) compared to at-the-money (ATM) options. Graphically, this forms a U-shape or smile.
    Interpretation: A volatility smile suggests that the market expects the underlying asset to experience significant volatility, with a non-zero probability of large price swings in either direction. This can be due to upcoming news, events, or general market uncertainty.
    Derivable Information: Indicates a leptokurtic distribution of asset returns, suggesting that the market is pricing in the possibility of significant moves away from the current price.
</p>
<p>
    2. Volatility Smirk/Skew:
    Shape: A smirk or skew is observed when OTM puts (lower strike prices) have higher IVs compared to ATM and OTM calls (higher strike prices), making the curve skew to one side.
    Interpretation: This indicates that the market is anticipating a potential downside more than an upside, often seen in equity markets due to the asymmetric risk of a market crash or correction.
    Derivable Information: Reflects market fear or bearish sentiment, pricing in higher probabilities of downward movements. This can influence hedging strategies, suggesting a demand for downside protection.
</p>
<p>
    3. Flat Surface:
    Shape: Implied volatilities are relatively uniform across strikes and maturities.
    Interpretation: Suggests a market consensus that the underlying asset will experience volatility consistent with historical levels, without expecting significant shifts.
    Derivable Information: Indicates a market in equilibrium or with balanced expectations. Arbitrage opportunities may be limited.
</p>
<p>
    4. Forward Skew:
    Shape: The implied volatility increases for options with higher strike prices.
    Interpretation: This is less common but can occur in markets or for assets where there's a greater expectation of an upward move (e.g., commodities experiencing supply shortages).
    Derivable Information: Reflects bullish sentiment or specific market conditions driving expectations of price increases.
</p>