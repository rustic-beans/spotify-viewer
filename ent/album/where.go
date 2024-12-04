// Code generated by ent, DO NOT EDIT.

package album

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rustic-beans/spotify-viewer/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Album {
	return predicate.Album(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Album {
	return predicate.Album(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Album {
	return predicate.Album(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Album {
	return predicate.Album(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Album {
	return predicate.Album(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Album {
	return predicate.Album(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Album {
	return predicate.Album(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Album {
	return predicate.Album(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Album {
	return predicate.Album(sql.FieldContainsFold(FieldID, id))
}

// TotalTracks applies equality check predicate on the "total_tracks" field. It's identical to TotalTracksEQ.
func TotalTracks(v int) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldTotalTracks, v))
}

// Href applies equality check predicate on the "href" field. It's identical to HrefEQ.
func Href(v string) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldHref, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldName, v))
}

// ReleaseDate applies equality check predicate on the "release_date" field. It's identical to ReleaseDateEQ.
func ReleaseDate(v string) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldReleaseDate, v))
}

// URI applies equality check predicate on the "uri" field. It's identical to URIEQ.
func URI(v string) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldURI, v))
}

// AlbumTypeEQ applies the EQ predicate on the "album_type" field.
func AlbumTypeEQ(v AlbumType) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldAlbumType, v))
}

// AlbumTypeNEQ applies the NEQ predicate on the "album_type" field.
func AlbumTypeNEQ(v AlbumType) predicate.Album {
	return predicate.Album(sql.FieldNEQ(FieldAlbumType, v))
}

// AlbumTypeIn applies the In predicate on the "album_type" field.
func AlbumTypeIn(vs ...AlbumType) predicate.Album {
	return predicate.Album(sql.FieldIn(FieldAlbumType, vs...))
}

// AlbumTypeNotIn applies the NotIn predicate on the "album_type" field.
func AlbumTypeNotIn(vs ...AlbumType) predicate.Album {
	return predicate.Album(sql.FieldNotIn(FieldAlbumType, vs...))
}

// TotalTracksEQ applies the EQ predicate on the "total_tracks" field.
func TotalTracksEQ(v int) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldTotalTracks, v))
}

// TotalTracksNEQ applies the NEQ predicate on the "total_tracks" field.
func TotalTracksNEQ(v int) predicate.Album {
	return predicate.Album(sql.FieldNEQ(FieldTotalTracks, v))
}

// TotalTracksIn applies the In predicate on the "total_tracks" field.
func TotalTracksIn(vs ...int) predicate.Album {
	return predicate.Album(sql.FieldIn(FieldTotalTracks, vs...))
}

// TotalTracksNotIn applies the NotIn predicate on the "total_tracks" field.
func TotalTracksNotIn(vs ...int) predicate.Album {
	return predicate.Album(sql.FieldNotIn(FieldTotalTracks, vs...))
}

// TotalTracksGT applies the GT predicate on the "total_tracks" field.
func TotalTracksGT(v int) predicate.Album {
	return predicate.Album(sql.FieldGT(FieldTotalTracks, v))
}

// TotalTracksGTE applies the GTE predicate on the "total_tracks" field.
func TotalTracksGTE(v int) predicate.Album {
	return predicate.Album(sql.FieldGTE(FieldTotalTracks, v))
}

// TotalTracksLT applies the LT predicate on the "total_tracks" field.
func TotalTracksLT(v int) predicate.Album {
	return predicate.Album(sql.FieldLT(FieldTotalTracks, v))
}

// TotalTracksLTE applies the LTE predicate on the "total_tracks" field.
func TotalTracksLTE(v int) predicate.Album {
	return predicate.Album(sql.FieldLTE(FieldTotalTracks, v))
}

// HrefEQ applies the EQ predicate on the "href" field.
func HrefEQ(v string) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldHref, v))
}

// HrefNEQ applies the NEQ predicate on the "href" field.
func HrefNEQ(v string) predicate.Album {
	return predicate.Album(sql.FieldNEQ(FieldHref, v))
}

// HrefIn applies the In predicate on the "href" field.
func HrefIn(vs ...string) predicate.Album {
	return predicate.Album(sql.FieldIn(FieldHref, vs...))
}

// HrefNotIn applies the NotIn predicate on the "href" field.
func HrefNotIn(vs ...string) predicate.Album {
	return predicate.Album(sql.FieldNotIn(FieldHref, vs...))
}

// HrefGT applies the GT predicate on the "href" field.
func HrefGT(v string) predicate.Album {
	return predicate.Album(sql.FieldGT(FieldHref, v))
}

// HrefGTE applies the GTE predicate on the "href" field.
func HrefGTE(v string) predicate.Album {
	return predicate.Album(sql.FieldGTE(FieldHref, v))
}

// HrefLT applies the LT predicate on the "href" field.
func HrefLT(v string) predicate.Album {
	return predicate.Album(sql.FieldLT(FieldHref, v))
}

// HrefLTE applies the LTE predicate on the "href" field.
func HrefLTE(v string) predicate.Album {
	return predicate.Album(sql.FieldLTE(FieldHref, v))
}

// HrefContains applies the Contains predicate on the "href" field.
func HrefContains(v string) predicate.Album {
	return predicate.Album(sql.FieldContains(FieldHref, v))
}

// HrefHasPrefix applies the HasPrefix predicate on the "href" field.
func HrefHasPrefix(v string) predicate.Album {
	return predicate.Album(sql.FieldHasPrefix(FieldHref, v))
}

// HrefHasSuffix applies the HasSuffix predicate on the "href" field.
func HrefHasSuffix(v string) predicate.Album {
	return predicate.Album(sql.FieldHasSuffix(FieldHref, v))
}

// HrefEqualFold applies the EqualFold predicate on the "href" field.
func HrefEqualFold(v string) predicate.Album {
	return predicate.Album(sql.FieldEqualFold(FieldHref, v))
}

// HrefContainsFold applies the ContainsFold predicate on the "href" field.
func HrefContainsFold(v string) predicate.Album {
	return predicate.Album(sql.FieldContainsFold(FieldHref, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Album {
	return predicate.Album(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Album {
	return predicate.Album(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Album {
	return predicate.Album(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Album {
	return predicate.Album(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Album {
	return predicate.Album(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Album {
	return predicate.Album(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Album {
	return predicate.Album(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Album {
	return predicate.Album(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Album {
	return predicate.Album(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Album {
	return predicate.Album(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Album {
	return predicate.Album(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Album {
	return predicate.Album(sql.FieldContainsFold(FieldName, v))
}

// ReleaseDateEQ applies the EQ predicate on the "release_date" field.
func ReleaseDateEQ(v string) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldReleaseDate, v))
}

// ReleaseDateNEQ applies the NEQ predicate on the "release_date" field.
func ReleaseDateNEQ(v string) predicate.Album {
	return predicate.Album(sql.FieldNEQ(FieldReleaseDate, v))
}

// ReleaseDateIn applies the In predicate on the "release_date" field.
func ReleaseDateIn(vs ...string) predicate.Album {
	return predicate.Album(sql.FieldIn(FieldReleaseDate, vs...))
}

// ReleaseDateNotIn applies the NotIn predicate on the "release_date" field.
func ReleaseDateNotIn(vs ...string) predicate.Album {
	return predicate.Album(sql.FieldNotIn(FieldReleaseDate, vs...))
}

// ReleaseDateGT applies the GT predicate on the "release_date" field.
func ReleaseDateGT(v string) predicate.Album {
	return predicate.Album(sql.FieldGT(FieldReleaseDate, v))
}

// ReleaseDateGTE applies the GTE predicate on the "release_date" field.
func ReleaseDateGTE(v string) predicate.Album {
	return predicate.Album(sql.FieldGTE(FieldReleaseDate, v))
}

// ReleaseDateLT applies the LT predicate on the "release_date" field.
func ReleaseDateLT(v string) predicate.Album {
	return predicate.Album(sql.FieldLT(FieldReleaseDate, v))
}

// ReleaseDateLTE applies the LTE predicate on the "release_date" field.
func ReleaseDateLTE(v string) predicate.Album {
	return predicate.Album(sql.FieldLTE(FieldReleaseDate, v))
}

// ReleaseDateContains applies the Contains predicate on the "release_date" field.
func ReleaseDateContains(v string) predicate.Album {
	return predicate.Album(sql.FieldContains(FieldReleaseDate, v))
}

// ReleaseDateHasPrefix applies the HasPrefix predicate on the "release_date" field.
func ReleaseDateHasPrefix(v string) predicate.Album {
	return predicate.Album(sql.FieldHasPrefix(FieldReleaseDate, v))
}

// ReleaseDateHasSuffix applies the HasSuffix predicate on the "release_date" field.
func ReleaseDateHasSuffix(v string) predicate.Album {
	return predicate.Album(sql.FieldHasSuffix(FieldReleaseDate, v))
}

// ReleaseDateEqualFold applies the EqualFold predicate on the "release_date" field.
func ReleaseDateEqualFold(v string) predicate.Album {
	return predicate.Album(sql.FieldEqualFold(FieldReleaseDate, v))
}

// ReleaseDateContainsFold applies the ContainsFold predicate on the "release_date" field.
func ReleaseDateContainsFold(v string) predicate.Album {
	return predicate.Album(sql.FieldContainsFold(FieldReleaseDate, v))
}

// ReleaseDatePrecisionEQ applies the EQ predicate on the "release_date_precision" field.
func ReleaseDatePrecisionEQ(v ReleaseDatePrecision) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldReleaseDatePrecision, v))
}

// ReleaseDatePrecisionNEQ applies the NEQ predicate on the "release_date_precision" field.
func ReleaseDatePrecisionNEQ(v ReleaseDatePrecision) predicate.Album {
	return predicate.Album(sql.FieldNEQ(FieldReleaseDatePrecision, v))
}

// ReleaseDatePrecisionIn applies the In predicate on the "release_date_precision" field.
func ReleaseDatePrecisionIn(vs ...ReleaseDatePrecision) predicate.Album {
	return predicate.Album(sql.FieldIn(FieldReleaseDatePrecision, vs...))
}

// ReleaseDatePrecisionNotIn applies the NotIn predicate on the "release_date_precision" field.
func ReleaseDatePrecisionNotIn(vs ...ReleaseDatePrecision) predicate.Album {
	return predicate.Album(sql.FieldNotIn(FieldReleaseDatePrecision, vs...))
}

// URIEQ applies the EQ predicate on the "uri" field.
func URIEQ(v string) predicate.Album {
	return predicate.Album(sql.FieldEQ(FieldURI, v))
}

// URINEQ applies the NEQ predicate on the "uri" field.
func URINEQ(v string) predicate.Album {
	return predicate.Album(sql.FieldNEQ(FieldURI, v))
}

// URIIn applies the In predicate on the "uri" field.
func URIIn(vs ...string) predicate.Album {
	return predicate.Album(sql.FieldIn(FieldURI, vs...))
}

// URINotIn applies the NotIn predicate on the "uri" field.
func URINotIn(vs ...string) predicate.Album {
	return predicate.Album(sql.FieldNotIn(FieldURI, vs...))
}

// URIGT applies the GT predicate on the "uri" field.
func URIGT(v string) predicate.Album {
	return predicate.Album(sql.FieldGT(FieldURI, v))
}

// URIGTE applies the GTE predicate on the "uri" field.
func URIGTE(v string) predicate.Album {
	return predicate.Album(sql.FieldGTE(FieldURI, v))
}

// URILT applies the LT predicate on the "uri" field.
func URILT(v string) predicate.Album {
	return predicate.Album(sql.FieldLT(FieldURI, v))
}

// URILTE applies the LTE predicate on the "uri" field.
func URILTE(v string) predicate.Album {
	return predicate.Album(sql.FieldLTE(FieldURI, v))
}

// URIContains applies the Contains predicate on the "uri" field.
func URIContains(v string) predicate.Album {
	return predicate.Album(sql.FieldContains(FieldURI, v))
}

// URIHasPrefix applies the HasPrefix predicate on the "uri" field.
func URIHasPrefix(v string) predicate.Album {
	return predicate.Album(sql.FieldHasPrefix(FieldURI, v))
}

// URIHasSuffix applies the HasSuffix predicate on the "uri" field.
func URIHasSuffix(v string) predicate.Album {
	return predicate.Album(sql.FieldHasSuffix(FieldURI, v))
}

// URIEqualFold applies the EqualFold predicate on the "uri" field.
func URIEqualFold(v string) predicate.Album {
	return predicate.Album(sql.FieldEqualFold(FieldURI, v))
}

// URIContainsFold applies the ContainsFold predicate on the "uri" field.
func URIContainsFold(v string) predicate.Album {
	return predicate.Album(sql.FieldContainsFold(FieldURI, v))
}

// HasImages applies the HasEdge predicate on the "images" edge.
func HasImages() predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ImagesTable, ImagesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImagesWith applies the HasEdge predicate on the "images" edge with a given conditions (other predicates).
func HasImagesWith(preds ...predicate.Image) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		step := newImagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasArtists applies the HasEdge predicate on the "artists" edge.
func HasArtists() predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ArtistsTable, ArtistsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArtistsWith applies the HasEdge predicate on the "artists" edge with a given conditions (other predicates).
func HasArtistsWith(preds ...predicate.Artist) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		step := newArtistsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTracks applies the HasEdge predicate on the "tracks" edge.
func HasTracks() predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TracksTable, TracksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTracksWith applies the HasEdge predicate on the "tracks" edge with a given conditions (other predicates).
func HasTracksWith(preds ...predicate.Track) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		step := newTracksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Album) predicate.Album {
	return predicate.Album(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Album) predicate.Album {
	return predicate.Album(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Album) predicate.Album {
	return predicate.Album(sql.NotPredicates(p))
}
