// Code generaTed by fileb0x at "2017-06-25 01:10:57.089955815 +0800 CST" from config file "b0x.yml" DO NOT EDIT.

package swaggerFiles

import (
	"log"
	"os"
)

// FileSwaggerUIStandalonePresetJsMap is "/swagger-ui-standalone-preset.js.map"
var FileSwaggerUIStandalonePresetJsMap = []byte("\x7b\x22\x76\x65\x72\x73\x69\x6f\x6e\x22\x3a\x33\x2c\x22\x66\x69\x6c\x65\x22\x3a\x22\x73\x77\x61\x67\x67\x65\x72\x2d\x75\x69\x2d\x73\x74\x61\x6e\x64\x61\x6c\x6f\x6e\x65\x2d\x70\x72\x65\x73\x65\x74\x2e\x6a\x73\x22\x2c\x22\x73\x6f\x75\x72\x63\x65\x73\x22\x3a\x5b\x22\x77\x65\x62\x70\x61\x63\x6b\x3a\x2f\x2f\x2f\x73\x77\x61\x67\x67\x65\x72\x2d\x75\x69\x2d\x73\x74\x61\x6e\x64\x61\x6c\x6f\x6e\x65\x2d\x70\x72\x65\x73\x65\x74\x2e\x6a\x73\x22\x5d\x2c\x22\x6d\x61\x70\x70\x69\x6e\x67\x73\x22\x3a\x22\x41\x41\x41\x41\x3b\x41\x41\x2b\x75\x46\x41\x3b\x3b\x3b\x3b\x3b\x41\x41\x79\x4f\x41\x3b\x41\x41\x6f\x37\x47\x41\x3b\x41\x41\x77\x30\x46\x41\x3b\x3b\x3b\x3b\x3b\x3b\x41\x41\x6d\x5a\x41\x3b\x41\x41\x2b\x75\x46\x41\x3b\x41\x41\x79\x2b\x43\x41\x3b\x41\x41\x6f\x2b\x43\x41\x3b\x41\x41\x69\x72\x43\x41\x3b\x41\x41\x75\x79\x45\x41\x22\x2c\x22\x73\x6f\x75\x72\x63\x65\x52\x6f\x6f\x74\x22\x3a\x22\x22\x7d")

func init() {

	f, err := FS.OpenFile(CTX, "/swagger-ui-standalone-preset.js.map", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(FileSwaggerUIStandalonePresetJsMap)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
