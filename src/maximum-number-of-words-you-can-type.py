

class Solution:
    def canBeTypedWords(self, text: str, brokenLetters: str) -> int:
        s = set(brokenLetters.lower())
        text = text.lower()
        array = text.split(" ")
        count = len(array)
        for t in array:
            for ch in t:
                if ch in s:
                    count = count - 1
                    break
        return count


s = Solution()
print(s.canBeTypedWords('leet code', 'e'))
