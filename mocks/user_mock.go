package mocks

import "github.com/cacing69/graphql/models"

var UserMock = []models.User{
	{
		Id:       1,
		UserName: "cacing69",
		History:  HistoryMock,
	},
	{
		Id:       2,
		UserName: "g9insane",
		History:  HistoryMock,
	},
	{
		Id:       3,
		UserName: "alucrad",
		History:  HistoryMock,
	},
}
