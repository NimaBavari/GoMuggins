package muggins

type Stream interface {
	ReadMessage() (int, []byte, error)
	WriteMessage(messageType int, msg []byte) error
	Close() error
}
