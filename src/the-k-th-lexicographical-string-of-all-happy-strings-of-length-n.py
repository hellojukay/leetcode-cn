class Solution:
    def getHappyString(self, n: int, k: int) -> str:
        def dfs(current: str,n: int, result: List[str]) -> List[str]:
            if n == 0:
                result.append(current)
                return result
            if len(current) == 0:
                dfs(current+"a",n-1,result)
                dfs(current+"b",n-1,result)
                dfs(current+"c",n-1,result)
            else:
                for c in "abc":
                    if c != current[-1]:
                        dfs(current + c,n-1,result)
            return result
        result = dfs("",n,[])
        result = sorted(result)
        if len(result) >= k:
            return result[k-1]
        return ""
