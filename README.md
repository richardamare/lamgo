# lamgo

_... as in Lambda + Go_

This is a simple AWS Lambda template for Go to write serverless functions as a single binary.

> Note: Still a work in progress

It uses [AWS SAM](https://aws.amazon.com/serverless/sam/) to build and deploy the Lambda function.

## Technologies

- [Go](https://golang.org/)
- [Echo](https://echo.labstack.com/)
- [AWS SAM](https://aws.amazon.com/serverless/sam/)
- [Make](https://www.gnu.org/software/make/)
- [AWS SDK for Go](https://aws.amazon.com/sdk-for-go/)
- [Zap](https://github.com/uber-go/zap)

## Prerequisites

- [AWS CLI](https://aws.amazon.com/cli/)
- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
- [Go](https://golang.org/doc/install)

## Usage

1. Clone this repository
    ```bash
    git clone https://github.com/richardamare/lamgo.git
    ```
2. Change directory to the cloned repository
    ```bash
    cd lamgo # or whatever you named the directory
    ```
3. Configure the AWS CLI
    ```bash
    aws configure
    ```
4. Configure the AWS SAM deployment
    ```bash
    sam deploy --guided
    target
    ```
   NOTE: Set the `SAM configuration environment` to `dev` or edit the Makefile deployment
5. Deploy the Lambda function to AWS
    ```bash
    make deploy
    ```

Run locally using:
```bash
go run cmd/api/main.go
```

## License

[MIT](./LICENSE)