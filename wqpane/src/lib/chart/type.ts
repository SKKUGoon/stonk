export type eachTrade = {
  avg_sll_unpr:       string;
  exrt:               string;
  frcr_pchs_amt1:     string;
  frcr_sll_amt_smtl1: string;
  frst_bltn_exrt:     string;
  ovrs_excg_cd:       string;
  ovrs_item_name:     string;
  ovrs_pdno:          string;
  ovrs_rlzt_pfls_amt: string;
  pchs_avg_pric:      string;
  pftrt:              string;
  slcl_qty:           string;
  stck_sll_tlex:      string;
  trad_day:           string;
}

export type periodProfit = {
  Output1: eachTrade[];
  
  msg1:   string;
  msg_cd: string;
  rt_cd:  string;
}