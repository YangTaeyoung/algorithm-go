def check(stones, minus, k):
    count = 0
    for stone in stones:
        if stone - minus <= 0:
            count += 1
            continue
        if count >= k:
            return True
        count = 0

    return count >= k


def solution(stones, k):
    maxStone = max(stones)
    minStone = 0

    answer = 0
    while minStone <= maxStone:
        mid = (maxStone + minStone)//2
        if check(stones, mid, k):
            answer = mid
            maxStone = mid - 1
        else:
            minStone = mid + 1

    return answer