package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"

	"github.com/xwb1989/sqlparser"

	"github.com/ehlxr/go-utils/utils/generator/tmpl"
	"github.com/ehlxr/go-utils/utils/log"
)

// 定义 sqlData 数据结构。
type SqlData struct {
	ClassName      string
	ClassNameLower string
	PackageName    string
	TableName      string
	ColumnList     []SqlColumn
}

// 表属性
type SqlColumn struct {
	Name       string // 英文小写名称
	CnName     string // 输入的中文名称
	Type       string // 类型
	GoName     string // Golang 名称
	GoType     string // Golang 类型
	IsNumber   bool   // 判断是否数字
	ColName    string
	ColType    string
	NameUpper  string
	IsLast     bool
	IsId       bool // 判断是否是 Id
	IsDateTime bool // 判断是否是时间类型
}

var (
	// Options
	flgSql  = flag.String("sql", "", "path to sql file or '-' to read from stdin")
	flagClz = flag.String("clz", "", "ClassName")
	flagPkg = flag.String("pkg", "", "PackageName")
)

func main() {
	log.SetLogLevel(logrus.InfoLevel)
	flag.Parse()

	if *flgSql == "" {
		log.Fatal("sql is nil")
	}
	if *flagClz == "" {
		log.Fatal("ClassName is nil")
	}
	if *flagPkg == "" {
		log.Fatal("PackageName is nil")
	}

	loadSql, err := loadData(*flgSql)
	if err != nil {
		log.Fatal(err)
	}

	sql := string(loadSql)
	log.Debugf("input sql %s", sql)

	// 规则引擎的要求，) 后面的不能解析。
	last := strings.LastIndex(sql, ")")
	if last > 0 {
		sql = sql[:last+1]
	}
	//AUTO_INCREMENT 不解析，只解析标准的 sql
	sql = strings.Replace(sql, "AUTO_INCREMENT", "", -1)
	sql = strings.Replace(sql, "auto_increment", "", -1)
	log.Debugf("new sql %s", sql)

	sqlNode, err := sqlparser.ParseStrictDDL(sql)
	if err != nil {
		log.Fatal(err)
	}

	sqlDataTmp, err := getSqlData(sqlNode)
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("sql data: %v", sqlDataTmp)
	for i, v := range sqlDataTmp.ColumnList {
		log.Debugf("index: %d, value: %v", i, v)
	}

	// path, err := filepath.Abs("")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pojo := filepath.Join(path, "tmpl/java/pojo.tmpl")
	// log.Info(pojo)

	data := make(map[string]interface{})
	data["SqlData"] = sqlDataTmp

	t := template.New("javapojo")
	t = template.Must(t.Parse(tmpl.Pojo))
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}

	t = template.New("javacontroller")
	t = template.Must(t.Parse(tmpl.Controller))
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}

// Helper func:  Read input from specified file or stdin
func loadData(p string) ([]byte, error) {
	if p == "" {
		return nil, fmt.Errorf("No path specified")
	}

	var rdr io.Reader
	if p == "-" {
		rdr = os.Stdin
	} else if p == "+" {
		return []byte("{}"), nil
	} else {
		if f, err := os.Open(p); err == nil {
			rdr = f
			defer f.Close()
		} else {
			return nil, err
		}
	}
	return ioutil.ReadAll(rdr)
}

// 解析 sql 并返回 sqlData 数据
func getSqlData(sqlNode sqlparser.SQLNode) (*SqlData, error) {
	node, ok := sqlNode.(*sqlparser.DDL)
	if !ok {
		return nil, errors.New("不是标准的创建 sql 语句")
	}

	//返回首字母小写
	first := string((*flagClz)[0])
	clzNameLower := strings.ToLower(first) + (*flagClz)[1:]

	sqlData := &SqlData{
		PackageName:    *flagPkg,
		ClassNameLower: clzNameLower,
		ClassName:      *flagClz,
		TableName:      node.Table.Name.String(),
		ColumnList:     getAllColumn(node)}

	return sqlData, nil
}

// 返回属性数组
func getAllColumn(node *sqlparser.DDL) []SqlColumn {
	columnList := []SqlColumn{}
	for index, col := range node.TableSpec.Columns {
		colName := col.Name.String()
		colType := col.Type.Type

		tmpColumn := SqlColumn{ColName: colName, ColType: colType}
		// 设置首字母大写，驼峰命名
		tmpColumn.Name, tmpColumn.GoName = getNameByTitle(colName)
		// 设置类型
		tmpColumn.Type, tmpColumn.GoType = getType(colType)
		// 返回首字母大写
		first := string(tmpColumn.Name[0])
		nameUpper := strings.ToUpper(first) + tmpColumn.Name[1:]
		tmpColumn.NameUpper = nameUpper
		// 判断是否是最后一个数据
		if index == (len(node.TableSpec.Columns) - 1) {
			tmpColumn.IsLast = true
		} else {
			tmpColumn.IsLast = false
		}
		//判断数据是否 ==  Id
		if strings.ToLower(colName) == "id" {
			tmpColumn.IsId = true
		} else {
			tmpColumn.IsId = false
		}
		// 判断数据是否 DateTime 类型
		if strings.ToLower(colType) == "datetime" {
			tmpColumn.IsDateTime = true
		} else {
			tmpColumn.IsDateTime = false
		}
		// 添加数据
		columnList = append(columnList, tmpColumn)
	}
	return columnList
}

// 获取 title 拆分的名称。java的 和 golang 的两个名称
func getNameByTitle(colName string) (string, string) {
	if colName == "" {
		return "", ""
	}
	tmp := strings.Replace(colName, "_", " ", -1)
	tmp = strings.Title(tmp) //title 支持空格分开的都title。
	tmp = strings.Replace(tmp, " ", "", -1)
	//返回首字母小写
	first := string(tmp[0])
	tmp2 := strings.ToLower(first) + tmp[1:]
	return tmp2, tmp
}

// 获取类型，java 和 golang 的两个类型
func getType(typeName string) (string, string) {
	if typeName == "" {
		return "", ""
	}
	var stringReg = regexp.MustCompile(`varchar|char`)    // Has digit(s)
	var dateTimeReg = regexp.MustCompile(`datetime|date`) // Has digit(s)
	var longReg = regexp.MustCompile(`bigint|long`)       // Has digit(s)
	var integerReg = regexp.MustCompile(`integer|int`)    // Has digit(s)
	var floatReg = regexp.MustCompile(`float`)            // Has digit(s)
	var doubleReg = regexp.MustCompile(`float`)           // Has digit(s)

	var tmp, goTmp string
	switch {
	case stringReg.MatchString(typeName):
		tmp = "String"
		goTmp = "string"
		break
	case dateTimeReg.MatchString(typeName):
		tmp = "Date"
		goTmp = "time.Time"
		break
	case longReg.MatchString(typeName):
		tmp = "Long"
		goTmp = "int64"
		break
	case integerReg.MatchString(typeName):
		tmp = "Integer"
		goTmp = "int32"
		break
	case floatReg.MatchString(typeName):
		tmp = "Float"
		goTmp = "float64"
		break
	case doubleReg.MatchString(typeName):
		tmp = "Double"
		goTmp = "float64"
		break
	}
	return tmp, goTmp
}
