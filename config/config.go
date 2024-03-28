package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

type ServerJson struct {
	Mysql  MysqlConfig  `json:"mysql"`
	Server ServerConfig `json:"server"`
}

func initPath() {
	sep := string(os.PathSeparator)
	ExecPath, _ = os.Executable()
	baseName := filepath.Base(ExecPath)
	ExecPath = filepath.Dir(ExecPath)
	if strings.Contains(baseName, "go_build") || strings.Contains(ExecPath, "go-build") || strings.Contains(baseName, "Test") {
		ExecPath, _ = os.Getwd()
	}
	if strings.Contains(baseName, "Test") {
		ExecPath = filepath.Dir(ExecPath)
	}
	length := utf8.RuneCountInString(ExecPath)
	lastChar := ExecPath[length-1:]
	if lastChar != sep {
		ExecPath = ExecPath + sep
	}
	ExecPath = strings.ReplaceAll(ExecPath, "\\", "/")
	log.Println(ExecPath)
}

func initJSON() {
	rawConfig, err := os.ReadFile(fmt.Sprintf("%sconfig.json", ExecPath))
	if err != nil {
		//未初始化
		tokenSecret := GenerateRandString(32)
		Server.Server.Env = "production"
		Server.Server.Port = 4001
		Server.Server.LogLevel = "error"
		Server.Server.TokenSecret = tokenSecret
	} else {
		if err = json.Unmarshal(rawConfig, &Server); err != nil {
			fmt.Println("Invalid Config: ", err.Error())
			logFile, err := os.OpenFile(ExecPath+"error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
			if nil == err {
				logFile.WriteString(fmt.Sprintln(time.Now().Format("2006-01-02 15:04:05"), "Invalid Config", err.Error()))
				defer logFile.Close()
			}
			os.Exit(-1)
		}
	}
}

var ExecPath string
var Server ServerJson
// var AnqiUser AnqiUserConfig

var GoogleValid bool // can visit google or not

// RestartChan 1 to restart app, 0 to reload template, 2 to exit app
var RestartChan = make(chan int)
var Languages = map[string]map[string]string{}

func init() {
	fmt.Printf("开始初始化")
	initPath()
	initJSON()
	initLanguage()
}

func WriteConfig() error {
	//将现有配置写回文件
	configFile, err := os.OpenFile(fmt.Sprintf("%sconfig.json", ExecPath), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer configFile.Close()

	buff := &bytes.Buffer{}

	buf, err := json.MarshalIndent(Server, "", "\t")
	if err != nil {
		return err
	}
	buff.Write(buf)

	_, err = io.Copy(configFile, buff)
	if err != nil {
		return err
	}

	return nil
}

func initLanguage() {
	// 重置
	// 读取language列表
	readerInfos, err := os.ReadDir(fmt.Sprintf("%slanguage", ExecPath))
	if err == nil {
		for _, info := range readerInfos {
			if strings.HasSuffix(info.Name(), ".yml") {
				lang := strings.TrimSuffix(info.Name(), ".yml")
				languagePath := ExecPath + "language/" + info.Name()
				languages := map[string]string{}
				yamlFile, err := os.ReadFile(languagePath)
				if err == nil {
					strSlice := strings.Split(strings.ReplaceAll(strings.ReplaceAll(string(yamlFile), "\r\n", "\n"), "\r", "\n"), "\n")
					for _, v := range strSlice {
						vSplit := strings.SplitN(strings.TrimSpace(v), ":", 2)
						if len(vSplit) == 2 {
							languages[strings.Trim(vSplit[0], "\" ")] = strings.Trim(vSplit[1], "\" ")
						}
					}
				}
				Languages[lang] = languages
			}
		}
	}
}

func GenerateRandString(length int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		b := r.Intn(26) + 65
		buf[i] = byte(b)
	}
	return strings.ToLower(string(buf))
}

// func main() {
// 	// init()
// 	// fmt.Printf(json.Marshal(Server))
// 	data, _ := json.Marshal(Server)
//     fmt.Printf("%s\n", data)
// }