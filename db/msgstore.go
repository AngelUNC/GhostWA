package db
import (
	"database/sql"
	"fmt"
)

const (
	MsgDBPath = "/data/data/com.whatsapp/databases/msgstore.db"
	WaDBPath = "/data/data/com.whatsapp/databases/wa.db"
)

var pollMs = 200
var cacheLookback = 500

func PollMs() int { return pollMs }
func CacheLookback() int { return cacheLookback }


func OpenMsgDB(path string) (*sql.DB, error) {
	return sql.Open("sqlite", fmt.Sprintf("file:%s?mode=ro&_journal_mode=WAL", path))
}

func OpenWaDB(path string) (*sql.DB, error) {
	return sql.Open("sqlite", fmt.Sprintf("file:%s?mode=ro", path))
}


func NullStringToString(ns sql.NullString) string {
	if !ns.Valid { return "" }
	return ns.String
}
func NullInt64ToInt(ni sql.NullInt64) int {
	if !ni.Valid { return 0 }
	return int(ni.Int64)
}

func StrOrEmpty(ps *string) string {
	if ps == nil { return "" }
	return *ps
}

func IntPtrToInt(pi *int) int {
	if pi == nil { return 0 }
	return *pi
}

func FormatTimestamp(ts *int64) string {
	if ts == nil { return "????/??/?? ??:??:??" }
	return time.UnixMilli(*ts).Local().Format("2006/01/02 15:04:05")
}