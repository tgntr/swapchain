#!/bin/sh

# set current working directory to file directory instead of caller directory
cd $(dirname $0)

# cleanup
rm -rf relayer

# install binary
git clone --branch=v2.2.0 https://github.com/cosmos/relayer.git
cd relayer && make build

# setup config
./build/rly config init --home=.relayer
./build/rly chains add-dir ../relayer-config/chains --home=.relayer
./build/rly paths add-dir ../relayer-config/paths --home=.relayer

# mnemonic for relayer account which is created and funded during local network setup.
mnemonic="stamp later develop betray boss ranch abstract puzzle calm right bounce march orchard edge correct canal fault miracle void dutch lottery lucky observe armed"

# add account to config
./build/rly keys restore swapchain acc "$mnemonic" --home=.relayer
./build/rly keys restore osmosis acc "$mnemonic" --home=.relayer

./build/rly tx link-then-start swapchain-osmosis --src-port interchainswap --dst-port icqhost --version=icq-1 --home=.relayer

