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
