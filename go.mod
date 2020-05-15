module github.com/go-msvc/examples

go 1.12

require (
	github.com/go-msvc/config v0.0.0-20191007220741-e8517c202db8
	github.com/go-msvc/errors v0.0.0-20191116111408-1c2c4914594f
	github.com/go-msvc/log v0.0.0-20200515104948-e039d1c2f30d
	github.com/go-msvc/logger v0.0.0-20191206150445-df0ebdb23067
	github.com/go-msvc/ms v0.0.0-20200225120321-51a419b6a727
	github.com/go-msvc/service v0.0.0-00010101000000-000000000000
)

replace github.com/go-msvc/config => ../config

replace github.com/go-msvc/log => ../log

replace github.com/go-msvc/errors => ../errors

replace github.com/go-msvc/service => ../service
