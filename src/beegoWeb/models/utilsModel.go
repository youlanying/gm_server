package models

import (
	"container/list"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gm_server/src/cfg"
	"gm_server/src/logger"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

//---- 一些其他项目可以用到的公共方法放在这里 ----

func HttpGet(url string) ([]byte, error) {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		println("HttpGet Error↓")
		println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		println("HttpGet Error resp.Body")
		return nil, err
	}

	println("HttpGet Success↓")
	var result = string(body)
	fmt.Println(result)
	return body, nil
}

func GetAuth(GmName string, SendTime string) string {
	GmKey := GetConfig("gmkey")
	GmKeyStr := Md5V(GmKey)
	StrList := []string{GmKeyStr, GmName, SendTime}
	StrCode := strings.Join(StrList, "|")
	Md5Code := Md5V(StrCode)
	return Md5Code
}

//interface ToJson
func ToJson(v interface{}) string {
	byteJson, errs := json.Marshal(v) //转换成JSON返回的是byte[]
	if errs != nil {
		logger.LogErr(" SaveJson ERROR", errs.Error())
		return ""
	} else {
		strJson := string(byteJson)
		return strJson
	}
}

//func Proc_make_auth(RoleId, MessageId)string{
//
//	NowTime := util_time:now_to_s(now())
//	NewMd5 := proc_make_md5(RoleId, MessageId, NowTime)
//	return message_pb:make_http_auth(RoleId, NewMd5, NowTime)
//}

func Proc_make_md5(RoleId int, MsgId int, NowTime int) string {
	Mac := "mac"
	StrMsgId := strconv.Itoa(MsgId)
	StrNowTime := strconv.Itoa(NowTime)
	StrMd5 := StrMsgId + "#" + strconv.Itoa(RoleId) + "#" + Mac + "#" + StrNowTime + "#" + strconv.Itoa(0) + "f6e5e27b278431fa343acadea46ad4e5"
	return Md5V(StrMd5)
}

/**
 * 获取当前 格林威治时间 秒
 */
func GetNow() int64 {
	now := time.Now().Unix() //获取时间戳
	return now
}

func GetNowNigxStr() string {
	now := time.Now().Unix() //获取时间戳
	str := strconv.FormatInt(now, 10)
	return str
}

func GetNowStr() string {
	str := TimeToString(time.Now())
	return str
}

func GetNowDateTime() time.Time {
	return time.Now()
}

func TimeToString(t time.Time) string {
	//2006-01-02 是 go诞生的时间
	return t.Format("2006-01-02 15:04:05")
}

func UnixToString(unixNum int) string {
	//2006-01-02 是 go诞生的时间
	curT := time.Unix(int64(unixNum), 0)
	return TimeToString(curT)
}

func StringToTime(s string) time.Time {
	//2006-01-02 是 go诞生的时间
	//t,_ :=  time.Parse("2006-01-02 15:04:05", s);
	//字符串转时间一定要通过 time.ParseInLocation,不能直接用Parse
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	return t
}

// "2017-08-30 16:40:12" 转 时间戳 秒
func StrToUnixTime(stringTime string) int64 {
	//stringTime := "2017-08-30 16:40:12"
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
	if err == nil {
		unixTime := theTime.Unix() //1504082441
		return unixTime
	}
	return 0
}

/**
 * yyyy-mm-dd string to yyyymmdd string
 */
func StringTimeToString(s string) string {
	return strings.Replace(s, "-", "", -1)
}

//---- 3种 md5实现方式 ----

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5V2(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func Md5V3(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

//---- proto_exchangenumber ----

type Exchangenumber struct {
	Id    string
	Name  string
	Item  int
	Mutex string
}

//---- 硬盘上文件的写 ----

var VersionMap map[string]string

/**
 * version.txt中的文本格式
 */
type VersionLocalMem struct {
	ServerId int
	Version  string
}

func InitVerison() {
	ReadVersionFile()

	//TestVersionWrite();
}

func WriteVersionFile() error {
	str := "["
	endStr := "]."

	isHead := true

	for key, mem := range VersionMap {

		if !isHead {
			str += ","
		} else {
			isHead = false
		}

		str += "{" + key + "," + mem + "}"
	}

	str += endStr

	bytes := []byte(str)
	var baseDir string
	if cfg.GetBasePath() == "." {
		baseDir = "./beegoWeb/server_version/version.txt"
	} else {
		baseDir = cfg.GetBasePath() + "/beegoWeb/server_version/version.txt"
	}

	err := ioutil.WriteFile(baseDir, bytes, 0644)

	if err != nil {
		logger.Logf("version_controller write version.txt error baseDir:%v", baseDir)
	}

	return err
}

func ReadVersionFile() error {

	logger.Log("ReadVersionFile")

	//这个200 以后可以改
	VersionMap = make(map[string]string, 200)

	var url string
	if cfg.GetBasePath() == "." {
		url = "./beegoWeb/server_version/version.txt"
	} else {
		url = cfg.GetBasePath() + "/beegoWeb/server_version/version.txt"
	}

	//url := "../beegoWeb/server_version/version.txt"

	bytes, err := ioutil.ReadFile(url)

	if err != nil {
		logger.Logf("version_controller read version.txt error url:%v", url)

		//如果是不存在的 错误
		if os.IsNotExist(err) {
			//那么就创建
			os.Create(url)
			logger.Log("version_controller created version.txt success")
		}

		return err
	}

	//[{2,[91,53,48,93]}].
	vStr := string(bytes)
	strLen := len(vStr)

	if strLen <= 0 {
		logger.Log("version_controller read version.txt is None")
		return nil
	}

	vStr1 := vStr[1 : strLen-2]

	logger.Log(vStr1)

	//按VersionSplitFun规则切割字符串为 字符串数组
	reArr := strings.FieldsFunc(vStr1, VersionSplitFun)

	logger.Log(reArr)

	arrlen := len(reArr)

	for i := 0; i < arrlen; i += 2 {
		key := reArr[i]
		val := reArr[i+1]
		VersionMap[key] = val
	}

	logger.Logf("ReadVersionFile Result = ", VersionMap)

	return nil
}

//版本管理 获取当前平台历史版本
func Proc_get_version(ServerId int32) string {
	vServer := strconv.Itoa(int(ServerId))
	for key, mem := range VersionMap {
		if vServer == key {
			return mem
		}
	}
	return ""
}

func VersionSplitFun(c rune) bool {
	if c == '{' || c == '}' || c == ',' {
		return true
	}
	return false
}

func CopyVersionFile(url string, tarurl string) {

}

func GetFullPath(path string) string {
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}

func GetFileChildNames(path string) []string {

	fullPath := GetFullPath(path)

	var reNameArr []string

	files, err := ioutil.ReadDir(fullPath) //读取目录下文件
	if err != nil {
		logger.Logf("======fullPath:%v", fullPath)
		return reNameArr
	}

	for _, file := range files {

		if file.IsDir() { //是文件夹

			name := file.Name()

			if name == ".svn" || name == ".git" || name == "logs" || name == "zipversion" {
				continue
			} else {
				reNameArr = append(reNameArr, name)
			}
		} else {
			continue
		}
	}

	sort.Strings(reNameArr)

	return reNameArr
}

func ConvertToSlice(listStr *list.List) []string {
	sli := []string{}
	for el := listStr.Front(); nil != el; el = el.Next() {
		sli = append(sli, el.Value.(string))
	}

	return sli
}

//---- end ----
