module github.com/aws/aws-sdk-go-v2/service/servicediscovery

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.17.0
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.24
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.18
	github.com/aws/smithy-go v1.13.4
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/
