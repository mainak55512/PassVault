package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	rn "math/rand"
	"os"
	"time"
)

type Authentication struct {
	Service_name string `json:"service_name"`
	User_id      string `json:"user_id"`
	Password     string `json:"password"`
}

type Config struct {
	PrivKey   string `json:"Privatekey"`
	PubKey    string `json:"Publickey"`
	Login_pin string `json:"login_pin"`
	Ques1 string `json:"sec_ques_1"`
	Ques2 string `json:"sec_ques_2"`
}

func PassStore(service_name, user_id, password string) {
	secretMessage := fmt.Sprint(password)
	fileContent := handleStorageFile()
	var OAuth []Authentication
	json.Unmarshal(fileContent, &OAuth)
	_, publicKey := Fetch_rsa_key_pair()
	encryptedMessage := RSA_OAEP_Encrypt(secretMessage, *publicKey)
	OAuth = append(OAuth, Authentication{Service_name: service_name, User_id: user_id, Password: encryptedMessage})
	result, e := json.MarshalIndent(OAuth, "", " ")
	if e != nil {
		fmt.Println("error", e)
	}
	os.WriteFile(".data/.storage/Storage.json", result, 0644)
}

func Fetch_cred_details() []string {
	fileContent := handleStorageFile()
	var OAuth []Authentication
	json.Unmarshal(fileContent, &OAuth)
	var options []string
	for _, i := range OAuth {
		options = append(options, i.Service_name)
	}
	return options
}

func Fetch_creds(service_id string) []string {
	fileContent := handleStorageFile()
	var OAuth []Authentication
	json.Unmarshal(fileContent, &OAuth)
	sname, userid, pwd := "Null", "Null", "Null"
	for _, i := range OAuth {
		if i.Service_name == service_id {
			sname = i.Service_name
			userid = i.User_id
			privateKey, _ := Fetch_rsa_key_pair()
			pwd = RSA_OAEP_Decrypt(i.Password, *privateKey)
		}
	}
	values := []string{sname, userid, pwd}
	return values
}

func Del_cred(service_id string) {
	fileContent := handleStorageFile()
	var OAuth []Authentication
	json.Unmarshal(fileContent, &OAuth)
	for i, val := range OAuth {
		if val.Service_name == service_id {
			OAuth = append(OAuth[:i], OAuth[i+1:]...)
			result, e := json.MarshalIndent(OAuth, "", " ")
			if e != nil {
				fmt.Println("error", e)
			}
			os.WriteFile(".data/.storage/Storage.json", result, 0644)
			break
		}
	}
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func RSA_OAEP_Encrypt(secretMessage string, key rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(secretMessage), label)
	CheckError(err)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func RSA_OAEP_Decrypt(cipherText string, privKey rsa.PrivateKey) string {
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &privKey, ct, label)
	CheckError(err)
	return string(plaintext)
}

func Generate_rsa_Key_pair() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	CheckError(err)
	publicKey := &privateKey.PublicKey
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatalln("MarshalPublicKey failed: ", err)
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)
	conf := handleKeyFile()
	conf.PrivKey = string(privkey_pem)
	conf.PubKey = string(pubkey_pem)
	result, e := json.MarshalIndent(conf, "", " ")
	if e != nil {
		fmt.Println("error", err)
	}
	os.WriteFile(".data/.keys/config.json", result, 0644)
}

func Fetch_rsa_key_pair() (*rsa.PrivateKey, *rsa.PublicKey) {
	conf := handleKeyFile()
	block1, _ := pem.Decode([]byte(conf.PrivKey))
	if block1 == nil {
		log.Fatalln("failed to parse PEM block containing the key")
	}
	priv_key, err := x509.ParsePKCS1PrivateKey(block1.Bytes)
	if err != nil {
		log.Fatalln("Can't decode key: ", err)
	}
	block2, _ := pem.Decode([]byte(conf.PubKey))
	if block2 == nil {
		log.Fatalln("failed to parse PEM block containing the key")
	}
	pub, err := x509.ParsePKIXPublicKey(block2.Bytes)
	if err != nil {
		log.Fatalln("failed to parse PEM block containing the key")
	}
	var pub_key *rsa.PublicKey
	switch pub := pub.(type) {
	case *rsa.PublicKey:
		pub_key = pub
	default:
		break
	}
	return priv_key, pub_key
}

func Generate_password(length int, upper_case, numbers, spl_char bool) string {
	rn.Seed(time.Now().UnixNano())
	digits, specials, uppers := "", "", ""
	lowers := "abcdefghijklmnopqrstuvwxyz"
	buf := make([]byte, length)
	if numbers {
		digits = "0123456789"
		buf[0] = digits[rn.Intn(len(digits))]
	} else {
		buf[0] = lowers[rn.Intn(len(lowers))]
	}
	if spl_char {
		specials = "~=+%^*/()[]{}\\!@#$?|"
		buf[1] = specials[rn.Intn(len(specials))]
	} else {
		buf[1] = lowers[rn.Intn(len(lowers))]
	}
	if upper_case {
		uppers = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		buf[2] = uppers[rn.Intn(len(uppers))]
	} else {
		buf[2] = lowers[rn.Intn(len(lowers))]
	}
	all := lowers + uppers + digits + specials
	for i := 3; i < length; i++ {
		buf[i] = all[rn.Intn(len(all))]
	}
	rn.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}

func ValidLogin(s string) bool {
	conf := handleKeyFile()
	result := fmt.Sprintf("%x",sha256.Sum256([]byte(s)))
	return result == conf.Login_pin
}

func Setpin(s,ques1,ques2 string){
	conf := handleKeyFile()
	result := fmt.Sprintf("%x",sha256.Sum256([]byte(s)))
	result1 := fmt.Sprintf("%x",sha256.Sum256([]byte(ques1)))
	result2 := fmt.Sprintf("%x",sha256.Sum256([]byte(ques2)))
	conf.Login_pin = result
	conf.Ques1 = result1
	conf.Ques2 = result2
	res, e := json.MarshalIndent(conf, "", " ")
	if e != nil {
		fmt.Println("error", e)
	}
	os.WriteFile(".data/.keys/config.json", res, 0644)
}

func Changepin(s string){
	conf := handleKeyFile()
	result := fmt.Sprintf("%x",sha256.Sum256([]byte(s)))
	conf.Login_pin = result
	res, e := json.MarshalIndent(conf, "", " ")
	if e != nil {
		fmt.Println("error", e)
	}
	os.WriteFile(".data/.keys/config.json", res, 0644)
}

func CheckConfig() bool {
	conf := handleKeyFile()
	return len(conf.Login_pin)==0 && len(conf.Ques1)==0 && len(conf.Ques2)==0
}

func checkScQues(s1, s2 string) bool {
	conf := handleKeyFile()
	result1 := fmt.Sprintf("%x",sha256.Sum256([]byte(s1)))
	result2 := fmt.Sprintf("%x",sha256.Sum256([]byte(s2)))

	return result1 == conf.Ques1 && result2 == conf.Ques2
}

func handleKeyFile() Config {
	file, err := os.Open(".data/.keys/config.json")
	if err != nil {
		log.Fatalln("error opening config file")
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("error reading config file")
	}
	var conf Config
	json.Unmarshal(data, &conf)
	return conf
}

func handleStorageFile() []byte {
	fileContent, err := os.ReadFile(".data/.storage/Storage.json")
	if err != nil {
		log.Fatal(err)
	}
	return fileContent
}
