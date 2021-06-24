from pathlib import Path
import numpy as np
from typing import Tuple
import torch

class Utils:
    @staticmethod
    def load_file(path : Path or str) -> list:
        with open(path, "r") as f:
            raw_data = f.readlines()
        return raw_data

    
    @staticmethod
    def preprocess(input : str) -> Tuple[torch.Tensor, torch.Tensor]:
        toarray = lambda x : np.array(list(map(int, list(x)))).reshape(9, 9)

        X, y = input.split()
        X, y = toarray(X), toarray(y)
        X = torch.tensor(X, dtype=torch.long)[None, :, :]
        y = torch.tensor(y, torch.long)[None, :, :]
        return X, y
    
    @staticmethod
    def pretty_print(board : np.array or torch.tensor or list):
        if isinstance(board, torch.Tensor):
            board = board.data.cpu().numpy()[0, ...].reshape(9, 9)
        else:
            board.reshape(9, 9)
            
        line = ''.join(['-' for _ in range(20)])
        print(line)
        for row in list(board):
            print('  '.join(row))
        print(line)        
