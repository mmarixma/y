#!/bin/bash

curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | \
jq -r --arg id "$HERO_ID" '.[] | select((.id | tostring) == $id) | .connections.relatives | gsub("\n"; "\\n")'