package models

type RailsResponse struct {
}

type Rail struct {
	ChainID     int             `json:"chain_id,omitempty"`
	Color       string          `json:"color,omitempty"`
	Name        string          `json:"name,omitempty"`
	IsLivenet   bool            `json:"is_livenet,omitempty"`
	NativeAsset string          `json:"native_asset,omitempty"`
	Faucet      string          `json:"faucet,omitempty"`
	Endpoints   []RailsEndpoint `json:"endpoints,omitempty"`
	Explorers   []RailsExplorer `json:"explorers,omitempty"`
	RPC         string          `json:"rpc,omitempty"`
	Definitions string          `json:"definitions,omitempty"`
	Icons       RailsIcon       `json:"icons,omitempty"`
}

type RailsEndpoint struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type RailsExplorer struct {
	Name        string `json:"name,omitempty"`
	URLTx       string `json:"url_tx,omitempty"`
	URLCTID     string `json:"url_ctid,omitempty"`
	URLAccount  string `json:"url_account,omitempty"`
	_URLAccount string `json:"_url_account,omitempty"`
}

type RailsIcon struct {
	IconSquare string `json:"icon_square,omitempty"`
	IconAsset  string `json:"icon_asset,omitempty"`
}
