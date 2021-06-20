import shutil
import torch
import torch.nn as nn
import pytorch_lightning as pl
from torch.optim import optimizer
import dataloader as dl

class NNSudoku(pl.LightningModule):
    def __init__(
        self,
        T_max = 0,
        eta_min = 0,
        batch_size = 0,
        num_workers = 0,
    ):
        super(NNSudoku, self).__init__()

        self.dl = dl.return_loaders(
            batch_size = batch_size, num_workers=num_workers, shuffle = True,
        )

        self.logger_kwargs = {
            "prog_bar": True, 
            "logger": True, 
            "on_step": True, 
            "on_epoch": True,
        }

        self.model_1 = nn.Sequential(
            nn.Conv2d(1, 512, 3, bias = True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True)
        )

        self.model_2 = nn.Sequential(
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
            nn.Conv2d(512, 512, 3, bias=True, padding=False),
            nn.BatchNorm2d(512),
            nn.ReLU(True),
        )

        self.model_out = nn.Conv2d(512, 9, 1)
    
    def forward(self, X):
        conv_1 = self.model_1(X)
        conv_2 = self.model_2(conv_1)
        out = self.model_out(conv_2)
        return out
    
    def training_step(self, batch, batch_idx):
        X, y = batch
        y_pred = self.forward(X)
        loss = nn.CrossEntropyLoss(y, y_pred)
        self.log(
            "train_loss", loss, **self.logger_kwargs
        )
        return loss
    
    def validation_step(self, batch, batch_idx):
        X, y = batch
        y_pred = self.forward(X)
        loss = nn.CrossEntropyLoss(y, y_pred)
        self.log(
            "train_loss", loss, **self.logger_kwargs
        )
        return loss
    
    def test_step(self, batch, batch_idx):
        X, y = batch
        y_pred = self.forward(X)
        loss = nn.CrossEntropyLoss(y, y_pred)
        self.log(
            "test_loss", loss, **self.logger_kwargs
        )
        return loss
    
    def configure_optimizers(self):
        optimizer = torch.optim.Adam(
            self.parameters(),
            self.optimizer_param['Adam'] 
        )
        scheduler = torch.optim.lr_scheduler.CosineAnnealingLR(
            optimizer=optimizer,
            eta_min=self.eta_min,
            T_max = self.T_max
        )
        return [optimizer], [scheduler]
    
    def on_epoch_end(self):
        return super().on_epoch_end()

    def train_dataloader(self):
        return self.data_loaders["train"]

    def test_dataloader(self):
        return self.data_loaders["test"]

    def val_dataloader(self):
        return self.data_loaders["validation"]
    

    


