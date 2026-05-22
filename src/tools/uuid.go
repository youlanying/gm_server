package tools

const (
	HIGHGUID_TYPE_MASK       = 0xFFF00000
	HIGHGUID_TYPE_UNIT       = 0x00000000 // 用于npc,monster
	HIGHGUID_TYPE_ROLE       = 0x00100000 // 用于roleid
	HIGHGUID_TYPE_ITEM       = 0x00200000 // 用于itemid
	HIGHGUID_TYPE_GUILD      = 0x00300000 // 用于guildid
	HIGHGUID_TYPE_MAIL       = 0x00400000 // 用于mailid
	HIGHGUID_TYPE_HERO       = 0x00500000 // 用于heroid
	HIGHGUID_TYPE_SYSTEMMAIL = 0x00600000 // 用于systemmailid
)

// MaskDefined
const (
	HIGHGUID_PLATFORM_MASK   = 0x000FF000 // platformid
	HIGHGUID_SERVER_MASK     = 0x00000FFF // serverid
	HIGHGUID_PLATSERVER_MASK = 0x000FFFFF // platform+server id
)

const MAX_TYPE_NUMS_BIT = 24

func GetHighGuid(tmpId uint64) uint64 {
	return tmpId >> MAX_TYPE_NUMS_BIT
}

func GetLowGuid(tmpId uint64) uint64 {
	return tmpId & 0xFFFFFFFF
}

// CalcPlatServerId 由平台id+服务id 生成 平台服务器id
func CalcPlatServerId(platId, serverId int32) uint64 {
	return ((uint64(platId) << 12) & HIGHGUID_PLATFORM_MASK) | (uint64(serverId) & HIGHGUID_SERVER_MASK)
}

// CalcMinIdByTypePlatServerId 生成各类型最小的id号
func CalcMinIdByTypePlatServerId(typeHigh, platServer uint64) uint64 {
	return (typeHigh | platServer) << MAX_TYPE_NUMS_BIT
}

// GetIdType 由各ID生成各类型
func GetIdType(tmpId uint64) uint64 {
	return GetHighGuid(tmpId) & HIGHGUID_TYPE_MASK
}

// CalcServerIdByTypeId 由各类型id生成服务器id
func CalcServerIdByTypeId(typeId uint64) uint64 {
	return GetHighGuid(typeId) & HIGHGUID_SERVER_MASK
}

// CalcPlatIdByTypeId 由各类型id生成平台id
func CalcPlatIdByTypeId(typeId uint64) uint64 {
	return (GetHighGuid(typeId) & HIGHGUID_PLATFORM_MASK) >> 12
}

func CalcFullIdByTypeLower(high, lower uint64) uint64 {
	return (high << MAX_TYPE_NUMS_BIT) | lower
}
