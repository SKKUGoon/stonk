export type account = {
    fxbalance: string;
    fxpnl: string;
    stocks: accountStocks[];
  };
  
  export type accountStocks = {
    code: string; // AAPL
    name: string; // 애플
  
    /* Below fields are actually numbers, but for precision turned into string */ 
  
    fxpnl:   string; // 1.74 USD
    pnlrate: string; // percentage
    avgprc:  string; 
    qty:     string;
    notsold: string;
  }