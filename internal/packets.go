package internal

import (
	"log"

	"github.com/Hucaru/Valhalla/common/mpacket"
	"github.com/Hucaru/Valhalla/common/opcode"
)

func PacketChannelPopUpdate(id byte, pop int16) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannelInfo)
	p.WriteByte(id)
	p.WriteByte(0) // 0 is population
	p.WriteInt16(pop)

	return p
}

func PacketChannelPlayerConnected(playerID int32, name string, channelID byte, channelChange bool, mapID int32) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannePlayerConnect)
	p.WriteInt32(playerID)
	p.WriteString(name)
	p.WriteByte(channelID)
	p.WriteBool(channelChange)
	p.WriteInt32(mapID)

	return p
}

func PacketChannelPlayerDisconnect(id int32, name string) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannePlayerDisconnect)
	p.WriteInt32(id)
	p.WriteString(name)

	return p
}

func PacketChannelBuddyEvent(op byte, recepientID, fromID int32, fromName string, channelID byte) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannelPlayerBuddyEvent)
	p.WriteByte(op)

	switch op {
	case 1: // add
		fallthrough
	case 2: // accept
		p.WriteInt32(recepientID)
		p.WriteInt32(fromID)
		p.WriteString(fromName)
		p.WriteByte(channelID)
	case 3: // delete / reject
		p.WriteInt32(recepientID)
		p.WriteInt32(fromID)
		p.WriteByte(channelID)
	default:
		log.Println("unkown internal buddy event type:", op)
	}

	return p
}

func PacketChannelWhispherChat(recepientName, fromName, msg string, channelID byte) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannelPlayerChatEvent)
	p.WriteByte(0) // whispher
	p.WriteString(recepientName)
	p.WriteString(fromName)
	p.WriteString(msg)
	p.WriteByte(channelID)

	return p
}

func PacketChannelPlayerChat(code byte, fromName string, buffer []byte) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannelPlayerChatEvent)
	p.WriteByte(code) // 1 buddy, 2 party, 3 guild
	p.WriteString(fromName)
	p.WriteBytes(buffer)

	return p
}

func PacketChannelPartyCreateRequest(playerID int32, channelID byte, mapID, job, level int32, name string) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannelPlayerPartyEvent)
	p.WriteByte(0) // new party request
	p.WriteInt32(playerID)
	p.WriteByte(channelID)
	p.WriteInt32(mapID)
	p.WriteInt32(job)
	p.WriteInt32(level)
	p.WriteString(name)

	return p
}

func PacketChannelPartyCreateApproved(playerID int32, success bool, party *Party) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannelPlayerPartyEvent)
	p.WriteByte(1) // new party
	p.WriteInt32(playerID)
	p.WriteBool(success)

	if success {
		p.WriteBytes(party.GeneratePacket())
	}

	return p
}

func PacketChannelPartyLeave(partyID, playerID int32, kicked bool) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannelPlayerPartyEvent)
	p.WriteByte(2) // leave party
	p.WriteInt32(partyID)
	p.WriteInt32(playerID)
	p.WriteBool(kicked)

	return p
}

func PacketWorldPartyLeave(partyID int32, destroy, kicked bool, index int32, party *Party) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannelPlayerPartyEvent)
	p.WriteByte(2) // leave party
	p.WriteInt32(partyID)
	p.WriteBool(destroy)
	p.WriteBool(kicked)
	p.WriteInt32(index)
	p.WriteBytes(party.GeneratePacket())

	return p
}

func PacketChannelPartyAccept(partyID, playerID, channelID, mapID, job, level int32, name string) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannelPlayerPartyEvent)
	p.WriteByte(3) // accept invite
	p.WriteInt32(partyID)
	p.WriteInt32(playerID)
	p.WriteInt32(channelID)
	p.WriteInt32(mapID)
	p.WriteInt32(job)
	p.WriteInt32(level)
	p.WriteString(name)

	return p
}

func PacketWorldPartyAccept(partyID, playerID, index int32, party *Party) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannelPlayerPartyEvent)
	p.WriteByte(3) // accept invite
	p.WriteInt32(partyID)
	p.WriteInt32(playerID)
	p.WriteInt32(index)
	p.WriteBytes(party.GeneratePacket())

	return p
}

func PacketChannelPartyUpdateInfo(partyID, playerID, job, level int32, name string) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannelPlayerPartyEvent)
	p.WriteByte(4) // update party window info
	p.WriteInt32(partyID)
	p.WriteInt32(playerID)
	p.WriteInt32(job)
	p.WriteInt32(level)
	p.WriteString(name)

	return p
}

func PacketWorldPartyUpdate(partyID, playerID, index int32, onlineStatus bool, party *Party) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.ChannelPlayerPartyEvent)
	p.WriteByte(4) // update party window info
	p.WriteInt32(partyID)
	p.WriteInt32(playerID)
	p.WriteInt32(index)
	p.WriteBool(onlineStatus)
	p.WriteBytes(party.GeneratePacket())

	return p
}

func PacketLoginDeletedCharacter(playerID int32) mpacket.Packet {
	p := mpacket.CreateInternal(opcode.LoginDeleteCharacter)
	p.WriteInt32(playerID)

	return p
}
