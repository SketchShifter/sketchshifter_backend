package utils

import (
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileUtils ファイル操作に関するインターフェース
type FileUtils interface {
	SaveFile(src io.Reader, destPath string) (string, error)
	DeleteFile(path string) error
}

// fileUtils FileUtilsの実装
type fileUtils struct {
	baseURL string
}

// NewFileUtils FileUtilsを作成
func NewFileUtils(baseURL string) FileUtils {
	return &fileUtils{
		baseURL: baseURL,
	}
}

// SaveFile ファイルを保存
func (f *fileUtils) SaveFile(src io.Reader, destPath string) (string, error) {
	// ディレクトリを確認
	dir := filepath.Dir(destPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	// ファイルを作成
	dest, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer dest.Close()

	// ファイルをコピー
	if _, err := io.Copy(dest, src); err != nil {
		return "", err
	}

	// 相対パスに変換
	relPath := strings.ReplaceAll(destPath, "\\", "/")
	if strings.HasPrefix(relPath, "./") {
		relPath = relPath[2:]
	}

	// URLを返す
	url := f.baseURL + "/" + relPath
	return url, nil
}

// DeleteFile ファイルを削除
func (f *fileUtils) DeleteFile(path string) error {
	// 絶対パスに変換
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// ファイルを削除
	return os.Remove(absPath)
}

// GenerateRandomString ランダムな文字列を生成
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数生成器を初期化
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}
