# ATG

ATG is a Go algorithm for automated trading.

More information about the algorithm and how it performs will be shared here: https;//jgislas.com/mscs.

The first version follows a 20 Day Moving Average (20DMA) strategy for Ethereum (ETH) on Bitfinex.

When the price of ETH closes above the 20DMA, the algorithm buys ETH. When the price of ETH closes below the 20DMA (i.e. a drop in the market is likely to occur) the algorithm sells ETH and buys USD. 

Everyday at 03:00 UTC the algorithm runs and the rebalance takes place automatically.

Current version: 1.3.

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
10. Update holdings
11. Reorganize code in CMD
12. Monitor via cronjob in an hourly basis if an order has been fulfilled
13. Create blank spreadsheet for the project
14. Reorganize spreadsheet to reflect current value of holdings

### V 1.3 - 10/07/2020
16. Test automation of MA algorithm
15. Create a new Btfinex account for the project
17. Send notification when the Rebalance has taken place (SMS, Email, etc)

### Other features
18. Build a simple website to show how the strategy performs against Buy and Hold.
19. Instead of rebalancing to USD, use CUSD (the Compound's version of USD) so that an interest can be accrued on the amount of USD.
20. Instead of rebalancing to ETH, find a way to accrue some interest in Compound or AAVE.
21. Testing the algorithm on other altcoins.
22. Generalize the algorithm so that it can be used for other periods of the Moving Average (i.e. 1 Week Moving Average for long term strategies in Bitcoin)
23. Use a Database instead of relying on Google Spreadsheet to store the data.
