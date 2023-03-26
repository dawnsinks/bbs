package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

func GenId() int64 {

	node, _ := snowflake.NewNode(1)

	return node.Generate().Int64()
}
