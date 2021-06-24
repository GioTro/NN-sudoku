from pathlib import Path
import numpy as np
from typing import Tuple
import torch

class Utils:
    @staticmethod
    def load_file(path : Path or str) -> list:
        with open(path, 'r') as f:
            raw_data = f.readlines()
        return raw_data

    
    @staticmethod
    def preprocess(inp : str) -> Tuple[torch.Tensor, torch.Tensor]:
        toarray = lambda x : np.array(list(map(int, list(x)))).reshape(9, 9)

        X, y = inp.split()
        X, y = toarray(X), toarray(y)
        X = torch.tensor(X, dtype=torch.long)[None, :, :]
        y = torch.tensor(y, torch.long)[None, :, :]
        return X, y
        
    @staticmethod
    def postprocess(X, y : torch.Tensor) -> np.array:
    	pass
    
    @staticmethod
    def pretty_print(board : np.array or torch.tensor or list):
        if isinstance(board, torch.Tensor):
            board = board.data.cpu().numpy()[0, ...].reshape(9, 9)
        else:
            board.reshape(9, 9)
            
        line = ''.join(['-' for _ in range(20)])
        print(line)
        for idx, row in enumerate(list(board)):
        	if idx % 3 && idx != 0:
        		print()
            print(''.join(row[:3]), '  ', ''.join(row[3:6]), '  ', ''.join(row[6:]))
        print(line)        
