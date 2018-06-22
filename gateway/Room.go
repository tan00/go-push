package gateway

import "sync"

// 房间
type Room struct {
	rwMutex sync.RWMutex
	roomId string
	id2Conn map[uint64]*WSConnection
}

func InitRoom(roomId string) (room *Room) {
	room = &Room {
		roomId: roomId,
		id2Conn: make(map[uint64]*WSConnection),
	}
	return
}

func (room *Room) Join(wsConn *WSConnection) (err error) {
	var (
		existed bool
	)

	room.rwMutex.Lock()
	defer room.rwMutex.Unlock()

	if _, existed = room.id2Conn[wsConn.connId]; existed {
		err = ERR_JOIN_ROOM_TWICE
		return
	}

	room.id2Conn[wsConn.connId] = wsConn
	return
}

func (room *Room) Leave(wsConn* WSConnection) (err error) {
	var (
		existed bool
	)

	room.rwMutex.Lock()
	defer room.rwMutex.Unlock()

	if _, existed = room.id2Conn[wsConn.connId]; !existed {
		err = ERR_NOT_IN_ROOM
		return
	}

	delete(room.id2Conn, wsConn.connId)
	return
}

func (room *Room) Count() int {
	room.rwMutex.RLock()
	defer room.rwMutex.RUnlock()

	return len(room.id2Conn)
}