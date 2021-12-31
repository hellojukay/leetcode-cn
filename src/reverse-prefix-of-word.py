class Solution:
    def reversePrefix(self, word: str, ch: str) -> str:
        i = -1
        j = 0
        length = len(word)
        while j < length:
            if word[j] == ch:
                i = j
                break
            j = j + 1
        if i > 0:
            return word[:j+1][::-1] + word[j+1:]
        return word


s = Solution()
print(s.reversePrefix("xyxzxe", "z"))
