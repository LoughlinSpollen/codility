import unittest

# A non-empty array A consisting of N integers is given.

# The leader of this array is the value that occurs in more than half of the elements of A.

# An equi leader is an index S such that 0 ≤ S < N − 1 and two sequences 
# A[0], A[1], ..., A[S] and A[S + 1], A[S + 2], ..., A[N − 1] 
# have leaders of the same value.

# For example, given array A such that:

#     A[0] = 4
#     A[1] = 3
#     A[2] = 4
#     A[3] = 4
#     A[4] = 4
#     A[5] = 2
# we can find two equi leaders:

# 0, because sequences: (4) and (3, 4, 4, 4, 2) have the same leader, whose value is 4.
# 2, because sequences: (4, 3, 4) and (4, 4, 2) have the same leader, whose value is 4.
# The goal is to count the number of equi leaders.

# Write a function:

# def solution(A)

# that, given a non-empty array A consisting of N integers, returns the number of equi leaders.

# For example, given:

#     A[0] = 4
#     A[1] = 3
#     A[2] = 4
#     A[3] = 4
#     A[4] = 4
#     A[5] = 2
# the function should return 2, as explained above.

# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [1..100,000];
# each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].



def solution(A):

    arr = A
    if len(arr) < 2:
        return len(arr)

    def most_common_val(arr):
        most_common_val = None
        # occurance count of values
        map = dict()
        for v in arr:
            if v in map:
                map[v] += 1
            else:
                map[v] = 1
            if most_common_val is None or map[v] > map[most_common_val]:
                most_common_val = v

        # unique leader
        unique_leader = True
        leader_occurance = map[most_common_val]
        for k in map:
            if k == most_common_val:
                continue
            if map[k] == leader_occurance:
                unique_leader = False
                break
        
        if not unique_leader:
            return None
        
        return most_common_val

    equi_leader_count = 0
    for k, _ in enumerate(arr):
        if (k+1) < len(arr):            
            left_leader = most_common_val(arr[:k+1]) 
            right_leader = most_common_val(arr[k+1:])

            if left_leader is None or right_leader is None:
                continue 

            if left_leader == right_leader:
                equi_leader_count += 1        
        
    return equi_leader_count
    


class TestEquiLeader(unittest.TestCase):
    def test_ex(self):
        self.assertEqual(solution([4, 3, 4, 4, 4, 2]), 2)

    def test_min(self):
        self.assertEqual(solution([1]), 1)
        self.assertEqual(solution([1, 1]), 1)

    def test_max(self):
        self.assertEqual(solution([1, 2, 3, 4, 5, 6, 7, 8, 9]), 0)
        self.assertEqual(solution([1, 1, 1, 1, 1, 1, 1, 1, 1]), 8)
