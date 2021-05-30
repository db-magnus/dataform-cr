#!/bin/sh
cat /secret/.df-credentials.json > .df-credentials.json
dataform run 
