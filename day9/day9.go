package day9

import (
	"bufio"
	"errors"
	"os"
	"slices"
	"strconv"

	u "github.com/bennyharvey/aoc2024/utils"
)

type BType int

const (
	BTYPE_FILE BType = iota
	BTYPE_EMPTY_SPACE
)

type File struct {
	// blocks []Block
	start int
	size  int
	id    int
}

type Block struct {
	btype BType
	id    int
}

type Disk struct {
	rawData   []rune
	blockData []rune
	blocks    []Block
	files     []File
}

func (d *Disk) print() {
	// u.Printf1("%c\n", d.rawData)
	// u.Printf1("%c\n", d.blockData)
	for _, block := range d.blocks {
		if block.btype == BTYPE_FILE {
			u.Printf1("%d", block.id)
		} else {
			u.Printf1(".")
		}
	}
	u.Println1("")
	u.PPrintSlice(d.files)
	u.Println1("")
}

func DiskFromLine(line string) *Disk {
	disk := Disk{rawData: []rune(line), blockData: make([]rune, 0), blocks: make([]Block, 0)}
	var btype BType
	btype = BTYPE_FILE
	blockId := 0

	for _, char := range line {
		len := int(char - '0')
		for range len {
			if btype == BTYPE_FILE {
				disk.blockData = append(disk.blockData, []rune(strconv.Itoa(blockId))[0])
				disk.blocks = append(disk.blocks, Block{btype: btype, id: blockId})
			}
			if btype == BTYPE_EMPTY_SPACE {
				disk.blockData = append(disk.blockData, '.')
				disk.blocks = append(disk.blocks, Block{btype: btype})
			}
		}

		if btype == BTYPE_FILE {
			btype = BTYPE_EMPTY_SPACE
		} else {
			btype = BTYPE_FILE
			blockId++
		}
	}

	return &disk
}

func DiskFromLine2(line string) *Disk {
	disk := Disk{blocks: make([]Block, 0), files: make([]File, 0)}
	var btype BType
	btype = BTYPE_FILE
	blockId := 0

	for _, char := range line {
		len := int(char - '0')
		for range len {
			if btype == BTYPE_FILE {
				disk.blocks = append(disk.blocks, Block{btype: btype, id: blockId})
			}
			if btype == BTYPE_EMPTY_SPACE {
				disk.blocks = append(disk.blocks, Block{btype: btype})
			}
		}

		if btype == BTYPE_FILE {
			btype = BTYPE_EMPTY_SPACE
		} else {
			btype = BTYPE_FILE
			blockId++
		}
	}

	parsingFile := false
	lastBlockId := -1
	file := File{}
	for addr, block := range disk.blocks {
		if block.btype == BTYPE_FILE {
			if !parsingFile {
				file.start = addr
				file.id = block.id
			}
			if lastBlockId != block.id && parsingFile {
				file.size = addr - file.start
				disk.files = append(disk.files, file)
				file = File{}
				file.start = addr
				file.id = block.id
			}
			parsingFile = true
			// file.blocks = append(file.blocks, block)
		}
		if block.btype == BTYPE_EMPTY_SPACE {
			if parsingFile {
				file.size = addr - file.start
				disk.files = append(disk.files, file)
				parsingFile = false
				file = File{}
			}
		}
		lastBlockId = block.id
	}

	file.size = len(disk.blocks) - file.start
	disk.files = append(disk.files, file)

	return &disk
}

func (d *Disk) lastElementIndex() (int, error) {
	for i := len(d.blocks) - 1; i > 0; i-- {
		if d.blocks[i].btype != BTYPE_EMPTY_SPACE {
			return i, nil
		}
	}
	return 0, errors.New("unreachable")
}

func isEmptySpace(blocks []Block) bool {
	for _, char := range blocks {
		if char.btype != BTYPE_EMPTY_SPACE {
			return false
		}
	}

	return true
}

func (d *Disk) compress() {
	for i := range d.blocks {
		if isEmptySpace(d.blocks[i:]) {
			return
		}
		j := u.Must(d.lastElementIndex())
		if d.blocks[i].btype == BTYPE_EMPTY_SPACE {
			d.blocks[i], d.blocks[j] = d.blocks[j], d.blocks[i]
		}
	}
}

func (d *Disk) defragment() {
	slices.Reverse(d.files)
	i := 0
	for _, file := range d.files[0 : len(d.files)-1] {
		freeBlocks := 0
		for addr, block := range d.blocks {
			if block.btype == BTYPE_EMPTY_SPACE {
				freeBlocks++
			}
			if block.btype == BTYPE_FILE {
				freeBlocks = 0
			}
			if freeBlocks == file.size && file.start > addr {
				offset := 0
				// u.Println3("found free space for file #", file.id, addr)
				// d.print()
				for i := addr - file.size + 1; i <= addr; i++ {
					// u.Println3("swapping", i, file.start+offset-1)
					d.blocks[i], d.blocks[file.start+offset] = d.blocks[file.start+offset], d.blocks[i]
					offset++
				}
				// d.print()
				break
			}
		}
		i++
	}

}

func (d *Disk) checksum() int {
	sum := 0
	for i, char := range d.blocks {
		if char.btype == BTYPE_EMPTY_SPACE {
			continue
		}
		sum += i * char.id
	}
	return sum
}

func SolvePart1(fileName string) int {
	file := u.Must(os.Open(fileName))
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	disk := DiskFromLine(line)
	disk.compress()
	checksum := disk.checksum()
	u.Println("Day 9, Part 1:", checksum)

	return checksum
}

func SolvePart2(fileName string) int {
	file := u.Must(os.Open(fileName))
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	disk := DiskFromLine2(line)
	disk.print()
	disk.defragment()
	disk.print()

	checksum := disk.checksum()
	u.Println("Day 9, Part 2:", checksum)

	return checksum
}
