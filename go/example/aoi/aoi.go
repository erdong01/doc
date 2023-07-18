package aoi

import "fmt"

//区域管理模块
type AOIManager struct {
	//区域的左边界坐标
	MinX int
	//区域的右边界坐标
	MaxX int

	CntsX int

	MinY int

	MaxY int

	CntsY int
	grids map[int]*Grid
}

func NewAOIManager(minX, maxX, cntsX, minY, maxY, cntsY int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		MinY:  minY,
		MaxY:  maxY,
		CntsX: cntsX,
		CntsY: cntsY,
		grids: make(map[int]*Grid),
	}

	//给AOI初始化区域的格子所有的格子进行编号和初始化
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			//计算格子ID 根据
			//格子编号： id = idy * cntY + idx
			gid := y*cntsX + x
			aoiMgr.grids[gid] = NewGrid(gid,
				aoiMgr.MinX+x*aoiMgr.gridWidth(),
				aoiMgr.MinX+(x+1)*aoiMgr.gridWidth(),
				aoiMgr.MinY+y*aoiMgr.gridLength(),
				aoiMgr.MinY+(y+1)*aoiMgr.gridLength(),
			)
		}
	}
	return aoiMgr
}

//得到每个格子在X轴反向的宽度
func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}

//得到每个格子在Y轴反向的高度
func (m *AOIManager) gridLength() int {
	return (m.MaxY - m.MinY) / m.CntsY
}

// 打印格子信息
func (m *AOIManager) String() string {
	//打印AOIManager信息
	s := fmt.Sprintf("AOIManager:\n  MinX%d, MaxX:%d, CntsX:%d, MinY:%d, MaxY:%d, CntsY%d \n Grids in AOIManager: \n",
		m.MinX, m.MaxX, m.CntsX, m.MinY, m.MaxY, m.CntsY,
	)
	//打印全部格子信息
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}
	return s
}

//根据格子GID得到周边九宫格ID的集合

func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	//判断
	if _, ok := m.grids[gID]; !ok {
		return
	}

	grids = append(grids, m.grids[gID])

	idx := gID % m.CntsX
	if idx > 0 {
		grids = append(grids, m.grids[gID-1])
	}

	if idx < m.CntsX-1 {
		grids = append(grids, m.grids[gID+1])
	}

	gidsX := make([]int, 0, len(grids))

	for _, v := range grids {
		gidsX = append(gidsX, v.GID)
	}

	for _, v := range gidsX {
		idy := v / m.CntsY
		if idy > 0 {
			grids = append(grids, m.grids[v-m.CntsX])
		}
		if idy < m.CntsY-1 {
			grids = append(grids, m.grids[v+m.CntsX])
		}
	}
	return
}

//通过x,y横纵坐标得到当前的GID的格子编号
func (m *AOIManager) GetGidByPos(x, y float32) (gid int) {
	idx := (int(x) - m.MinX) / m.gridWidth()

	idy := (int(y) - m.MinY) / m.gridLength()
	gid = idy*m.CntsX + idx
	return
}

//通过横纵坐标得到周边九宫格内全部的PlayerIDs
func (m *AOIManager) GetPidsByPos(x, y float32) (playerIDs []int) {
	//得到当前玩家的Gid格子id
	gID := m.GetGidByPos(x, y)
	//通过GID得到周边九宫格信息
	grids := m.GetSurroundGridsByGid(gID)
	//将九宫格的信息里面的全部的player的id 累加到 playerIDs
	for _, grid := range grids {
		playerIDs = append(playerIDs, grid.GetPlayerIDs()...)
		fmt.Printf("===>grid ID:%d,pids:%v ===", grid.GID, grid.GetPlayerIDs())
	}
	return
}

//添加一个PlayerID到一个格子中
func (m *AOIManager) AddPidToGrid(pID, gID int) {
	m.grids[gID].Add(pID)
}

//移除一个格子的playerID
func (m *AOIManager) RemovePidFromGrid(pID, gID int) {
	m.grids[gID].Remove(pID)
}

//通过坐标将Player添加到一个格子中
func (m *AOIManager) GetPidsByGid(gID int) (playerIDs []int) {
	playerIDs = m.grids[gID].GetPlayerIDs()
	return
}

// 通过坐标将Player添加到一个格子中
func (m *AOIManager) AddToGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Add(pID)
}

//通过坐标把一个Player从一个格子删除
func (m *AOIManager) RemoveFromGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Remove(pID)
}
