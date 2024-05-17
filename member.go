package main

import (
	"github.com/oklog/ulid/v2"
)

type MemberID string

func NewMemberID() MemberID {
	return MemberID(ulid.Make().String())
}

type MemberPriviledgeLevel uint

const (
	MemberPrivledgeLevel1 MemberPriviledgeLevel = 1
	MemberPrivledgeLevel2 MemberPriviledgeLevel = 2
	MemberPrivledgeLevel3 MemberPriviledgeLevel = 3
)

type Member struct {
	ID              MemberID
	Name            string
	PriviledgeLevel MemberPriviledgeLevel
}

var MemberSeeds = []*Member{
	{
		ID:              NewMemberID(),
		Name:            "Jack",
		PriviledgeLevel: MemberPrivledgeLevel1,
	},
	{
		ID:              NewMemberID(),
		Name:            "Bob",
		PriviledgeLevel: MemberPrivledgeLevel2,
	},
	{
		ID:              NewMemberID(),
		Name:            "Stephen",
		PriviledgeLevel: MemberPrivledgeLevel3,
	},
}
