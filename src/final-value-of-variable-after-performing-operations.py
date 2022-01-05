class Solution:
    def finalValueAfterOperations(self, operations: List[str]) -> int:
        n = 0
        for word in operations:
            if "++" in word:
                n = n + 1
                continue
            if "--" in word:
                n = n - 1
        return n
