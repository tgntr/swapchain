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

### Start swapchain network

```bash 
ignite chain serve --reset-once
```

### Start local osmosis network

```bash 
make start-local-osmosis
```

### Wait until osmosis node starts producing blocks (~10 secs) then create a pool

```bash
osmosisd tx gamm create-pool --pool-file=contrib/osmosis-pool.json --from validator --home=contrib/osmosis/.osmosis --node=tcp://localhost:26662 --chain-id=osmosis --fees=875stake -y
```

### Start the relayer

```bash
make start-relayer
```

### Wait until relayer makes connection between the networks (~30 secs) then query osmosis pool spot price from swapchain

```bash
swapchaind tx interchainswap send-query-osmosis-spot-price channel-0 1 stake uosmo --from=validator --home=.swapchain --node=tcp://localhost:26660 -y
```

### See the result of the query

```bash
swapchaind query interchainswap query-state 1 --home=.swapchain --node=tcp://localhost:26660
```                                         
