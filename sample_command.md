**复现步骤**

1. SASRec

   ```python
   # 训练模型
   python train_SASRec.py --device=cuda:7 --num_epochs=10000 --dropout_rate=0.2 --dataset=ml-1m --save_dir=saved --inference_only=false --maxlen=200 --hidden_units=64
   # 存在过拟合，程度比较轻
   # 读出训练好的模型进行测试
   python train_SASRec.py --device=cuda:7 --num_epochs=10000 --dataset=ml-1m --save_dir=saved --state_dict_path=saved/SASRec.epoch=10000.lr=0.001.layer=2.head=1.hidden=64.maxlen=200.reg_weight=0.0.pth --inference_only=true --maxlen=200 --hidden_units=64
     
   # old
   python train_SASRec.py --device=cuda:6 --num_epochs=10000 --dataset=ml-1m --save_dir=saved --state_dict_path=saved/SASRec.epoch=10000.lr=0.001.layer=2.head=1.hidden=50.maxlen=200.pth --inference_only=true --maxlen=200
   ```

2. LightGCN

   ```python
   # 训练模型
   python train_LightGCN.py --device=cuda:3 --epochs=1000 --datapath=dataset/movielens1m_m1 --save_dir=saved --decay=1e-4 --lr=0.001 --layer=3 --seed=2020 --dataset="ml-1m" --topks="[10]" --recdim=64 --pop_reg_weight=0.0
   # 读出训练好的模型进行测试
   python train_LightGCN.py --mode=test --device=cuda:3 --epochs=1000 --save_model=False --datapath=dataset/movielens1m_m1 --save_dir=saved --decay=1e-4 --lr=0.001 --layer=3 --seed=2020 --dataset="ml-1m" --topks="[10]" --recdim=64 --LOAD=1
   ```

3. SimpleX

   ```python
   # 训练模型
   python train_SimpleX.py --config ./config/SimpleX_movielens1m_m1 --expid SimpleX_movielens1m_m1 --gpu 6
   # 读出训练好的模型进行测试
   python train_SimpleX.py --mode test --config ./config/SimpleX_movielens1m_m1 --expid SimpleX_movielens1m_m1 --gpu 6
   ```


