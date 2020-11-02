#!/bin/bash
ips=$(hostname -I)
read -r -a ipArr <<< "$ips"
#for ip in "${ipArr[@]}"
#    do
#        #printf "\n%s assisted-api.local.openshift.io\n" "$ip" >> /etc/hosts
#    done
ips_cs=`echo $ips | xargs | sed -e 's/ /,/g'`
printf "\nSERVICE_IPS=%s" "$ips_cs" >> onprem-environment
