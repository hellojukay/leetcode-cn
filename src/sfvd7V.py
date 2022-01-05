class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        m = {}
        for word in strs:
            s = "".join(sorted(word))
            tmp = m.get(s, [])
            tmp.append(word)
            m[s] = tmp
        return [m.get(k) for k in m.keys()]
