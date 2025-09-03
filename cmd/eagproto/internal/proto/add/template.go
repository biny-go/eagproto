package add

import (
	"bytes"
	"strings"
	"text/template"
)

const protoTemplate = `
syntax = "proto3";

package {{.Package}};

option go_package = "{{.GoPackage}}";

import "google/api/annotations.proto";
import "common/common.proto"; // --proto_path=eagproto路径/common/common.proto  proto文件包名+结构体：common.BaseResult

service {{.Service}} {
	rpc Save{{.Service}} (Save{{.Service}}Request) returns (common.BaseResult) {
		option (google.api.http) = {
			post: "/{{.ServicePath}}"
			body: "*"
		};
	};
	rpc Delete{{.Service}} (Delete{{.Service}}Request) returns (common.BaseResult) {
		option (google.api.http) = {
			delete: "/{{.ServicePath}}"
		};
	};
	rpc Get{{.Service}} (Get{{.Service}}Request) returns (common.BaseResult) {
		option (google.api.http) = {
			get: "/{{.ServicePath}}"
		};
	};
	rpc List{{.Service}} (List{{.Service}}Request) returns (common.BaseResultArray) {
		option (google.api.http) = {
			post: "/{{.ServicePath}}/list"
			body: "*"
		};
	};
}

message Save{{.Service}}Request {}

message Delete{{.Service}}Request {}

message Get{{.Service}}Request {}

message List{{.Service}}Request {}

message {{.Service}}POList {
	repeated {{.Service}}PO items = 1;
}

message {{.Service}}PO {
	string SYS_ID = 1;
}
`

func (p *Proto) execute() ([]byte, error) {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("proto").Parse(strings.TrimSpace(protoTemplate))
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
