from pathlib import Path
import numpy as np
from typing import Tuple
import torch

class Utils:
    @staticmethod
    def load_file(path : Path or str) -> Tuple[np.array, np.array]:
        with open(path, 'r') as f:
            X = f.readline()
            y = f.readline()
        
        X = np.array(list(map(int, X.split(','))))
        y = np.array(list(map(int, X.split(','))))

        return X.reshape(9, 9), y.reshape(9, 9)
    
    @staticmethod
    def pre_process(X : np.array, y : np.array) -> Tuple[torch.Tensor, torch.Tensor]:
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
