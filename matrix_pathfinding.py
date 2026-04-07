def multiply_matrices(A, B):
    """Multiplies two matrices without external libraries."""
    result = [[0 for _ in range(len(B[0]))] for _ in range(len(A))]
    for i in range(len(A)):
        for j in range(len(B[0])):
            for k in range(len(B)):
                result[i][j] += A[i][k] * B[k][j]
    return result

def simple_greedy_path(grid, start, end):
    """A basic greedy approach to find a path in a coordinate grid."""
    curr_x, curr_y = start
    path = [start]
    while (curr_x, curr_y) != end:
        if curr_x < end[0]: curr_x += 1
        elif curr_y < end[1]: curr_y += 1
        path.append((curr_x, curr_y))
    return path
