def solution(m, n, puddles):
    MOD = 10**9 + 7

    # 1-indexed 접근을 위해 (n+1)x(m+1) 크기의 배열 생성
    dp = [[0] * (m+1) for _ in range(n+1)]
    dp[1][1] = 1

    # 물웅덩이 좌표를 (row, col) 형태의 set으로 변환
    blocked = {(y, x) for x, y in puddles}

    for i in range(1, n+1):
        for j in range(1, m+1):
            # 시작점 또는 물웅덩이는 건너뜀
            if (i, j) == (1, 1) or (i, j) in blocked:
                continue
            dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % MOD

    return dp[n][m]