package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func main() {
	homedir := UserHomeDir()
	homedir = homedir + "/.pwp"
	fmt.Println("Home Directory is: ", homedir)
	if _, err := os.Stat(homedir); err == nil {
		//home directory exist, check for db
	} else {
		//home dir doesn't exist, create
		fmt.Println("Creating Directory")
	}
	key := []byte("yellow submarineyellow submarine") // any 128-, 192-, or 256-bit key
	//fmt.Printf("key size: %d\n", len(key))
	b, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//initialize GCM
	gcm, err := cipher.NewGCM(b)
	if err != nil {
		panic(err)
	}

	//read plaintext from sql template file
	plaintext, err := ioutil.ReadFile("sql-template.sql")
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	if ioutil.WriteFile("pwp.enc", ciphertext, 0600) != nil {
		panic(err)
	}

	ciphertext, err = ioutil.ReadFile("pwp.enc")
	if err != nil {
		panic(err)
	}
	//decrypt
	v, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(v))

}
