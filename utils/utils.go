package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log/slog"
	"mime/multipart"
)

func GenerateSHA1(file multipart.File) (string, error) {
	defer file.Seek(0, io.SeekStart)
	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		slog.Error("计算文件哈希失败", "err", err)
		return "", err
	}
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}
