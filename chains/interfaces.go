// SPDX-License-Identifier: LGPL-3.0-only

package chains

import (
	"github.com/NftTeamHackaton/bridge-utils/msg"
)

type Router interface {
	Send(message msg.Message) error
}

//type Writer interface {
//	ResolveMessage(message msg.Message) bool
//}
