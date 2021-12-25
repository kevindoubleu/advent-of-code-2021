package polymer

import (
	"strings"
)

type Polymer struct {
	Template	string
	rules		map[string]byte
}

func NewPolymer(template string, rules []string) *Polymer {
	p := Polymer{
		Template: template,
		rules: make(map[string]byte),
	}

	// parse the rules into a map
	for _, rule := range rules {
		from, to := parseRule(rule)
		p.rules[from] = to
	}

	return &p
}

func (p *Polymer) Polymerize() {
	strLen := len(p.Template)
	if strLen < 2 { return }

	// add the first char
	newTemplate := strings.Builder{}
	newTemplate.WriteByte(p.Template[0])

	// then for every pair
	for i := 1; i < strLen; i++ {
		pair := p.Template[i-1 : i+1]
		
		// add the middle char we get from the rule
		newTemplate.WriteByte(p.rules[pair])
		// then only the second char of the pair
		// bcs the first char of the pair was already printed before
		newTemplate.WriteByte(p.Template[i])
	}

	p.Template = newTemplate.String()
}

func (p *Polymer) PolymerizeMultiple(times int) {
	for i := 0; i < times; i++ {
		p.Polymerize()
	}
}

func (p Polymer) MostCommonCharCount() int {
	charMap := toCharMap(p.Template)
	
	max := 0
	for _, count := range charMap {
		if count > max {
			max = count
		}
	}

	return max
}

func (p Polymer) LeastCommonCharCount() int {
	charMap := toCharMap(p.Template)

	min := len(p.Template)
	for _, count := range charMap {
		if count < min {
			min = count
		}
	}

	return min
}
