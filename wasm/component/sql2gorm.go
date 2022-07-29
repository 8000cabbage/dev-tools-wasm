package component

import (
	"github.com/cascax/sql2gorm/parser"
)

func Sql2Gorm(sql string) ([]string, error) {
	code, err := parser.ParseSql(sql)
	return code.StructCode, err
}
