class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        m = {}
        i = 0 
        while i < len(numbers):
            m[numbers[i]] = i
        i = 0
        while i < len(numbers):
            j = m.get(target-numbers[i],None)
            if j and j > i:
                return [i,j]
            i = i + 1
        return []