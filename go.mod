module open-cluster-management.io/sdk-go

go 1.22.5

toolchain go1.22.10

replace open-cluster-management.io/api => github.com/haoqing0110/api v0.0.0-20250103062352-f89d6a235b5d

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/cloudevents/sdk-go/protocol/kafka_confluent/v2 v2.0.0-20240413090539-7fef29478991
	github.com/cloudevents/sdk-go/protocol/mqtt_paho/v2 v2.0.0-20241008145627-6bcc075b5b6c
	github.com/cloudevents/sdk-go/v2 v2.15.3-0.20240911135016-682f3a9684e4
	github.com/confluentinc/confluent-kafka-go/v2 v2.3.0
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc
	github.com/eclipse/paho.golang v0.21.0
	github.com/evanphx/json-patch v5.9.0+incompatible
	github.com/golang/protobuf v1.5.4
	github.com/google/cel-go v0.17.8
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.6.0
	github.com/mochi-mqtt/server/v2 v2.6.5
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.34.1
	github.com/openshift/build-machinery-go v0.0.0-20240419090851-af9c868bcf52
	github.com/openshift/library-go v0.0.0-20240621150525-4bb4238aef81
	github.com/prometheus/client_golang v1.18.0
	github.com/prometheus/client_model v0.5.0
	github.com/stretchr/testify v1.9.0
	golang.org/x/oauth2 v0.20.0
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.30.3
	k8s.io/apimachinery v0.30.3
	k8s.io/client-go v0.30.3
	k8s.io/klog/v2 v2.130.1
	k8s.io/utils v0.0.0-20240310230437-4693a0247e57
	open-cluster-management.io/api v0.15.0
	open-cluster-management.io/ocm v0.15.2
	sigs.k8s.io/controller-runtime v0.18.5
	sigs.k8s.io/yaml v1.4.0
)

require (
	cloud.google.com/go/compute/metadata v0.3.0 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230305170008-8188dc5388df // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/evanphx/json-patch/v5 v5.9.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.23.0 // indirect
	golang.org/x/term v0.23.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240701130421-f6361c86f094 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240701130421-f6361c86f094 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.30.3 // indirect
	k8s.io/apiserver v0.30.3 // indirect
	k8s.io/component-base v0.30.3 // indirect
	k8s.io/kube-openapi v0.0.0-20240228011516-70dd3763d340 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
)
