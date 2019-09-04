#!/bin/bash

MILESTONE="$1"

LABELS=("experimental" "breaking change" "release:add" "release:change" "release:fix" "release:remove")
CATEGORIES=("Experimental" "Breaking Changes" "Added" "Changed" "Fixed" "Removed")

PR=$(hub pr list -s merged -f "%Mt %L|- ([%i](%U)) %t.%n" --sort-ascending | grep $MILESTONE)

printf "## [$MILESTONE](https://github.com/mesg-foundation/engine/releases/tag/$MILESTONE)\n\n"

for i in "${!LABELS[@]}"; do 
  LABEL=${LABELS[$i]}
  CATEGORY=${CATEGORIES[$i]}
  printf "#### $CATEGORY\n\n"
  echo "$PR" | grep "$LABEL" | cut -d'|' -f2
  printf "\n"
done
