class Solution:
    def checkAlmostEquivalent(self, word1: str, word2: str) -> bool:
        def check(m1: dict,m2: dict) -> bool:
            for key in m1.keys():
                v1 = m1.get(key)
                v2 = m2.get(key,0)
                if abs(v1-v2) > 3:
                    return False
            return True
        m1 = {}
        for ch in word1:
            n = m1.get(ch,0)
            m1[ch] = n + 1
        m2 = {}
        for ch in word2:
            n = m2.get(ch,0)
            m2[ch] = n + 1
        if check(m1,m2)  and check(m2,m1):
            return True
        return False
