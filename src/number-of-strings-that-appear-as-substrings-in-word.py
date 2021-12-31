class Solution:
    def numOfStrings(self, patterns: List[str], word: str) -> int:
        count = 0
        for s in patterns:
            if s in word:
                count = count + 1
        return count
