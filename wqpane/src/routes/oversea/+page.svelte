<script lang="ts">
  import LineChart from "$lib/chart/LineChart.svelte";
  import Account from "$lib/table/Account.svelte";
  import { onMount } from "svelte";
  import type { account } from "$lib/table/type";

  const unitedStatesAccountUrl = "http://0.0.0.0:10501/api/v1/kis/account/us";
  const unitedStatesPnl30 = "http://0.0.0.0:10501/api/v1/kis/periodpnl/us";

  let accntBal: number = 0;
  let accnt: { [ key: string ]: account };
  let pp: { [ key: string ]: any };

  onMount(async () => {
    const account = await fetch(unitedStatesAccountUrl);
    if (account.ok) {
      let data = await account.json();
      accnt = { ...data };
    }

    const period = await fetch(unitedStatesPnl30);
    if (period.ok) {
      let data = await period.json();
      pp = { ...data };
    }
  });
</script>

<Account bind:accnt bind:accntBal></Account>
<LineChart bind:pp bind:accntBal></LineChart>