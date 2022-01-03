class Solution:
    def findRelativeRanks(self, score: List[int]) -> List[str]:
        s = sorted(score, reverse=True)
        i = 0
        m = {}
        while i < len(s):
            m[s[i]] = i
            i = i + 1

        def get(m: dict[int, str], n: int) -> str:
            x = m.get(n)
            if x == 0:
                return "Gold Medal"
            if x == 1:
                return "Silver Medal"
            if x == 2:
                return "Bronze Medal"
            return str(x + 1)

        result = []
        for n in score:
            s = get(m, n)
            result.append(s)
        return result
