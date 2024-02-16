<script lang="ts">
  import { onMount } from "svelte";
  import Grid from "gridjs-svelte";
  import type { account } from "./type";


  export let accnt: { [ key: string ]: account } | undefined;
  export let accntBal: number;

  let accntCol: { name: string }[] = [];
  let accntDat: string[][] = [];
  let stockCol: { name: string }[] = [];
  let stockDat: string[][] = [];

  const trimZeroRe = /0+$/;

  const accountBalance = (data: account) => {
    return +data.fxbalance;
  }

  const accountTable = (data: account) => {
    const columns = ['Total Balance', 'Pnl(Estimate)']
      .map((str) => { return { name: str } });

    const tabledata = [[ 
      `${data.fxbalance.replace(trimZeroRe, "")} $`, 
      `${data.fxpnl.replace(trimZeroRe, "")} $` 
    ]];

    return { columns, tabledata };
  }

  const stockTable = (data: account) => {
    const columns = ['Name', 'PnL(Estimate)', 'Avg. Prc', 'Qty']
      .map((str) => { return { name: str } });

    let tabledata: string[][] = [];

    for (let stock of data.stocks) {
      tabledata.push([
        stock.code, 
        `${stock.fxpnl.replace(trimZeroRe, "")}$ (${stock.pnlrate.replace(trimZeroRe, "")}%)`,
        `${stock.avgprc.replace(trimZeroRe, "")}$`,
        stock.qty,
      ]);
    }

    return { columns, tabledata };
  }

  $: if (accnt && "payload0" in accnt) {
    let { columns: sc, tabledata: sd } = stockTable(accnt['payload0']);
    stockCol = sc;
    stockDat = sd;

    let { columns: ac, tabledata: ad } = accountTable(accnt['payload0']);
    accntCol = ac;
    accntDat = ad;

    accntBal = accountBalance(accnt['payload0']);
  }

  onMount(async () => {});
</script>

<style global>
  @import "https://cdn.jsdelivr.net/npm/gridjs/dist/theme/mermaid.min.css";
</style>

<div>
  <Grid data={accntDat} columns={accntCol} />
  <br />
  <Grid data={stockDat} columns={stockCol} />    
</div>

