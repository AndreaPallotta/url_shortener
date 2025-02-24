package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"github.com/itchyny/base58-go"
)

func generateHash(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func encodeBase58(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortLink(ogUrl string, userId string) string {
	hashBytes := generateHash(ogUrl + userId)
	randomNum := new(big.Int).SetBytes(hashBytes).Uint64()

	shortLink := encodeBase58([]byte(fmt.Sprintf("%d", randomNum)))
	return shortLink[:8]
}