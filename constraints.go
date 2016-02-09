package qbit

import (
	"fmt"
	"strings"
)

type PrimaryKey Constraint
type ForeignKey Constraint

type Constraint struct {
	Name     string
}

func Null() Constraint {
	return Constraint{"NULL"}
}

func NotNull() Constraint {
	return Constraint{"NOT NULL"}
}

func Default(value interface{}) Constraint {
	return Constraint{fmt.Sprintf("DEFAULT `%v`", value)}
}

func Unique(cols ...string) Constraint {
	if len(cols) == 0 {
		return Constraint{"UNIQUE"}
	}
	return Constraint{fmt.Sprintf("UNIQUE(%s)", strings.Join(cols, ", "))}
}

func Key() Constraint {
	return Constraint{"KEY"}
}

// TODO: Determine if these are needed
//func PrimaryKey(cols ...string) Constraint {
//	if len(cols) == 0 {
//		return Constraint{"PRIMARY KEY"}
//	}
//	constraint := Constraint{fmt.Sprintf("PRIMARY KEY(%s)", strings.Join(cols, ", "))}
//	constraint.Delegate = true
//	return constraint
//}

//func ForeignKey(cols string, reftable string, refcols string) Constraint {
//	constraint := Constraint{
//		fmt.Sprintf(
//			"FOREIGN KEY (%s) REFERENCES %s(%s)",
//			cols,
//			reftable,
//			refcols,
//		),
//	}
//	constraint.Delegate = true
//	return constraint
//}

//func References(table string, refcol string) Constraint {
//	return Constraint{
//		fmt.Sprintf(
//			"REFERENCES %s(%s)",
//			table,
//			refcol,
//		),
//	}
//}

func Index() Constraint {
	return Constraint{"INDEX"}
}
