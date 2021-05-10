#!/bin/bash
api=https://gateway.elrond.com
output=/tmp/$RANDOM.txt
>$output
function base64_to_hexa {
        echo $1 | base64 -d -i | hexdump -v -e '/1 "%02x"'
}

function hexa_to_decimal {
        x=$(echo ${$1^^})
        echo "obase=10; ibase=16; $x" | bc
}

function hexa_to_string {
        echo $1 | base64 -di
}

function getQueueIndex {
        index=$(curl -s -d '{
                "scAddress": "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqllls0lczs7",
                "funcName":  "getQueueIndex",
                "args":      ["'"$1"'"],
                "caller":    "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqplllst77y4l"
        }' -X POST $api/vm-values/query | jq .data.data.returnData[] 2>/dev/null)
        if [ $(echo $index | wc -c) -eq 1 ]
        then
                index=0
        else
                index=$index
        fi
        echo $index
}

keys=$(curl -s 'https://api.elrond.com/nodes?status=queued&size=3200' | jq .[].bls)
data=$(curl -s $api/node/heartbeatstatus)
echo -n Processing .
for i in $(echo $keys)
do
        key=$(echo $i | tr -d '"')
        index=$(getQueueIndex $key)
        if [ $index != "0" ]
        then
                pos=$(hexa_to_string $(getQueueIndex $key))
                identity=$(echo $data | jq -r --arg KEY "$key" '.data.heartbeats[] | select(.publicKey | contains ($KEY))' | jq '.identity' | tr -d '"')
                displayName=$(echo $data | jq -r --arg KEY "$key" '.data.heartbeats[] | select(.publicKey | contains ($KEY))' | jq '.nodeDisplayName' | tr -d '"')
                if [ ! $identity ]
                then
                        identity=N/A
                fi
                echo -n "."
                echo "$pos $identity $displayName" >> $output
        fi
        echo -n "."
done
echo ""
sort -n $output
rm -rf $output