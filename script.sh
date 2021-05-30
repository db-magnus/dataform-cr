#!/bin/sh
#wait and retry logic
for i in $(seq 1 5); do cp /secret/.df-credentials.json . && s=0 && break || s=$? && sleep 5; done; (exit $s)
dataform run
