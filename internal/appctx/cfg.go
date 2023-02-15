// Package appctx
package appctx

import (
	"fmt"

	"github.com/bachtiarfr/pubsub-logger/v2/internal/consts"
	"github.com/bachtiarfr/pubsub-logger/v2/pkg/file"
)

// NewConfig initialize config object
func NewConfig() (*Config, error) {
	fpath := []string{consts.ConfigPath}
	cfg, err := readCfg("app.yaml", fpath...)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

//go:generate easytags $GOFILE yaml,json
// Config object contract
type Config struct {
	PubSub  PubSub       `yaml:"pub_sub" json:"pub_sub"`
}

// readCfg reads the configuration from file
// args:
//	fname: filename
//	ps: full path of possible configuration files
// returns:
//	*config.Configuration: configuration ptr object
//	error: error operation
func readCfg(fname string, ps ...string) (*Config, error) {
	var cfg *Config
	var errs []error

	for _, p := range ps {
		f := fmt.Sprint(p, fname)

		err := file.ReadFromYAML(f, &cfg)
		if err != nil {
			errs = append(errs, fmt.Errorf("file %s error %s", f, err.Error()))
			continue
		}
		break
	}

	if cfg == nil {
		return nil, fmt.Errorf("file config parse error %v", errs)
	}

	return cfg, nil
}

// PubSub config
type PubSub struct {
	ProjectID      string `yaml:"project_id" json:"project_id"`
	CredentialFile string `yaml:"credential_file" json:"credential_file"`
	TopicID        string `yaml:"topic_id" json:"topic_id"`
	SubscriptionID string `yaml:"subscription_id" json:"subscription_id"`
}
