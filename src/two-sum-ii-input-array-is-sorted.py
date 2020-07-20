class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        dic = {}
        for i , n in enumerate(numbers):
            dic[n] = i + 1
        for i , n in enumerate(numbers):
            index  = dic.get(target -n,-1)
            if index == -1:
                continue
            if i + 1 == index:
                continue
            else:
                return [i+1,index]

