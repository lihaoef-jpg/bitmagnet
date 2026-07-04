package adapter

import (
	"testing"

	"github.com/bitmagnet-io/bitmagnet/internal/database/search"
	"github.com/bitmagnet-io/bitmagnet/internal/model"
	"github.com/bitmagnet-io/bitmagnet/internal/torznab"
	"github.com/stretchr/testify/assert"
)

func TestTorrentContentResultToTorznabResultSetsResponseTotal(t *testing.T) {
	t.Parallel()

	result := torrentContentResultToTorznabResult(
		torznab.SearchRequest{
			Profile: torznab.ProfileDefault,
			Offset:  model.NewNullUint(100),
		},
		search.TorrentContentResult{
			TotalCount: 250,
		},
	)

	assert.Equal(t, uint(100), result.Channel.Response.Offset)
	assert.Equal(t, uint(250), result.Channel.Response.Total)
}
