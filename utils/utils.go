package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log/slog"
	"mime/multipart"
)

func GenerateSHA1(file *multipart.FileHeader) (string, multipart.File, error) {
	File, err := file.Open()
	if err != nil {
		slog.Error("打开文件失败", "err", err)
		return "", nil, err
	}
	defer File.Seek(0, io.SeekStart)
	hash := sha1.New()
	if _, err = io.Copy(hash, File); err != nil {
		slog.Error("计算文件哈希失败", "err", err)
		return "", nil, err
	}
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes), File, nil
}
