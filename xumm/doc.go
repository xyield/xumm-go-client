/*

This client provides an SDK to interact with the XUMM API in a golang application.

Detailed examples can be found in GitHub repository.

Quick Start

To use this library you will need to configure a new Config object e.g.
	cfg, _ := xumm.NewConfig()

If you do not provide any arguments, reasonable defaults are used such as:
	* standard Golang library HTTP client
	* XUMM api base URL of https://xumm.app/api/v1
	* check your environment variables for XUMM_API_KEY and XUMM_API_SECRET - use the WithAuth() method to set these manually instead

The environment variables are used to set authentication headers.

Next, create a client using this config:
	client := client.New(cfg)

Call methods as follows:
	pong, err := client.Meta.Ping()
*/
package xumm
