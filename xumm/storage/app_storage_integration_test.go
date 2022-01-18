// +build integration

package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestStorageIntegration(t *testing.T) {
	cfg, _ := xumm.NewConfig()
	s := Storage{
		Cfg: cfg,
	}
	json, err := s.DeleteAppStorage()

	assert.NoError(t, err)
	assert.Equal(t, &models.AppStorageResponse{
		Application: models.Application{
			UUIDV4:   "eda1fbb4-8641-47bd-91c8-3adca27cd6e3",
			Name:     "Test Xumm App",
			Disabled: 0,
		},
		Stored: true,
	}, json)

	json, err = s.GetAppStorage()

	assert.NoError(t, err)
	assert.Equal(t, &models.AppStorageResponse{
		Application: models.Application{
			UUIDV4:   "eda1fbb4-8641-47bd-91c8-3adca27cd6e3",
			Name:     "Test Xumm App",
			Disabled: 0,
		},
		Stored: false,
	}, json)

	json, err = s.SetAppStorage(anyjson.AnyJson{
		"test": "value",
	})

	assert.NoError(t, err)
	assert.Equal(t, &models.AppStorageResponse{
		Application: models.Application{
			UUIDV4:   "eda1fbb4-8641-47bd-91c8-3adca27cd6e3",
			Name:     "Test Xumm App",
			Disabled: 0,
		},
		Stored: true,
		Data: anyjson.AnyJson{
			"test": "value",
		},
	}, json)

	json, err = s.GetAppStorage()
	assert.NoError(t, err)
	assert.Equal(t, &models.AppStorageResponse{
		Application: models.Application{
			UUIDV4:   "eda1fbb4-8641-47bd-91c8-3adca27cd6e3",
			Name:     "Test Xumm App",
			Disabled: 0,
		},
		Stored: false,
		Data: anyjson.AnyJson{
			"test": "value",
		},
	}, json)

}
