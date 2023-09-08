package ent

import (
	"context"
	"fmt"
	"todo/ent/exercise"
	"todo/ent/musclesgroup"
	"todo/ent/schema/pksuid"
	"todo/ent/todo"
)

// prefixMap maps PKSUID prefixes to table names.
var prefixMap = map[pksuid.ID]string{
	"TD": todo.Table,
	"EX": exercise.Table,
	"MG": musclesgroup.Table,
}

// IDToType maps a pksuid.ID to the underlying table.
func IDToType(ctx context.Context, id pksuid.ID) (string, error) {
	if len(id) < 2 {
		return "", fmt.Errorf("IDToType: id too short")
	}
	prefix := id[:2]
	typ := prefixMap[prefix]
	if typ == "" {
		return "", fmt.Errorf("IDToType: could not map prefix '%s' to a type", prefix)
	}
	return typ, nil
}
