#!/bin/bash

HOSTNUM=100000
MAXTRY=10000

i=0
while [ $i -lt $HOSTNUM ]; do 
    let i++

    # find unique random private IP address
    dup=1
    try=$MAXTRY
    while [[ $dup -ne 0 && $try -gt 0 ]]; do 
        let try--
        hidx1=$(($RANDOM*255/32767))
        hidx2=$(($RANDOM*255/32767))
        hidx3=$(($RANDOM*255/32767))

        ip="10.$hidx1.$hidx2.$hidx3"
        
        if ! egrep -q "^$ip," hosts.csv; then
            dup=0
        fi
    done

    rtt=$(($RANDOM*120/32767))

    echo "$ip,$rtt" >> hosts.csv   
done
