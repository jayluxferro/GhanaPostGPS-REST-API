package ghanapostgps

import (
  "bytes"
  "crypto/aes"
  "crypto/cipher"
  "encoding/base64"
  "io/ioutil"
  mrand "math/rand"
  "net/http"
  "net/url"
  "strings"
  "time"
  "fmt"
)

const (
	ALLOWED_CHARACTERS = "0123456789qwertyuiopasdfghjklzxcvbnm!@$#^&*()"
)

type Params struct {
	ApiKey         string
	UUID           string
	ApiURL         string
	AsaaseAPI      string
	Language       string
	LanguageCode   string
	AndroidCert    string
	AndroidPackage string
	Country        string
	CountryName    string
}

func APIRequest(method string, params *Params, payload *strings.Reader) string {
	client := &http.Client{}
	req, err := http.NewRequest(method, params.ApiURL, payload)

	if err != nil {
		print(err)
	}
	req.Header.Add("Language", params.Language)
	req.Header.Add("X-Android-Cert", params.AndroidCert)
	req.Header.Add("X-Android-Package", params.AndroidPackage)
	req.Header.Add("DeviceID", params.UUID)
	req.Header.Add("LanguageCode", params.LanguageCode)
	req.Header.Add("Country", params.Country)
	req.Header.Add("CountryName", params.CountryName)
	req.Header.Add("AsaaseUser", params.AsaaseAPI)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	return string(body)
}

func RandomString(number int) string {
	allowedCharacters := []byte(ALLOWED_CHARACTERS)
	data := make([]byte, number)
	mrand.Seed(time.Now().UnixNano())
	for i := 0; i < number; i++ {
		data[i] = allowedCharacters[mrand.Intn(len(allowedCharacters))]
	}
	return string(data)
}

func GetLocation(code string, defaults *Params) string {
	params := url.Values{
		"AsaaseLogs": {""},
		"Action":     {"GetLocation"},
		"GPSName":    {code},
	}
	dataRequest := GetDataRequest(&params, defaults)
	return GPGPSDecrypt(APIRequest("POST", defaults, dataRequest), defaults)
}

func GetDataRequest(v *url.Values, defaults *Params) *strings.Reader {
	data := GPGPSEncrypt(v.Encode(), defaults)
	params := url.Values{
		"DataRequest": {data},
	}
	return strings.NewReader(params.Encode() + "&")
}

func GPGPSEncrypt(data string, params *Params) string {
	decryptionKey := []byte(params.ApiKey)
	iv := []byte(RandomString(len(decryptionKey)))
	encryptedMsg := AESEncrypt(iv, decryptionKey, data)
	payload := []byte{}
	payload = append(payload, iv...)
	payload = append(payload, encryptedMsg...)
	// encode payload to base64

	return base64.StdEncoding.EncodeToString(payload)
}

func GPGPSDecrypt(encodedData string, params *Params) string {
	decodedData, _ := base64.StdEncoding.DecodeString(encodedData)
	// remove IV (16 byte)
	decryptionKey := []byte(params.ApiKey)
	iv := decodedData[:len(decryptionKey)]
	msg := decodedData[len(decryptionKey):]

	return string(AESDecrypt(iv, decryptionKey, msg))
}

func print(data ...interface{}) (n int, err error) {
	return fmt.Println(data...)
}

func AESEncrypt(iv []byte, key []byte, src string) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		print("key error1", err)
	}

	if src == "" {
		print("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, iv)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func AESDecrypt(iv []byte, key []byte, crypt []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		print("key error1", err)
	}
	if len(crypt) == 0 {
		print("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return PKCS5Trimming(decrypted)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func IsValidGPAddress(address string) (bool, string){
  isValid := true
  address = strings.Join(strings.Split(strings.ToUpper(strings.Trim(address, "")), "-"), "")

  if len(address) < 9 {
    isValid = false
  }

  return isValid, address
}
