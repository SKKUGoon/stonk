<script lang="ts">
  import { onMount } from "svelte";
  import Grid from "gridjs-svelte";


  type account = {
    fxbalance: string;
    fxpnl: string;
    stocks: accountStocks[];
  };

  type accountStocks = {
    code: string; // AAPL
    name: string; // 애플

    /* Below fields are actually numbers, but for precision turned into string */ 

    fxpnl:   string; // 1.74 USD
    pnlrate: string; // percentage
    avgprc:  string; 
    qty:     string;
    notsold: string;
  }

  const unitedStatesAccountUrl = "http://0.0.0.0:10501/api/v1/account/us";
  let respData: { [ key: string ]: account };

  let accntCol: { name: string }[] = [];
  let accntDat: string[][] = [];
  let stockCol: { name: string }[] = [];
  let stockDat: string[][] = [];

  const accountTable = (data: account) => {
    const columns = ['Total Balance', 'Pnl(Estimate)']
      .map((str) => { return { name: str } });

    const tabledata = [[ `${data.fxbalance}$`, `${data.fxpnl}$` ]];

    return { columns, tabledata };
  }

  const stockTable = (data: account) => {
    const columns = ['Name', 'PnL(Estimate)', 'Avg. Prc', 'Qty']
      .map((str) => { return { name: str } });

    let tabledata: string[][] = [];

    for (let stock of data.stocks) {
      tabledata.push([
        stock.code, 
        `${stock.fxpnl}$(${stock.pnlrate}%)`,
        `${stock.avgprc}$`,
        stock.qty,
      ]);
    }

    return { columns, tabledata };
  }

  onMount(async () => {
    const response = await fetch(unitedStatesAccountUrl);
    if (response.ok) {
      respData = await response.json();
      
      let { columns: sc, tabledata: sd } = stockTable(respData['payload0']);
      stockCol = sc;
      stockDat = sd;

      let { columns: ac, tabledata: ad } = accountTable(respData['payload0']);
      accntCol = ac;
      accntDat = ad;
    } else {
      console.error("failed to fetch data", response.statusText);
    }
  });
</script>

<style global>
    @import "https://cdn.jsdelivr.net/npm/gridjs/dist/theme/mermaid.min.css";
  </style>

<div>
  <Grid data={accntDat} columns={accntCol} />
  <br />
  <Grid data={stockDat} columns={stockCol} />    
</div>

