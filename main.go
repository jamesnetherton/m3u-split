package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jamesnetherton/m3u"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("Usage m3u-split <m3u file>")
	}

	playList, err := m3u.Parse(os.Args[1])
	if err == nil {
		for _, track := range playList.Tracks {
			m3u := fmt.Sprintf("#EXTM3U\n#EXTINF:%d,%s\n%s", track.Length, track.Name, track.Path)
			ioutil.WriteFile(track.Name+".m3u", []byte(m3u), 0644)
		}
	} else {
		log.Fatalf(err.Error())
	}
}
