package main

type Note struct {
	entries	[]NoteEntry
}

func newNote(entries []string) Note {
	note := Note{entries: make([]NoteEntry, 0)}

	for _, entry := range entries {
		note.entries = append(note.entries, newNoteEntry(entry))
	}

	return note
}

// 1
func (n Note) getEasyOutputCount() int {
	total := 0

	for _, entry := range n.entries {
		for _, output := range entry.outputs {
			if len(output) == 2 ||
				len(output) == 3 ||
				len(output) == 4 ||
				len(output) == 7 {
				total++
			}
		}
	}

	return total
}
