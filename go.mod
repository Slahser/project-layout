module github.com/Slahser/coup-de-grace

go 1.14

require (
	github.com/c-bata/go-prompt v0.2.3
	github.com/containous/traefik/v2 v2.2.8
	github.com/coreos/etcd v3.3.24+incompatible
	github.com/deckarep/gosx-notifier v0.0.0-20180201035817-e127226297fb
	github.com/dimiro1/banner v1.0.0
	github.com/felixge/fgprof v0.9.0
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.6.3
	github.com/go-openapi/strfmt v0.19.5 // indirect
	github.com/go-playground/validator/v10 v10.3.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/goproxy/goproxy v0.1.5
	github.com/guptarohit/asciigraph v0.5.0
	github.com/jedib0t/go-pretty v4.3.0+incompatible
	github.com/jpillora/opts v1.2.0
	github.com/json-iterator/go v1.1.10
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mattn/go-tty v0.0.3 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/pkg/term v0.0.0-20200520122047-c3ffed290a03 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/v3io/version-go v0.0.2
	go.uber.org/zap v1.10.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	k8s.io/client-go v0.17.3
	k8s.io/kubernetes v1.13.0
	sigs.k8s.io/yaml v1.2.0 // indirect
)

// Docker v19.03.6
replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20200204220554-5f6d6f3f2203

// Containous forks
replace (
	github.com/abbot/go-http-auth => github.com/containous/go-http-auth v0.4.1-0.20200324110947-a37a7636d23e
	github.com/go-check/check => github.com/containous/check v0.0.0-20170915194414-ca0bf163426a
	github.com/gorilla/mux => github.com/containous/mux v0.0.0-20181024131434-c33f32e26898
	github.com/mailgun/minheap => github.com/containous/minheap v0.0.0-20190809180810-6e71eb837595
	github.com/mailgun/multibuf => github.com/containous/multibuf v0.0.0-20190809014333-8b6c9a7e6bba
)
