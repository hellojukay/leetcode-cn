class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        if root == None:
            return True
        currentLevel = []
        nextLevel = []
        currentLevel.append(root)
        while currentLevel:
            # 检查一下当前的层是否对称
            i = 0
            j = len(currentLevel)-1
            while i <= j:
                if currentLevel[i] == None and currentLevel[j] != None:
                    return False
                if currentLevel[i] != None and currentLevel[j] == None:
                    return False
                if currentLevel[i] == None and currentLevel[j] == None:
                    i = i + 1
                    j = j - 1
                    continue
                if currentLevel[i].val != currentLevel[j].val:
                    return False
                i = i + 1
                j = j - 1
            for node in currentLevel:
                if not node:
                    continue
                nextLevel.append(node.left)
                nextLevel.append(node.right)
            currentLevel = nextLevel
            nextLevel = []
        return True



