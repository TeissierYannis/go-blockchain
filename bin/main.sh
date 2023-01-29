#!/bin/bash

blockchain createwallet > ./addresses/node_address_$NODE_ID
# Remove ew address is: from the output
sed -i 's/New address is: //g' ./addresses/node_address_$NODE_ID

blockchain createblockchain -address $(cat ./addresses/node_address_$NODE_ID)

sleep 5

# From env var NODE_LIST, copy folder blocks_$FIRST_NODE to the current node (split by comma)
IFS=',' read -ra NODE_LIST <<< "$NODE_LIST"
for i in "${NODE_LIST[@]}"; do
  cp -r ./tmp/blocks_$NODE_ID ./tmp/blocks_$i
done

# Count number of files in the directory addresses (while there is less than 3 addresses, wait 5 seconds and try again)
while [ $(ls ./addresses | wc -l) -lt 3 ]; do
    sleep 5
done

# Get the addresses of the other nodes and send them coins
for node_id in $(ls ./addresses | grep -v node_address_$NODE_ID); do
    blockchain send -from $(cat ./addresses/node_address_$NODE_ID) -to $(cat ./addresses/$node_id) -amount 10 -mine
done



sleep 10

# Start the node
blockchain startnode -miner $(cat ./addresses/node_address_$NODE_ID)