#!/bin/bash
bold=$(tput bold)
norm=$(tput sgr0)
for folder in ./*/
do
    file="$folder"glide.yaml
    if [ -f "$file" ]; then
        echo "======== ${bold} $d ${norm} glide install... ========"
        (cd "$folder" && glide install);
    fi
done
