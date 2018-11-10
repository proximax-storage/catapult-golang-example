package catapult_config

import (
	"github.com/proximax-storage/nem2-sdk-go/sdk"
	"math/big"
)

const (
	CatapultUrl = "http://privatetest1.proximax.io:3000"
	NetworkType = sdk.MijinTest

	// Sender Account (XPX)
	SenderPrivateKey = "d6c89092e210501c9c2c91a9c173b51a07b3b74144715145d3761ff8a15d77ee"
	SenderPublicKey  = "bd199f4739b033c06655b6a54c4daa2caa296a0f1bf8841a328b85ea04c3c117"
	SenderAddress    = "SCV5LMNUL7QDDMMYYU6BNHK3OD6GF7KWRU63CGLJ"

	// Recipient Account (XPX)
	RecipientPrivateKey = "de75f17634e74afe8968e09a227bc9002c3ac4802d9799d7507e5286f3cd3f93"
	RecipientPublicKey  = "41ee278c043f23ac007772a3d95f1e4795b73e243d90c48b86ff2dd3f17f0db8"
	RecipientAddress    = "SBMN7CCXQLQMAGZ3LFNHIHAXKAY5Z4FZ46RUA5KU"
)

// XPX helper functions
var XpxMosaicId, _ = sdk.NewMosaicIdFromName("prx:xpx")

// Create xem with using xem as unit
func Xpx(amount int64) *sdk.Mosaic {
	return &sdk.Mosaic{XpxMosaicId, big.NewInt(amount)}
}

func XpxRelative(amount int64) *sdk.Mosaic {
	return Xpx(big.NewInt(0).Mul(big.NewInt(1000000), big.NewInt(amount)).Int64())
}
