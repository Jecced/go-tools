package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DbModel struct {
	DbEngine *sql.DB
	DbInfo   *DbInfo
}

type DbInfo struct {
	DbType   string
	Host     string
	Port     string
	UserName string
	Password string
	Charset  string
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

// 表字段类型映射
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

func NewDbModel(info *DbInfo) *DbModel {
	return &DbModel{DbInfo: info}
}

func (m *DbModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s:%s)/information_schema?" +
		"charset=%s&parseTime=True&loc=Local"

	dsn := fmt.Sprintf(
		s,
		m.DbInfo.UserName,
		m.DbInfo.Password,
		m.DbInfo.Host,
		m.DbInfo.Port,
		m.DbInfo.Charset,
	)
	m.DbEngine, err = sql.Open(m.DbInfo.DbType, dsn)
	if err != nil {
		return err
	}
	return nil
}

func (m *DbModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "select column_name, data_type, column_key, is_nullable, column_type, column_comment " +
		"from columns " +
		"where table_schema = ? and table_name = ?"
	rows, err := m.DbEngine.Query(query, dbName, tableName)
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
		err = rows.Scan(
			&column.ColumnName,
			&column.DataType,
			&column.ColumnKey,
			&column.IsNullable,
			&column.ColumnType,
			&column.ColumnComment,
		)
		columns = append(columns, &column)
	}
	return columns, nil
}
