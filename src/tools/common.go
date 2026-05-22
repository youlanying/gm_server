package tools

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	MILLISECONDS_OF_DAY = 86400000

	MILLISECONDS_OF_HOUR = 3600000

	MILLISECONDS_OF_MINUTE = 60000

	MILLISECONDS_OF_SECOND = 1000

	RATE_BASE_NUM = 10000
)

func CheckError(err error) {
	if err != nil {
		//logger.LogFatalf("Fatal error: %v", err)
		os.Exit(-1)
	}
}

// 返回毫秒时间戳。
func msy(t time.Time) int32 {
	return (int32)((t.UnixNano() / 1000000) % 1000000000)
}

func MS(t time.Time) int64 {
	return int64(t.UnixNano() / 1000000)
}

// 返回unix时间戳毫秒
func CurrentMS() int64 {
	return int64(time.Now().UnixNano() / 1000000)
}

//返回unix时间戳秒
func CurrentS() int64 {
	return int64(time.Now().UnixNano() / 1000000 / 1000)
}

// NowStrTime 2006-01-02 15:04:05
func NowStrTime() string {
	strCreateTime := time.Now().Format("2006-01-02 15:04:05")
	return strCreateTime
}

func Ms2Time(ms int64) time.Time {
	sec := ms / 1e3
	nsec := (ms % 1e3) * 1e6
	return time.Unix(sec, nsec).UTC()
}

// date format: "2006-01-02 13:04:00"
func S2UnixTime(value string) int64 {
	re := regexp.MustCompile(`([\d]+)-([\d]+)-([\d]+) ([\d]+):([\d]+):([\d]+)`)
	slices := re.FindStringSubmatch(value)
	if slices == nil || len(slices) != 7 {
		//logger.LogErrf("time[%s] format error, expect format: 2006-01-02 13:04:00...", value)
		return 0
	}
	year, _ := strconv.Atoi(slices[1])
	month, _ := strconv.Atoi(slices[2])
	day, _ := strconv.Atoi(slices[3])
	hour, _ := strconv.Atoi(slices[4])
	min, _ := strconv.Atoi(slices[5])
	sec, _ := strconv.Atoi(slices[6])
	loc, _ := time.LoadLocation("UTC") // use UTC instend of Local
	t := time.Date(year, time.Month(month), day, hour, min, sec, 0, loc)
	return int64(t.UnixNano() / 1000000)
}

func Time2Midnight(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// 从一个毫秒时间戳获得当前时区的本日凌晨时间。
func Ms2Midnight(t int64) time.Time {
	midTime := Time2Midnight(time.Unix(t/1000, t%1000))
	return midTime
}

func CurMidnight() int64 {
	return MS(Time2Midnight(time.Now()))
}

func NextMidnight(t int64) int64 {
	midTime := Time2Midnight(time.Unix(t/1000, t%1000))
	return midTime.UnixNano()/1e6 + MILLISECONDS_OF_DAY
}

//na:=time.Now().UnixNano()//获取时间戳包含纳秒
//fmt.Println(na / 1e6) //纳秒转毫秒
//fmt.Println(na / 1e9) //纳秒转秒
func NextMidnightS(t int64) int64 {
	midTime := Time2Midnight(time.Unix(t, 0))
	return midTime.UnixNano()/1e9 + MILLISECONDS_OF_DAY/1000
}

// 从一个毫秒时间戳获取下一个准点时间。
func NextHour(t int64) int64 {
	t1 := time.Unix(t/1000, t%1000)
	year, month, day := t1.Date()
	hour, _, _ := t1.Clock()
	t2 := time.Date(year, month, day, hour+1, 0, 0, 0, t1.Location())
	return t2.UnixNano() / 1e6
}

// 同一天毫秒 true不是同一天
func OtherDay(curTime, lstTime int64) bool {
	return curTime >= NextMidnight(lstTime)
}

// 同一天 秒 true不是同一天
func OtherDayS(curTime, lstTime int64) bool {
	return curTime >= NextMidnightS(lstTime)
}

// OtherWeek 同一周
func OtherWeek(curTime, lstTime int64) bool {
	t1 := time.Unix(curTime/1000, curTime%1000)
	t2 := time.Unix(lstTime/1000, lstTime%1000)
	year1, week1 := t1.ISOWeek()
	year2, week2 := t2.ISOWeek()
	if year1 == year2 && week1 == week2 {
		return false
	}
	return true
}

// OtherWeek 秒 同一周
func OtherWeekS(curTime, lstTime int64) bool {
	t1 := time.Unix(curTime, 0)
	t2 := time.Unix(lstTime, 0)
	year1, week1 := t1.ISOWeek()
	year2, week2 := t2.ISOWeek()
	if year1 == year2 && week1 == week2 {
		return false
	}
	return true
}

// 表格时间转时间戳 ymd：年月日，hms：时分秒
func Int32TimeStampMS(ymd, hms int32) uint64 {
	year := ymd / 10000
	month := (ymd % 10000) / 100
	day := ymd % 100
	hour := hms / 10000
	min := (hms % 10000) / 100
	sec := hms % 100
	loc, _ := time.LoadLocation("UTC") // use UTC instend of Local
	t := time.Date(int(year), time.Month(int(month)), int(day), int(hour), int(min), int(sec), 0, loc)
	return uint64(t.UnixNano() / 1000000)
}

/*
  a simple 32 bit checksum that can be upadted from either end
  (inspired by Mark Adler's Adler-32 checksum)
*/
func Adler32(data []byte) uint32 {
	size := len(data)
	s1 := uint32(0)
	s2 := uint32(0)
	i := 0
	for i < size-4 {
		s2 += 4*(s1+uint32(data[i])) + 3*uint32(data[i+1]) + 2*uint32(data[i+2]) + uint32(data[i+3])
		s1 += uint32(data[i+0]) + uint32(data[i+1]) + uint32(data[i+2]) + uint32(data[i+3])
		i += 4
	}
	for i < size {
		s1 += uint32(data[i])
		s2 += s1
		i++
	}
	return (s1 & 0xffff) + (s2 << 16)
}

//bit
func SetBit(src int16, offset uint) int16 {
	return src | (1 << offset)
}

func ClearBit(src int16, offset uint) int16 {
	return src &^ (1 << offset)
}

func CheckBit(src int16, offset uint) bool {
	return (src & (1 << offset)) != 0
}
func SetBit32(src int32, offset uint) int32 {
	return src | (1 << offset)
}

func ClearBit32(src int32, offset uint) int32 {
	return src &^ (1 << offset)
}

func CheckBit32(src int32, offset uint) bool {
	return (src & (1 << offset)) != 0
}

func CheckBit64(src int64, offset uint) bool {
	return (src & (1 << offset)) != 0
}

func formatMapKey(values []reflect.Value) string {
	report := ""
	v := values
	if len(values) > 64 {
		v = values[:64]
	}

	for _, v := range v {
		if v.CanInterface() {
			report += fmt.Sprintf("%v, ", v.Interface())
		} else if v.Kind() == reflect.Ptr {
			e := v.Elem()
			if e.CanInterface() {
				report += fmt.Sprintf("%v, ", e.Interface())
			} else {
				report += fmt.Sprintf("NO SUPPORT, ")
			}
		}
	}

	if len(values) > 64 {
		report += "..."
	}

	return report
}

func formatStruct(s reflect.Value, deep int16) string {
	var report string
	if s.Kind() == reflect.Interface {
		s = s.Elem()
	}
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	prefix := ""
	for strdeep := deep; strdeep >= 0; strdeep-- {
		prefix += "\t"
	}

	typeOfT := s.Type()
	if s.Kind() == reflect.Struct {
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			if f.Kind() == reflect.Map {
				report += fmt.Sprintf("%s%s keys: {%v}\n", prefix,
					typeOfT.Field(i).Name, formatMapKey(f.MapKeys()))
			} else if (f.Kind() == reflect.Slice) || (f.Kind() == reflect.Array) {
				report += fmt.Sprintf("%s%s len: %d\n", prefix,
					typeOfT.Field(i).Name, f.Len())
			} else if f.Kind() == reflect.Struct {
				if deep > 1 {
					report += fmt.Sprintf("%s%s=%v\n", prefix,
						typeOfT.Field(i).Name, f.Interface())
				} else {
					report += fmt.Sprintf("%s%s:\n", prefix, typeOfT.Field(i).Name)
					report += formatStruct(f, deep+1)
				}
			} else if f.Kind() == reflect.Interface {
				if deep > 1 {
					report += fmt.Sprintf("%s%s=%v\n", prefix,
						typeOfT.Field(i).Name, f.Interface())
				} else {
					report += fmt.Sprintf("%s%s:\n", prefix, typeOfT.Field(i).Name)
					report += formatStruct(f, deep+1)
				}
			} else if f.CanInterface() {
				report += fmt.Sprintf("%s%s=%v\n", prefix,
					typeOfT.Field(i).Name, f.Interface())
			} else if f.Kind() == reflect.Ptr {
				e := f.Elem()
				if f.CanInterface() {
					report += fmt.Sprintf("%s%s=%v\n", prefix,
						typeOfT.Field(i).Name, e.Interface())
				} else {
					report += fmt.Sprintf("%s%s=NO SUPPORT\n", prefix,
						typeOfT.Field(i).Name)
				}
			}
		}
	} else {
		report += fmt.Sprintf("%s%s=%v\n", prefix,
			typeOfT.Name(), s.Interface())
	}
	return report
}

func FormatStruct(obj interface{}) string {
	return formatStruct(reflect.ValueOf(obj), 0)
}

func GetSplitField(src string, delim string, index int) string {
	var ret string
	fields := strings.Split(src, delim)
	if index >= 0 && index < len(fields) {
		ret = fields[index]
	}

	return ret
}

// 获取本地内网地址。
func GetLocalInternalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1", err
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "127.0.0.1", nil
}

// GetLocalInternalIpList 获取本地内网地址list
func GetLocalInternalIpList() []string {
	ipList := make([]string, 0)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ipList
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址 && !ipnet.IP.IsLoopback()
		if ipnet, ok := address.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				ipList = append(ipList, ipnet.IP.String())
			}
		}
	}
	return ipList
}

// 获取本地外网地址。
func GetLocalExternalIp() (string, error) {
	resp, e := http.Get("http://myexternalip.com/raw")
	if e != nil {
		return "127.0.0.1", e
	}
	defer resp.Body.Close()

	result, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return "127.0.0.1", e
	}
	reg := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)
	return reg.FindString(string(result)), nil
}

//随机概率生成值为1-10000，传人参数为0时随机结果始终为false
func GetRandRateResult(rate int) bool {
	rand.Seed(time.Now().UnixNano())
	randNum := 1 + rand.Intn(RATE_BASE_NUM)
	return rate >= randNum
}

//生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := random.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

func GetRandomFromGd(max int64, min int64) (ret int64) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	if max-min == 0 {
		ret = min
	} else if max-min > 0 {
		ret = int64(random.Intn(int(max-min)) + int(min))
	} else {
		// max-min < 0
		min = min + 1
		ret = int64(random.Intn(int(min-max)) + int(max))
	}
	return
}

//------------------------------------------------ 返回当天的0点
func PM12() int64 {
	t := time.Now().Unix()
	return t - t%86400
}

//------------------------------------------------ 明天utc0点
func NextPM12MS() int64 {
	return (PM12() + 86400) * 1000
}

func NextSundayMS() int64 {
	t := time.Now()
	weekday := t.Weekday()
	if weekday == time.Sunday {
		return (PM12() + 7*86400) * 1000
	}
	return (PM12() + int64(7-int(weekday))*86400) * 1000
}

//------------------------------------------------ 返回当前的整点时间
func SharpClock() int64 {
	t := time.Now().Unix()
	return t - t%3600
}

//------------------------------------------------ 返回从整点到现在的差值
func NowToSharpClock() int64 {
	t := time.Now().Unix()
	return t % 3600
}

//-------------------------------------------------------------------------------
//  随机数相关
//-------------------------------------------------------------------------------
func RandomValue(BaseRadio int32) int32 {
	rndSource := rand.NewSource(CurrentMS())
	rnd := rand.New(rndSource)
	return rnd.Int31n(BaseRadio)
}
