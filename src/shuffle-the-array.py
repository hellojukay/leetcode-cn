class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        result = []
        list1 = nums[0:n]
        list2 = nums[n:]
        for index in range(len(list1)):
            result.append(list1[index])
            result.append(list2[index])
        return result
