package proto

import (
	"github.com/spf13/cobra"

	"github.com/biny-go/eagproto/cmd/eagproto/internal/proto/add"
	"github.com/biny-go/eagproto/cmd/eagproto/internal/proto/client"
	"github.com/biny-go/eagproto/cmd/eagproto/internal/proto/server"
)

// CmdProto represents the proto command.
var CmdProto = &cobra.Command{
	Use:   "proto",
	Short: "Generate the proto files",
	Long:  "Generate the proto files.",
}

func init() {
	CmdProto.AddCommand(add.CmdAdd)
	CmdProto.AddCommand(client.CmdClient)
	CmdProto.AddCommand(server.CmdServer)
}
