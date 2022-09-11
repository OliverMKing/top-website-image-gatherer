#!/bin/bash

out=./screenshots # directory to put screenshots in
total=1000 # total number of top website screenshots to gather 
num=100 # batch size
for ((skip = 0; i <= $total; i += $num))
do
    twig gather -o $out -n $num -s $skip

    # TODO: some action here to use the screenshots

    echo "cleaning screenshots"
    rm -rf $out
done
