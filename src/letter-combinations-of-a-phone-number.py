class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not len(digits):
            return []
        def chars(n: int)-> List[str]:
            if n == 1:
                return [] 
            if n == 2:
                return ["a","b","c"]
            if n == 3:
                return ["d","e","f"]
            if n == 4:
                return ["g","h","i"];
            if n == 5:
                return ["j","k","l"]
            if n == 6:
                return ["m","n","o"]
            if n == 7:
                return ["p","q","r","s"]
            if n == 8:
                return ["t","u","v"]
            if n == 9:
                return ["w","x","y","z"]
            return []
        def dfs(current: str,last: str,result: List[str]) -> None:
            if not len(last):
                result.append(current)
                return
            n = int(last[0])
            letters = chars(n)
            for s in letters:
                dfs(current + s,last[1:],result)

        result = []
        dfs("",digits,result)
        return result
