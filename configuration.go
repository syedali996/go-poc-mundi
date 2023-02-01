package pagarmeapisdk

import (
    "github.com/apimatic/go-core-runtime/https"
    "os"
)

type Configuration struct {
    Environment       Environment
    HttpConfiguration https.HttpConfiguration
    BasicAuthUserName string
    BasicAuthPassword string
}

func (c *Configuration) LoadFromEnvironment() {
    environment := os.Getenv("PAGARMEAPISDK_ENVIRONMENT")
    if environment != "" {
        c.Environment = Environment(environment)
    }
    basicAuthUserName := os.Getenv("PAGARMEAPISDK_BASIC_AUTH_USER_NAME")
    if basicAuthUserName != "" {
        c.BasicAuthUserName = basicAuthUserName
    }
    basicAuthPassword := os.Getenv("PAGARMEAPISDK_BASIC_AUTH_PASSWORD")
    if basicAuthPassword != "" {
        c.BasicAuthPassword = basicAuthPassword
    }
}

// This is a type declaration.
func DefaultConfig() Configuration {

    return Configuration{
        Environment:       PRODUCTION,                       
        HttpConfiguration: https.DefaultHttpConfiguration(), 
        BasicAuthUserName: "TODO",                           
        BasicAuthPassword: "TODO",                           
    }
}

// Available Servers
type Server string

const (
    ENUMDEFAULT Server = "default"
)

// Available Environments
type Environment string

const (
    PRODUCTION Environment = "production"
)
