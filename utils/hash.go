package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/wumansgy/goEncrypt/rsa"
)

func MD5(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

// 检查当前目录下的公私钥是否合法
func Islegal(public string, private string) {
	msg := "cored is the core service of super nodes and cloud platforms"
	//验证加密
	cipherText, err := rsa.RsaEncryptToBase64([]byte(msg), public)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	plainText, err := rsa.RsaDecryptByBase64(cipherText, private)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if !strings.EqualFold(msg, string(plainText)) {
		fmt.Println("decrypt error")
		os.Exit(1)
	}
	//验证签名
	sign, err := rsa.RsaSignBase64([]byte(msg), private)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b := rsa.RsaVerifySignBase64([]byte(msg), sign, public)
	if !b {
		fmt.Println("verify error")
		os.Exit(1)
	}
}

func CheckAndCreateFiles(publicPath, privatePath string) {

	// 检查a.txt是否存在
	_, errA := os.Stat(publicPath)
	aExists := !os.IsNotExist(errA)

	// 检查b.txt是否存在
	_, errB := os.Stat(privatePath)
	bExists := !os.IsNotExist(errB)

	// 如果a.txt和b.txt都存在，则不创建
	if aExists && bExists {
		publicKey, _ := os.ReadFile(publicPath)
		privateKey, _ := os.ReadFile(privatePath)
		Islegal(string(publicKey), string(privateKey))
		return
	}

	// 删除存在的文件
	if aExists {
		os.Remove(publicPath)
	}
	if bExists {
		os.Remove(privatePath)
	}
	pair, err := rsa.GenerateRsaKeyBase64(2048)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 重新创建a.txt和b.txt
	createFile(publicPath, pair.PublicKey)
	createFile(privatePath, pair.PrivateKey)
	Islegal(string(pair.PublicKey), string(pair.PrivateKey))
}

// createFile 创建文件并写入内容
func createFile(filePath, content string) {
	// 将内容写入文件
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func ExtractCmdValue(jsonStr string) (string, error) {
	// 创建一个空的map来存储解析后的JSON数据
	var data map[string]interface{}
	// 解析JSON字符串
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", err
	}
	// 获取Command字段的值
	cmdValue, ok := data["command"].(string)
	if !ok {
		//fmt.Printf("the command value is not string : %s", err.Error())
		return "", fmt.Errorf("cmd value not string")
	}
	return cmdValue, nil
}
func TimeParse() time.Time {
	timeStr := "2006-01-02 15:04:05.888"
	// 格式字符串中，'2006-01-02 15:04:05.999999' 表示：
	// 2006 年，01 月，02 日，15 时，04 分，05 秒，999999 微秒
	layout := "2006-01-02 15:04:05.999999"

	// 使用 time.Parse 将字符串转换为 time.Time 类型
	parsedTime, _ := time.Parse(layout, timeStr)
	return parsedTime
}

// Bcrypt 加密密码
func Bcrypt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// BcryptVerify  验证密码
func BcryptVerify(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return false
	}
	return true
}
