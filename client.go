/*
Package pagarmeapisdk

This file was automatically generated by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package pagarmeapisdk

import (
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "net/http"
    "pagarmeapisdk/controllers"
)

type client struct {
    callBuilderFactory      https.CallBuilderFactory
    config                  Configuration
    userAgent               string
    SubscriptionsController controllers.SubscriptionsController
    OrdersController        controllers.OrdersController
    PlansController         controllers.PlansController
    InvoicesController      controllers.InvoicesController
    CustomersController     controllers.CustomersController
    ChargesController       controllers.ChargesController
    RecipientsController    controllers.RecipientsController
    TokensController        controllers.TokensController
    TransactionsController  controllers.TransactionsController
    TransfersController     controllers.TransfersController
}

// Constructor for client.
func NewClient(config Configuration) *client {
    client := &client{
        config: config,
    }
    
    client.userAgent = utilities.UpdateUserAgent("PagarmeApiSDK - Go 6.7.2")
    client.callBuilderFactory = callBuilderHandler(
    	func(server string) string {
    		if server == "" {
    			server = "default"
    		}
    		return getBaseUri(Server(server), client.config)
    	},
    	BasicAuthentication(config),
    	https.NewHttpClient(
    		https.WithTimeout(config.HttpConfiguration.Timeout),
    		https.WithTransport(config.HttpConfiguration.Transport),
    	),
        withUserAgent(client.userAgent),
    )
    
    baseController := controllers.NewBaseController(client)
    client.SubscriptionsController = *controllers.NewSubscriptionsController(*baseController)
    client.OrdersController = *controllers.NewOrdersController(*baseController)
    client.PlansController = *controllers.NewPlansController(*baseController)
    client.InvoicesController = *controllers.NewInvoicesController(*baseController)
    client.CustomersController = *controllers.NewCustomersController(*baseController)
    client.ChargesController = *controllers.NewChargesController(*baseController)
    client.RecipientsController = *controllers.NewRecipientsController(*baseController)
    client.TokensController = *controllers.NewTokensController(*baseController)
    client.TransactionsController = *controllers.NewTransactionsController(*baseController)
    client.TransfersController = *controllers.NewTransfersController(*baseController)
    return client
}

func (c *client) GetCallBuilder() https.CallBuilderFactory {
    return c.callBuilderFactory
}

func getBaseUri(
    server Server,
    config Configuration) string {
    if config.Environment == Environment(PRODUCTION) {
        if server == Server(ENUMDEFAULT) {
            return "https://api.pagar.me/core/v5"
        }
    }
    return "TODO: Select a valid server."
}

type clientOptions func(cb https.CallBuilder)

func callBuilderHandler(
    baseUrlProvider func(server string) string,
    auth https.Authenticator,
    httpClient https.HttpClient,
    opts ...clientOptions) https.CallBuilderFactory {
    callBuilderFactory := https.CreateCallBuilderFactory(baseUrlProvider, auth, httpClient)
    return tap(callBuilderFactory, opts...)
}

func tap(
    callBuilderFactory https.CallBuilderFactory,
    opts ...clientOptions) https.CallBuilderFactory {
    return func(httpMethod, path string) https.CallBuilder {
    	callBuilder := callBuilderFactory(httpMethod, path)
    	for _, opt := range opts {
    		opt(callBuilder)
    	}
    	return callBuilder
    }
}

func withUserAgent(userAgent string) clientOptions {
    f := func(request *http.Request) *http.Request {
    	request.Header.Add("user-agent", userAgent)
    	return request
    }
    return func(cb https.CallBuilder) {
    	cb.InterceptRequest(f)
    }
}
