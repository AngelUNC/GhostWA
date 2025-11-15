package core
import (
	"database/sql"
	"log"
	"github.com/AngelUNC/GhostWA/db"
)

var messageCache = make(map[int64]MessageSnapshot)

func InitializeSnapshot(msgDB *sql.DB) {
	rows, err := msgDB.Query("SELECT _id, text_data, message_type FROM message
ORDER BY _id DESC LIMIT ?", db.CacheLookback)
	if err != nil {
		log.Printf(" initializeSnapshot query falló: %v", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
	var id sql.NullInt64
	var text sql.NullString
	var mtype sql.NullInt64
	if err := rows.Scan(&id, &text, &mtype); err != nil {
		continue
	}
	messageCache[id.Int64] = MessageSnapshot{Text:
db.NullStringToString(text), Type: db.NullInt64ToInt(mtype)}
	}
}