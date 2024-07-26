package config

import (
	"coredx/utils"
	"crypto/rand"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/kardianos/osext"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"github.com/wumansgy/goEncrypt/rsa"
	"gopkg.in/yaml.v2"
	// "honnef.co/go/tools/config"
	// "honnef.co/go/tools/config"
	// "honnef.co/go/tools/config"
)

var Config *config

func Init(cfg, privateFile, publicFile string) {
	dir, err := osext.ExecutableFolder()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("path.Join(dir, cfg): %v\n", filepath.Join(dir, cfg))
	_, err = os.Stat(filepath.Join(dir, cfg))
	if err != nil {
		fmt.Println("not find config.yaml, create default config.yaml")
		err = utils.CreateAppendFile(filepath.Join(dir, cfg), []byte(templateConfig))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(1)
	}
	var privateErr error
	var publicErr error

	//获取公私密钥算法

	conf, err := os.ReadFile(filepath.Join(dir, cfg))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(conf, &Config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, privateErr = os.Stat(filepath.Join(dir, privateFile))
	_, publicErr = os.Stat(filepath.Join(dir, publicFile))
	//公私钥都存在
	if privateErr == nil && publicErr == nil {
		switch Config.Node.NodeAlgo {
		case "rsa":
			private, err := os.ReadFile(path.Join(dir, privateFile))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			public, err := os.ReadFile(filepath.Join(dir, publicFile))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			utils.Islegal(string(public), string(private))
			publicMd5 := utils.MD5(public)
			privateMd5 := utils.MD5(private)
			Config.Node.NodePublicMd5 = publicMd5
			Config.Node.NodePrivateMd5 = privateMd5
			return
		case "sm2":
			private, err := os.ReadFile(path.Join(dir, privateFile))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			public, err := os.ReadFile(filepath.Join(dir, publicFile))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			sm2PrivateKey, err := x509.ReadPrivateKeyFromPem(private, []byte(Config.Node.NodeStoreKey))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			sm2PublicKey, err := x509.ReadPublicKeyFromPem(public)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			utils.Sm2IsLegal(sm2PrivateKey, sm2PublicKey)
			publicMd5 := utils.MD5(public)
			privateMd5 := utils.MD5(private)
			Config.Node.NodePublicMd5 = publicMd5
			Config.Node.NodePrivateMd5 = privateMd5
		}

	}
	if os.IsExist(privateErr) {
		err = os.Remove(filepath.Join(dir, privateFile))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	if os.IsExist(publicErr) {
		err = os.Remove(filepath.Join(dir, publicFile))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	// fmt.Println(Config.Node.NodeAlgo)
	switch Config.Node.NodeAlgo {
	case "rsa":
		keyPair, err := rsa.GenerateRsaKeyBase64(2048)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		utils.Islegal(keyPair.PublicKey, keyPair.PrivateKey)
		err = utils.CreateAppendFile(filepath.Join(dir, privateFile), []byte(keyPair.PrivateKey))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = utils.CreateAppendFile(filepath.Join(dir, publicFile), []byte(keyPair.PublicKey))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "sm2":
		sm2PrivateKey, err := sm2.GenerateKey(rand.Reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		utils.Sm2IsLegal(sm2PrivateKey, &sm2PrivateKey.PublicKey)
		sm2Private, err := x509.WritePrivateKeyToPem(sm2PrivateKey, []byte(Config.Node.NodeStoreKey))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sm2Public, err := x509.WritePublicKeyToPem(&sm2PrivateKey.PublicKey)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = utils.CreateAppendFile(filepath.Join(dir, privateFile), sm2Private)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = utils.CreateAppendFile(filepath.Join(dir, publicFile), sm2Public)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// fmt.Println(string(sm2Private))
	default:
		fmt.Println("not support other algo")
		os.Exit(1)
	}

}

// _, privateErr = os.Stat(filepath.Join(dir, privateFile))
// _, publicErr = os.Stat(filepath.Join(dir, publicFile))
// //公私钥都存在
// if privateErr == nil && publicErr == nil {
// 	private, err := os.ReadFile(path.Join(dir, privateFile))
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	public, err := os.ReadFile(filepath.Join(dir, publicFile))
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	utils.Islegal(string(public), string(private))
// 	//读取文配置文件

// 	conf, err := os.ReadFile(filepath.Join(dir, cfg))
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(Config)
// 	return
// }

// keyPair, err := rsa.GenerateRsaKeyBase64(2048)
// if err != nil {
// 	fmt.Println(err)
// 	os.Exit(1)
// }
// utils.Islegal(keyPair.PublicKey, keyPair.PrivateKey)
// err = utils.CreateAppendFile(filepath.Join(dir, privateFile), []byte(keyPair.PrivateKey))
// if err != nil {
// 	fmt.Println(err)
// 	os.Exit(1)
// }
// err = utils.CreateAppendFile(filepath.Join(dir, publicFile), []byte(keyPair.PublicKey))
// if err != nil {
// 	fmt.Println(err)
// 	os.Exit(1)
// }
// conf, err = os.ReadFile(filepath.Join(dir, cfg))
// if err != nil {
// 	fmt.Println(err)
// 	os.Exit(1)
// }

// // var config Config
// err = yaml.Unmarshal(conf, Config)
// if err != nil {
// 	fmt.Println(err)
// 	os.Exit(1)
// }

// }

const templateConfig = `
db:
  host: "127.0.0.1"
  name: "te"
  port: 3306
  user: "root"
  password: "root"

jwt:
  exp: 90
  iss: "cored"
  sub: "cored"
  secret: "cored"

cache:
  expire: 3
  use_ssl: false
  access_id: "minio_admin"
  end_point: "10.10.1.40:9000"
  secret_accessKey: "9ijnBHU*@123"

store:
  expire: 3
  use_ssl: false
  access_id: "minio_admin"
  end_point: "10.10.1.40:9000"
  default_size: 500
  secret_accessKey: "9ijnBHU*@123"

backup:
  addr: "http://10.10.254.11:5001"

logging:
  file: "logs/cored.log"
  level: "debug"
  max_age: 30
  max_size: 10
  max_backups: 5
ws:
  listen_address:
dr:
  del-data: "delData"
  encrypt: "encrypt"
  endpoint: "https://sqwalletapi.smartdatachain.com"
  upload-data: "uploadData"
  dr-user-create: "drUserCreate"
  dr-user-modify: "drUserModify"
  dr-account-info: "drAccountInfo"
  dr-account-sync: "drAccountSync"
  get-dr-user-info: "getDrUserInfo"
  authentication: "authentication"
  select-upload-data: "selectUploadData"

api:
  key-file: "x"
  cert_file: ''
  enable-tls: false
  listen-address: 0.0.0.0:8080

water:
  grid: "security/grid"
  first: "security/l1"
  plain: "security/plain"
  third: "security/l3"
  protocol: "http"
  end_point: "123.60.56.112:9001"
  generate-stamps: "stamp/circle/create"
  stamp-circle-extract: "stamp/circle/extract"
  stamp-elipse-extract: "stamp/elipse/extract"
  generate-stamp-elipse: "stamp/elipse/create"
  generate-bs-water-mark: "screenBS"
  weak-water-mark-insert: "weak"
  weak-water-mark-extract: "weak/extract"
  electronic-watermark-bright: "flow/plain"
  electronic-watermark-underflow: "flow"
  electronic-watermark-underflow-extract: "flow/extract"

manage:
  listen-addr: "0.0.0.0"
  listen-port: 6000

node:
  node-number: "请填写节点编号"
  node-unit: "请填写节点单位"
  node-deploy: "请填写节点部署地址"
  node-oper: "请填写节点运营商"
  node-login-code: "请填写节点登录码"
  node-public-md5: ""
  node-private-md5: ""
  node-algo: "rsa"
  node-store-key: "1111111111111111"
  node-platform: "ws://0.0.0.0:8083/api/v1/1"
`
