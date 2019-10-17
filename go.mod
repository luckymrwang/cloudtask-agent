module cloudtask-agent

go 1.12

require (
	cloudtask v0.0.1
	github.com/cloudtask/libtools v0.0.0-20180622030929-385f94132c66
	github.com/coreos/etcd v3.3.17+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/mux v1.7.3
	go.etcd.io/etcd v3.3.17+incompatible // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.2.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/net v0.0.0-20191014212845-da9a3fd4c582 // indirect
	google.golang.org/grpc v1.24.0 // indirect
	gopkg.in/eapache/queue.v1 v1.1.0 // indirect
	gopkg.in/yaml.v2 v2.2.4
)

replace cloudtask v0.0.1 => ../cloudtask
