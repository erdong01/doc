package aoi

import (
	"fmt"
	"sync"
)

type Grid struct {
	GID int //格子ID

	MinX int

	MaxX int

	MinY int

	MaxY int

	playerIDs map[int]bool

	pIDLock sync.RWMutex
}

func NewGrid(gID, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID:       gID,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		playerIDs: make(map[int]bool),
	}
}

func (g *Grid) Add(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	g.playerIDs[playerID] = true
}

func (g *Grid) Remove(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.playerIDs, playerID)
}

func (g *Grid) GetPlayerIDs() (palyerIDs []int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	for k, _ := range g.playerIDs {
		palyerIDs = append(palyerIDs, k)
	}
	return
}

func (g *Grid) String() string {
	return fmt.Sprintf("Grid id: %d, minX%d, maxX%d, minY%d, maxY%d, playerIDs:%v\n",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.playerIDs)
}
