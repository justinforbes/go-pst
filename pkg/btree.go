// go-pst is a library for reading Personal Storage Table (.pst) files (written in Go/Golang).
//
// Copyright (C) 2022  Marten Mooij
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package pst

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

// GetNodeBTreeOffset returns the file offset to the node b-tree.
// References "The 64-bit header data", "The 32-bit header data".
func (file *File) GetNodeBTreeOffset() (int64, error) {
	var outputBuffer []byte
	var offset int64

	switch file.FormatType {
	case FormatTypeUnicode:
		outputBuffer = make([]byte, 8)
		offset = 224
	case FormatTypeUnicode4k:
		outputBuffer = make([]byte, 8)
		offset = 224
	case FormatTypeANSI:
		outputBuffer = make([]byte, 4)
		offset = 188
	default:
		return 0, errors.WithStack(ErrFormatTypeUnsupported)
	}

	if _, err := file.ReadAt(outputBuffer, offset); err != nil {
		return 0, errors.WithStack(err)
	}

	switch file.FormatType {
	case FormatTypeANSI:
		return int64(binary.LittleEndian.Uint32(outputBuffer)), nil
	default:
		return int64(binary.LittleEndian.Uint64(outputBuffer)), nil
	}
}

// GetBlockBTreeOffset returns the file offset to the block b-tree.
// References "The 64-bit header data", "The 32-bit header data".
func (file *File) GetBlockBTreeOffset() (int64, error) {
	var outputBuffer []byte
	var offset int64

	switch file.FormatType {
	case FormatTypeUnicode:
		outputBuffer = make([]byte, 8)
		offset = 240
	case FormatTypeUnicode4k:
		outputBuffer = make([]byte, 8)
		offset = 240
	case FormatTypeANSI:
		outputBuffer = make([]byte, 4)
		offset = 196
	default:
		return 0, errors.WithStack(ErrFormatTypeUnsupported)
	}

	if _, err := file.ReadAt(outputBuffer, offset); err != nil {
		return 0, errors.WithStack(err)
	}

	switch file.FormatType {
	case FormatTypeANSI:
		return int64(binary.LittleEndian.Uint32(outputBuffer)), nil
	default:
		return int64(binary.LittleEndian.Uint64(outputBuffer)), nil
	}
}

// GetBTreeNodeEntryCount returns the amount of entries in the b-tree.
// References "The node and block b-tree".
func (file *File) GetBTreeNodeEntryCount(btreeNodeOffset int64) (uint16, error) {
	var outputBuffer []byte
	var offset int64

	switch file.FormatType {
	case FormatTypeUnicode:
		outputBuffer = make([]byte, 1)
		offset = btreeNodeOffset + 488
	case FormatTypeUnicode4k:
		outputBuffer = make([]byte, 2)
		offset = btreeNodeOffset + 4056
	case FormatTypeANSI:
		outputBuffer = make([]byte, 1)
		offset = btreeNodeOffset + 496
	default:
		return 0, errors.WithStack(ErrFormatTypeUnsupported)
	}

	if _, err := file.ReadAt(outputBuffer, offset); err != nil {
		return 0, errors.WithStack(err)
	}

	switch file.FormatType {
	case FormatTypeUnicode4k:
		return binary.LittleEndian.Uint16(outputBuffer), nil
	default:
		return uint16(outputBuffer[0]), nil
	}
}

// GetBTreeNodeEntrySize returns the size of an entry in the b-tree.
// References "The node and block b-tree".
func (file *File) GetBTreeNodeEntrySize(btreeNodeOffset int64) (uint8, error) {
	outputBuffer := make([]byte, 1)
	var offset int64

	switch file.FormatType {
	case FormatTypeUnicode:
		offset = btreeNodeOffset + 490
	case FormatTypeUnicode4k:
		offset = btreeNodeOffset + 4060
	case FormatTypeANSI:
		offset = btreeNodeOffset + 498
	default:
		return 0, errors.WithStack(ErrFormatTypeUnsupported)
	}

	if _, err := file.ReadAt(outputBuffer, offset); err != nil {
		return 0, errors.WithStack(err)
	}

	return outputBuffer[0], nil
}

// GetBTreeNodeLevel returns the level of the b-tree node.
// References "The node and block b-tree".
func (file *File) GetBTreeNodeLevel(btreeNodeOffset int64) (uint8, error) {
	outputBuffer := make([]byte, 1)
	var offset int64

	switch file.FormatType {
	case FormatTypeUnicode:
		offset = btreeNodeOffset + 491
	case FormatTypeUnicode4k:
		offset = btreeNodeOffset + 4061
	case FormatTypeANSI:
		offset = btreeNodeOffset + 499
	default:
		return 0, errors.WithStack(ErrFormatTypeUnsupported)
	}

	if _, err := file.ReadAt(outputBuffer, offset); err != nil {
		return 0, errors.WithStack(err)
	}

	return outputBuffer[0], nil
}

// BTreeNode represents an entry in a b-tree node.
type BTreeNode struct {
	// Identifier is only unique to the node level.
	Identifier                 Identifier `json:"identifier"`
	FileOffset                 int64      `json:"fileOffset"`
	DataIdentifier             Identifier `json:"dataIdentifier"`
	LocalDescriptorsIdentifier Identifier `json:"localDescriptorsIdentifier"`
	Size                       uint16     `json:"size"`
	NodeLevel                  uint8      `json:"nodeLevel"`
}

// NewBTreeNodeReader is used by the Heap-on-Node.
func NewBTreeNodeReader(btreeNode BTreeNode, file *File) *io.SectionReader {
	return io.NewSectionReader(file, btreeNode.FileOffset, int64(btreeNode.Size))
}

// BTreeNodeLessFunc tests whether the current node is less than the given argument.
//
// This must provide a strict weak ordering.
// If !a.Less(b) && !b.Less(a), we treat this to mean a == b (we can only hold one of either a or b in the tree).
func BTreeNodeLessFunc(a BTreeNode, b BTreeNode) bool {
	if a.Identifier == b.Identifier {
		// We don't return the first identifier because there may be two or more nodes with
		// the same identifier, so prefer leaf nodes (which there is always only one of).
		return a.NodeLevel < b.NodeLevel
	} else {
		return a.Identifier < b.Identifier
	}
}

// GetBTreeNodeRawEntries returns the raw b-tree node entries in bytes.
// Used by GetBTreeNodeEntries.
func (file *File) GetBTreeNodeRawEntries(btreeNodeOffset int64) ([]byte, error) {
	var outputBuffer []byte

	switch file.FormatType {
	case FormatTypeUnicode:
		outputBuffer = make([]byte, 488)
	case FormatTypeUnicode4k:
		outputBuffer = make([]byte, 4056)
	case FormatTypeANSI:
		outputBuffer = make([]byte, 496)
	default:
		return nil, errors.WithStack(ErrFormatTypeUnsupported)
	}

	if _, err := file.ReadAt(outputBuffer, btreeNodeOffset); err != nil {
		return nil, errors.WithStack(err)
	}

	return outputBuffer, nil
}

// GetBTreeNodeEntries returns the entries in the b-tree node.
// References "The node and block b-tree".
func (file *File) GetBTreeNodeEntries(btreeNodeOffset int64, btreeType BTreeType, btreeNodeLevel uint8) ([]BTreeNode, error) {
	btreeNodeRawEntries, err := file.GetBTreeNodeRawEntries(btreeNodeOffset)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	btreeNodeEntryCount, err := file.GetBTreeNodeEntryCount(btreeNodeOffset)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	btreeNodeEntrySize, err := file.GetBTreeNodeEntrySize(btreeNodeOffset)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	btreeNodeEntries := make([]BTreeNode, btreeNodeEntryCount)

	for i := 0; i < int(btreeNodeEntryCount); i++ {
		btreeNodeEntryData := btreeNodeRawEntries[i*int(btreeNodeEntrySize) : (i*int(btreeNodeEntrySize))+int(btreeNodeEntrySize)]

		if btreeNodeLevel > 0 && (btreeType == BTreeTypeNode || btreeType == BTreeTypeBlock) {
			// Branch node or block b-tree node.
			btreeNodeEntries[i] = BTreeNode{
				Identifier: GetBTreeNodeEntryIdentifier(btreeNodeEntryData, file.FormatType),
				FileOffset: GetBTreeNodeEntryFileOffset(btreeNodeEntryData, true, file.FormatType),
				NodeLevel:  btreeNodeLevel,
			}
		} else if btreeNodeLevel == 0 && btreeType == BTreeTypeNode {
			// Leaf node b-tree node.
			btreeNodeEntries[i] = BTreeNode{
				Identifier:                 GetBTreeNodeEntryIdentifier(btreeNodeEntryData, file.FormatType),
				DataIdentifier:             GetBTreeNodeEntryDataIdentifier(btreeNodeEntryData, file.FormatType),
				LocalDescriptorsIdentifier: GetBTreeNodeEntryLocalDescriptorsIdentifier(btreeNodeEntryData, file.FormatType),
				NodeLevel:                  btreeNodeLevel,
			}
		} else if btreeNodeLevel == 0 && btreeType == BTreeTypeBlock {
			// Leaf block b-tree node.
			btreeNodeEntries[i] = BTreeNode{
				Identifier: GetBTreeNodeEntryIdentifier(btreeNodeEntryData, file.FormatType),
				FileOffset: GetBTreeNodeEntryFileOffset(btreeNodeEntryData, false, file.FormatType),
				Size:       GetBTreeNodeEntrySize(btreeNodeEntryData, file.FormatType),
				NodeLevel:  btreeNodeLevel,
			}
		}
	}

	return btreeNodeEntries, nil
}

// GetBTreeNodeEntryIdentifier returns the Identifier of this b-tree node entry.
// References "The b-tree entries".
func GetBTreeNodeEntryIdentifier(btreeNodeEntryData []byte, formatType FormatType) Identifier {
	return GetIdentifierFromBytes(btreeNodeEntryData[:GetIdentifierSize(formatType)], formatType)
}

// GetIdentifierFromBytes returns the Identifier type from bytes.
func GetIdentifierFromBytes(identifierBytes []byte, formatType FormatType) Identifier {
	switch formatType {
	case FormatTypeANSI:
		return Identifier(binary.LittleEndian.Uint32(identifierBytes))
	default:
		// TODO - Reference [MS-PDF] that this is actually 32-bit
		return Identifier(binary.LittleEndian.Uint32(identifierBytes))
	}
}

// GetIdentifierSize returns the size of an Identifier.
func GetIdentifierSize(formatType FormatType) uint8 {
	switch formatType {
	case FormatTypeANSI:
		return 4
	default:
		return 8
	}
}

// Identifier represents a b-tree node identifier.
// TODO - Document the int types per use case and use separate types.
type Identifier int64

// Constants defining the special b-tree node identifiers.
const (
	IdentifierRootFolder   Identifier = 290
	IdentifierMessageStore Identifier = 33
	IdentifierNameToIDMap  Identifier = 97
)

// GetType returns the IdentifierType of this Identifier.
func (identifier Identifier) GetType() IdentifierType {
	// Bit-masking:
	// Use bitwise ANDing in order to extract a subset of the bits in the value.
	// 11111 (binary) = 0x1F (hex), which with bitwise ANDing extracts the first 5 bits.
	// See: https://www.rapidtables.com/convert/number/binary-to-hex.html
	return IdentifierType(identifier & 0x1F)
}

// IdentifierType represents the type of Identifier.
type IdentifierType uint8

// Constants defining the identifier types.
// References "Identifier types".
const (
	IdentifierTypeHID                     IdentifierType = 0
	IdentifierTypeInternal                IdentifierType = 1
	IdentifierTypeNormalFolder            IdentifierType = 2
	IdentifierTypeSearchFolder            IdentifierType = 3
	IdentifierTypeNormalMessage           IdentifierType = 4
	IdentifierTypeAttachment              IdentifierType = 5
	IdentifierTypeSearchUpdateQueue       IdentifierType = 6
	IdentifierTypeSearchCriteriaObject    IdentifierType = 7
	IdentifierTypeAssociatedMessage       IdentifierType = 8
	IdentifierTypeContentsTableIndex      IdentifierType = 10
	IdentifierTypeReceiveFolderTable      IdentifierType = 11
	IdentifierTypeOutgoingQueueTable      IdentifierType = 12
	IdentifierTypeHierarchyTable          IdentifierType = 13
	IdentifierTypeContentsTable           IdentifierType = 14
	IdentifierTypeAssociatedContentsTable IdentifierType = 15
	IdentifierTypeSearchContentsTable     IdentifierType = 16
	IdentifierTypeAttachmentTable         IdentifierType = 17
	IdentifierTypeRecipientTable          IdentifierType = 18
	IdentifierTypeSearchTableIndex        IdentifierType = 19
	IdentifierTypeLTP                     IdentifierType = 31
)

// GetBTreeNodeEntryFileOffset returns the file offset for this b-tree branch or leaf node.
// References "The b-tree entries".
func GetBTreeNodeEntryFileOffset(btreeNodeEntryData []byte, isBranchNode bool, formatType FormatType) int64 {
	if isBranchNode {
		switch formatType {
		case FormatTypeANSI:
			return int64(binary.LittleEndian.Uint32(btreeNodeEntryData[8 : 8+4]))
		default:
			return int64(binary.LittleEndian.Uint64(btreeNodeEntryData[16 : 16+8]))
		}
	} else {
		switch formatType {
		case FormatTypeANSI:
			return int64(binary.LittleEndian.Uint32(btreeNodeEntryData[4 : 4+4]))
		default:
			return int64(binary.LittleEndian.Uint64(btreeNodeEntryData[8 : 8+8]))
		}
	}
}

// GetBTreeNodeEntryDataIdentifier returns the node identifier of the data (in the block b-tree).
// References "The b-tree entries".
func GetBTreeNodeEntryDataIdentifier(btreeNodeEntryData []byte, formatType FormatType) Identifier {
	switch formatType {
	case FormatTypeANSI:
		return GetIdentifierFromBytes(btreeNodeEntryData[4:4+GetIdentifierSize(formatType)], formatType)
	default:
		return GetIdentifierFromBytes(btreeNodeEntryData[8:8+GetIdentifierSize(formatType)], formatType)
	}
}

// GetBTreeNodeEntryLocalDescriptorsIdentifier returns the identifier to the local descriptors in the block b-tree.
func GetBTreeNodeEntryLocalDescriptorsIdentifier(btreeNodeEntryData []byte, formatType FormatType) Identifier {
	switch formatType {
	case FormatTypeANSI:
		return GetIdentifierFromBytes(btreeNodeEntryData[8:8+GetIdentifierSize(formatType)], formatType)
	default:
		return GetIdentifierFromBytes(btreeNodeEntryData[16:16+GetIdentifierSize(formatType)], formatType)
	}
}

// GetBTreeNodeEntrySize returns the size of the data in the block b-tree leaf node entry.
// References "The b-tree entries".
func GetBTreeNodeEntrySize(btreeNodeEntryData []byte, formatType FormatType) uint16 {
	switch formatType {
	case FormatTypeANSI:
		return binary.LittleEndian.Uint16(btreeNodeEntryData[8 : 8+2])
	default:
		return binary.LittleEndian.Uint16(btreeNodeEntryData[16 : 16+2])
	}
}

// BTreeType represents either the node b-tree or block b-tree.
type BTreeType uint8

// Constants defining the b-tree types.
const (
	BTreeTypeNode BTreeType = iota
	BTreeTypeBlock
)

// WalkAndCreateBTree walks the b-tree and updates the given b-tree store.
func (file *File) WalkAndCreateBTree(btreeOffset int64, btreeType BTreeType, btreeStore BTreeStore) error {
	nodeLevel, err := file.GetBTreeNodeLevel(btreeOffset)

	if err != nil {
		return errors.WithStack(err)
	}

	nodeEntries, err := file.GetBTreeNodeEntries(btreeOffset, btreeType, nodeLevel)

	if err != nil {
		return errors.WithStack(err)
	}

	if nodeLevel > 0 {
		// Branch node entries.
		for i := 0; i < len(nodeEntries); i++ {
			nodeEntry := nodeEntries[i]

			if _, exists := btreeStore.Load(nodeEntry); exists {
				return errors.WithStack(ErrBTreeNodeConflict)
			}

			err = file.WalkAndCreateBTree(nodeEntry.FileOffset, btreeType, btreeStore)

			if err != nil {
				return errors.WithStack(err)
			}
		}
	} else {
		// Leaf node entries
		for i := 0; i < len(nodeEntries); i++ {
			if _, exists := btreeStore.Load(nodeEntries[i]); exists {
				return errors.WithStack(ErrBTreeNodeConflict)
			}
		}
	}

	return nil
}

// GetBTreeNode returns the node in the node or block b-tree with the given identifier.
func (file *File) GetBTreeNode(identifier Identifier, btreeStore BTreeStore) (BTreeNode, error) {
	btreeNode, found := btreeStore.Get(BTreeNode{Identifier: identifier})

	if !found {
		return BTreeNode{}, errors.WithStack(ErrBTreeNodeNotFound)
	}

	return btreeNode, nil
}

// GetNodeBTreeNode returns the node with the given identifier in the node b-tree.
func (file *File) GetNodeBTreeNode(identifier Identifier) (BTreeNode, error) {
	return file.GetBTreeNode(identifier, file.NodeBTree)
}

// GetBlockBTreeNode returns the node with the given identifier in the block b-tree.
func (file *File) GetBlockBTreeNode(identifier Identifier) (BTreeNode, error) {
	// Clear the least significant bit (LSB), which is reserved, but sometimes set.
	return file.GetBTreeNode(identifier&0xfffffffe, file.BlockBTree)
}

// GetDataBTreeNode searches the identifier in the node b-tree, then searches the data identifier in the block b-tree.
func (file *File) GetDataBTreeNode(identifier Identifier) (BTreeNode, error) {
	nodeBTreeNode, err := file.GetNodeBTreeNode(identifier)

	if err != nil {
		return BTreeNode{}, errors.WithStack(err)
	}

	return file.GetBlockBTreeNode(nodeBTreeNode.DataIdentifier)
}
