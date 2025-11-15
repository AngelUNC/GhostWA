package types

type WhatsAppMessage struct {
	RowID          int64
	Timestamp      *int64
	TextData       *string
	MessageType    *int
	ChatRowID      *int64
	FromMe         *int
	SenderJidRowID *int64
}

type ContactInfo struct {
	ID          int64
	JID         string
	Number      string
	DisplayName string
	WaName      string
}
