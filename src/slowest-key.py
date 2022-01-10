class Solution:
    def slowestKey(self, releaseTimes: List[int], keysPressed: str) -> str:
        m =  0
        c = 0
        i = 0
        result = ""
        while i < len(keysPressed):
            if i == 0:
                c = releaseTimes[i]
            else:
                c = releaseTimes[i] - releaseTimes[i-1]
            if c > m:
                m = c
                result = keysPressed[i]
            else:
                if c ==m:
                    if ord(keysPressed[i]) > ord(result):
                        m = c 
                        result = keysPressed[i]
            i = i + 1 
        return result
