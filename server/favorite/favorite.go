package favorite

import "github.com/RaymondCode/simple-demo/server/favorite/db"

type FavReq struct {
	token    string
	video_id string
}
type FavoriteServer struct {
	err error
}

func Favorite(token string, video_id string) {
	fav := &db.FavoriteDb{
		Token:    token,
		Video_id: video_id,
	}
	db.Favorite(fav)
}
func UnFavorite(token string, video_id string) {
	fav := &db.FavoriteDb{
		Token:    token,
		Video_id: video_id,
	}
	db.UnFavorite(fav)
}
