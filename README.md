# lamgo

This is a simple AWS Lambda template for Go to write serverless functions as a single binary.

It uses [AWS SAM](https://aws.amazon.com/serverless/sam/) to build and deploy the Lambda function.

## Prerequisites

- [AWS CLI](https://aws.amazon.com/cli/)
- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
- [Go](https://golang.org/doc/install)

## Usage

1. Clone this repository
2. Run `make` to build the binary
3. Run `make deploy` to deploy the Lambda function


## Technologies

- [Go](https://golang.org/)
- [Echo](https://echo.labstack.com/)
- [AWS SAM](https://aws.amazon.com/serverless/sam/)
- [Make](https://www.gnu.org/software/make/)
- [AWS SDK for Go](https://aws.amazon.com/sdk-for-go/)
- [Logrus](https://github.com/sirupsen/logrus)

## License

[MIT](./LICENSE)