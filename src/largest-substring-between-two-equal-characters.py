class Solution:
    def maxLengthBetweenEqualCharacters(self, s: str) -> int:
        m = {}
        i = 0
        result = -1
        while i < len(s):
            index = m.get(s[i],-1)
            if index >= 0:
                if i - index - 1 > result:
                    result = i - index -1
            else:
                m[s[i]] = i
            i = i + 1
        return result

