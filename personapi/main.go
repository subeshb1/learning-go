package main

import (
	"personapi/bindings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	gy "github.com/graniticio/granitic-yaml/v2"
)

func main() {

	gy.StartGraniticWithYaml(bindings.Components())
}
