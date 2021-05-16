package Global

import "github.com/streadway/amqp"

var Ch *amqp.Channel
var Conn *amqp.Connection

type ImagesInfo struct {
	Image1      string
	Image2      string
	Image3      string
	FolderId    string
	Instruction string
	Color       string
	BorderWidth string
}
type WorkCompletion struct {
	Err bool
	Msg string
	FolderId string
}