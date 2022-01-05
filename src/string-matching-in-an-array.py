class Solution:
    def stringMatching(self, words: List[str]) -> List[str]:
        result = set()
        i = 0
        while i < len(words):
            j = 0
            while j < len(words):
                if i == j:
                    j = j + 1
                    continue
                if words[i] in words[j]:
                    result.add(words[i])
                j = j + 1
            i = i + 1
        return list(result)
