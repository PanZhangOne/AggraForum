package avatar

import (
	"forum/conf"
	"github.com/issue9/identicon"
	"image/color"
	"image/png"
	"os"
)

func GenerateAvatarFromUsername(username string) {
	img, _ := identicon.Make(128, color.RGBA{255, 0, 0, 100}, color.RGBA{0, 255, 255, 100}, []byte(username))
	fi, _ := os.Create(conf.SystemConfig.UserAvatarPath + username + ".png")
	_ = png.Encode(fi, img)
	_ = fi.Close()
}
