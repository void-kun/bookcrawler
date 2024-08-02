package main

type ClientOption func(options *ClientOptions)

type ClientOptions struct {
	Uri                     string            `json:"uri,omitempty"`
	Host                    string            `json:"host,omitempty"`
	Port                    string            `json:"port,omitempty"`
	Db                      string            `json:"db,omitempty"`
	Hosts                   []string          `json:"hosts,omitempty"`
	Username                string            `json:"username,omitempty"`
	Password                string            `json:"password,omitempty"`
	AuthSource              string            `json:"auth_source,omitempty"`
	AuthMechanism           string            `json:"auth_mechanism,omitempty"`
	AuthMechanismProperties map[string]string `json:"auth_mechanism_properties,omitempty"`
}

func WithUri(value string) ClientOption {
	return func(options *ClientOptions) {
		options.Uri = value
	}
}

func WithHost(value string) ClientOption {
	return func(options *ClientOptions) {
		options.Host = value
	}
}

func WithPort(value string) ClientOption {
	return func(options *ClientOptions) {
		options.Port = value
	}
}

func main() {
	var opts []ClientOption
	opts = append(opts, WithUri("Url"))
	opts = append(opts, WithHost("Host"))
	opts = append(opts, WithPort("Port"))

	_opts := &ClientOptions{}
}
