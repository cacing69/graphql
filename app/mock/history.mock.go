package mock

import "github.com/cacing69/graphql/app/model"

var HistoryMock = []model.History{
	{
		Id:     1,
		UserId: 1,
		Desc:   "lorem 1",
		Viewer: ViewerMock,
	},
	{
		Id:     2,
		UserId: 2,
		Desc:   "lorem 2",
		Viewer: ViewerMock,
	},
	{
		Id:     3,
		UserId: 1,
		Desc:   "lorem 3",
		Viewer: ViewerMock,
	},
	{
		Id:     4,
		UserId: 1,
		Desc:   "lorem 4",
		Viewer: ViewerMock,
	},
	{
		Id:     5,
		UserId: 3,
		Desc:   "lorem 5",
		Viewer: ViewerMock,
	},
}
