# interchainquerydemo
**ICQ Demo** 

## Disclaimer

The following repository and [`x/interchainswap`](./x/interchainswap/) module serves as an example and is used to exercise the functionality of Interchain Accounts end-to-end for development purposes only.
This module **SHOULD NOT** be used in production systems

## Overview 

The following repository contains a basic example of an Interchain Queries module and serves as a developer guide for teams that wish to use interchain queries functionality.

### Developer Documentation

Interchain Queries developer docs can be found here
https://github.com/strangelove-ventures/async-icq/blob/main/README.md

Medium blog - 
https://link.medium.com/a70uOK1cFwb

## Demo

### Start the two instances of demo chain with following commands

```bash 
ignite chain serve -c sender.yml --reset-once
```

```bash 
ignite chain serve -c receiver.yml --reset-once
```

### Configure and start the relayer

```bash
rm -rf ~/.ignite/relayer
```


```bash
ignite relayer configure -a \
--source-rpc "http://localhost:26659" \
--source-faucet "http://localhost:4500" \
--source-port "interchainswap" \
--source-gasprice "0.0stake" \
--source-gaslimit 5000000 \
--source-prefix "cosmos" \
--source-version "icq-1" \
--target-rpc "http://localhost:26559" \
--target-faucet "http://localhost:4501" \
--target-port "icqhost" \
--target-gasprice "0.0stake" \
--target-gaslimit 300000 \
--target-prefix "cosmos"  \
--target-version "icq-1"
```

```bash
ignite relayer connect
```

### Send the query to "receiver" chain

```bash
swapchaind tx interchainswap send-query-all-balances channel-0 cosmos1ez43ye5qn3q2zwh8uvswppvducwnkq6w6mthgl --chain-id=sender --node=tcp://localhost:26659 --home ~/.sender --from alice
```

### See the result of packet 1

```bash
swapchaind query interchainswap query-state 1 --chain-id=sender --node=tcp://localhost:26659
```                                         

Output:

```
    request:
    '@type': /cosmos.bank.v1beta1.QueryAllBalancesRequest
    address: cosmos1ez43ye5qn3q2zwh8uvswppvducwnkq6w6mthgl
    pagination:
        count_total: false
        key: null
        limit: "100"
        offset: "0"
        reverse: false
    response:
    '@type': /cosmos.bank.v1beta1.QueryAllBalancesResponse
    balances:
    - amount: "8000000"
        denom: atom
    - amount: "800000000"
        denom: stake
    pagination:
        next_key: null
        total: "0"
```

Relayer Output:

```
Relay 1 acks from receiver => sender
```
