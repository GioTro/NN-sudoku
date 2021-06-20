from genericpath import exists
import pathlib
from typing import Tuple
import random
import numpy as np
from pathlib import Path
from tqdm import tqdm

"""
Notes:
	- This is probably slow will only run ones.

	- Consider skipping making multiple files and dump everything in one.

	- The difficulty setting is not 'true' sudoku rated.
"""


class Board:
    def __init__(self):
        self.board = np.zeros(shape=(9, 9))
        self.solved = False
        self.solution_count = 0

    def reset(self):
        self.__init__()

    def valid_move(self, row : int, col : int, n : int) -> bool:
        r_o, c_o = (row//3)*3, (col//3)*3
        if n == 0:
            return False
        if n in self.board[row, :]:
            return False
        if n in self.board[:, col]:
            return False
        if n in self.board[r_o:r_o+3, c_o:c_o+3]:
            return False
        return True

    def check_correct(self) -> bool:
        for row in 9:
            for col in 9:
                if not self.valid_move(row, col, self.board[row, col]):
                    return False
        return True

    def copy(self) -> np.array:
        return np.copy(self.board)


class SudokuSolver(Board):
    # self.board : np.array
    def __init__(self, difficulty : int):
        Board.__init__(self)
        self.difficulty = difficulty

    def recursive(self, condition):
        if condition():
            return

        for row in range(9):
            for col in range(9):
                if super().board[row, col] == 0:
                    for n in self.rgen(range(1, 10)):
                        if condition() or not super().valid_move(row, col, n):
                            continue

                        super().board[row, col] = n
                        self.recursive()
                        super().board[row, col] = 0 # else backtrack
                    return
        super().solution_count += 1

    def find_first(self):
        self.recursive(lambda : super().solution_count > 0)
        return super().solution_count

    def find_all(self):
        self.recursive(lambda : False)
        return super().solution_count

    def generate_valid_board(self):
        self.recursive(self.find_first())

    def remove_elements(self):
        tries = self.idxgen()
        n_left = 9*9
        condition = n_left - self.difficulty

        for (x, y) in tries:
            if condition > n_left:
                return

            before = super().board[x, y]
            super().board[x, y] = 0
            if self.find_all() > 1:
                super().board[x, y] = before
            else:
                n_left -= 1

        return (condition - n_left) == 0

    @ staticmethod
    def idxgen(shape = (9, 9))-> iter:
        x, y = shape
        idx = [(xx, yy) for xx in range(x) for yy in range(y)]
        random.shuffle(idx)
        return iter(idx)

    @staticmethod
    def rgen(r : range) -> iter:
        g = [i for i in r]
        random.shuffle(g)
        return iter(g)

    def new_instance(self) -> Tuple[np.array, np.array]:
        super().reset()
        self.generate_valid_board()
        assert super().check_correct()
        y = super().copy()
        self.remove_elements()
        X = super().board
        return X, y


class Boardgenerator(SudokuSolver):
    def __init__(self, difficulty : int, end_at : int):
        super(Boardgenerator, self).__init__(difficulty)
        self.end_at = end_at
        # self.encoder = lambda x : (np.eye(10)[x.flatten()]).reshape(-1, 1)

    def __iter__(self) -> iter:
        return self

    def __next__(self) -> Tuple[list, list]:
        if self.end_at > 0:
            X, y = super().new_instance()
            self.end_at -= 1
            return X.flatten().tolist(), y.flatten().tolist()
        else:
            raise StopIteration

if __name__ == '__main__':
    root = Path(__file__).parent.absolute()
    do = {'train': int(1e5), 'validation': int(1e3), 'test': int(1e3)}

    paths = [
        root / '..' / 'dataset' / 'train',
        root / '..' / 'dataset' / 'test',
        root / '..' / 'dataset' / 'val',
    ]

    for p in paths:
        p.mkdir(exist_ok = True)
        b = Boardgenerator(40, do[p.stem])
        b = iter(b)
        idx = 0
        for X, y in tqdm(next(b)):
            fname = p / f'{p.stem}_{idx}.txt'
            with open(fname, 'w') as f:
                f.write(', '.join(map(str, X)) + '\n')
                f.write(', '.join(map(str, y)))
