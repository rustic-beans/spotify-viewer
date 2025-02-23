package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rustic-beans/spotify-viewer/internal/database"
	"golang.org/x/oauth2"
)

type Token = database.Token
type UpsertTokenParams = database.UpsertTokenParams

func IntoOauth2Token(token *Token) *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry.Time,
	}
}

func FromTokenToUpsertParams(token *oauth2.Token) *UpsertTokenParams {
	timestamp := pgtype.Timestamp{
		Time:  token.Expiry,
		Valid: true,
	}

	return &UpsertTokenParams{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       timestamp,
	}
}
