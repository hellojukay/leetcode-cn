# https://leetcode-cn.com/problems/reverse-only-letters/
class Solution:
    def reverseOnlyLetters(self, S: str) -> str:
        chars = []
        for ch in S:
            if ch.isalpha():
                chars.append(ch)
        chars.reverse()
        for i in range(len(S)):
            if not S[i].isalpha():
                chars.insert(i,S[i])
        return ''.join(chars)

s = Solution()
print(s.reverseOnlyLetters("ab-cd"))
