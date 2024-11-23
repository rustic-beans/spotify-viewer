package spotify

import (
	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/zmb3/spotify/v2"
)

func ImageToModel(i *spotify.Image) *models.Image {
	return &models.Image{
		URL:    i.URL,
		Width:  int(i.Width),
		Height: int(i.Height),
	}
}
