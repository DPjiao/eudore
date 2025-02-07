package eudore

/*
用于根据配置信息来初始化程序。
*/

import (
	"encoding/xml"
)

type (
	seteudore struct {
		XMLName    xml.Name        `xml:"eudore"`
		Logger     setlogger       `xml:"logger" json:"logger"`
		Reload     []setreload     `xml:"reload" json:"reload"`
		Middleware []setmiddleware `xml:"middleware" json:"middleware"`
		Server     []setserver     `xml:"server" json:"server"`
		Router     []setrouter     `xml:"router" json:"router"`
	}
	setlogger struct {
		Type string `xml:"type" json:"type"`
	}
	setreload struct {
		Name  string `xml:"name" json:"name"`
		Index string `xml:"index" json:"index"`
		Func  string `xml:"func" json:"func"`
	}
	setmiddleware struct {
	}
	setserver struct {
		Port     int  `xml:"port" json:"port"`
		Https    bool `xml:"https" json:"https"`
		Http2    bool `xml:"http2" json:"http2"`
		Keyfile  string
		Cretfile string
	}
	setrouter struct {
		Router  []setrouter  `xml:"router" json:"router"`
		Handler []sethandler `xml:"handler" json:"handler"`
		Type    string       `xml:"type" json:"type"`
		Path    string       `xml:"path" json:"path"`
	}
	sethandler struct {
		Path string `xml:"path" json:"path"`
		Func string `xml:"func" json:"func"`
	}
)
