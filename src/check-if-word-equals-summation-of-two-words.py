class Solution:
    def isSumEqual(self, firstWord: str, secondWord: str, targetWord: str) -> bool:
        def toInt(s: str) -> int:
            result = ""
            for ch in s:
                result = result + str(ord(ch) - ord('a'))
            return int(result)
        return toInt(firstWord) + toInt(secondWord) == toInt(targetWord)
