package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"mime/multipart"
)

func SHA1FileHash(file *multipart.FileHeader) string {
	src, _ := file.Open()
	defer src.Close()

	s := sha1.New()
	io.Copy(s, src)
	r := hex.EncodeToString(s.Sum(nil))
	return r
}
