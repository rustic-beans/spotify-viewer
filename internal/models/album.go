package models

import (
	"fmt"

	"github.com/rustic-beans/spotify-viewer/ent"
	"github.com/rustic-beans/spotify-viewer/ent/album"
)

type Album = ent.Album
type CreateAlbumInput = ent.CreateAlbumInput

type AlbumType = album.AlbumType

const (
	AlbumTypeAlbum       = album.AlbumTypeAlbum
	AlbumTypeCompilation = album.AlbumTypeCompilation
	AlbumTypeSingle      = album.AlbumTypeSingle
)

var albumMap = map[string]AlbumType{
	"album":       AlbumTypeAlbum,
	"single":      AlbumTypeSingle,
	"compilation": AlbumTypeCompilation,
}

func StringToAlbumType(s string) (albumType AlbumType, _ error) {
	albumType, ok := albumMap[s]
	if !ok {
		return AlbumTypeAlbum, fmt.Errorf("invalid album type: %s", s)
	}

	return albumType, nil
}

type AlbumReleaseDatePrecision = album.ReleaseDatePrecision

const (
	ReleaseDatePrecisionDay   = album.ReleaseDatePrecisionDay
	ReleaseDatePrecisionMonth = album.ReleaseDatePrecisionMonth
	ReleaseDatePrecisionYear  = album.ReleaseDatePrecisionYear
)

var releaseDatePrecisionMap = map[string]AlbumReleaseDatePrecision{
	"day":   ReleaseDatePrecisionDay,
	"month": ReleaseDatePrecisionMonth,
	"year":  ReleaseDatePrecisionYear,
}

func StringToAlbumReleaseDatePrecision(s string) (releaseDatePrecision AlbumReleaseDatePrecision, _ error) {
	releaseDatePrecision, ok := releaseDatePrecisionMap[s]
	if !ok {
		return ReleaseDatePrecisionDay, fmt.Errorf("invalid album release date precision: %s", s)
	}

	return releaseDatePrecision, nil
}
