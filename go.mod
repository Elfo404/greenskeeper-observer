module github.com/Elfo404/greenskeeper-observer

go 1.16

require (
	go.uber.org/zap v1.19.1
	tinygo.org/x/bluetooth v0.4.0
)

require (
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/godbus/dbus/v5 v5.0.6 // indirect
	github.com/muka/go-bluetooth v0.0.0-20211122080231-b99792bbe62a // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/sys v0.0.0-20211204120058-94396e421777 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace tinygo.org/x/bluetooth v0.4.0 => github.com/lacendarko/bluetooth v0.3.1-0.20210924082245-3503977da2c2
