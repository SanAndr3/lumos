#!/bin/bash

KEY="dev0"
CHAINID="lumos_7798-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t lumos-datadir.XXXXX)

echo "create and add new keys"
./lumosd keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Lumos with moniker=$MONIKER and chain-id=$CHAINID"
./lumosd init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./lumosd add-genesis-account \
"$(./lumosd keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000alumos,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./lumosd gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./lumosd collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./lumosd validate-genesis --home $DATA_DIR

echo "starting lumos node $i in background ..."
./lumosd start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started lumos node"
tail -f /dev/null