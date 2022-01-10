class Solution:
    def isPrefixOfWord(self, sentence: str, searchWord: str) -> int:
        words = sentence.split(" ")
        i = 0
        while i < len(words):
            if words[i].startswith(searchWord):
                return i + 1
            i = i + 1
        return -1

