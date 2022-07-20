package hash

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"

	"golang.org/x/crypto/sha3"
)

func SHAFile(shaType string, encoding string , fileName string ) string {
	f, err := os.Open(fileName)
	if err != nil {
	  fmt.Printf("Unable to Open file %s %s \n", fileName, err )
	}
	defer f.Close()
  return doSha(shaType, encoding, f)
}

func SHA3File(shaType string, encoding string , fileName string ) string {
	f, err := os.Open(fileName)
	if err != nil {
	  fmt.Printf("Unable to Open file %s %s \n", fileName, err )
	}
	defer f.Close()
  return doSha3(shaType, encoding, f)
}
func SHA(shaType string, encoding string, input []byte) string {
	reader :=  bytes.NewReader(input)
	return doSha(shaType, encoding, reader)
}

func doSha(shaType string , encoding string, reader io.Reader) string {

	var shaResult string
	switch shaType {
	case "1":
		shaResult = sumHash(sha1.New(), encoding, reader)
	case "256":
		shaResult = sumHash(sha256.New(), encoding, reader)
	default:
		shaResult = sumHash(sha512.New(), encoding, reader)
		break
	}
	return shaResult
}

func SHA3(shaType string, encoding string, input []byte) string {
   reader := bytes.NewReader(input)
   return doSha3(shaType, encoding, reader)
 }

func doSha3(shaType string ,encoding string, reader io.Reader) string {

	var sha3Result string
	switch shaType {
	case "224":
		sha3Result = sumHash(sha3.New224(), encoding, reader)
		break
	case "256":
		sha3Result = sumHash(sha3.New256(), encoding, reader)
		break
	case "384":
		sha3Result = sumHash(sha3.New384(), encoding, reader)
		break

	default:
		sha3Result = sumHash(sha3.New512(), encoding, reader)
	}

	return sha3Result

}

func sumHash(digester hash.Hash, encoding string, input io.Reader) string {
	io.Copy(digester, input)
	sumResult := digester.Sum(nil)
	var sumStringResult string
	switch encoding {
	case "hex":
		sumStringResult = hex.EncodeToString(sumResult)
	case "base32":
		sumStringResult = base32.StdEncoding.EncodeToString(sumResult)

	case "base64url":
		sumStringResult = base64.URLEncoding.EncodeToString(sumResult)
	default:
		sumStringResult = base64.StdEncoding.EncodeToString(sumResult)
	}
	return sumStringResult
}
