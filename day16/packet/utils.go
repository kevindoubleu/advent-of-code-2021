package packet

import (
	"errors"
	"strconv"
	"strings"
)

func hex2bin(hexchar byte) string {
	hexchar = byte(strings.ToUpper(string(hexchar))[0])

	switch hexchar {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'A':
		return "1010"
	case 'B':
		return "1011"
	case 'C':
		return "1100"
	case 'D':
		return "1101"
	case 'E':
		return "1110"
	case 'F':
		return "1111"
	}

	panic(errors.New("invalid hex char"))
}

// calls hex2bin on each char and concats them together
func hexstr2bin(hexStr string) string {
	result := strings.Builder{}

	for _, hexchar := range hexStr {
		result.WriteString(hex2bin(byte(hexchar)))
	}

	return result.String()
}

// parsing packet body (groups) for literal value type packets
func (p *Packet) parseLiteralGroups() {
	if len(p.binary) < 5 { return }
	
	binaryResult := strings.Builder{}
	start, end := 6, 11 // first quintet start at 6

	group := p.binary[start : end]
	
	// keep parsing quintets until terminator group
	// a terminator group is a quintet that starts with 0
	for {
		binaryResult.WriteString(group[1 :])
		
		// first group could be terminator group, we must check immediately
		if group[0] == '0' {
			break
		}

		start += 5
		end += 5
		group = p.binary[start : end]
	}

	// parse the resulting binary into decimal
	result, err := strconv.ParseInt(binaryResult.String(), 2, 0)
	if err != nil {
		panic(err)
	}
	
	// assign it to packet value
	p.Value = int(result)

	// throw away the remaining bits after terminator group
	p.binary = p.binary[: end]
}

func (p *Packet) parseSubpacketWithLen(length int) {
	// the first subpacket starts at idx 22
	currentSubpacketHead := 22
	processedLength := 0

	// we keep processing new packets as long as we havent reached len
	for processedLength < length {
		subp := NewPacketWithBinary(p.binary[currentSubpacketHead :])
		
		currentSubpacketHead += len(subp.binary)
		processedLength += len(subp.binary)
		p.SubPackets = append(p.SubPackets, &subp)
	}

	// after done, throw away the unused bits
	// this is to make our len match only the bits that we processed
	p.binary = p.binary[: currentSubpacketHead]
}

func (p *Packet) parseSubpacketWithCount(count int) {
	// the first subpacket starts at idx 18
	currentSubpacketHead := 18
	processed := 0

	for processed < count {
		subp := NewPacketWithBinary(p.binary[currentSubpacketHead :])
		
		currentSubpacketHead += len(subp.binary)
		processed++
		p.SubPackets = append(p.SubPackets, &subp)
	}

	// after done, throw away the unused bits
	// this is to make our len match only the bits that we processed
	p.binary = p.binary[: currentSubpacketHead]
}
