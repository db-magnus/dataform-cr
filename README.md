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

Enable apis
```
gcloud services enable cloudbuild.googleapis.com \
run.googleapis.com \
secretmanager.googleapis.com \
sourcerepo.googleapis.com \
artifactregistry.googleapis.com \
cloudresourcemanager.googleapis.com
```

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
Upload this file to Secret Manager
```
gcloud secrets create dataform-sa-key
gcloud secrets versions add dataform-sa-key --data-file=.df-credentials.json
```
Go to Secret Manager in the console, permissions. Add the default compute service account as secretmanager.secretAccesor

In Cloud Source, make a trigger to run cloud build on push to git.
In cloud build, Settings, allow Cloud Build to manage Cloud Run by clicking enable.  


## Usage

Call Cloud Run with
```
$ curl -H "Authorization: Bearer $(gcloud auth print-identity-token)" https://url-to-cloud-run

```
