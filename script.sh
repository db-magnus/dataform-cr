#!/bin/sh
cat /secret/.df-credentials.json > .df-credentials.json
dataform run >> /var/log/dataform`date "+%FT%T"`.log
