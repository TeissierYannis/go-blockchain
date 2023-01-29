#!/bin/bash

blockchain createwallet > ./addresses/node_address_$NODE_ID
# Remove ew address is: from the output
sed -i 's/New address is: //g' ./addresses/node_address_$NODE_ID

sleep 30

# Start the node
blockchain startnode -miner $(cat ./addresses/node_address_$NODE_ID)