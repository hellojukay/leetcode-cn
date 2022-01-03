class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        m = {}
        for ch in magazine:
            n = m.get(ch, 0)
            m[ch] = n + 1
        for ch in ransomNote:
            n = m.get(ch, 0)
            if n == 0:
                return False
            m[ch] = n - 1
        return True
