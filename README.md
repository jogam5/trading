# ATG

ATG is a Go algorithm for automated trading.

The first version follows a 20 Day Moving Average (20DMA) strategy for Ethereum (ETH) on Bitfinex.

When the price of ETH closes above the 20DMA, the algorithm buys ETH. When the price of ETH is below the 20DMA (i.e. a drop in the market is likely to occur) the algorithm sells ETH and buys USD. 

Everyday at 11:00 UTC the rebalance takes place automatically.

Current version: 1.2.

## Roadmap 

### V 1.0 - 09/21/2020
1. Choose where to fetch the prices (Source)
2. Make connection between Source and Google Spredsheets
3. Automate the store of data each hour (hourly price from Source)
4. Enter formulas for computing the Moving Average 20 day

### V 1.1 - 09/25/2020
5. Read actionable data from Spreadsheet
6. Connect the exchange with the API permissions to create orders
7. Set up BUY/SELL via API
8. Improve README

### V 1.2 - 09/29/2020
9. Write orderID in the Spreadsheet and monitors its status on Bitfinex
10. Send notification when the Rebalance has taken place (SMS, Email, etc)

### V1.x
11. Build a simple website to show how the strategy performs against Buy and Hold.
12. Instead rebalancing to USD, use CUSD (the Compound's version of USD) so that an interest can be accrued on the amount of USD. 
13. Instead of rebalancing to ETH, find a way to accrue some interest in Compound or AAVE.
14. Testing the algorithm on other altcoins.
15. Generalize the algorithm so that it can be used for other periods of the Moving Average (i.e. 1 Week Moving Average for long term strategies in Bitcoin)
16. Use a Database instead of relying on Google Spreadsheet to store the data.
