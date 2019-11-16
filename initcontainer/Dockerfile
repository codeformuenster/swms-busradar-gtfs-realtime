FROM alpine

VOLUME /gtfsfeed
ENV GTFSFEED_URL=https://www.stadtwerke-muenster.de/fileadmin/stwms/busverkehr/kundencenter/dokumente/GTFS/stadtwerke_feed_20191028.zip

RUN apk --no-cache add zip wget curl bash dos2unix

COPY fetch-gtfsfeed.sh /usr/local/bin/fetch-gtfsfeed

CMD ["fetch-gtfsfeed"]

