package main

import (
	"io/ioutil"
	"log"
)

func _handleDir(origRoot string, root string, parent string, tree *RadixTree) {
	if len(parent) > 0 {
		root = parent + "/" + root
	}
	entries, err := ioutil.ReadDir(root)
	if err != nil {
		log.Printf("Unable to iterate directory: %s\n", root)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			_handleDir(origRoot, entry.Name(), root, tree)
		} else {
			name := root + "/" + entry.Name()
			name = name[len(origRoot):]

			log.Printf("Adding file: %s", name)
			contents, err := ioutil.ReadFile(entry.Name())
			if err == nil {
				tree.insert(name, contents)
			}
		}
	}
}

func createRadixTreeFromDir(root string) *RadixTree {
	log.Printf("Creating radix tree from root: %s\n", root)

	tree := radixTreeCreate(1)
	_handleDir(root, root, "", tree)
	tree.finalise()

	return tree
}
