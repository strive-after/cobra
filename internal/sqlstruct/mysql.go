package sqlstruct


import (
"database/sql"
"errors"
"fmt"
//导入驱动
_ "github.com/go-sql-driver/mysql"
)
// 要获取表涨列的信息 需要访问数据库的information_schema数据库的COLUMNS表 做连接查询

//连接mysql的核心对象
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}
//存放mysql的基础信息
type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}
//村粗COLUMNS表中我们需要的一些字段
type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}
//由于datatype字段类型与go结构体不是完全一致的 如varcher等 所以要座一层简单的类型转换 这里用枚举然后用map做映射获取
var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?" +
		"charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}

	return nil
}
//需要针对COLUMNS进行查询和数据组装
func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, " +
		"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}

	return columns, nil
}

