//go:build windows

package chromium

import (
	"encoding/base64"
	"errors"
	"os"

	"github.com/tidwall/gjson"

	"github.com/nowhszh/hack-browser-data/common/decrypter"
	"github.com/nowhszh/hack-browser-data/common/item"
	"github.com/nowhszh/hack-browser-data/common/log"
	"github.com/nowhszh/hack-browser-data/common/utils/fileutil"
)

var errDecodeMasterKeyFailed = errors.New("decode master key failed")

func (c *chromium) GetMasterKey() ([]byte, error) {
	keyFile, err := fileutil.ReadFile(item.TempChromiumKey)
	if err != nil {
		return nil, err
	}
	defer os.Remove(keyFile)
	encryptedKey := gjson.Get(keyFile, "os_crypt.encrypted_key")
	if !encryptedKey.Exists() {
		return nil, nil
	}
	pureKey, err := base64.StdEncoding.DecodeString(encryptedKey.String())
	if err != nil {
		return nil, errDecodeMasterKeyFailed
	}
	c.masterKey, err = decrypter.DPAPI(pureKey[5:])
	log.Infof("%s initialized master key success", c.name)
	return c.masterKey, err
}
