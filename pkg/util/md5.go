package util

import (
	"crypto/md5"
	"encoding/hex"
)

//该方法是对上传后的文件名进行格式化，简单来说，就是对文件名用 MD5 加密后再进行写入，避免直接暴露原始名称。

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
