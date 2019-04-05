# A Trie is a Tree used to store prefixes of words

class Node:
    def __init__(self, value=None):
        self.children = {}
        self.value = value

    @property
    def end(self):
        return self.value is None


class Trie:
    def __init__(self):
        self.roots = {}

    def insert_word(self, word):
        c = word[0]
        node = self.roots.setdefault(c, Node(c))

        for c in word[1:]:
            node = node.children.setdefault(c, Node(c))

        node.children.setdefault("*", Node())

    def all_words(self):
        return self.begins_with("")

    def begins_with(self, word):
        children = self.roots
        i = 0
        while children and i < len(word):
            node = children.get(word[i], Node())
            children = node.children
            i += 1

        words = []
        for child in children.values():
            self._dfs(child, words, word)

        return words

    def _dfs(self, node, words, word):
        if not node.end:
            for child in node.children.values():
                self._dfs(child, words, word + node.value)
        else:
            words.append(word)


trie = Trie()
words = ["ola", "olar", "omar", "stenio", "elson", "stenioelson", "alo"]
for word in words:
    trie.insert_word(word)

print("all words inserted")
print(trie.all_words())

print("all words starting with o")
print(trie.begins_with("o"))

print("all words starting with ola")
print(trie.begins_with("ola"))

print("all words starting with t")
print(trie.begins_with("t"))
