// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/rustic-beans/spotify-viewer/ent/album"
	"github.com/rustic-beans/spotify-viewer/ent/artist"
	"github.com/rustic-beans/spotify-viewer/ent/image"
	"github.com/rustic-beans/spotify-viewer/ent/predicate"
	"github.com/rustic-beans/spotify-viewer/ent/schema"
	"github.com/rustic-beans/spotify-viewer/ent/track"
)

// AlbumUpdate is the builder for updating Album entities.
type AlbumUpdate struct {
	config
	hooks    []Hook
	mutation *AlbumMutation
}

// Where appends a list predicates to the AlbumUpdate builder.
func (au *AlbumUpdate) Where(ps ...predicate.Album) *AlbumUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetAlbumType sets the "album_type" field.
func (au *AlbumUpdate) SetAlbumType(at album.AlbumType) *AlbumUpdate {
	au.mutation.SetAlbumType(at)
	return au
}

// SetNillableAlbumType sets the "album_type" field if the given value is not nil.
func (au *AlbumUpdate) SetNillableAlbumType(at *album.AlbumType) *AlbumUpdate {
	if at != nil {
		au.SetAlbumType(*at)
	}
	return au
}

// SetTotalTracks sets the "total_tracks" field.
func (au *AlbumUpdate) SetTotalTracks(i int) *AlbumUpdate {
	au.mutation.ResetTotalTracks()
	au.mutation.SetTotalTracks(i)
	return au
}

// SetNillableTotalTracks sets the "total_tracks" field if the given value is not nil.
func (au *AlbumUpdate) SetNillableTotalTracks(i *int) *AlbumUpdate {
	if i != nil {
		au.SetTotalTracks(*i)
	}
	return au
}

// AddTotalTracks adds i to the "total_tracks" field.
func (au *AlbumUpdate) AddTotalTracks(i int) *AlbumUpdate {
	au.mutation.AddTotalTracks(i)
	return au
}

// SetExternalUrls sets the "external_urls" field.
func (au *AlbumUpdate) SetExternalUrls(sm *schema.StringMap) *AlbumUpdate {
	au.mutation.SetExternalUrls(sm)
	return au
}

// SetHref sets the "href" field.
func (au *AlbumUpdate) SetHref(s string) *AlbumUpdate {
	au.mutation.SetHref(s)
	return au
}

// SetNillableHref sets the "href" field if the given value is not nil.
func (au *AlbumUpdate) SetNillableHref(s *string) *AlbumUpdate {
	if s != nil {
		au.SetHref(*s)
	}
	return au
}

// SetName sets the "name" field.
func (au *AlbumUpdate) SetName(s string) *AlbumUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AlbumUpdate) SetNillableName(s *string) *AlbumUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetReleaseDate sets the "release_date" field.
func (au *AlbumUpdate) SetReleaseDate(s string) *AlbumUpdate {
	au.mutation.SetReleaseDate(s)
	return au
}

// SetNillableReleaseDate sets the "release_date" field if the given value is not nil.
func (au *AlbumUpdate) SetNillableReleaseDate(s *string) *AlbumUpdate {
	if s != nil {
		au.SetReleaseDate(*s)
	}
	return au
}

// SetReleaseDatePrecision sets the "release_date_precision" field.
func (au *AlbumUpdate) SetReleaseDatePrecision(adp album.ReleaseDatePrecision) *AlbumUpdate {
	au.mutation.SetReleaseDatePrecision(adp)
	return au
}

// SetNillableReleaseDatePrecision sets the "release_date_precision" field if the given value is not nil.
func (au *AlbumUpdate) SetNillableReleaseDatePrecision(adp *album.ReleaseDatePrecision) *AlbumUpdate {
	if adp != nil {
		au.SetReleaseDatePrecision(*adp)
	}
	return au
}

// SetURI sets the "uri" field.
func (au *AlbumUpdate) SetURI(s string) *AlbumUpdate {
	au.mutation.SetURI(s)
	return au
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (au *AlbumUpdate) SetNillableURI(s *string) *AlbumUpdate {
	if s != nil {
		au.SetURI(*s)
	}
	return au
}

// SetGenres sets the "genres" field.
func (au *AlbumUpdate) SetGenres(s []string) *AlbumUpdate {
	au.mutation.SetGenres(s)
	return au
}

// AppendGenres appends s to the "genres" field.
func (au *AlbumUpdate) AppendGenres(s []string) *AlbumUpdate {
	au.mutation.AppendGenres(s)
	return au
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (au *AlbumUpdate) AddImageIDs(ids ...string) *AlbumUpdate {
	au.mutation.AddImageIDs(ids...)
	return au
}

// AddImages adds the "images" edges to the Image entity.
func (au *AlbumUpdate) AddImages(i ...*Image) *AlbumUpdate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return au.AddImageIDs(ids...)
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (au *AlbumUpdate) AddArtistIDs(ids ...string) *AlbumUpdate {
	au.mutation.AddArtistIDs(ids...)
	return au
}

// AddArtists adds the "artists" edges to the Artist entity.
func (au *AlbumUpdate) AddArtists(a ...*Artist) *AlbumUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddArtistIDs(ids...)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (au *AlbumUpdate) AddTrackIDs(ids ...string) *AlbumUpdate {
	au.mutation.AddTrackIDs(ids...)
	return au
}

// AddTracks adds the "tracks" edges to the Track entity.
func (au *AlbumUpdate) AddTracks(t ...*Track) *AlbumUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTrackIDs(ids...)
}

// Mutation returns the AlbumMutation object of the builder.
func (au *AlbumUpdate) Mutation() *AlbumMutation {
	return au.mutation
}

// ClearImages clears all "images" edges to the Image entity.
func (au *AlbumUpdate) ClearImages() *AlbumUpdate {
	au.mutation.ClearImages()
	return au
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (au *AlbumUpdate) RemoveImageIDs(ids ...string) *AlbumUpdate {
	au.mutation.RemoveImageIDs(ids...)
	return au
}

// RemoveImages removes "images" edges to Image entities.
func (au *AlbumUpdate) RemoveImages(i ...*Image) *AlbumUpdate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return au.RemoveImageIDs(ids...)
}

// ClearArtists clears all "artists" edges to the Artist entity.
func (au *AlbumUpdate) ClearArtists() *AlbumUpdate {
	au.mutation.ClearArtists()
	return au
}

// RemoveArtistIDs removes the "artists" edge to Artist entities by IDs.
func (au *AlbumUpdate) RemoveArtistIDs(ids ...string) *AlbumUpdate {
	au.mutation.RemoveArtistIDs(ids...)
	return au
}

// RemoveArtists removes "artists" edges to Artist entities.
func (au *AlbumUpdate) RemoveArtists(a ...*Artist) *AlbumUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveArtistIDs(ids...)
}

// ClearTracks clears all "tracks" edges to the Track entity.
func (au *AlbumUpdate) ClearTracks() *AlbumUpdate {
	au.mutation.ClearTracks()
	return au
}

// RemoveTrackIDs removes the "tracks" edge to Track entities by IDs.
func (au *AlbumUpdate) RemoveTrackIDs(ids ...string) *AlbumUpdate {
	au.mutation.RemoveTrackIDs(ids...)
	return au
}

// RemoveTracks removes "tracks" edges to Track entities.
func (au *AlbumUpdate) RemoveTracks(t ...*Track) *AlbumUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTrackIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AlbumUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AlbumUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AlbumUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AlbumUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AlbumUpdate) check() error {
	if v, ok := au.mutation.AlbumType(); ok {
		if err := album.AlbumTypeValidator(v); err != nil {
			return &ValidationError{Name: "album_type", err: fmt.Errorf(`ent: validator failed for field "Album.album_type": %w`, err)}
		}
	}
	if v, ok := au.mutation.Href(); ok {
		if err := album.HrefValidator(v); err != nil {
			return &ValidationError{Name: "href", err: fmt.Errorf(`ent: validator failed for field "Album.href": %w`, err)}
		}
	}
	if v, ok := au.mutation.Name(); ok {
		if err := album.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Album.name": %w`, err)}
		}
	}
	if v, ok := au.mutation.ReleaseDate(); ok {
		if err := album.ReleaseDateValidator(v); err != nil {
			return &ValidationError{Name: "release_date", err: fmt.Errorf(`ent: validator failed for field "Album.release_date": %w`, err)}
		}
	}
	if v, ok := au.mutation.ReleaseDatePrecision(); ok {
		if err := album.ReleaseDatePrecisionValidator(v); err != nil {
			return &ValidationError{Name: "release_date_precision", err: fmt.Errorf(`ent: validator failed for field "Album.release_date_precision": %w`, err)}
		}
	}
	if v, ok := au.mutation.URI(); ok {
		if err := album.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Album.uri": %w`, err)}
		}
	}
	return nil
}

func (au *AlbumUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(album.Table, album.Columns, sqlgraph.NewFieldSpec(album.FieldID, field.TypeString))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.AlbumType(); ok {
		_spec.SetField(album.FieldAlbumType, field.TypeEnum, value)
	}
	if value, ok := au.mutation.TotalTracks(); ok {
		_spec.SetField(album.FieldTotalTracks, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedTotalTracks(); ok {
		_spec.AddField(album.FieldTotalTracks, field.TypeInt, value)
	}
	if value, ok := au.mutation.ExternalUrls(); ok {
		_spec.SetField(album.FieldExternalUrls, field.TypeJSON, value)
	}
	if value, ok := au.mutation.Href(); ok {
		_spec.SetField(album.FieldHref, field.TypeString, value)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(album.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.ReleaseDate(); ok {
		_spec.SetField(album.FieldReleaseDate, field.TypeString, value)
	}
	if value, ok := au.mutation.ReleaseDatePrecision(); ok {
		_spec.SetField(album.FieldReleaseDatePrecision, field.TypeEnum, value)
	}
	if value, ok := au.mutation.URI(); ok {
		_spec.SetField(album.FieldURI, field.TypeString, value)
	}
	if value, ok := au.mutation.Genres(); ok {
		_spec.SetField(album.FieldGenres, field.TypeJSON, value)
	}
	if value, ok := au.mutation.AppendedGenres(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, album.FieldGenres, value)
		})
	}
	if au.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   album.ImagesTable,
			Columns: album.ImagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedImagesIDs(); len(nodes) > 0 && !au.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   album.ImagesTable,
			Columns: album.ImagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   album.ImagesTable,
			Columns: album.ImagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.ArtistsTable,
			Columns: album.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedArtistsIDs(); len(nodes) > 0 && !au.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.ArtistsTable,
			Columns: album.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.ArtistsTable,
			Columns: album.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   album.TracksTable,
			Columns: []string{album.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTracksIDs(); len(nodes) > 0 && !au.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   album.TracksTable,
			Columns: []string{album.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   album.TracksTable,
			Columns: []string{album.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{album.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AlbumUpdateOne is the builder for updating a single Album entity.
type AlbumUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AlbumMutation
}

// SetAlbumType sets the "album_type" field.
func (auo *AlbumUpdateOne) SetAlbumType(at album.AlbumType) *AlbumUpdateOne {
	auo.mutation.SetAlbumType(at)
	return auo
}

// SetNillableAlbumType sets the "album_type" field if the given value is not nil.
func (auo *AlbumUpdateOne) SetNillableAlbumType(at *album.AlbumType) *AlbumUpdateOne {
	if at != nil {
		auo.SetAlbumType(*at)
	}
	return auo
}

// SetTotalTracks sets the "total_tracks" field.
func (auo *AlbumUpdateOne) SetTotalTracks(i int) *AlbumUpdateOne {
	auo.mutation.ResetTotalTracks()
	auo.mutation.SetTotalTracks(i)
	return auo
}

// SetNillableTotalTracks sets the "total_tracks" field if the given value is not nil.
func (auo *AlbumUpdateOne) SetNillableTotalTracks(i *int) *AlbumUpdateOne {
	if i != nil {
		auo.SetTotalTracks(*i)
	}
	return auo
}

// AddTotalTracks adds i to the "total_tracks" field.
func (auo *AlbumUpdateOne) AddTotalTracks(i int) *AlbumUpdateOne {
	auo.mutation.AddTotalTracks(i)
	return auo
}

// SetExternalUrls sets the "external_urls" field.
func (auo *AlbumUpdateOne) SetExternalUrls(sm *schema.StringMap) *AlbumUpdateOne {
	auo.mutation.SetExternalUrls(sm)
	return auo
}

// SetHref sets the "href" field.
func (auo *AlbumUpdateOne) SetHref(s string) *AlbumUpdateOne {
	auo.mutation.SetHref(s)
	return auo
}

// SetNillableHref sets the "href" field if the given value is not nil.
func (auo *AlbumUpdateOne) SetNillableHref(s *string) *AlbumUpdateOne {
	if s != nil {
		auo.SetHref(*s)
	}
	return auo
}

// SetName sets the "name" field.
func (auo *AlbumUpdateOne) SetName(s string) *AlbumUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AlbumUpdateOne) SetNillableName(s *string) *AlbumUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetReleaseDate sets the "release_date" field.
func (auo *AlbumUpdateOne) SetReleaseDate(s string) *AlbumUpdateOne {
	auo.mutation.SetReleaseDate(s)
	return auo
}

// SetNillableReleaseDate sets the "release_date" field if the given value is not nil.
func (auo *AlbumUpdateOne) SetNillableReleaseDate(s *string) *AlbumUpdateOne {
	if s != nil {
		auo.SetReleaseDate(*s)
	}
	return auo
}

// SetReleaseDatePrecision sets the "release_date_precision" field.
func (auo *AlbumUpdateOne) SetReleaseDatePrecision(adp album.ReleaseDatePrecision) *AlbumUpdateOne {
	auo.mutation.SetReleaseDatePrecision(adp)
	return auo
}

// SetNillableReleaseDatePrecision sets the "release_date_precision" field if the given value is not nil.
func (auo *AlbumUpdateOne) SetNillableReleaseDatePrecision(adp *album.ReleaseDatePrecision) *AlbumUpdateOne {
	if adp != nil {
		auo.SetReleaseDatePrecision(*adp)
	}
	return auo
}

// SetURI sets the "uri" field.
func (auo *AlbumUpdateOne) SetURI(s string) *AlbumUpdateOne {
	auo.mutation.SetURI(s)
	return auo
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (auo *AlbumUpdateOne) SetNillableURI(s *string) *AlbumUpdateOne {
	if s != nil {
		auo.SetURI(*s)
	}
	return auo
}

// SetGenres sets the "genres" field.
func (auo *AlbumUpdateOne) SetGenres(s []string) *AlbumUpdateOne {
	auo.mutation.SetGenres(s)
	return auo
}

// AppendGenres appends s to the "genres" field.
func (auo *AlbumUpdateOne) AppendGenres(s []string) *AlbumUpdateOne {
	auo.mutation.AppendGenres(s)
	return auo
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (auo *AlbumUpdateOne) AddImageIDs(ids ...string) *AlbumUpdateOne {
	auo.mutation.AddImageIDs(ids...)
	return auo
}

// AddImages adds the "images" edges to the Image entity.
func (auo *AlbumUpdateOne) AddImages(i ...*Image) *AlbumUpdateOne {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return auo.AddImageIDs(ids...)
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (auo *AlbumUpdateOne) AddArtistIDs(ids ...string) *AlbumUpdateOne {
	auo.mutation.AddArtistIDs(ids...)
	return auo
}

// AddArtists adds the "artists" edges to the Artist entity.
func (auo *AlbumUpdateOne) AddArtists(a ...*Artist) *AlbumUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddArtistIDs(ids...)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (auo *AlbumUpdateOne) AddTrackIDs(ids ...string) *AlbumUpdateOne {
	auo.mutation.AddTrackIDs(ids...)
	return auo
}

// AddTracks adds the "tracks" edges to the Track entity.
func (auo *AlbumUpdateOne) AddTracks(t ...*Track) *AlbumUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTrackIDs(ids...)
}

// Mutation returns the AlbumMutation object of the builder.
func (auo *AlbumUpdateOne) Mutation() *AlbumMutation {
	return auo.mutation
}

// ClearImages clears all "images" edges to the Image entity.
func (auo *AlbumUpdateOne) ClearImages() *AlbumUpdateOne {
	auo.mutation.ClearImages()
	return auo
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (auo *AlbumUpdateOne) RemoveImageIDs(ids ...string) *AlbumUpdateOne {
	auo.mutation.RemoveImageIDs(ids...)
	return auo
}

// RemoveImages removes "images" edges to Image entities.
func (auo *AlbumUpdateOne) RemoveImages(i ...*Image) *AlbumUpdateOne {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return auo.RemoveImageIDs(ids...)
}

// ClearArtists clears all "artists" edges to the Artist entity.
func (auo *AlbumUpdateOne) ClearArtists() *AlbumUpdateOne {
	auo.mutation.ClearArtists()
	return auo
}

// RemoveArtistIDs removes the "artists" edge to Artist entities by IDs.
func (auo *AlbumUpdateOne) RemoveArtistIDs(ids ...string) *AlbumUpdateOne {
	auo.mutation.RemoveArtistIDs(ids...)
	return auo
}

// RemoveArtists removes "artists" edges to Artist entities.
func (auo *AlbumUpdateOne) RemoveArtists(a ...*Artist) *AlbumUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveArtistIDs(ids...)
}

// ClearTracks clears all "tracks" edges to the Track entity.
func (auo *AlbumUpdateOne) ClearTracks() *AlbumUpdateOne {
	auo.mutation.ClearTracks()
	return auo
}

// RemoveTrackIDs removes the "tracks" edge to Track entities by IDs.
func (auo *AlbumUpdateOne) RemoveTrackIDs(ids ...string) *AlbumUpdateOne {
	auo.mutation.RemoveTrackIDs(ids...)
	return auo
}

// RemoveTracks removes "tracks" edges to Track entities.
func (auo *AlbumUpdateOne) RemoveTracks(t ...*Track) *AlbumUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTrackIDs(ids...)
}

// Where appends a list predicates to the AlbumUpdate builder.
func (auo *AlbumUpdateOne) Where(ps ...predicate.Album) *AlbumUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AlbumUpdateOne) Select(field string, fields ...string) *AlbumUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Album entity.
func (auo *AlbumUpdateOne) Save(ctx context.Context) (*Album, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AlbumUpdateOne) SaveX(ctx context.Context) *Album {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AlbumUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AlbumUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AlbumUpdateOne) check() error {
	if v, ok := auo.mutation.AlbumType(); ok {
		if err := album.AlbumTypeValidator(v); err != nil {
			return &ValidationError{Name: "album_type", err: fmt.Errorf(`ent: validator failed for field "Album.album_type": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Href(); ok {
		if err := album.HrefValidator(v); err != nil {
			return &ValidationError{Name: "href", err: fmt.Errorf(`ent: validator failed for field "Album.href": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Name(); ok {
		if err := album.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Album.name": %w`, err)}
		}
	}
	if v, ok := auo.mutation.ReleaseDate(); ok {
		if err := album.ReleaseDateValidator(v); err != nil {
			return &ValidationError{Name: "release_date", err: fmt.Errorf(`ent: validator failed for field "Album.release_date": %w`, err)}
		}
	}
	if v, ok := auo.mutation.ReleaseDatePrecision(); ok {
		if err := album.ReleaseDatePrecisionValidator(v); err != nil {
			return &ValidationError{Name: "release_date_precision", err: fmt.Errorf(`ent: validator failed for field "Album.release_date_precision": %w`, err)}
		}
	}
	if v, ok := auo.mutation.URI(); ok {
		if err := album.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Album.uri": %w`, err)}
		}
	}
	return nil
}

func (auo *AlbumUpdateOne) sqlSave(ctx context.Context) (_node *Album, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(album.Table, album.Columns, sqlgraph.NewFieldSpec(album.FieldID, field.TypeString))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Album.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, album.FieldID)
		for _, f := range fields {
			if !album.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != album.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.AlbumType(); ok {
		_spec.SetField(album.FieldAlbumType, field.TypeEnum, value)
	}
	if value, ok := auo.mutation.TotalTracks(); ok {
		_spec.SetField(album.FieldTotalTracks, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedTotalTracks(); ok {
		_spec.AddField(album.FieldTotalTracks, field.TypeInt, value)
	}
	if value, ok := auo.mutation.ExternalUrls(); ok {
		_spec.SetField(album.FieldExternalUrls, field.TypeJSON, value)
	}
	if value, ok := auo.mutation.Href(); ok {
		_spec.SetField(album.FieldHref, field.TypeString, value)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(album.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.ReleaseDate(); ok {
		_spec.SetField(album.FieldReleaseDate, field.TypeString, value)
	}
	if value, ok := auo.mutation.ReleaseDatePrecision(); ok {
		_spec.SetField(album.FieldReleaseDatePrecision, field.TypeEnum, value)
	}
	if value, ok := auo.mutation.URI(); ok {
		_spec.SetField(album.FieldURI, field.TypeString, value)
	}
	if value, ok := auo.mutation.Genres(); ok {
		_spec.SetField(album.FieldGenres, field.TypeJSON, value)
	}
	if value, ok := auo.mutation.AppendedGenres(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, album.FieldGenres, value)
		})
	}
	if auo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   album.ImagesTable,
			Columns: album.ImagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedImagesIDs(); len(nodes) > 0 && !auo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   album.ImagesTable,
			Columns: album.ImagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   album.ImagesTable,
			Columns: album.ImagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.ArtistsTable,
			Columns: album.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedArtistsIDs(); len(nodes) > 0 && !auo.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.ArtistsTable,
			Columns: album.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.ArtistsTable,
			Columns: album.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   album.TracksTable,
			Columns: []string{album.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTracksIDs(); len(nodes) > 0 && !auo.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   album.TracksTable,
			Columns: []string{album.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   album.TracksTable,
			Columns: []string{album.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Album{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{album.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}