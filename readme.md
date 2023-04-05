# Overview 

The following repository contains a basic example of a module that interacts with Osmosis through IBC.

<br>

### Prerequisites

- Go 1.19+
https://go.dev/doc/install

- Linux users need to make sure they have GCC installed

<br>

### Developer Documentation

- IBC docs
https://ibc.cosmos.network/

- Osmosis docs
https://docs.osmosis.zone/

- ICQ docs
https://github.com/strangelove-ventures/async-icq#readme

<br>

# Demo

### Prerequisites


- Ignite CLI v0.26+
https://docs.ignite.com/welcome/install

<br>

### 1. Start local swapchain network

```bash 
make start-local-swapchain
```

<br>

### 2. Start local osmosis network

```bash 
make start-local-osmosis
```

<br>

### 3. Wait until osmosis node starts producing blocks (~15 secs)

<br>

### 4. Create osmosis liquidity pool

```bash
# create liquidity pool for stake-uosmo with equal amounts
osmosisd tx gamm create-pool --pool-file=contrib/osmosis-pool.json --from validator --home=contrib/osmosis/.osmosis --node=tcp://localhost:26662 --chain-id=osmosis --fees=875stake -y
```

<br>

### 5. Start relayer

```bash
make start-relayer
```

<br>

### 6. Wait until relayer makes connection between the networks (~30 secs)

<br>

### 7. Query osmosis pool spot price through ICQ

```bash
# query spot price for stake
swapchaind tx interchainswap send-query-osmosis-spot-price channel-0 1 stake uosmo --from=validator --home=.swapchain --node=tcp://localhost:26660 -y
```

<br>

### 8. Wait until relayer relays packet (~10sec)

<br>

### 9. Query the result of the ICQ query

```bash
# price should be 1.00 since the pool has equal amounts for both tokens
swapchaind query interchainswap query-state 1 --home=.swapchain --node=tcp://localhost:26660
```

<br>

### 10. Create osmosis account through ICA

```bash
swapchaind tx interchainswap register-interchain-account connection-0 --from=validator --home=.swapchain --node=tcp://localhost:26660 -y
```

<br>

### 11. Wait until relayer relays packet (~10sec)

<br>

### 12. Query the created ICA account address 

```bash
swapchaind q interchain-accounts controller interchain-account cosmos14y0kdvznkssdtal2r60a8us266n0mm97r2xju8 connection-0 --home=.swapchain --node=tcp://localhost:26660
```

<br>

### 13. Fund the ICA address in osmosis
- Replace $INTERCHAIN_ADDRESS with the address queried in step 12

```bash
# fund address with 10000uosmo
osmosisd tx bank send faucet $INTERCHAIN_ADDRESS 10000uosmo --home=contrib/osmosis/.osmosis --node=tcp://localhost:26662 --chain-id=osmosis --fees=875stake -y
```

<br>

### 14. Swap tokens through ICA
- Replace $INTERCHAIN_ADDRESS with the address queried in step 12

```bash
# swap 10uosmo for stake
swapchaind tx interchainswap send-msg-osmosis-swap '{"@type":"/osmosis.gamm.v1beta1.MsgSwapExactAmountIn","sender":"$INTERCHAIN_ADDRESS","routes":[{"poolId":"1","tokenOutDenom":"stake"}],"tokenIn":{"denom":"uosmo","amount":"10"},"tokenOutMinAmount":"1"}' connection-0 --from=validator --home=.swapchain --node=tcp://localhost:26660 -y
```

<br>

### 15. Wait until relayer relays packet (~10sec)

<br>

### 16. Verify swap was successful
- Replace $INTERCHAIN_ADDRESS with the address queried in step 12

```bash
# interchain account should have 9990uosmo and some stake in their balance
osmosisd q bank balances $INTERCHAIN_ADDRESS --home=contrib/osmosis/.osmosis --node=tcp://localhost:26662 --chain-id=osmosis

# pool should have more uosmo than stake, therefore the spot price for uosmo should be less than 1.00 stake
osmosisd q gamm spot-price 1 stake uosmo --home=contrib/osmosis/.osmosis --node=tcp://localhost:26662 --chain-id=osmosis
```

<br>

# Tests

[./interchaintest](./interchaintest/) contains example end-to-end IBC test using [`interchaintest`](https://github.com/strangelove-ventures/interchaintest) - a package that provides test suite for complete interchain simulation

<br>

### Prerequisites

- Docker
https://docs.docker.com/get-docker/

<br>

### 1. Build binary image

```bash
make build-docker
```

<br>

### 2. Run tests

```bash
make interchaintest
```

<br>

# Working with Docker

### 1. Build binary image

```bash
# create image named 'swapchaind'
make build-docker
```

<br>

### 2. Example interaction

```bash
# init config and persist config folder (/swapchain/.swapchain) to volume 'my-vol'
docker run --rm -it -v=my-vol:/swapchain/.swapchain swapchaind init testmoniker --chain-id=testchainid

# add account key named 'test'
docker run --rm -it -v=my-vol:/swapchain/.swapchain swapchaind keys add test --keyring-backend=test

# fund account and add to genesis
docker run --rm -it -v=my-vol:/swapchain/.swapchain swapchaind add-genesis-account test 999999999999999stake --keyring-backend=test

# generate genesis tx that sets 'test' as validator
docker run --rm -it -v=my-vol:/swapchain/.swapchain swapchaind gentx test 2222222222stake --keyring-backend=test --chain-id=testchainid

# collect the genesis tx and add to genesis
docker run --rm -it -v=my-vol:/swapchain/.swapchain swapchaind collect-gentxs

# start node
docker run --rm -it -v=my-vol:/swapchain/.swapchain swapchaind start
```
