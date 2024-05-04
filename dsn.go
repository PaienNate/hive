package hive

import "sqlflow.org/gohive"

type DSNConfig gohive.Config

func ParseDSN(dsn string) (*DSNConfig, error) {
	config, err := gohive.ParseDSN(dsn)
	if err != nil {
		return nil, err
	}
	return (*DSNConfig)(config), nil
}

func (c *DSNConfig) FormatDSN() string {
	return (*gohive.Config)(c).FormatDSN()
}
