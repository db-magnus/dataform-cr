# Dataform example for Cloud Run

```June 2021: This is a work-in-progress.```

An example pipeline for running Dataform pipelines in Cloud Run.

Inspired by Mete Atamel and his example of [dbt in Cloud Run](https://github.com/meteatamel/cloudrun-tutorial/blob/master/docs/scheduled-dbt-service-bigquery.md)

Code is stored in Cloud Source repositories, on push to git it will build a new version of the service with Cloud Build.

Credentials are stored in Secret Manager.

## How it works
There is a webserver made in nodejs that listens for a request on 8080. Request is authenticated by cloud run, so no extra authentication is done in nodejs.

The webserver will then spawn a bash-script that runs dataform.


## Installation

Add service account and policy binding for dataform:
```
gcloud iam service-accounts create dataform-cr --display-name "Dataform BigQuery SA"

gcloud projects add-iam-policy-binding $PROJECT_NAME --member=serviceAccount:dataform-cr@$(gcloud config get-value project).iam.gserviceaccount.com   --role=roles/bigquery.admin
```
Create artifact repository to store docker images
```
gcloud artifacts repositories create dataform-cr-docker --repository-format=docker --location=europe-west1 --description="Repository for Cloud Run Dataform images"
```


Download credential from the service account, use dataform to create the credential file:
```
dataform init-creds bigquery

[1] US (default)
[2] EU
[3] other

Enter the location of your datasets [1, 2, 3]: 2

[1] ADC (default)
[2] JSON Key

Do you wish to use Application Default Credentials or JSON Key [1/2]: 2
Please follow the instructions at https://docs.dataform.co/dataform-cli#create-a-credentials-file/
to create and download a private key from the Google Cloud Console in JSON format.
(You can delete this file after credential initialization is complete.)

Enter the path to your Google Cloud private key file:
> /path/to/5a9138bc207f.json            

Running connection test...

Warehouse test query completed successfully.

Credentials file successfully written:
 ~/src/dataform-testing/.df-credentials.json
To change connection settings, edit this file directly.
```
Upload this file to Secret Manager, grant permission for cloud run to use it.

In Cloud Source, make a trigger to run cloud build on push to git.


## Usage

Call Cloud Run with
```
$ curl -H "Authorization: Bearer $(gcloud auth print-identity-token)" https://url-to-cloud-run

```

## Current Error
Dataform will compile and dry-run, but with run it will do an exit 0 with no error message....

### Error log
<details>
  <summary>Click to expand</summary>

  ```
2021-06-03T21:56:53.651464Zstdout: -rw-rw-r-- 1 root root 2474 Jun 3 21:56 .df-credentials.json
2021-06-03T21:56:53.651828Zstdout: starting dataform compile\n
2021-06-03T21:56:54.932605Zstdout: Compiling...
2021-06-03T21:56:55.274948Zstdout: [32mCompiled 0 action(s).[0m
2021-06-03T21:56:55.305195Zstdout: dry run\n
2021-06-03T21:56:56.585766Zstdout: Compiling...
2021-06-03T21:56:56.938360Zstdout: [32mCompiled successfully.
2021-06-03T21:56:56.938377Z[0m
2021-06-03T21:56:56.943341Zstdout: Dry run (--dry-run) mode is turned on; not running the following actions against your warehouse:
2021-06-03T21:56:56.961267Zstdout: run
2021-06-03T21:56:58.239315Zstdout: Compiling...
2021-06-03T21:56:58.581361Zstdout: [32mCompiled successfully.
2021-06-03T21:56:58.599821Zstdout: Running...
D2021-06-03T21:56:58.615214Zchild process finished
2021-06-03T21:56:58.622666Zchild process exited with code 0
2021-06-03T21:56:58.632653ZGET200793 B5.6 scurl/7.64.0 https://url
```
</details>
