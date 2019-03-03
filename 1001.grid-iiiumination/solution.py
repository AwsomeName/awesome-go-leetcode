# -*- coding: utf-8 -*-

# class Solution(object):
#     def gridIllumination(self, N, lamps, queries):
#         """
#         :type N: int
#         :type lamps: List[List[int]]
#         :type queries: List[List[int]]
#         :rtype: List[int]
#         """
class Solution(object):
    def gridIllumination(self, N, lamps, queries):
        from collections import Counter
        lamps = map(tuple, lamps)
        light_set = set(lamps)
        # one cell may be illuminated by many lamps
        horizontal = Counter()
        vertical = Counter()
        l_oblique = Counter()
        r_oblique = Counter()
        for x, y in lamps:
            horizontal[x] += 1
            vertical[y] += 1
            l_oblique[x+y] += 1
            r_oblique[y-x] += 1

        res = []
        for x,y in queries:
            if x in horizontal or y in vertical or x+y in l_oblique or y-x in r_oblique:
                res.append(1)
            else:
                res.append(0)
            for dx,dy in [[-1,-1], [-1,0], [-1,1], [0,-1], [0,0], [0,1], [1,-1], [1,0], [1,1]]:
                xpos, ypos = x + dx, y + dy
                if (xpos, ypos) in light_set:
                    light_set.remove((xpos, ypos))
                    horizontal[xpos] -= 1
                    if horizontal[xpos] == 0: del horizontal[xpos]

                    vertical[ypos] -= 1
                    if vertical[ypos] == 0: del vertical[ypos]

                    l_oblique[xpos+ypos] -= 1
                    if l_oblique[xpos+ypos] == 0: del l_oblique[xpos+ypos]

                    r_oblique[ypos-xpos] -= 1
                    if r_oblique[ypos-xpos] == 0: del r_oblique[ypos-xpos]
        return res
