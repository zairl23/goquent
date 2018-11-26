package tag

import (
        "fmt"
        "reflect"
)

type Query struct {
        Type string
        Sql  string
}

// read the static tag, convent to sql
func MakeQuery(data interface{}) map[string]Query {
        queries := make(map[string]Query)
        s := reflect.TypeOf(data).Elem()
        values := reflect.ValueOf(data).Elem()

        for i := 0; i < s.NumField(); i++ {
                field := s.Field(i)
                statics := field.Tag.Get("statics")
                if statics == "" {
                        continue
                }

                queries[field.Name] = Query{
                        Sql:  fmt.Sprintf(statics, values.FieldByName("StartDay").String()),
                        Type: field.Type.String(),
                }
        }

        return queries

}
