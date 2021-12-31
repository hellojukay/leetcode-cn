class Solution:
    def count(self, s: str) -> int:
        ct = 0
        for c in s:
            if c == 'a' or c == 'o' or c == 'e' or c == 'i' or c == 'u':
                ct = ct + 1
        return ct

    def halvesAreAlike(self, s: str) -> bool:
        length = len(s)
        if length == 0:
            return False
        s = s.lower()
        s1 = s[0:int(length/2)]
        return self.count(s1) * 2 == self.count(s)
