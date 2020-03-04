package ipfs_proxy

import (
	"github.com/turt2live/matrix-media-repo/common/rcontext"
	"github.com/turt2live/matrix-media-repo/ipfs_proxy/ipfs_models"
)

type IPFSImplementation interface {
	GetObject(contentId string, ctx rcontext.RequestContext) (*ipfs_models.IPFSObject, error)
	Stop()
}
