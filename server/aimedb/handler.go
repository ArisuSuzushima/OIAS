package aimedb

type CommandHandler interface {
	Handle(header *AimeDbHeader, payload []byte) (resultCode uint16, response []byte, err error)
}
