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
unzip -o /gtfsfeed/feed.zip -d /gtfsfeed

# convert files to unix format
dos2unix /gtfsfeed/*.txt

# for usage in opentripplanner, create a non standard feed-info.txt
# See https://github.com/HSLdevcom/OpenTripPlanner-data-container/blob/b14f9b96d968940af7c9d52e57a838bace2ca710/task/setFeedId.js
cat <<- EOF > /gtfsfeed/feed-info.txt
feed_publisher_name,feed_publisher_url,feed_lang,feed_id
Code for Münster,https://swms-busradar-gtfs-realtime.codeformuenster.org/stadtwerke_feed_20191028-with-feed_id.zip,de,STWMS
EOF

# create zip file
zip -9 -j /gtfsfeed/stadtwerke_feed_20191028-with-feed_id.zip /gtfsfeed/*.txt
