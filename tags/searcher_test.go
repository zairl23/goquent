package tags

import (
	"fmt"
	"testing"
)

type UserSearch struct {
    Name   string `search:"col=name;opt=like"`
    UserId uint64 `search:"col=user_id;con=and"`
    Phone  string `search:"col=phone;con=and"`
}

func TestToSql(t *testing.T) {
	user := &UserSearch{
		Name: "chronos", 
		Phone:"123456789",
	}
	
	searcher := NewMysqlSearcher()

    search := searcher.ToSql(user)

    fmt.Println(search)
}