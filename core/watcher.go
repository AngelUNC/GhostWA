package core

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
	"github.com/AngelUNC/GhostWA/db"
)

func WatchMessages(msgDB *sql.DB, waDB *sql.DB) {
	for {
		rows, err := msgDB.Query(`
			SELECT _id, timestamp, text_data, message_type, chat_row_id, from_me, sender_jid_row_id
			FROM message
			ORDER BY _id DESC
			LIMIT ?
		`, db.CacheLookback)
		if err != nil {
			log.Printf("⚠️ Error ejecutando query de mensajes: %v", err)
			time.Sleep(time.Millisecond * time.Duration(db.PollMs()))
			continue
		}

		var msgs []WhatsAppMessage
		for rows.Next() {
			var (
				_id, timestamp, message_type, chat_row_id, from_me, sender_jid_rowid sql.NullInt64
				text_data sql.NullString
			)

			if err := rows.Scan(&_id, &timestamp, &text_data, &message_type, &chat_row_id, &from_me, &sender_jid_rowid); err != nil {
			log.Printf(" Error scanning row: %v", err)
			continue
			}

			var msg WhatsAppMessage
			msg.RowID = _id.Int64
			if timestamp.Valid { ts := timestamp.Int64; msg.Timestamp = &ts }
			if text_data.Valid { td := text_data.String; msg.TextData = &td }
			if message_type.Valid { mt := int(message_type.Int64); msg.MessageType = &mt }
			if chat_row_id.Valid { cr := chat_row_id.Int64; msg.ChatRowID = &cr }
			if from_me.Valid { fm := int(from_me.Int64); msg.FromMe = &fm }
			if sender_jid_rowid.Valid { sj := sender_jid_rowid.Int64; msg.SenderJidRowID = &sj }
			msgs = append(msgs, msg)
		}
		rows.Close()


		for i := len(msgs) - 1; i >= 0; i-- {
			msg := msgs[i]
			current := MessageSnapshot{Text: db.StrOrEmpty(msg.TextData), Type: db.IntPtrToInt(msg.MessageType)}
			prev, exists := messageCache[msg.RowID]

			senderName, chatLabel, isGroup := db.ResolveContext(msg, waDB)

			var prefix string
			if isGroup {
				prefix = fmt.Sprintf("[Grupo: %s] Remitente: %s", chatLabel, senderName)
			} else { 
				prefix = senderName
			}

			if !exists {
			messageCache[msg.RowID] = current
			fmt.Printf("%s - %s\n", db.FormatTimestamp(msg.Timestamp), prefix)
			continue
			}

			if prev.Text != current.Text || prev.Type != current.Type {
			if current.Type == 15 || strings.EqualFold(current.Text, "(null)") {
				fmt.Printf("%s - Mensaje eliminado: %s\n", time.Now().Format("2006/01/02 15:04:05"), prev.Text)
			} else if prev.Text != current.Text && prev.Text != "" {
				fmt.Printf("%s - Mensaje editado:\nOriginal: %s\nEditado: %s\n", time.Now().Format("2006/01/02 15:04:05"), prev.Text, current.Text)
			}
			messageCache[msg.RowID] = current
			}
		}

		time.Sleep(time.Millisecond * time.Duration(db.PollMs()))
	}
}
