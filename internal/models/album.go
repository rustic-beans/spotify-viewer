package models

import (
	"github.com/cockroachdb/errors"
	"github.com/rustic-beans/spotify-viewer/internal/database"
)

type Album = database.Album
type CreateAlbumParams = database.CreateAlbumParams
type AlbumType = database.AlbumType

const (
	AlbumTypeAlbum       = database.AlbumTypeAlbum
	AlbumTypeCompilation = database.AlbumTypeCompilation
	AlbumTypeSingle      = database.AlbumTypeSingle
)

var albumMap = map[string]AlbumType{
	"album":       AlbumTypeAlbum,
	"single":      AlbumTypeSingle,
	"compilation": AlbumTypeCompilation,
}

func StringToAlbumType(s string) (albumType AlbumType, _ error) {
	albumType, ok := albumMap[s]
	if !ok {
		return AlbumTypeAlbum, errors.Newf("invalid album type: %s", s)
	}

	return albumType, nil
}

type AlbumReleaseDatePrecision = database.AlbumReleaseDatePrecision

const (
	ReleaseDatePrecisionDay   = database.AlbumReleaseDatePrecisionDay
	ReleaseDatePrecisionMonth = database.AlbumReleaseDatePrecisionMonth
	ReleaseDatePrecisionYear  = database.AlbumReleaseDatePrecisionYear
)

var releaseDatePrecisionMap = map[string]AlbumReleaseDatePrecision{
	"day":   ReleaseDatePrecisionDay,
	"month": ReleaseDatePrecisionMonth,
	"year":  ReleaseDatePrecisionYear,
}

func StringToAlbumReleaseDatePrecision(s string) (releaseDatePrecision AlbumReleaseDatePrecision, _ error) {
	releaseDatePrecision, ok := releaseDatePrecisionMap[s]
	if !ok {
		return ReleaseDatePrecisionDay, errors.Newf("invalid album release date precision: %s", s)
	}

	return releaseDatePrecision, nil
}
