class Solution:
    def customSortString(self, order: str, s: str) -> str:
        m = {}
        i = 0 
        while i < len(order):
            m[order[i]] = i
            i = i + 1
        suffix = ""
        prefix = ""
        i = 0 
        while i < len(s):
            if m.get(s[i],None) != None:
                prefix = prefix + s[i]
            else:
                suffix = suffix + s[i]
            i = i + 1
        # sort prefix
        def mysort(word: str,m: dict[str,int]):
            if len(word) <= 1:
                return word
            if len(word) == 2:
                if m.get(word[0],-1) < m.get(word[1],-1):
                    return word
                return word[1] + word[0]
            mid = ""
            left = ""
            right = ""
            x = int(len(word)/2)
            i = 0
            while i < len(word):
                if m.get(word[i],-1) < m.get(word[x],-1):
                    left = left + word[i]
                    i = i + 1
                    continue
                if m.get(word[i],-1) > m.get(word[x],-1):
                    right = right + word[i]
                    i = i + 1
                    continue
                mid = mid + word[i]
                i = i + 1
            return mysort(left,m) + mid + mysort(right,m)

        return mysort(prefix,m) + suffix
s = Solution()
print(s.customSortString("cba","abcd"))
