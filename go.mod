module github.com/Slahser/coup-de-grace

go 1.15

require (
	github.com/c-bata/go-prompt v0.2.3
	github.com/containous/traefik/v2 v2.2.11
	github.com/coreos/etcd v3.3.24+incompatible
	github.com/deckarep/gosx-notifier v0.0.0-20180201035817-e127226297fb
	github.com/dimiro1/banner v1.0.0
	github.com/felixge/fgprof v0.9.1
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.6.3
	github.com/go-errors/errors v1.1.1
	github.com/go-openapi/strfmt v0.19.5 // indirect
	github.com/go-playground/validator/v10 v10.3.0
	github.com/goproxy/goproxy v0.2.1
	github.com/guptarohit/asciigraph v0.5.0
	github.com/hashicorp/go-multierror v1.1.0
	github.com/inconshreveable/go-update v0.0.0-20160112193335-8152e7eb6ccf
	github.com/jedib0t/go-pretty v4.3.0+incompatible
	github.com/jpillora/opts v1.2.0
	github.com/json-iterator/go v1.1.10
	github.com/magefile/mage v1.10.0
	github.com/mattn/go-tty v0.0.3 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/moby/buildkit v0.7.2
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/nuclio/nuclio v0.0.0-00010101000000-000000000000
	github.com/nuclio/zap v0.0.3
	github.com/pkg/errors v0.9.1
	github.com/pkg/term v0.0.0-20200520122047-c3ffed290a03 // indirect
	github.com/serialx/hashring v0.0.0-20200727003509-22c0c7ab6b1b
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/v3io/version-go v0.0.2
	go.uber.org/zap v1.10.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	k8s.io/client-go v0.19.0
	k8s.io/apimachinery v0.19.0
)

replace (
	// Containous forks
	github.com/abbot/go-http-auth => github.com/containous/go-http-auth v0.4.1-0.20200324110947-a37a7636d23e
	github.com/containerd/containerd => github.com/containerd/containerd v1.4.0

	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	github.com/docker/docker => github.com/docker/engine v1.4.2-0.20200204220554-5f6d6f3f2203
	github.com/go-check/check => github.com/containous/check v0.0.0-20170915194414-ca0bf163426a
	github.com/golang/protobuf => github.com/golang/protobuf v1.3.5
	github.com/gorilla/mux => github.com/containous/mux v0.0.0-20181024131434-c33f32e26898
	github.com/hashicorp/go-immutable-radix => github.com/tonistiigi/go-immutable-radix v0.0.0-20170803185627-826af9ccf0fe
	github.com/jaguilar/vt100 => github.com/tonistiigi/vt100 v0.0.0-20190402012908-ad4c4a574305
	github.com/mailgun/minheap => github.com/containous/minheap v0.0.0-20190809180810-6e71eb837595
	github.com/mailgun/multibuf => github.com/containous/multibuf v0.0.0-20190809014333-8b6c9a7e6bba
	github.com/nuclio/nuclio => ./third_party/nuclio

	//replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

	google.golang.org/grpc => google.golang.org/grpc v1.27.0

	k8s.io/api => k8s.io/api v0.16.9
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.9
	k8s.io/client-go => k8s.io/client-go v0.16.9
)
