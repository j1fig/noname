#!/bin/bash
# script to update the open source data on public transports.
# this blows away all the previously existing data.

rm -rf data/
mkdir -p data/tmp

wget -O data/tmp/bus.zip http://www.transporlis.pt/Portals/0/OpenData/gtfs/zip/1/gtfs_1.zip
unzip data/tmp/bus.zip -d data/

# everuthing is actually a CSV so we might as well be honest.
for file in data/*.txt
do
  mv "$file" "${file/.txt/.csv}"
done
