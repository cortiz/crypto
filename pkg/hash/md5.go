package hash

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"os"
)


func MD5(encoding string , input []byte) string {

 return sumHash(md5.New(),encoding, bytes.NewReader(input))
}

func MD5File(encoding string , fileName string) string { 
	f, err := os.Open(fileName)
	if err != nil {
	  fmt.Printf("Unable to Open file %s %s \n", fileName, err )
	}
	defer f.Close()
  return sumHash(md5.New(), encoding ,f)

}
