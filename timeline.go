package goinsta

import (
	"encoding/json"
)

type Timeline struct {
	inst *Instagram
}

func newTimeline(inst *Instagram) *Timeline {
	time := &Timeline{
		inst: inst,
	}
	return time
}

// Get returns latest media from timeline.
//
// For pagination use FeedMedia.Next()
func (time *Timeline) Get() (*FeedMedia, error) {
	insta := time.inst

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlTimeline,
			Query: map[string]string{
				"max_id":         "",
				"rank_token":     insta.rankToken,
				"ranked_content": "true",
			},
		},
	)
	if err == nil {
		media := &FeedMedia{}
		err = json.Unmarshal(body, media)
		media.inst = insta
		media.endpoint = urlTimeline
		return media, err
	}
	return nil, err
}

// Stories returns slice of StoryMedia
func (time *Timeline) Stories() (*Tray, error) {
	body, err := time.inst.sendSimpleRequest(urlStories)
	if err == nil {
		tray := &Tray{}
		err = json.Unmarshal(body, tray)
		if err != nil {
			return nil, err
		}
		tray.set(time.inst, urlStories)
		return tray, nil
	}
	return nil, err
}