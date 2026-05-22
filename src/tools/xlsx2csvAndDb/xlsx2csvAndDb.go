// refer to https://github.com/tealeg/xlsx2csv/archive/master.zip
package main

import (
	"bufio"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xuri/excelize/v2"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var xlsxPath = flag.String("i", "./", "Path to an XLSX file or directory")
var sheetIndex = flag.Int("n", 0, "Index of sheet to convert, zero based")
var outPath = flag.String("o", "./", "Path to store result file")
var outType = flag.String("t", "d", "Type of the result file: c:.csv d:.db")

const DB_NAME string = "./slg_sqlite.db"

var dbFile string

//var delimiter = flag.String("d", ",", "Delimiter to use between fields")
type outputer func(s string)

func removeExistedFile(fileName string) {
	//如果指定的文件或目录存在
	_, err := os.Stat(fileName)
	if err == nil || os.IsExist(err) {
		//删除文件
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println("file remove Error!")
			fmt.Printf("%s", err)
		} else {
			fmt.Printf("Remove old file [%s] completely!\n", fileName)
		}
	}
}

func generateCSVFromXLSXFile(excelFileName string, sheetIndex int, outputf outputer) error {
	xlFile, error := excelize.OpenFile(excelFileName)
	if error != nil {
		return error
	}
	sheetLen := len(xlFile.GetSheetList())
	switch {
	case sheetLen == 0:
		return errors.New("This XLSX file contains no sheets.")
	case sheetIndex >= sheetLen:
		return fmt.Errorf("No sheet %d available, please select a sheet between 0 and %d\n", sheetIndex, sheetLen-1)
	}
	sheetName := xlFile.GetSheetList()[sheetIndex]
	//仅取文件名，去掉原文件的后缀，修改为目标文件类型csv的后缀
	csvFile := *outPath + strings.Split(filepath.Base(excelFileName), ".")[0] + ".csv"
	removeExistedFile(csvFile)
	//新导出的文件直接替换旧文件，不使用O_APPEND
	fi, err := os.OpenFile(csvFile, os.O_CREATE|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	defer fi.Close()
	checkErr(err)
	// rowCnt记录行数
	rowCnt := 0
	//colNum记录列数
	colNum := -1
	rows, err := xlFile.GetRows(sheetName)
	for _, row := range rows {
		var vals []string
		rowCnt++
		colCnt := 1

		if row != nil {
			for _, str := range row {
				//首列即主键，主键为“”时不再读excel
				if (str == "" || len(str) == 0) && colCnt == 1 {
					break
				}
				//第一行为表的字段名，从第一行获取列数
				if rowCnt == 1 {
					if str == "" {
						colNum = colCnt
						break
					}
				}
				// 从第二行开始为字段的值
				vals = append(vals, fmt.Sprintf("%q", str))

				if colCnt == colNum {
					break
				}
				colCnt++
			} //读完一行
			//读完第一行后，记录列数
			if rowCnt == 1 {
				colNum = colCnt - 1
			}
			//读完一行后，写入目标文件
			fi.WriteString(strings.Join(vals, `,`) + "\n")
		}
	}
	fmt.Printf("Executed successfully: Convert %s to .csv completely.\n", excelFileName)
	return nil
}

func generateDbFromXLSXFile(excelFileName string, sheetIndex int, outputf outputer) error {
	xlFile, error := excelize.OpenFile(excelFileName)
	checkErr(error)
	sheetLen := len(xlFile.GetSheetList())
	switch {
	case sheetLen == 0:
		return errors.New("This XLSX file contains no sheets.")
	case sheetIndex >= sheetLen:
		return fmt.Errorf("No sheet %d available, please select a sheet between 0 and %d\n", sheetIndex, sheetLen-1)
	}
	sheetName := xlFile.GetSheetList()[sheetIndex]

	dbFile = *outPath + DB_NAME

	db, err := sql.Open("sqlite3", dbFile)
	checkErr(err)
	defer db.Close()
	//仅取文件名(去掉原文件的后缀)——文件名即表名
	tableName := strings.Split(filepath.Base(excelFileName), ".")[0]
	//keys记录插入的字段名，仅读excel的第一行时写入
	var keys []string
	//sqlIns记录全部插入语句
	var sqlIns []string
	//rowCnt记录行数: 第一行为字段名，从第二行开始为字段的值
	rowCnt := 0

	sqlCreate := `create table  ` + tableName + ` (`
	rows, err := xlFile.GetRows(sheetName)
	for _, row := range rows {
		//vals 记录插入的字段值
		var vals []string
		//colCnt记录列数， 将字段值填入对应的字段
		colCnt := 0

		rowCnt++
		if row != nil {
			//按行读取：逐个读取各个cell的值
			for _, str := range row {
				//第一行为表的字段名
				if rowCnt == 1 {
					if str == "" {
						break
					}
					keys = append(keys, fmt.Sprintf("%v", str))
				} else {
					// 从第二行开始为字段的值
					if strings.Contains(keys[colCnt], "INT") {
						vals = append(vals, fmt.Sprintf("%v", str))
					} else {
						//						vals = append(vals, fmt.Sprintf("%#q", str))
						vals = append(vals, "\""+str+"\"")
					}
					colCnt++
					if colCnt == len(keys) {
						break
					}
				}
			} //读完一行
			//第一行只读字段名keys，没有字段值vals，只做create操作，不做insert操作
			if rowCnt == 1 {
				for _, v := range keys {
					if strings.Contains(v, "INT") {
						sqlCreate += v + ` integer, `
					} else {
						sqlCreate += v + ` text, `
					}
				}
				//截掉最后一个空格和逗号
				sqlCreate = sqlCreate[0 : len(sqlCreate)-2]
				sqlCreate += `);`

				if _, err := db.Exec(sqlCreate); nil == err {
					fmt.Printf("Create table[%s] completely, performing insert-operation ...\n", tableName)
				} else {
					fmt.Printf("Error: Create table[%s] failed, please check the 'Create' instruction below:\n[%v]\n", tableName, sqlCreate)
					return err
				}
				continue
			}
			//拼接 insert 语句
			sqlInsert := `insert into ` + tableName + `(`
			for _, v := range keys {
				sqlInsert += v + `, `
			}
			//截掉最后一个空格和逗号
			sqlInsert = sqlInsert[0 : len(sqlInsert)-2]
			sqlInsert += `) values(`
			for _, v := range vals {
				sqlInsert += v + `, `
			}
			sqlInsert = sqlInsert[0 : len(sqlInsert)-2]

			sqlInsert += `);`
			sqlIns = append(sqlIns, fmt.Sprintf("%v", sqlInsert))
		}
	}
	// "outputSql"是输出.sql文件的开关, 用于测试导出的sql
	outputSql := false
	//var fi *os.File
	if outputSql {
		sqlFile := *outPath + strings.Split(filepath.Base(excelFileName), ".")[0] + ".sql"
		removeExistedFile(sqlFile)
		fi, err := os.OpenFile(sqlFile, os.O_CREATE|os.O_RDWR, os.ModePerm|os.ModeTemporary)
		defer fi.Close()
		checkErr(err)
		for _, v := range sqlIns {
			fi.WriteString(v + "\n")
		}
	}
	// 部分excel行数太多，以事务方式处理
	db.Exec("begin;")
	for _, v := range sqlIns {
		db.Exec(v + "\n")
	}
	db.Exec("commit;")

	fmt.Printf("Executed successfully: Convert %s to .db completely.\n", excelFileName)
	return nil
}

func main() {
	flag.Parse()
	fmt.Println("Start dealing xlsx in processing, wait a moment ...\n")
	//所有表写到同一个DB中，因此在遍历前删除已有DB
	if strings.Contains(*outType, "d") {
		dbFile = *outPath + DB_NAME
		removeExistedFile(dbFile)
	}
	os.MkdirAll(path.Dir(*outPath), os.ModePerm)

	getExcelFiles(*xlsxPath)

	fmt.Println("Dealing completely. Press ENTER to exit... ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Printf("%v", string([]byte(input)[0]))
}

func getExcelFiles(path string) {
	err := filepath.Walk(path, func(path string, file os.FileInfo, err error) error {
		if file == nil {
			return err
		}
		//去除零时文件和文件夹文件
		if file.IsDir() || strings.Contains(file.Name(), "~") {
			return nil
		}
		//仅处理xlsx格式的excel文件，因为："github.com/tealeg/xlsx"目前不支持xls格式
		if !strings.HasSuffix(file.Name(), ".xlsx") {
			return nil
		}
		printer := func(str string) {
			fmt.Printf("%s", str)
		}

		if strings.Contains(*outType, "d") {
			fmt.Printf("Convert [%s] to .db file ...\n", path)
			if err := generateDbFromXLSXFile(path, *sheetIndex, printer); err != nil {
				fmt.Println(err)
			}
		}

		if strings.Contains(*outType, "c") {
			fmt.Printf("Convert [%s] to .csv file ...\n", path)
			if err := generateCSVFromXLSXFile(path, *sheetIndex, printer); err != nil {
				fmt.Println(err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
