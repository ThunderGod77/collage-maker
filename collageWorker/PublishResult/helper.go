package PublishResult

import (
	"bytes"
	"collageWorker/Global"
	"encoding/gob"
	"log"
)

func OnError(msg string,folderName string) {
	r := Global.WorkCompletion{
		Err: true,
		Msg: msg,
		FolderId: folderName,
	}
	var mqCarrier bytes.Buffer
	enc := gob.NewEncoder(&mqCarrier)
	err := enc.Encode(r)
	if err != nil {
		log.Println(err)
	}
	wPublish(mqCarrier.Bytes(),folderName)
}
func OnCompletion(folderName string){
	r := Global.WorkCompletion{
		Err: false,
		Msg: "Work Done!",
		FolderId: folderName,
	}
	var mqCarrier bytes.Buffer
	enc := gob.NewEncoder(&mqCarrier)
	err := enc.Encode(r)
	if err != nil {
		log.Println(err)
	}
	wPublish(mqCarrier.Bytes(),folderName)
}

