package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	MyTimeInterval = 60 * 10
	MyTimeExpired  = 60 * 60 * 24
)

var (
	// myKey token Key
	mykey = []byte("@#$yymmxxkkyzilm")
)

func timeNow() int {
	var now = time.Now()
	var unix = now.Unix()
	unix = unix / (int64(MyTimeInterval))
	return int(unix)
}

func verifyToken(r *http.Request) (string, bool) {
	var tk = r.Header.Get("tk")

	if tk == "" {
		return "", false
	}

	v, e := ParseTK(tk)
	if e == nil {
		return v, true
	}

	return v, false
}

// GenTK 生成一个加密的token
func GenTK(account string) string {
	var plainTK = fmt.Sprintf("%s@%d", account, time.Now().Unix())
	log.Println("GenTK, plainTK is:", plainTK)
	return encrypt(mykey, plainTK)
}

func ParseTK(token string) (string, error) {
	// log.Printf("ParseTk, tok:%s, len:%d\n", token, len(token))
	if token == "" {
		return "", errors.New("ErrTokenEmpty")
	}

	var plainTK, err = decrypt(mykey, token)
	if err != nil {
		log.Println("ParseTK, err:", err)
		return "", errors.New("ErrTokenDecrypt")
	}

	//log.Println("ParseTK, plainTK is:", plainTK)

	var splits = strings.Split(plainTK, "@")
	if len(splits) != 2 {
		log.Println("ParseTK, err: no @ at text")
		return "", errors.New("ErrTokenFormat")
	}

	timestamp, err := strconv.Atoi(splits[1])
	if err != nil {
		log.Println("ParseTK, err: ", err)
		return "", errors.New("ErrTokenFormat")
	}

	var now = int(time.Now().Unix())
	//log.Printf("ParseTK, account:%s, timestamp:%d, now:%d", splits[0], timestamp, now)

	if now-timestamp > (MyTimeExpired) {
		log.Println("ParseTK, token has been expired")
		return "", errors.New("ErrTokenExpired")
	}
	//if timestamp < servercfg.StartTime {
	//	//log.Println("Server restart,need re-login")
	//	return "", errTokenExpired
	//}

	return splits[0], nil
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) (string, error) {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil
}
