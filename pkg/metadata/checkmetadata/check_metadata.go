// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package checkmetadata

import "sync"

var (
	// checkMetadataCache maps checkID -> latest JSON metadata dump
	checkMetadataCache = make(map[string]string)
	cacheMutex         = &sync.Mutex{}
)

// SetCheckMetadata adds a JSON string to the cache.
func SetCheckMetadata(checkID, metadata string) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	checkMetadataCache[checkID] = metadata
}

// GetPayload fills and returns the check metadata payload
func GetPayload() *Payload {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	payload := Payload{}
	for _, metadata := range checkMetadataCache {
		payload = append(payload, metadata)
	}

	// clear the cache
	checkMetadataCache = make(map[string]string)
	return &payload
}
