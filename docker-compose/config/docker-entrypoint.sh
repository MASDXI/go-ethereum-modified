#!/bin/sh

if [ ! -d /geth-node/devnet/data/geth ]; then
    echo "Initial genesis.json!"
    geth --datadir /geth-node/devnet/data init /geth-node/devnet/config/genesis.json
fi

geth $@