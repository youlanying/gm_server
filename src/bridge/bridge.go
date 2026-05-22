package bridge

import (
	"github.com/golang/protobuf/proto"
	"gm_server/src/network/message"
	"strings"
)

const MSG_PACKAGE_NAME = "network.message."

func GetFSMessageData(pbData proto.Message) ([]byte, error) {
	msgName := proto.MessageName(pbData)
	cmd := strings.TrimPrefix(string(msgName), MSG_PACKAGE_NAME)

	body, err := proto.Marshal(pbData)
	if err != nil {
		return nil, err
	}

	fsMsg := &network_message.FSMessage{
		Head: cmd,
		Body: body,
	}

	return proto.Marshal(fsMsg)
}
