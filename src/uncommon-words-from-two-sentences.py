class Solution:
    def uncommonFromSentences(self, s1: str, s2: str) -> List[str]:
        def toMap(sentence: str) -> dict[str, int]:
            m = {}
            for s in sentence.split(" "):
                n = m.get(s, 0)
                m[s] = n + 1
            return m

        def get(m1: dict[str, int], m2: dict[str, int], result):
            for key in m1.keys():
                if m1.get(key) == 1 and (not key in m2):
                    result.append(key)
            return result

        m1 = toMap(s1)
        m2 = toMap(s2)
        result = []
        result = get(m1, m2, result)
        result = get(m2, m1, result)
        return result
