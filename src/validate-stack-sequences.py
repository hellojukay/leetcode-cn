class Solution:
    def validateStackSequences(self, pushed: List[int], popped: List[int]) -> bool:
        if len(pushed) != len(popped):
            return False
        stack = []
        i = 0
        j = 0
        while True:
            if len(stack):
                if stack[-1] == popped[j]:
                    j = j + 1
                    stack.pop()
                else:
                    if i == len(pushed):
                        return False
                    else:
                        stack.append(pushed[i])
                        i = i + 1
            else:
                if j == len(pushed):
                    break
                stack.append(pushed[i])
                i = i + 1
        return len(stack) == 0 and j == i and j == len(pushed)
