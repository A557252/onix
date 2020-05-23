// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitializeScript() (*RInfo, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}
	client := NewSource(config)
	script, err := NewRInfo(config, client)
	if err != nil {
		return nil, err
	}
	return script, nil
}

func InitializeConfig() (*Config, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func InitializeClient() (*Source, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}
	client := NewSource(config)
	return client, nil
}