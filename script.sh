#!/bin/bash
cat /secret/.df-credentials.json > .df-credentials.json
ls -la .df-credential*
echo "starting dataform compile\n"
dataform compile 
echo "dry run\n"
dataform run --dry-run
echo "run"
dataform run 
