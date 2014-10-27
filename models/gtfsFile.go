package models

import (
	"github.com/akinsella/go-playground/utils"
)

type GTFSFile struct {
	Filename string
}

func (gf GTFSFile) LinesIterator() <- chan []byte {
	channel := make(chan []byte)
	go func() {
		utils.ReadCsvFile(gf.Filename, channel)
		defer close(channel)
	}()
	return channel
}