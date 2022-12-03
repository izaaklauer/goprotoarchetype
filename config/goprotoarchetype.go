package config

type Goprotoarchetype struct {
	HelloWorldMessage string `hcl:"hello_world_message,attr"`

	// ... your config here
}

// DefaultGoprotoarchetypeConfig returns default config values
func DefaultGoprotoarchetypeConfig() Goprotoarchetype {
	return Goprotoarchetype{
		HelloWorldMessage: "hello from the default config",
	}
}
