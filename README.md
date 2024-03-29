# Project Repo Title

A little introduction here...

<hr />

## **Requirements**
- Golang 1.21 or higher
- Openssl3

<hr />

## **Dependencies**
- [PymigrateDB](https://pypi.org/project/pymigratedb/)
- [Gcloud CLI](https://cloud.google.com/sdk/docs/install)

<hr />

## Setup
- To install project run ```make```.
#- It will create two new files in the project root, called: ".env" and "config.yaml"
#- Fill it with correct values to procceed with development

<hr />

## Run
- To run project execute ```make run``` into the terminal. It will start the API and serve the requests connected to the resources filled in the config files.

<hr />

## Tests
- To run and perform test cases, run the following command: ```make test```. It will begin the tests execution.

<hr />

## Deploy

The deploy is performed using some [GCP](https://cloud.google.com/?hl=pt-br) services.

1. Register service account and principal on **Innovation IAM** with the correct permissions.

    Obs.: Use the correct **Innovation IAM Backoffice** accordingly to the destiny environment.

2. Upload all necessary files to right folder in [Bucket](https://console.cloud.google.com/storage) such as `private.pem` and `innovation.json` (download it in **Innovation IAM Backoffice**).

3. Map all environment variables in `.env.example` and provide them in "docker build" step on `cloudbuild.yaml`.

4. Check if `Dockerfile` is receiving all arguments provided in "docker build" step on `cloudbuild.yaml`.

5. Map all configuration variables in `config.example.yaml` and add them to the [Secret Manager](https://console.cloud.google.com/security/secret-manager).

    Obs.: All variables should be prefixed with the project name, to prevent conflict with other projects. Example: **db_string** turns into **`project_name`_db_string**.

6. Create a trigger in [Cloud Build](https://console.cloud.google.com/cloud-build) pointing to the project repository, prefer to use the second gen **Source**.

   * Add all the **substitution variables** provided in `cloudbuild.yaml`.

7. Check if worker pool is created in [Cloud Build](https://console.cloud.google.com/cloud-build).

8. Check if repository is created in [Google Artifact Registry](https://console.cloud.google.com/artifacts).

9. If it's not configured to run automatically when a commit is made to develop or main, run the trigger by commit hash or branch.

While CI/CD is running pay attention if the specified files will be copied and if all the necessary migrations will run.

## Accepted Methods and Content-Types

| Method | Content-Type |
|:------:|:------------:|
|POST    |application/json|
|GET     |
|OPTIONS |

<hr />

## Relevant Environment Variables

| Variable | Possible Values | Description |
|:--------:|:---------------:|:-----------:|
|ENVIRONMENT|local, staging,  production|Environment where deployment is working|
|USE_SECRETMANAGER|true,false,""|Decide to use or not Google Cloud Secret Manager|
|INNOVATION_CREDENTIALS|path to credentials.json|Indico IAM Credentials file to integrate with authorization and authentication functionalities|
|SEC_PREFIX|desirable prefix|Prefix to use inside Secret Manager|
|AUTH_SERVER|iam auth server uri|IAM Auth Server URI|
|INSECURE|true,false,""|Use IAM Server in insecure mode|
|USE_TLS|true,false,""|Use HTTPS connections|
|PREFORK|true,false,""|Use prefork threads|

<hr />

## API Structure

```bash
.
├── adapters
│   ├── database
│   │   └── database.go
│   ├── logging
│   │   └── logging.go
│   ├── storage
│   │   └── storage.go
│   └── totp
│       └── totp.go
├── app
│   ├── errors
│   │   └── errors.go
│   ├── repository
│   │   └── .gitkeep
│   └── usecases
│       └── .gitkeep
├── clients
│   ├── google
│   │   ├── logging
│   │   │   └── logging.go
│   │   └── storage
│   │       └── storage.go
│   ├── iam
│   │   └── client.go
│   └── postgres
│       └── postgres.go
├── config
│   ├── constants
│   │   └── constants.go
│   └── config.go
├── entity
│   ├── http_response.go
│   └── log.go
├── handler
│   └── health
│       └── health.go
├── middleware
│   ├── auth.go
│   ├── content.go
│   └── security.go
├── pkg
│   ├── app
│   │   └── app.go
│   ├── crypt
│   │   └── crypt.go
│   ├── helpers
│   │   ├── http.go
│   │   ├── json.go
│   │   └── utils.go
│   └── jwt
│       └── jwt.go
├── .dockerignore
├── .env.example
├── .gitignore
├── .innovation_env
├── config.example.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── README.md
└── route.go
```