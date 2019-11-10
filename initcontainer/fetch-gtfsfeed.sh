#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

feed_url=${GTFSFEED_URL}
if [[ -z "${feed_url}" ]]; then
    echo "Missing or empty GTFSFEED_URL environment variable"
    exit 1
fi

# download the feed file
curl -o /gtfsfeed/feed.zip "${feed_url}"

# extract the files
unzip /gtfsfeed/feed.zip -d /gtfs

# convert files to unix format
dos2unix /gtfsfeed/*.txt
