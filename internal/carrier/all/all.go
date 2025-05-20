/*
©cyber-atharv | 2026
all.go

Side-effect imports that register every carrier implementation into the registry
*/

package all

import (
	_ "github.com/cyber-atharv/crypha/internal/carrier/audio"
	_ "github.com/cyber-atharv/crypha/internal/carrier/image"
	_ "github.com/cyber-atharv/crypha/internal/carrier/pdf"
	_ "github.com/cyber-atharv/crypha/internal/carrier/qr"
	_ "github.com/cyber-atharv/crypha/internal/carrier/text"
)
