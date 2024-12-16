# go-disoc
 
The experiment program for [DISOC](https://ieeexplore.ieee.org/abstract/document/10438405/), structure of the repository is as follows:

```
├─cmd
│  ├─generator              # generate key pairs for different application
│  ├─node-wd                # simulate node program without DISOC
│  ├─node                   # DISOC simulate node program
│  ├─oracle-wd              # oracle program without DISOC
│  ├─oracle                 # oracle program
│  └─ring                   # ring signature test program
├─contracts                 # evm contracts
├─internal
│  ├─crypto                 # signature and encrypt repo
│  └─ethereum               # ethereum call client
└─pkg
    └─totp                  # TOTP repo
```

`node-wd` and `oracle-wd` are a pair of programs that test the overhead **without DISOC framework**, corresponding to the experiment in Section VI-B.

Before running the program, you need to deploy the OracleProxy or OracleProxyWithoutDISOC contract on any EVM blockchain(depending on your specific experimental needs)

If you used the code from this repository, please cite it using the following BibTeX entry:

```
@article{xiong2024disoc,
  title={DISOC: A Trustworthy Decentralized Oracle Architecture With Strong Privacy Preservation for IIoT Data Sharing},
  author={Xiong, Chenxi and Yang, Ting and Wang, Xiangyu and Cao, Yunwei and Mao, Gang},
  journal={IEEE Internet of Things Journal},
  year={2024},
  publisher={IEEE}
}
```
