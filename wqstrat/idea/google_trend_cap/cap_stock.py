import pandas as pd
from pytrends.request import TrendReq

import requests
import copy

# Yahoo limit
MAX = 2000

df = pd.read_csv("small_cap_proxy.csv")

health = df.loc[df['Sector Classification'] == 'Health Care']
industry = df.loc[df['Sector Classification'] == 'Industrials']
finance = df.loc[df['Sector Classification'] == 'Financials']
itech = df.loc[df['Sector Classification'] == 'Information Technology']

# Initialize a pytrends request object
# pytrend = TrendReq(
#     hl='en-US',
#     proxies=['https://34.203.233.13:80',],
#     retries=2,
#     backoff_factor=0.1,
#     requests_args={'verify': False}
# )
#
# keyword = 'Tenet Healthcare Corporation'
#
# pytrend.build_payload(kw_list=[keyword], timeframe='all')
#
# hist = pytrend.interest_over_time()




# tickers = []
# for stock in data[:10]:
#     isin_code = stock['ISIN']
#     original_row = copy.deepcopy(stock)
#
#     if len(isin_code) != 12:
#         tickers.append(None)
#
#     yahoo_isin_to_ticker = f"https://query1.finance.yahoo.com/v1/finance/search?q={isin_code}&lang=en-US&region=US&quotesCount=6&newsCount=2&listsCount=2&enableFuzzyQuery=false&quotesQueryId=tss_match_phrase_query&multiQuoteQueryId=multi_quote_single_token_query&newsQueryId=news_cie_vespa&enableCb=true&enableNavLinks=true&enableEnhancedTrivialQuery=true&enableResearchReports=false&enableCulturalAssets=true&enableLogoUrl=true&researchReportsCount=2"
#     header = {"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"}
#     resp = requests.get(yahoo_isin_to_ticker, headers=header)
#
#     if resp.status_code == 200:
#         body = resp.json()
#         if len(body['quotes'][0]) == 0:
#             original_row['Ticker'] = None
#             continue
#         original_row['Ticker'] = body['quotes'][0]['symbol']
#     else:
#         original_row['Ticker'] = None
#
#     tickers.append(original_row)
#
# df = pd.DataFrame().from_dict(tickers)
# df.to_csv("small_cap_proxy_with_ticks.csv")
