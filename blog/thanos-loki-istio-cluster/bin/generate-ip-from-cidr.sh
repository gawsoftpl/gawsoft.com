#!/bin/bash
ip=$( echo $1 | cut -d"/" -f1 )
no=$( echo $1 | cut -d"/" -f2 )
num=1
new_ip=$( echo $ip | cut -d"." -f1-3 )
echo $new_ip".$2"