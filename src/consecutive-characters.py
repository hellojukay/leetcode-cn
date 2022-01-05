class Solution:
    def maxPower(self, s: str) -> int:
        m = 1
        t = 1
        length = len(s)
        if length < 1:
            return length
        i = 1
        while i < length:
            if s[i] == s[i-1]:
                t = t + 1
                if t > m:
                    m = t
            else:
                t = 1
            i = i + 1
        if t > m:
            m = t
        return m
