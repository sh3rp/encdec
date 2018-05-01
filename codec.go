package encdec

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"net/url"
)

type Codec interface {
	Decode([]byte) ([]byte, error)
	Encode([]byte) ([]byte, error)
}

var Codecs map[string]Codec

func init() {
	Codecs = make(map[string]Codec)
	Codecs["b64"] = b64{}
	Codecs["b32"] = b32{}
	Codecs["url"] = urlescape{}
	Codecs["hex"] = hexcode{}
}

type b64 struct{}

func (b b64) Decode(data []byte) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(data))
	return []byte(decoded), err
}

func (b b64) Encode(data []byte) ([]byte, error) {
	encoded := base64.StdEncoding.EncodeToString(data)
	return []byte(encoded), nil
}

type b32 struct{}

func (b b32) Decode(data []byte) ([]byte, error) {
	decoded, err := base32.StdEncoding.DecodeString(string(data))
	return []byte(decoded), err
}

func (b b32) Encode(data []byte) ([]byte, error) {
	encoded := base32.StdEncoding.EncodeToString(data)
	return []byte(encoded), nil
}

type urlescape struct{}

func (u urlescape) Decode(data []byte) ([]byte, error) {
	decoded, err := url.PathUnescape(string(data))
	return []byte(decoded), err
}

func (u urlescape) Encode(data []byte) ([]byte, error) {
	encoded := url.PathEscape(string(data))
	return []byte(encoded), nil
}

type hexcode struct{}

func (h hexcode) Decode(data []byte) ([]byte, error) {
	decoded, err := hex.DecodeString(string(data))
	return []byte(decoded), err
}

func (h hexcode) Encode(data []byte) ([]byte, error) {
	encoded := hex.EncodeToString(data)
	return []byte(encoded), nil
}

type epoch struct{}
