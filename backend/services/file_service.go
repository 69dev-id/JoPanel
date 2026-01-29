package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type FileService interface {
	ListFiles(basePath, requestPath string) ([]FileInfo, error)
	ReadFile(basePath, requestPath string) (string, error)
	UploadFile(basePath, requestPath string, file *multipart.FileHeader) error
	CreateDirectory(basePath, requestPath string) error
	DeletePath(basePath, requestPath string) error
	MovePath(basePath, srcPath, destPath string) error
}

type fileService struct{}

func NewFileService() FileService {
	return &fileService{}
}

type FileInfo struct {
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	IsDir   bool   `json:"is_dir"`
	ModTime string `json:"mod_time"`
}

// SecureJoin prevents path traversal
func SecureJoin(basePath, requestPath string) (string, error) {
	fullPath := filepath.Join(basePath, requestPath)
	if !strings.HasPrefix(fullPath, filepath.Clean(basePath)) {
		return "", fmt.Errorf("invalid path: path traversal attempt")
	}
	return fullPath, nil
}

func (s *fileService) ListFiles(basePath, requestPath string) ([]FileInfo, error) {
	fullPath, err := SecureJoin(basePath, requestPath)
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, err
	}

	var files []FileInfo
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		files = append(files, FileInfo{
			Name:    entry.Name(),
			Size:    info.Size(),
			IsDir:   entry.IsDir(),
			ModTime: info.ModTime().Format("2006-01-02 15:04:05"),
		})
	}
	return files, nil
}

func (s *fileService) ReadFile(basePath, requestPath string) (string, error) {
	fullPath, err := SecureJoin(basePath, requestPath)
	if err != nil {
		return "", err
	}
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (s *fileService) UploadFile(basePath, requestPath string, file *multipart.FileHeader) error {
	fullPath, err := SecureJoin(basePath, requestPath)
	if err != nil {
		return err
	}
	
	// Create destination (join with filename if path is dir, or assumes full path provided? Usually upload is to a Dir)
	// Strategy: requestPath is the DIR.
	destPath := filepath.Join(fullPath, file.Filename)
	
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func (s *fileService) CreateDirectory(basePath, requestPath string) error {
	fullPath, err := SecureJoin(basePath, requestPath)
	if err != nil {
		return err
	}
	return os.MkdirAll(fullPath, 0755)
}

func (s *fileService) DeletePath(basePath, requestPath string) error {
	fullPath, err := SecureJoin(basePath, requestPath)
	if err != nil {
		return err
	}
	return os.RemoveAll(fullPath)
}

func (s *fileService) MovePath(basePath, srcPath, destPath string) error {
	fullSrc, err := SecureJoin(basePath, srcPath)
	if err != nil {
		return err
	}
	fullDest, err := SecureJoin(basePath, destPath)
	if err != nil {
		return err
	}
	return os.Rename(fullSrc, fullDest)
}
