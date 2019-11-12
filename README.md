# swms-busradar-gtfs-realtime

Stadtwerke Münster Busradar to GTFS-realtime. This software grabs messages from the [Stadtwerke Münster Busradar] [`/fahrzeuge`] endpoint and converts them to a GTFS realtime feed.

A hosted version of this realtime feed is hosted at [https://swms-busradar-gtfs-realtime.codeformuenster.org/feed](https://swms-busradar-gtfs-realtime.codeformuenster.org/feed).

## Running

The program will periodically fetch the [`/fahrzeuge`] endpoint of the [Stadtwerke Münster Busradar], translate the values into a GTFS realtime feed and store everything into a file called `feed` next to the executable binary. Set `GTFS_REALTIME_FEED_PATH` in your environment to override the path where the feed will be stored.

Running requires a [Stadtwerke Münster GTFS feed] from the [Stadtwerke Münster GTFS feed download page]. **Attention**: The files from file `stadtwerke_feed_20191028.zip` need to be converted to unix format using `dos2unix`. Extract the feed zip next to the executable binary or specify the path to the static feed (zip or folder) by setting `GTFS_FEED_PATH` in your environment.

### Container images

Container image [`quay.io/codeformuenster/swms-busradar-gtfs-realtime`] contains the main program of this project and is built from the root of this repository.

Container image [`quay.io/codeformuenster/swms-busradar-gtfs-realtime-init`] can be used to download, extract and convert the [Stadtwerke Münster GTFS feed] for usage in the main container.

### Kubernetes deployment

You'll find kubernetes manifests for running this project in your cluster in the [Code for Münster kubernetes-deployments repository].

## Development requirements

- Go >= 1.13

[Stadtwerke Münster GTFS feed]: https://www.stadtwerke-muenster.de/fileadmin/stwms/busverkehr/kundencenter/dokumente/GTFS/stadtwerke_feed_20191028.zip
[Stadtwerke Münster GTFS feed download page]: https://www.stadtwerke-muenster.de/privatkunden/mobilitaet/fahrplaninfos/fahr-netzplaene-downloads/open-data-gtfs/gtfs-download.html
[Stadtwerke Münster Busradar]: http://api.busradar.conterra.de/
[`/fahrzeuge`]: https://rest.busradar.conterra.de/prod/fahrzeuge
[Code for Münster kubernetes-deployments repository]: https://github.com/codeformuenster/kubernetes-deployments/
[`quay.io/codeformuenster/swms-busradar-gtfs-realtime`]: https://quay.io/repository/codeformuenster/swms-busradar-gtfs-realtime
[`quay.io/codeformuenster/swms-busradar-gtfs-realtime-init`]: https://quay.io/repository/codeformuenster/swms-busradar-gtfs-realtime-init
