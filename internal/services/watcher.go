package services

import (
	"context"
	"time"

	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	"go.uber.org/zap"
)

const (
	sleepTime                     = 5 * time.Second
	lastTrackDurationPercentage   = int64(50)
	replayTrackDurationPercentage = int64(5)
)

type Watcher struct {
	sharedService *Shared

	playerStateWebsocketHandler *models.PlayerStateWebsocketHandler
	lastPlayerState             *models.PlayerState
	ctx                         context.Context
}

func NewWatcher(sharedService *Shared, playerStateWebsocketHandler *models.PlayerStateWebsocketHandler) *Watcher {
	return &Watcher{
		sharedService:               sharedService,
		playerStateWebsocketHandler: playerStateWebsocketHandler,
	}
}

func (w *Watcher) StartPlayerStateLoop() {
	//TODO: Add a loop restart if error occours
	w.ctx = context.Background()

	for {
		// Sleep for 5 seconds to give the server time to start
		time.Sleep(sleepTime)

		playerState, err := w.sharedService.GetPlayerState(w.ctx)
		if err != nil {
			// This will more than likely happen in the case where nothing is playing or authentication
			// fails
			// TODO: Do something else here
			utils.Logger.Error("Error getting playerstate", zap.Error(err))
			continue
		}

		// Check if player is not nil and that the player has an item
		if playerState != nil && playerState.Track != nil {
			if w.lastPlayerState == nil {
				w.lastPlayerState = playerState
			}

			// This function requires data from the previous loop so it needs to be called before the update to the playerstate
			// This is to check if the track has changed and if so add it to the db or if the track has been replayed
			_ = w.checkUpdate(w.ctx, playerState)
			w.playerStateWebsocketHandler.Broadcast(playerState)
			w.lastPlayerState = playerState
		}

		// For testing to see if the loop is working
		utils.Logger.Debug("Playerstate receieved", zap.Any("playerState", playerState))
	}
}

func (w *Watcher) checkUpdate(ctx context.Context, playerState *models.PlayerState) bool {
	lastPlayerState := w.lastPlayerState

	// Check if the track has just changed and if so add it to the db
	if lastPlayerState.Track.Name != playerState.Track.Name {
		// TODO: Add a row to the "plays" table
		utils.Logger.Debug("Track has changed")
		return true
	}

	trackDuration := playerState.Track.DurationMs
	lastPlayerProgress := lastPlayerState.ProgressMs
	playerProgress := playerState.ProgressMs

	// Check for replays
	// TODO: Maybe find a better way to do this but works for now
	// Check if last track update duration is more than 50% done and if current progress is less than 05% into the track
	// This is what constitutes as a replay
	if (trackDuration/lastTrackDurationPercentage)*100 < lastPlayerProgress &&
		playerProgress <= (trackDuration/replayTrackDurationPercentage)*int64(100) {
		utils.Logger.Debug("Track has been replayed")

		return true
	}

	return false
}
