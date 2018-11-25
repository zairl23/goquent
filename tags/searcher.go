package tags

import (
	"fmt"
	"strings"
    "reflect"
)

var (
	Opts = map[string]string{
		"uint": "%d",
		"int": "%d",
		"int64": "%d",
		"uint64": "%d", 
		"string": "'%s'",
		"float": "%f",
		"float64": "%f",
	}
)

type Searcher interface {
	ToSql(interface{}) string
}

type MysqlSearcher struct {

}

func NewMysqlSearcher() MysqlSearcher {
	return MysqlSearcher{}
}

type Where struct {
	Sql string 
	Conn string
}

func (ms MysqlSearcher) ToSql(data interface{}) string {
	var whereSlice = make([]Where, 0) 
	s := reflect.TypeOf(data).Elem()
	values := reflect.ValueOf(data).Elem()

    for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		search := field.Tag.Get("search")
		if search == "" {
			continue
		}
		
		where := toSqlExpression(search, field.Type.String(), values.Field(i))
		whereSlice = append(whereSlice, where)
	}
	
	var sql string
	if len(whereSlice) > 0 {
		for index, where := range whereSlice {
			if index == 0 {
				sql += "where " + where.Sql
			} else  {
				sql += " " + whereSlice[index].Conn + " " + where.Sql
			}
		}
	}

    return sql
}

// Convert search tag string to sql string
func toSqlExpression(search, t string, v interface{}) Where {
	var where Where
	fields := strings.Split(search, ";")

	if len(fields) > 0 {
		var (
			col string = ""
			con string = "and"
			opt string = "="
		)

		for _, f := range fields {
			detail := strings.Split(f, "=")

			switch detail[0] {
			case "col":
				col = detail[1]
			case "con":
				con = detail[1]
			case "opt":	
				opt = detail[1]
			}
		}

		if col == "" {
			panic("col not set")
		}

		switch opt {
		case "like":
			v = fmt.Sprintf("%%%s%%", v)
		}

		sql := fmt.Sprintf("%s %s " + Opts[t], col, opt, v);

		where = Where{
			Sql: sql,
			Conn: con,
		}
	}

	return where
}