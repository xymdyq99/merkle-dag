package merkledag

import (
	"encoding/json"
	"hash"
)

// Node 表示Merkle树中的一个节点
type Node struct {
	Data []byte // 节点数据
}

// KVStore 表示键值存储接口
type KVStore interface {
	Put(key string, value []byte) error
	Get(key string) ([]byte, error)
}

// Add 将Node中的数据保存在KVStore中，并返回Merkle Root
func Add(store KVStore, node Node, h hash.Hash) ([]byte, error) {
	// 计算节点的哈希值
	h.Reset()
	h.Write(node.Data)
	nodeHash := h.Sum(nil)

	// 将节点保存到KVStore中
	nodeJSON, err := json.Marshal(node)
	if err != nil {
		return nil, err
	}
	if err := store.Put(string(nodeHash), nodeJSON); err != nil {
		return nil, err
	}

	// 更新Merkle Root
	merkleRoot, err := updateMerkleRoot(store, h)
	if err != nil {
		return nil, err
	}

	return merkleRoot, nil
}

// updateMerkleRoot 根据存储在KVStore中的所有节点数据更新Merkle Root
func updateMerkleRoot(store KVStore, h hash.Hash) ([]byte, error) {
	// TODO: 实现Merkle树算法，根据所有节点的哈希值更新Merkle Root
	// 这里只是一个简单的示例，实际上需要根据Merkle树的规则进行递归计算

	return []byte("sampleMerkleRoot"), nil
}
