package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config is the top-level gateway configuration loaded from YAML.
type Config struct {
	Server         ServerConfig    `yaml:"server"`
	DefaultMethods []string        `yaml:"default_methods"`
	Services       []ServiceConfig `yaml:"services"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

// ServiceConfig describes a single upstream service routed by the gateway.
type ServiceConfig struct {
	Name    string   `yaml:"name"`
	Prefix  string   `yaml:"prefix"`
	URL     string   `yaml:"url"`
	Methods []string `yaml:"methods"`

	target *url.URL
}

var validMethods = map[string]bool{
	http.MethodGet: true, http.MethodPost: true, http.MethodPut: true,
	http.MethodPatch: true, http.MethodDelete: true, http.MethodHead: true,
	http.MethodOptions: true,
}

// LoadConfig reads the YAML file at path, expands environment variables,
// applies defaults and validates the result.
func LoadConfig(path string) (*Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config %q: %w", path, err)
	}

	expanded := os.Expand(string(raw), expandWithDefault)

	var cfg Config
	dec := yaml.NewDecoder(strings.NewReader(expanded))
	dec.KnownFields(true)
	if err := dec.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("parsing config %q: %w", path, err)
	}

	if err := cfg.normalize(); err != nil {
		return nil, fmt.Errorf("invalid config %q: %w", path, err)
	}
	return &cfg, nil
}

func (c *Config) normalize() error {
	if c.Server.Port == "" {
		c.Server.Port = "8080"
	}
	if len(c.DefaultMethods) == 0 {
		c.DefaultMethods = []string{
			http.MethodGet, http.MethodPost, http.MethodPut,
			http.MethodPatch, http.MethodDelete,
		}
	}
	if len(c.Services) == 0 {
		return fmt.Errorf("at least one service must be defined")
	}

	seenPrefixes := map[string]string{}
	for i := range c.Services {
		s := &c.Services[i]
		if s.Name == "" {
			return fmt.Errorf("services[%d]: name is required", i)
		}
		if s.URL == "" {
			return fmt.Errorf("service %q: url is required", s.Name)
		}
		if !strings.HasPrefix(s.Prefix, "/") {
			return fmt.Errorf("service %q: prefix must start with '/'", s.Name)
		}
		s.Prefix = strings.TrimSuffix(s.Prefix, "/")
		if other, ok := seenPrefixes[s.Prefix]; ok {
			return fmt.Errorf("service %q: prefix %q already used by %q", s.Name, s.Prefix, other)
		}
		seenPrefixes[s.Prefix] = s.Name

		u, err := url.Parse(s.URL)
		if err != nil || u.Scheme == "" || u.Host == "" {
			return fmt.Errorf("service %q: invalid url %q", s.Name, s.URL)
		}
		s.target = u

		if len(s.Methods) == 0 {
			s.Methods = append([]string(nil), c.DefaultMethods...)
		}
		for j, m := range s.Methods {
			m = strings.ToUpper(m)
			if !validMethods[m] {
				return fmt.Errorf("service %q: unsupported method %q", s.Name, m)
			}
			s.Methods[j] = m
		}
	}
	return nil
}

// expandWithDefault resolves ${VAR} and ${VAR:-default} references.
func expandWithDefault(key string) string {
	name, def, hasDefault := strings.Cut(key, ":-")
	if v := os.Getenv(name); v != "" {
		return v
	}
	if hasDefault {
		return def
	}
	return ""
}
