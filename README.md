# memorize efficiently

This is a sample template for memorize efficiently - Below is a brief explanation of what we have generated for you:

```bash
.
```

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)



### Installing dependencies & building the target

build.sh:
    deploy: upload build files to S3 bucket
    api: run api on local host

### Local development

**Invoking function locally through local API Gateway**

```bash
cd container && docker-compose up -d
```


**SAM CLI** is used to emulate both Lambda and API Gateway locally and uses our `template.yaml` to understand how to bootstrap this environment (runtime, where the source code is, etc.) - The following excerpt is what the CLI will read in order to initialize an API and its routes:

```yaml
...
Events:
    HelloWorld:
        Type: Api # More info about API Event Source: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
        Properties:
            Path: /hello
            Method: get
```

You can find your API Gateway Endpoint URL in the output values displayed after deployment.

