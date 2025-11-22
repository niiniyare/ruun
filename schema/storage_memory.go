package schema

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// MemoryStorage provides in-memory storage
type MemoryStorage struct {
	schemas  map[string][]byte
	versions map[string]map[string][]byte
	metadata map[string]*StorageMetadata
	mu       sync.RWMutex
}

// NewMemoryStorage creates memory storage
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		schemas:  make(map[string][]byte),
		versions: make(map[string]map[string][]byte),
		metadata: make(map[string]*StorageMetadata),
	}
}

func (s *MemoryStorage) Get(ctx context.Context, id string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, ok := s.schemas[id]
	if !ok {
		return nil, fmt.Errorf("schema not found: %s", id)
	}

	// Update metadata
	if meta, ok := s.metadata[id]; ok {
		meta.AccessCount++
		meta.LastAccess = time.Now()
	}

	return data, nil
}

func (s *MemoryStorage) Set(ctx context.Context, id string, data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()

	// Store current version if versioning enabled
	if existing, ok := s.schemas[id]; ok {
		if s.versions[id] == nil {
			s.versions[id] = make(map[string][]byte)
		}
		version := fmt.Sprintf("%d", time.Now().Unix())
		s.versions[id][version] = existing
	}

	s.schemas[id] = data

	// Update metadata
	if meta, ok := s.metadata[id]; ok {
		meta.UpdatedAt = now
		meta.Size = int64(len(data))
	} else {
		s.metadata[id] = &StorageMetadata{
			ID:        id,
			Size:      int64(len(data)),
			CreatedAt: now,
			UpdatedAt: now,
		}
	}

	return nil
}

func (s *MemoryStorage) Delete(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.schemas, id)
	delete(s.versions, id)
	delete(s.metadata, id)

	return nil
}

func (s *MemoryStorage) Exists(ctx context.Context, id string) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.schemas[id]
	return ok, nil
}

func (s *MemoryStorage) List(ctx context.Context, filter *StorageFilter) ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ids := make([]string, 0, len(s.schemas))
	for id := range s.schemas {
		// Apply filters if needed
		ids = append(ids, id)
	}

	// Apply limit/offset
	if filter != nil {
		if filter.Offset < len(ids) {
			ids = ids[filter.Offset:]
		}
		if filter.Limit > 0 && filter.Limit < len(ids) {
			ids = ids[:filter.Limit]
		}
	}

	return ids, nil
}

func (s *MemoryStorage) GetVersion(ctx context.Context, id, version string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	versions, ok := s.versions[id]
	if !ok {
		return nil, fmt.Errorf("no versions for schema: %s", id)
	}

	data, ok := versions[version]
	if !ok {
		return nil, fmt.Errorf("version not found: %s", version)
	}

	return data, nil
}

func (s *MemoryStorage) ListVersions(ctx context.Context, id string) ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	versions, ok := s.versions[id]
	if !ok {
		return []string{}, nil
	}

	list := make([]string, 0, len(versions))
	for version := range versions {
		list = append(list, version)
	}

	return list, nil
}

func (s *MemoryStorage) GetBatch(ctx context.Context, ids []string) (map[string][]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(map[string][]byte)
	for _, id := range ids {
		if data, ok := s.schemas[id]; ok {
			result[id] = data
		}
	}

	return result, nil
}

func (s *MemoryStorage) SetBatch(ctx context.Context, items map[string][]byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	for id, data := range items {
		s.schemas[id] = data
		s.metadata[id] = &StorageMetadata{
			ID:        id,
			Size:      int64(len(data)),
			CreatedAt: now,
			UpdatedAt: now,
		}
	}

	return nil
}

func (s *MemoryStorage) GetMetadata(ctx context.Context, id string) (*StorageMetadata, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	meta, ok := s.metadata[id]
	if !ok {
		return nil, fmt.Errorf("metadata not found: %s", id)
	}

	return meta, nil
}

func (s *MemoryStorage) Health(ctx context.Context) error {
	return nil
}
