```
					╔═╗╦═╗╔═╗╦
					╠═╣╠╦╝╠═╝║
					╩ ╩╩╚═╩  ╩
```

<div align="center">

Basic network scanner based on nmap for Raspberry Pi. 

<img src="./assets/demo.gif"/>


</div>

---

## Install 

**Note**: Prebuilt binaries (32-bit) doesn't require Go.

- Prerequisite : arpi is built on top of nmap. 
```bash
sudo apt update && sudo apt install nmap
```

- Installer : 
```bash 
curl -sSL https://raw.githubusercontent.com/PierreKieffer/arpi/master/install/install_arpi32.sh | bash
```

## Run 
To take advantage of deep network scan, arpi must run as root user. 
```bash
sudo arpi
```
**Note**: Default network is `192.168.1.0/24`. To scan another network : 
```
sudo arpi -net=192.168.0.0/24
```




