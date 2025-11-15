package db

import (
    "database/sql"
    "log"
    "strings"

    "github.com/AngelUNC/GhostWA/types"
)

var contactCache = make(map[int64]types.ContactInfo)
var jidToName = make(map[string]string)

func PreloadAll(waDB *sql.DB) {
    if waDB == nil { return }

    // Load wa_contacts
    r, err := waDB.Query("SELECT _id, jid, number, display_name, wa_name FROM wa_contacts")
    if err == nil {
        defer r.Close()
        for r.Next() {
            var id sql.NullInt64
            var jid sql.NullString
            var number sql.NullString
            var display sql.NullString
            var waName sql.NullString

            if err := r.Scan(&id, &jid, &number, &display, &waName); err == nil {
                contactCache[id.Int64] = types.ContactInfo{
                    ID:          id.Int64,
                    JID:         jid.String,
                    Number:      number.String,
                    DisplayName: display.String,
                    WaName:      waName.String,
                }
            }
        }
    }

    // wa_address_book
    r2, err2 := waDB.Query("SELECT jid, display_name FROM wa_address_book")
    if err2 == nil {
        defer r2.Close()
        for r2.Next() {
            var jid sql.NullString
            var display sql.NullString

            if err := r2.Scan(&jid, &display); err == nil {
                if jid.Valid && display.Valid && display.String != "" {
                    jidToName[jid.String] = display.String
                }
            }
        }
    }

    log.Printf("✅ Preloaded wa contacts: %d, addressbook: %d",
        len(contactCache), len(jidToName))
}

func ResolveContext(msg types.WhatsAppMessage, waDB *sql.DB) (senderName string, chatLabel string, isGroup bool) {

    senderName = "(desconocido)"
    chatLabel = ""
    isGroup = false

    // group sender
    if msg.SenderJidRowID != nil {
        if ci, ok := contactCache[*msg.SenderJidRowID]; ok {
            if ci.DisplayName != "" {
                senderName = ci.DisplayName
            } else if name, ok2 := jidToName[ci.JID]; ok2 {
                senderName = name
            } else {
                senderName = ci.Number
            }
        }
        isGroup = true
        return
    }

    // private chat
    if msg.ChatRowID != nil && waDB != nil {
        var jid sql.NullString
        err := waDB.QueryRow(
            "SELECT jid FROM chat_list WHERE _id = ? LIMIT 1",
            *msg.ChatRowID).Scan(&jid)

        if err == nil && jid.Valid {
            j := jid.String

            // Try address book name
            if name, ok := jidToName[j]; ok {
                senderName = name
            } else {
                // Try wa_contacts by JID
                for _, ci := range contactCache {
                    if ci.JID == j {
                        senderName = ci.DisplayName
                        break
                    }
                }
            }

            isGroup = strings.HasSuffix(strings.ToLower(j), "@g.us")
        }
    }

    return
}
