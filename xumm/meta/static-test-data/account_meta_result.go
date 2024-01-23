package statictestdata

import "github.com/xyield/xumm-go-client/xumm/models"

var AccountMetaTestResult = &models.AccountMetaResponse{
	Account:     "rwietsevLFg8XSmG3bEZzFein1g8RBqWDZ",
	KycApproved: true,
	XummPro:     true,
	Blocked:     false,
	ForceDtag:   false,
	Avatar:      "https://xumm.app/avatar/rwietsevLFg8XSmG3bEZzFein1g8RBqWDZ.png",
	XummProfile: models.XummProfile{
		AccountAlias: "XRPL Labs - Wietse Wind",
		OwnerAlias:   "Wietse Wind",
		Slug:         "wietsewind",
		ProfileURL:   "https://xumm.me/wietsewind",
		AccountSlug:  "",
		PayString:    "wietsewind$xumm.me",
	},
	ThirdPartyProfiles: []models.ThirdPartyProfile{
		{
			AccountAlias: "Wietse Wind",
			Source:       "xumm.app",
		},
		{
			AccountAlias: "wietse.com",
			Source:       "xrpl",
		},
		{
			AccountAlias: "XRPL-Labs",
			Source:       "bithomp.com",
		},
	},
	GlobalID: models.GlobalID{
		Linked:          "2021-06-29T10:22:25.000Z",
		ProfileURL:      "https://app.global.id/u/wietse",
		SufficientTrust: true,
	},
}
