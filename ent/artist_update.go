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

// ArtistUpdate is the builder for updating Artist entities.
type ArtistUpdate struct {
	config
	hooks    []Hook
	mutation *ArtistMutation
}

// Where appends a list predicates to the ArtistUpdate builder.
func (au *ArtistUpdate) Where(ps ...predicate.Artist) *ArtistUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetExternalUrls sets the "external_urls" field.
func (au *ArtistUpdate) SetExternalUrls(sm *schema.StringMap) *ArtistUpdate {
	au.mutation.SetExternalUrls(sm)
	return au
}

// SetHref sets the "href" field.
func (au *ArtistUpdate) SetHref(s string) *ArtistUpdate {
	au.mutation.SetHref(s)
	return au
}

// SetNillableHref sets the "href" field if the given value is not nil.
func (au *ArtistUpdate) SetNillableHref(s *string) *ArtistUpdate {
	if s != nil {
		au.SetHref(*s)
	}
	return au
}

// SetName sets the "name" field.
func (au *ArtistUpdate) SetName(s string) *ArtistUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *ArtistUpdate) SetNillableName(s *string) *ArtistUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetURI sets the "uri" field.
func (au *ArtistUpdate) SetURI(s string) *ArtistUpdate {
	au.mutation.SetURI(s)
	return au
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (au *ArtistUpdate) SetNillableURI(s *string) *ArtistUpdate {
	if s != nil {
		au.SetURI(*s)
	}
	return au
}

// SetGenres sets the "genres" field.
func (au *ArtistUpdate) SetGenres(s []string) *ArtistUpdate {
	au.mutation.SetGenres(s)
	return au
}

// AppendGenres appends s to the "genres" field.
func (au *ArtistUpdate) AppendGenres(s []string) *ArtistUpdate {
	au.mutation.AppendGenres(s)
	return au
}

// AddAlbumIDs adds the "albums" edge to the Album entity by IDs.
func (au *ArtistUpdate) AddAlbumIDs(ids ...string) *ArtistUpdate {
	au.mutation.AddAlbumIDs(ids...)
	return au
}

// AddAlbums adds the "albums" edges to the Album entity.
func (au *ArtistUpdate) AddAlbums(a ...*Album) *ArtistUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddAlbumIDs(ids...)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (au *ArtistUpdate) AddTrackIDs(ids ...string) *ArtistUpdate {
	au.mutation.AddTrackIDs(ids...)
	return au
}

// AddTracks adds the "tracks" edges to the Track entity.
func (au *ArtistUpdate) AddTracks(t ...*Track) *ArtistUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTrackIDs(ids...)
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (au *ArtistUpdate) AddImageIDs(ids ...string) *ArtistUpdate {
	au.mutation.AddImageIDs(ids...)
	return au
}

// AddImages adds the "images" edges to the Image entity.
func (au *ArtistUpdate) AddImages(i ...*Image) *ArtistUpdate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return au.AddImageIDs(ids...)
}

// Mutation returns the ArtistMutation object of the builder.
func (au *ArtistUpdate) Mutation() *ArtistMutation {
	return au.mutation
}

// ClearAlbums clears all "albums" edges to the Album entity.
func (au *ArtistUpdate) ClearAlbums() *ArtistUpdate {
	au.mutation.ClearAlbums()
	return au
}

// RemoveAlbumIDs removes the "albums" edge to Album entities by IDs.
func (au *ArtistUpdate) RemoveAlbumIDs(ids ...string) *ArtistUpdate {
	au.mutation.RemoveAlbumIDs(ids...)
	return au
}

// RemoveAlbums removes "albums" edges to Album entities.
func (au *ArtistUpdate) RemoveAlbums(a ...*Album) *ArtistUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveAlbumIDs(ids...)
}

// ClearTracks clears all "tracks" edges to the Track entity.
func (au *ArtistUpdate) ClearTracks() *ArtistUpdate {
	au.mutation.ClearTracks()
	return au
}

// RemoveTrackIDs removes the "tracks" edge to Track entities by IDs.
func (au *ArtistUpdate) RemoveTrackIDs(ids ...string) *ArtistUpdate {
	au.mutation.RemoveTrackIDs(ids...)
	return au
}

// RemoveTracks removes "tracks" edges to Track entities.
func (au *ArtistUpdate) RemoveTracks(t ...*Track) *ArtistUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTrackIDs(ids...)
}

// ClearImages clears all "images" edges to the Image entity.
func (au *ArtistUpdate) ClearImages() *ArtistUpdate {
	au.mutation.ClearImages()
	return au
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (au *ArtistUpdate) RemoveImageIDs(ids ...string) *ArtistUpdate {
	au.mutation.RemoveImageIDs(ids...)
	return au
}

// RemoveImages removes "images" edges to Image entities.
func (au *ArtistUpdate) RemoveImages(i ...*Image) *ArtistUpdate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return au.RemoveImageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArtistUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArtistUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArtistUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArtistUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ArtistUpdate) check() error {
	if v, ok := au.mutation.Href(); ok {
		if err := artist.HrefValidator(v); err != nil {
			return &ValidationError{Name: "href", err: fmt.Errorf(`ent: validator failed for field "Artist.href": %w`, err)}
		}
	}
	if v, ok := au.mutation.Name(); ok {
		if err := artist.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Artist.name": %w`, err)}
		}
	}
	if v, ok := au.mutation.URI(); ok {
		if err := artist.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Artist.uri": %w`, err)}
		}
	}
	return nil
}

func (au *ArtistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(artist.Table, artist.Columns, sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.ExternalUrls(); ok {
		_spec.SetField(artist.FieldExternalUrls, field.TypeJSON, value)
	}
	if value, ok := au.mutation.Href(); ok {
		_spec.SetField(artist.FieldHref, field.TypeString, value)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(artist.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.URI(); ok {
		_spec.SetField(artist.FieldURI, field.TypeString, value)
	}
	if value, ok := au.mutation.Genres(); ok {
		_spec.SetField(artist.FieldGenres, field.TypeJSON, value)
	}
	if value, ok := au.mutation.AppendedGenres(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, artist.FieldGenres, value)
		})
	}
	if au.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AlbumsTable,
			Columns: artist.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAlbumsIDs(); len(nodes) > 0 && !au.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AlbumsTable,
			Columns: artist.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AlbumsTable,
			Columns: artist.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.TracksTable,
			Columns: artist.TracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTracksIDs(); len(nodes) > 0 && !au.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.TracksTable,
			Columns: artist.TracksPrimaryKey,
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
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.TracksTable,
			Columns: artist.TracksPrimaryKey,
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
	if au.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.ImagesTable,
			Columns: artist.ImagesPrimaryKey,
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
			Table:   artist.ImagesTable,
			Columns: artist.ImagesPrimaryKey,
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
			Table:   artist.ImagesTable,
			Columns: artist.ImagesPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArtistUpdateOne is the builder for updating a single Artist entity.
type ArtistUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArtistMutation
}

// SetExternalUrls sets the "external_urls" field.
func (auo *ArtistUpdateOne) SetExternalUrls(sm *schema.StringMap) *ArtistUpdateOne {
	auo.mutation.SetExternalUrls(sm)
	return auo
}

// SetHref sets the "href" field.
func (auo *ArtistUpdateOne) SetHref(s string) *ArtistUpdateOne {
	auo.mutation.SetHref(s)
	return auo
}

// SetNillableHref sets the "href" field if the given value is not nil.
func (auo *ArtistUpdateOne) SetNillableHref(s *string) *ArtistUpdateOne {
	if s != nil {
		auo.SetHref(*s)
	}
	return auo
}

// SetName sets the "name" field.
func (auo *ArtistUpdateOne) SetName(s string) *ArtistUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *ArtistUpdateOne) SetNillableName(s *string) *ArtistUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetURI sets the "uri" field.
func (auo *ArtistUpdateOne) SetURI(s string) *ArtistUpdateOne {
	auo.mutation.SetURI(s)
	return auo
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (auo *ArtistUpdateOne) SetNillableURI(s *string) *ArtistUpdateOne {
	if s != nil {
		auo.SetURI(*s)
	}
	return auo
}

// SetGenres sets the "genres" field.
func (auo *ArtistUpdateOne) SetGenres(s []string) *ArtistUpdateOne {
	auo.mutation.SetGenres(s)
	return auo
}

// AppendGenres appends s to the "genres" field.
func (auo *ArtistUpdateOne) AppendGenres(s []string) *ArtistUpdateOne {
	auo.mutation.AppendGenres(s)
	return auo
}

// AddAlbumIDs adds the "albums" edge to the Album entity by IDs.
func (auo *ArtistUpdateOne) AddAlbumIDs(ids ...string) *ArtistUpdateOne {
	auo.mutation.AddAlbumIDs(ids...)
	return auo
}

// AddAlbums adds the "albums" edges to the Album entity.
func (auo *ArtistUpdateOne) AddAlbums(a ...*Album) *ArtistUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddAlbumIDs(ids...)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (auo *ArtistUpdateOne) AddTrackIDs(ids ...string) *ArtistUpdateOne {
	auo.mutation.AddTrackIDs(ids...)
	return auo
}

// AddTracks adds the "tracks" edges to the Track entity.
func (auo *ArtistUpdateOne) AddTracks(t ...*Track) *ArtistUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTrackIDs(ids...)
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (auo *ArtistUpdateOne) AddImageIDs(ids ...string) *ArtistUpdateOne {
	auo.mutation.AddImageIDs(ids...)
	return auo
}

// AddImages adds the "images" edges to the Image entity.
func (auo *ArtistUpdateOne) AddImages(i ...*Image) *ArtistUpdateOne {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return auo.AddImageIDs(ids...)
}

// Mutation returns the ArtistMutation object of the builder.
func (auo *ArtistUpdateOne) Mutation() *ArtistMutation {
	return auo.mutation
}

// ClearAlbums clears all "albums" edges to the Album entity.
func (auo *ArtistUpdateOne) ClearAlbums() *ArtistUpdateOne {
	auo.mutation.ClearAlbums()
	return auo
}

// RemoveAlbumIDs removes the "albums" edge to Album entities by IDs.
func (auo *ArtistUpdateOne) RemoveAlbumIDs(ids ...string) *ArtistUpdateOne {
	auo.mutation.RemoveAlbumIDs(ids...)
	return auo
}

// RemoveAlbums removes "albums" edges to Album entities.
func (auo *ArtistUpdateOne) RemoveAlbums(a ...*Album) *ArtistUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveAlbumIDs(ids...)
}

// ClearTracks clears all "tracks" edges to the Track entity.
func (auo *ArtistUpdateOne) ClearTracks() *ArtistUpdateOne {
	auo.mutation.ClearTracks()
	return auo
}

// RemoveTrackIDs removes the "tracks" edge to Track entities by IDs.
func (auo *ArtistUpdateOne) RemoveTrackIDs(ids ...string) *ArtistUpdateOne {
	auo.mutation.RemoveTrackIDs(ids...)
	return auo
}

// RemoveTracks removes "tracks" edges to Track entities.
func (auo *ArtistUpdateOne) RemoveTracks(t ...*Track) *ArtistUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTrackIDs(ids...)
}

// ClearImages clears all "images" edges to the Image entity.
func (auo *ArtistUpdateOne) ClearImages() *ArtistUpdateOne {
	auo.mutation.ClearImages()
	return auo
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (auo *ArtistUpdateOne) RemoveImageIDs(ids ...string) *ArtistUpdateOne {
	auo.mutation.RemoveImageIDs(ids...)
	return auo
}

// RemoveImages removes "images" edges to Image entities.
func (auo *ArtistUpdateOne) RemoveImages(i ...*Image) *ArtistUpdateOne {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return auo.RemoveImageIDs(ids...)
}

// Where appends a list predicates to the ArtistUpdate builder.
func (auo *ArtistUpdateOne) Where(ps ...predicate.Artist) *ArtistUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArtistUpdateOne) Select(field string, fields ...string) *ArtistUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Artist entity.
func (auo *ArtistUpdateOne) Save(ctx context.Context) (*Artist, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArtistUpdateOne) SaveX(ctx context.Context) *Artist {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArtistUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArtistUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ArtistUpdateOne) check() error {
	if v, ok := auo.mutation.Href(); ok {
		if err := artist.HrefValidator(v); err != nil {
			return &ValidationError{Name: "href", err: fmt.Errorf(`ent: validator failed for field "Artist.href": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Name(); ok {
		if err := artist.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Artist.name": %w`, err)}
		}
	}
	if v, ok := auo.mutation.URI(); ok {
		if err := artist.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Artist.uri": %w`, err)}
		}
	}
	return nil
}

func (auo *ArtistUpdateOne) sqlSave(ctx context.Context) (_node *Artist, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(artist.Table, artist.Columns, sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Artist.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, artist.FieldID)
		for _, f := range fields {
			if !artist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != artist.FieldID {
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
	if value, ok := auo.mutation.ExternalUrls(); ok {
		_spec.SetField(artist.FieldExternalUrls, field.TypeJSON, value)
	}
	if value, ok := auo.mutation.Href(); ok {
		_spec.SetField(artist.FieldHref, field.TypeString, value)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(artist.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.URI(); ok {
		_spec.SetField(artist.FieldURI, field.TypeString, value)
	}
	if value, ok := auo.mutation.Genres(); ok {
		_spec.SetField(artist.FieldGenres, field.TypeJSON, value)
	}
	if value, ok := auo.mutation.AppendedGenres(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, artist.FieldGenres, value)
		})
	}
	if auo.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AlbumsTable,
			Columns: artist.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAlbumsIDs(); len(nodes) > 0 && !auo.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AlbumsTable,
			Columns: artist.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AlbumsTable,
			Columns: artist.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.TracksTable,
			Columns: artist.TracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTracksIDs(); len(nodes) > 0 && !auo.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.TracksTable,
			Columns: artist.TracksPrimaryKey,
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
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.TracksTable,
			Columns: artist.TracksPrimaryKey,
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
	if auo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.ImagesTable,
			Columns: artist.ImagesPrimaryKey,
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
			Table:   artist.ImagesTable,
			Columns: artist.ImagesPrimaryKey,
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
			Table:   artist.ImagesTable,
			Columns: artist.ImagesPrimaryKey,
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
	_node = &Artist{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}