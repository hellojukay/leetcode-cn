class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        length = len(s)
        m = {}
        i = 0
        while i < length:
            c = m.get(s[i],None)
            if not c:
                if t[i] in set(m.values()):
                    return False
                m[s[i]] = t[i]
                i = i + 1
                continue
            if c == t[i]:
                i = i + 1
                continue
            if  c != t[i]:
                return False
            i = i + 1
        return True
               
