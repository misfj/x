package exe

import (
	"fmt"
	"os"
	"os/exec"
)

// var mapExeList *MapExeList

// type SinglgExe struct {
// 	Protocol         string
// 	IP               string
// 	Port             uint
// 	StopFunc         func(StopMsg []byte) error
// 	RestartFunc      func(RestartMsg []byte) error
// 	UpdateConfigFunc func(UpdateConfigMsg []byte) error
// 	GetConfigFunc    func(GetConfigMsg []byte) error
// }

// type MapExeList struct {
// 	mutex sync.Mutex
// 	Exes  map[string]*SinglgExe
// }

// func NewMapExeList() {
// 	mapExeList = &MapExeList{
// 		Exes: make(map[string]*SinglgExe, 0),
// 	}
// }
// func (MapExeList *MapExeList) Add(programName string, exe *SinglgExe) {
// 	mapExeList.mutex.Lock()
// 	defer mapExeList.mutex.Unlock()
// 	mapExeList.Exes[programName] = exe
// }

// func NewSingleExe(protocol string, ip string, port uint) *SinglgExe {
// 	return &SinglgExe{
// 		Protocol: protocol,
// 		IP:       ip,
// 		Port:     port,
// 		StopFunc: func(killMsg []byte) error {
// 			// TODO: 实现kill函数
// 			udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", protocol, port))
// 			if err != nil {
// 				return err
// 			}
// 			conn, err := net.DialUDP(protocol, nil, udpAddr)
// 			if err != nil {
// 				return err
// 			}
// 			defer conn.Close()
// 		},
// 	}
// }

// const WaterKillMsg = `{
// 	 "Command": "serviceControl",
// 	 "ID": 0,
// 	 "Payload": {
// 	   "Action": 1,
// 	   "Config": {
// 	     "PlatConfig": {
// 	       "watermark-service": {
// 	         "ServerAddr": "123.60.55.112:9001"
// 	       },
// 	       "ai-service": {
// 	         "ServerAddr": "123.60.55.112:9001"
// 	       }
// 	     },
// 	     "MinioConfig": {
// 	       "MinioEndpoint": "10.10.1.41:9000",
// 	       "MinioAccessKey": "LYVbwGtWft1T4Bdbjqq8",
// 	       "MinioSecretKey": "nAsq2o9Bc4EhSgVvLItIoEJDWTqS2zfOudZUjACi"
// 	     },
// 	     "DbConfig": {
// 	       "Host": "10.10.1.41",
// 	       "Port": 3307,
// 	       "Password": "9ijnBHU*@123",
// 	       "DbName": "nwyl",
// 	       "User": "root"
// 	     }
// 	   }
// 	 }
// 	}`

// func Kill(programType string) {
// 	switch programType {
// 	case "water":

// 	}
// }
// func water_kill()
// func Exec(exePath string, args []string, path string) error {
// 	cmd := "a.exe"
// 	args := []string{":9090"}

// }
func Init(exeName string, args []string) {
	waterlog, err := os.Create("water.log")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func() {
		cmd := exec.Command(exeName, args...)
		cmd.Stdout = os.Stdout // 标准输出

		// cmd.Stderr = os.Stderr // 标准错误
		cmd.Stderr = waterlog
		err := cmd.Run()
		if err != nil {
			fmt.Printf("命令执行失败: %v\n", err)
		} else {
			fmt.Println("命令执行成功")
		}
	}()

}
