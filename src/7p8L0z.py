class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        def dfs(arr: List[int], current: List[int],result: List[List[int]]) -> List[List[int]]:
            if len(arr) == 0:
                result.append(current)
                return result
            length = len(arr)
            i = 0
            while i < length:
                tmp = current.copy()
                tmp.append(arr[i])
                tmp2 = arr.copy()
                tmp2 = arr[:i] + arr[i+1:]
                dfs(tmp2,tmp,result)
                i = i + 1
            return result
        result = []
        dfs(nums,[],result)
        m = {}
        for s in result:
            m[str(s)] = s
        return list(m.values())
