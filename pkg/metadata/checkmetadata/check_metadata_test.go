// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package checkmetadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPayload(t *testing.T) {
	// empty cache, empty payload
	p := *GetPayload()
	assert.Len(t, p, 0)

	checkID := "directory:12345"
	metadata1 := "{\"option\":true}"
	metadata2 := "{\"option\":false}"

	// add a few payloads to the cache
	SetCheckMetadata(checkID, metadata1)
	SetCheckMetadata(checkID, metadata2)

	p = *GetPayload()
	assert.Len(t, p, 1)
	assert.Equal(t, p[0], metadata2)

	// GetPayload is supposed to empty the cache
	assert.Len(t, checkMetadataCache, 0)
}
