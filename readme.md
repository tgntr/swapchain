## Overview 

The following repository contains a basic example of a module that interacts with Osmosis through IBC.

### Developer Documentation

IBC docs
https://ibc.cosmos.network/

Osmosis docs
https://docs.osmosis.zone/

Interchain Queries docs
https://github.com/strangelove-ventures/async-icq/blob/main/README.md

## Demo

### Prerequisites

Go 1.19+
https://go.dev/doc/install

Ignite CLI v0.26+
https://docs.ignite.com/welcome/install

### 1. Start swapchain network

```bash 
ignite chain serve --reset-once
```

### 2. Start local osmosis network

```bash 
make start-local-osmosis
```

### 3. Wait until osmosis node starts producing blocks (~15 secs)

### 4. Create osmosis liquidity pool

```bash
# create liquidity pool for stake-uosmo with equal amounts
osmosisd tx gamm create-pool --pool-file=contrib/osmosis-pool.json --from validator --home=contrib/osmosis/.osmosis --node=tcp://localhost:26662 --chain-id=osmosis --fees=875stake -y
```

### 5. Start the relayer

```bash
make start-relayer
```

### 6. Wait until relayer makes connection between the networks (~30 secs)

### 7. Query osmosis pool spot price through ICQ

```bash
# query spot price for stake
swapchaind tx interchainswap send-query-osmosis-spot-price channel-0 1 stake uosmo --from=validator --home=.swapchain --node=tcp://localhost:26660 -y
```

### 8. Wait until relayer relays packet (~10sec)

### 9. Query the result of the ICQ query

```bash
# price should be 1.00 since the pool has equal amounts for both tokens
swapchaind query interchainswap query-state 1 --home=.swapchain --node=tcp://localhost:26660
```

### 10. Create osmosis account through ICA

```bash
swapchaind tx interchainswap register-interchain-account connection-0 --from=validator --home=.swapchain --node=tcp://localhost:26660 -y
```

### 11. Wait until relayer relays packet (~10sec)

### 12. Query the created ICA account address 

```bash
swapchaind q interchain-accounts controller interchain-account cosmos14y0kdvznkssdtal2r60a8us266n0mm97r2xju8 connection-0 --home=.swapchain --node=tcp://localhost:26660
```

### 13. Fund the ICA address in osmosis
- Replace INTERCHAIN_ADDRESS with the address queried in step 12

```bash
# fund address with 10000uosmo
osmosisd tx bank send faucet INTERCHAIN_ADDRESS 10000uosmo --home=contrib/osmosis/.osmosis --node=tcp://localhost:26662 --chain-id=osmosis --fees=875stake -y
```

### 14. Swap tokens through ICA
- Replace INTERCHAIN_ADDRESS with the address queried in step 12

```bash
# swap 10uosmo for stake
swapchaind tx interchainswap send-msg-osmosis-swap '{"@type":"/osmosis.gamm.v1beta1.MsgSwapExactAmountIn","sender":"INTERCHAIN_ADDRESS","routes":[{"poolId":"1","tokenOutDenom":"stake"}],"tokenIn":{"denom":"uosmo","amount":"10"},"tokenOutMinAmount":"1"}' connection-0 --from=validator --home=.swapchain --node=tcp://localhost:26660 -y
```

### 15. Wait until relayer relays packet (~10sec)

### 16. Verify swap was successful
- Replace INTERCHAIN_ADDRESS with the address queried in step 12

```bash
# interchain account should have 9990uosmo and some stake in their balance
osmosisd q bank balances INTERCHAIN_ADDRESS --home=contrib/osmosis/.osmosis --node=tcp://localhost:26662 --chain-id=osmosis

# pool should have more stake than uosmo, therefore the spot price for stake should be less than 1.00
osmosisd q gamm spot-price 1 stake uosmo --home=contrib/osmosis/.osmosis --node=tcp://localhost:26662 --chain-id=osmosis
```