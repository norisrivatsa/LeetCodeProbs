class Solution(object):
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        dp = [[] for _ in range(n + 1)]
        dp[0] = [""]

        for i in range(1, n + 1):
            current = []

            for left in range(i):
                right = i - 1 - left

                for inside in dp[left]:
                    for outside in dp[right]:
                        current.append("(" + inside + ")" + outside)

            dp[i] = current

        return dp[n]