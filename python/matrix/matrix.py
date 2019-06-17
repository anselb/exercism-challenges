class Matrix(object):
    def __init__(self, matrix_string):
        rows = matrix_string.splitlines()
        self.matrix = []
        for row in rows:
            # Split row into number strings e.g. "10"
            row = row.split()
            # Convert each value in a row from int to str
            for index in range(len(row)):
                row[index] = int(row[index])
            # Append row of numbers
            self.matrix.append(row)

    def row(self, index):
        return self.matrix[index - 1]

    def column(self, index):
        return [row[index - 1] for row in self.matrix]
