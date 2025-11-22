package schema

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// FileStorage provides file-based storage
type FileStorage struct {
	basePath string
	mu       sync.RWMutex
}

// NewFileStorage creates file storage
func NewFileStorage(basePath string) (*FileStorage, error) {
	// Create directory if not exists
	if err := os.MkdirAll(basePath, 0o755); err != nil {
		return nil, fmt.Errorf("create base directory: %w", err)
	}

	return &FileStorage{
		basePath: basePath,
	}, nil
}

func (s *FileStorage) Get(ctx context.Context, id string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	path := s.getSchemaPath(id)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("schema not found: %s", id)
		}
		return nil, fmt.Errorf("read file: %w", err)
	}

	return data, nil
}

func (s *FileStorage) Set(ctx context.Context, id string, data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	path := s.getSchemaPath(id)

	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	// Write file
	if err := ioutil.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

func (s *FileStorage) Delete(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	path := s.getSchemaPath(id)
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("remove file: %w", err)
	}

	return nil
}

func (s *FileStorage) Exists(ctx context.Context, id string) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	path := s.getSchemaPath(id)
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (s *FileStorage) List(ctx context.Context, filter *StorageFilter) ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var ids []string
	err := filepath.Walk(s.basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".json") {
			// Extract ID from path
			rel, _ := filepath.Rel(s.basePath, path)
			id := strings.TrimSuffix(rel, ".json")
			ids = append(ids, id)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk directory: %w", err)
	}

	return ids, nil
}

func (s *FileStorage) GetVersion(ctx context.Context, id, version string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	path := s.getVersionPath(id, version)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read version: %w", err)
	}

	return data, nil
}

func (s *FileStorage) ListVersions(ctx context.Context, id string) ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	versionDir := filepath.Join(s.basePath, ".versions", id)
	files, err := ioutil.ReadDir(versionDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, fmt.Errorf("read version directory: %w", err)
	}

	versions := make([]string, 0, len(files))
	for _, file := range files {
		if !file.IsDir() {
			version := strings.TrimSuffix(file.Name(), ".json")
			versions = append(versions, version)
		}
	}

	return versions, nil
}

func (s *FileStorage) GetBatch(ctx context.Context, ids []string) (map[string][]byte, error) {
	result := make(map[string][]byte)

	for _, id := range ids {
		data, err := s.Get(ctx, id)
		if err == nil {
			result[id] = data
		}
	}

	return result, nil
}

func (s *FileStorage) SetBatch(ctx context.Context, items map[string][]byte) error {
	for id, data := range items {
		if err := s.Set(ctx, id, data); err != nil {
			return fmt.Errorf("set %s: %w", id, err)
		}
	}

	return nil
}

func (s *FileStorage) GetMetadata(ctx context.Context, id string) (*StorageMetadata, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	path := s.getSchemaPath(id)
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("stat file: %w", err)
	}

	return &StorageMetadata{
		ID:        id,
		Size:      info.Size(),
		UpdatedAt: info.ModTime(),
	}, nil
}

func (s *FileStorage) Health(ctx context.Context) error {
	// Check if base path is accessible
	if _, err := os.Stat(s.basePath); err != nil {
		return fmt.Errorf("base path inaccessible: %w", err)
	}
	return nil
}

func (s *FileStorage) getSchemaPath(id string) string {
	return filepath.Join(s.basePath, id+".json")
}

func (s *FileStorage) getVersionPath(id, version string) string {
	return filepath.Join(s.basePath, ".versions", id, version+".json")
}
