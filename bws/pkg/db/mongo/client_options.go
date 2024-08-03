package mongo

import "context"

type ClientOption = func(*ClientOptions)

type ClientOptions struct {
	Context                 context.Context
	Uri                     string
	Host                    string
	Port                    string
	Db                      string
	Hosts                   []string
	Username                string
	Password                string
	AuthSource              string
	AuthMechanism           string
	AuthMechanismProperties map[string]string
}

// Client options builder
func WithContext(ctx context.Context) ClientOption {
	return func(opts *ClientOptions) {
		opts.Context = ctx
	}
}

func WithUri(uri string) ClientOption {
	return func(opts *ClientOptions) {
		opts.Uri = uri
	}
}

func WithHost(host string) ClientOption {
	return func(opts *ClientOptions) {
		opts.Host = host
	}
}

func WithPort(port string) ClientOption {
	return func(opts *ClientOptions) {
		opts.Port = port
	}
}

func WithDb(db string) ClientOption {
	return func(opts *ClientOptions) {
		opts.Db = db
	}
}

func WithHosts(hosts []string) ClientOption {
	return func(opts *ClientOptions) {
		opts.Hosts = hosts
	}
}

func WithUsername(username string) ClientOption {
	return func(opts *ClientOptions) {
		opts.Username = username
	}
}

func WithPassword(password string) ClientOption {
	return func(opts *ClientOptions) {
		opts.Password = password
	}
}

func WithAuthSource(authSource string) ClientOption {
	return func(opts *ClientOptions) {
		opts.AuthSource = authSource
	}
}

func WithAuthMechanism(authMechanism string) ClientOption {
	return func(opts *ClientOptions) {
		opts.AuthMechanism = authMechanism
	}
}

func WithAuthMechanismProperties(authMechanismProperties map[string]string) ClientOption {
	return func(opts *ClientOptions) {
		opts.AuthMechanismProperties = authMechanismProperties
	}
}
