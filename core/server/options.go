package server

var (
	DefaultEtcdAddress = []string{"127.0.0.1:2379"}
	DefaultServerName  = "real.micro"
	DefaultVersion     = "1.0.0"
)

type Options struct {
	Address          []string
	RegisterTTL      int64
	RegisterInterval int64

	ServerName string
	Version    string
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		Address:    DefaultEtcdAddress,
		ServerName: DefaultServerName,
		Version:    DefaultVersion,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func ServerName(ServerName string) Option {
	return func(o *Options) {
		o.ServerName = ServerName
	}
}

func Version(Version string) Option {
	return func(o *Options) {
		o.Version = Version
	}
}

func EtcdAddress(address []string) Option {
	return func(o *Options) {
		o.Address = address
	}
}

func EtcdRegisterTTL(RegisterTTL int64) Option {
	return func(o *Options) {
		o.RegisterTTL = RegisterTTL
	}
}

func EtcdRegisterInterval(RegisterInterval int64) Option {
	return func(o *Options) {
		o.RegisterInterval = RegisterInterval
	}
}
