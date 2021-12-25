package packet

import (
	"fmt"
	"strconv"
)

type Packet struct {
	// just a regular metadata
	Version		int
	
	// determines if this packet has value or more subpackets
	TypeID		int

	// the value if this is a literal value packet
	Value		int
	// the nested subpacket if this is an operator packet
	SubPackets	[]*Packet


	// implementation details

	// binary representation of HexStr
	binary		string
}

const (
	sum = iota
	product
	minimum
	maximum
	literalValue
	greaterThan
	lessThan
	equal
)

func NewPacket(hexStr string) Packet {
	p := Packet{
		SubPackets : make([]*Packet, 0),
		binary     : hexstr2bin(hexStr),
	}
	p.Parse()

	return p
}

func NewPacketWithBinary(binaryStr string) Packet {
	p := Packet{
		SubPackets : make([]*Packet, 0),
		binary     : binaryStr,
	}
	p.Parse()

	return p
}

func (p Packet) version() int {
	// the first 3 bits make up the version
	binaryStr := p.binary[0 : 3]

	result, err := strconv.ParseInt(binaryStr, 2, 0)
	if err != nil {
		panic(err)
	}

	return int(result)
}

func (p Packet) typeID() int {
	// the 4th, 5th, 6th bit make up the type id
	binaryStr := p.binary[3 : 6]
	
	result, err := strconv.ParseInt(binaryStr, 2, 0)
	if err != nil {
		panic(err)
	}

	return int(result)
}

// recursively parses the packet and its subpackets
func (p *Packet) Parse() {
	p.Version = p.version()
	p.TypeID = p.typeID()

	// packet is literal value
	if p.TypeID == literalValue {
		p.parseLiteralGroups()
		// fmt.Println("parsing literal", p.binary, "got", p.Value)
		return
	}

	// packet is operator
	// check the length type id
	if p.binary[6] == '0' {
		// use len of subpacket bits

		// it is 15 bits after the length type id
		lengthInBinary := p.binary[7 : 22]
		length, err := strconv.ParseInt(lengthInBinary, 2, 0)
		if err != nil {
			panic(err)
		}

		fmt.Println("operator with length type id LENGTH", lengthInBinary, length)

		p.parseSubpacketWithLen(int(length))
	} else if p.binary[6] == '1' {
		// use subpacket count

		// it is 11 bits after the length type id
		countInBinary := p.binary[7 : 18]
		count, err := strconv.ParseInt(countInBinary, 2, 0)
		if err != nil {
			panic(err)
		}

		fmt.Println("operator with length type id COUNT", countInBinary, count)

		p.parseSubpacketWithCount(int(count))
	}

	// part 2
	p.calculateValue()
}

// 1
func (p Packet) TotalVersionNumber() int {
	return p.addVersionNumber(0)
}

// recursively adds all version numbers
func (p Packet) addVersionNumber(total int) int {
	total += p.Version

	if p.SubPackets != nil {
		for _, subpacket := range p.SubPackets {
			total = subpacket.addVersionNumber(total)
		}
	}

	return total
}

// 2
func (p *Packet) calculateValue() {
	switch p.TypeID {
	case sum:
		total := 0
		for _, subp := range p.SubPackets {
			total += subp.Value
		}
		p.Value = total
	
	case product:
		total := 1
		for _, subp := range p.SubPackets {
			total *= subp.Value
		}
		p.Value = total
	
	case minimum:
		min := p.SubPackets[0].Value
		for _, subp := range p.SubPackets {
			if subp.Value < min {
				min = subp.Value
			}
		}
		p.Value = min

	case maximum:
		max := p.SubPackets[0].Value
		for _, subp := range p.SubPackets {
			if subp.Value > max {
				max = subp.Value
			}
		}
		p.Value = max
	
	case greaterThan:
		// from the problem:
		// These packets always have exactly two sub-packets.
		if p.SubPackets[0].Value > p.SubPackets[1].Value {
			p.Value = 1
		}
		// go int default is already 0
	
	case lessThan:
		// from the problem:
		// These packets always have exactly two sub-packets.
		if p.SubPackets[0].Value < p.SubPackets[1].Value {
			p.Value = 1
		}
		// go int default is already 0

	case equal:
		// from the problem:
		// These packets always have exactly two sub-packets.
		if p.SubPackets[0].Value == p.SubPackets[1].Value {
			p.Value = 1
		}
		// go int default is already 0
	}
}
