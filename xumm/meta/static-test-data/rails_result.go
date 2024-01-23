package statictestdata

import "github.com/xyield/xumm-go-client/xumm/models"

var RailsResponseResult = &models.RailsResponse{
	models.MAINNET: models.Rail{
		ChainID:     models.MainnetChainID,
		Color:       models.MainnetColor,
		Name:        models.MainnetName,
		IsLivenet:   models.MainnetIsLivenet,
		NativeAsset: models.MainnetNativeAsset,
		Endpoints: []models.RailsEndpoint{
			{
				Name: models.MainnetEndpointsXRPLClusterName,
				Url:  models.MainnetEndpointsXRPLClusterUrl,
			},
			{
				Name: models.MainnetEndpointsRippleS2Name,
				Url:  models.MainnetEndpointsRippleS2Url,
			},
			{
				Name: models.MainnetEndpointsXRPLFallbackName,
				Url:  models.MainnetEndpointsXRPLFallbackUrl,
			},
		},
		Explorers: []models.RailsExplorer{
			{
				Name:       models.MainnetExplorersBithompName,
				URLTx:      models.MainnetExplorersBithompURLTx,
				URLCTID:    models.MainnetExplorersBithompURLCTID,
				URLAccount: models.MainnetExplorersBithompURLAccount,
			},
			{
				Name:       models.MainnetExplorersXRPScanName,
				URLTx:      models.MainnetExplorersXRPScanURLTx,
				URLAccount: models.MainnetExplorersXRPScanURLAccount,
			},
			{
				Name:       models.MainnetExplorersXRPLOrgName,
				URLTx:      models.MainnetExplorersXRPLOrgURLTx,
				URLAccount: models.MainnetExplorersXRPLOrgURLAccount,
			},
			{
				Name:       models.MainnetExplorersXRPLFTechnicalName,
				URLTx:      models.MainnetExplorersXRPLFTechnicalURLTx,
				URLCTID:    models.MainnetExplorersXRPLFTechnicalCTID,
				URLAccount: models.MainnetExplorersXRPLFTechnicalURLAccount,
			},
		},
		RPC:         models.MainnetRPC,
		Definitions: models.MainnetDefinitions,
		Icons: models.RailsIcon{
			IconSquare: models.MainnetIconSquare,
			IconAsset:  models.MainnetIconAsset,
		},
	},
	models.TESTNET: models.Rail{
		ChainID:     models.TestnetChainID,
		Color:       models.TestnetColor,
		Name:        models.TestnetName,
		IsLivenet:   models.TestnetIsLivenet,
		NativeAsset: models.TestnetNativeAsset,
		Faucet:      models.TestnetFaucet,
		XpopEndpoints: []models.XpopEndpoints{
			models.TestnetXpopEndpointXRPLLabs,
		},
		Endpoints: []models.RailsEndpoint{
			{
				Name: models.TestnetEndpointsXRPLLabsTestnetName,
				Url:  models.TestnetEndpointsXRPLLabsTestnetUrl,
			},
			{
				Name: models.TestnetEndpointsRippleXRPLTestnetName,
				Url:  models.TestnetEndpointsRippleXRPLTestnetUrl,
			},
		},
		Explorers: []models.RailsExplorer{
			{
				Name:       models.TestnetExplorersBithompName,
				URLTx:      models.TestnetExplorersBithompURLTx,
				URLCTID:    models.TestnetExplorersBithompURLCTID,
				URLAccount: models.TestnetExplorersBithompURLAccount,
			},
			{
				Name:       models.TestnetExplorersXRPLOrgName,
				URLTx:      models.TestnetExplorersXRPLOrgURLTx,
				URLAccount: models.TestnetExplorersXRPLOrgURLAccount,
			},
			{
				Name:       models.TestnetExplorersXRPLFTechnicalName,
				URLTx:      models.TestnetExplorersXRPLFTechnicalURLTx,
				URLAccount: models.TestnetExplorersXRPLFTechnicalURLAccount,
				URLCTID:    models.TestnetExplorersXRPLFTechnicalURLCTID,
			},
		},
		RPC:         models.TestnetRPC,
		Definitions: models.TestnetDefinitions,
		Icons: models.RailsIcon{
			IconSquare: models.TestnetIconSquare,
			IconAsset:  models.TestnetIconAsset,
		},
	},
	models.DEVNET: models.Rail{
		ChainID:     models.DevnetChainID,
		Color:       models.DevnetColor,
		Name:        models.DevnetName,
		IsLivenet:   models.DevnetIsLivenet,
		NativeAsset: models.DevnetNativeAsset,
		Faucet:      models.DevnetFaucet,
		Endpoints: []models.RailsEndpoint{
			{
				Name: models.DevnetEndpointsRippleXRPLDevnetName,
				Url:  models.DevnetEndpointsRippleXRPLDevnetUrl,
			},
		},
		Explorers: []models.RailsExplorer{
			{
				Name:       models.DevnetExplorersXRPLOrgName,
				URLTx:      models.DevnetExplorersXRPLOrgURLTx,
				URLAccount: models.DevnetExplorersXRPLOrgURLAccount,
			},
		},
		RPC:         models.DevnetRPC,
		Definitions: models.DevnetDefinitions,
		Icons: models.RailsIcon{
			IconSquare: models.DevnetIconSquare,
			IconAsset:  models.DevnetIconAsset,
		},
	},
	models.XAHAU: models.Rail{
		ChainID:     models.XahauChainID,
		Color:       models.XahauColor,
		Name:        models.XahauName,
		IsLivenet:   models.XahauIsLivenet,
		NativeAsset: models.XahauNativeAsset,
		Endpoints: []models.RailsEndpoint{
			{
				Name: models.XahauEndpointsXahauNetworkName,
				Url:  models.XahauEndpointsXahauNetworkUrl,
			},
		},
		Explorers: []models.RailsExplorer{
			{
				Name:       models.XahauExplorersBithompName,
				URLTx:      models.XahauExplorersBithompURLTx,
				URLCTID:    models.XahauExplorersBithompURLCTID,
				URLAccount: models.XahauExplorersBithompURLAccount,
			},
			{
				Name:             models.XahauExplorersXRPLFTechnicalName,
				URLTx:            models.XahauExplorersXRPLFTechnicalURLTx,
				URLAccount:       models.XahauExplorersXRPLFTechnicalURLAccount,
				URLAccountNonExp: models.XahauExplorersXRPLFTechnicalURLAccountNonExp,
			},
		},
		RPC:         models.XahauRPC,
		Definitions: models.XahauDefinitions,
		Icons: models.RailsIcon{
			IconSquare: models.XahauIconSquare,
			IconAsset:  models.XahauIconAsset,
		},
	},
	models.XAHAUTESTNET: models.Rail{
		ChainID:     models.XahauTestnetChainID,
		Color:       models.XahauTestnetColor,
		Name:        models.XahauTestnetName,
		IsLivenet:   models.XahauTestnetIsLivenet,
		NativeAsset: models.XahauTestnetNativeAsset,
		Faucet:      models.XahauTestnetFaucet,
		Endpoints: []models.RailsEndpoint{
			{
				Name: models.XahauTestnetEndpointsXahauTestnetName,
				Url:  models.XahauTestnetEndpointsXahauTestnetUrl,
			},
		},
		Explorers: []models.RailsExplorer{
			{
				Name:       models.XahauTestnetExplorersBithompName,
				URLTx:      models.XahauTestnetExplorersBithompURLTx,
				URLCTID:    models.XahauTestnetExplorersBithompURLCTID,
				URLAccount: models.XahauTestnetExplorersBithompURLAccount,
			},
			{
				Name:       models.XahauTestnetExplorersXRPLFTechnicalName,
				URLTx:      models.XahauTestnetExplorersXRPLFTechnicalURLTx,
				URLAccount: models.XahauTestnetExplorersXRPLFTechnicalURLAccount,
				URLCTID:    models.XahauTestnetExplorersXRPLFTechnicalURLCTID,
			},
		},

		RPC:         models.XahauTestnetRPC,
		Definitions: models.XahauTestnetDefinitions,
		Icons: models.RailsIcon{
			IconSquare: models.XahauTestnetIconSquare,
			IconAsset:  models.XahauTestnetIconAsset,
		},
	},
}
