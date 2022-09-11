#!/bin/bash

out=./screenshots
total=1000
num=100
for ((skip = 0; i <= $total; i += $num))
do
    twig gather -o $out -n $num -s $skip

    # TODO: some action here to use the screenshots

    echo "cleaning screenshots"
    rm -rf $out
done
