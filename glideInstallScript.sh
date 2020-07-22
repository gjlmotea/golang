#!/bin/bash
bold=$(tput bold)
norm=$(tput sgr0)
for d in ./*/
do
    file="$d"glide.yaml
    if [ -f "$file" ]; then
        echo "======== ${bold} $d ${norm} glide install... ========"
        (cd "$d" && glide install);
    fi
done
