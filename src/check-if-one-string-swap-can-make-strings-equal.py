class Solution:
    def areAlmostEqual(self, s1: str, s2: str) -> bool:
        if s1 == s2:
            return True
        if len(s1) != len(s2):
            return False
        if "".join(sorted(s1)) != "".join(sorted(s2)):
            return False
        length = len(s2)
        i = 0
        count = 0
        while i < length:
            if s1[i] != s2[i]:
                count = count + 1
            i = i + 1
        if count <= 2:
            return True
        return False
