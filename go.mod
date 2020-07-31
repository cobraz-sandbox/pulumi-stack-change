module github.com/cobraz-sandbox/pulumi-stack-change

go 1.14

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.5.0
)

require (
	github.com/pulumi/pulumi/pkg/v2 v2.7.1 // indirect
	github.com/pulumi/pulumi/sdk/v2 v2.7.1 // indirect
)
