module github.com/wercker/stern

go 1.14

require (
	github.com/Azure/go-autorest/autorest v0.10.2 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.8.3 // indirect
	github.com/fatih/color v1.7.1-0.20180516100307-2d684516a886
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mitchellh/go-homedir v0.0.0-20161203194507-b8bc1bf76747
	github.com/pkg/errors v0.8.1-0.20180311214515-816c9085562c
	github.com/spf13/cobra v0.0.4-0.20180629152535-a114f312e075
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/net v0.0.0-20200528225125-3c3fba18258b // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	k8s.io/api v0.0.0-20191016110408-35e52d86657a
	k8s.io/apimachinery v0.0.0-20191004115801-a2eda9f80ab8
	k8s.io/client-go v12.0.0+incompatible
)

replace (
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.3.0
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191016111102-bec269661e48
)
