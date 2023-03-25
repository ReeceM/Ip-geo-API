#! /bin/bash
rm -rf *.csv
curl -L "https://cdn.jsdelivr.net/npm/@ip-location-db/geo-whois-asn-country/geo-whois-asn-country-ipv4-num.csv" -o "ipv4_db.csv"
curl -L "https://cdn.jsdelivr.net/npm/@ip-location-db/geo-asn-country/geo-asn-country-ipv6-num.csv" -o "ipv6_db.csv"
