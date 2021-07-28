#!/bin/bash
echo " ---------- Downloading arpi ---------- "
wget https://github.com/PierreKieffer/arpi/raw/master/bin/arpi

chmod +x arpi
sudo mv arpi /usr/local/bin
echo " ---------- arpi is installed ---------- "
echo " ---------- usage ---------- "
echo " Note : arpi is built on top nmap, to install nmap :"
echo " sudo apt update && sudo apt install nmap"
echo ""
echo " - default network (192.168.1.0/24) : "
echo "     sudo arpi"
echo ""
echo " - custom network  : "
echo "     sudo arpi -net=192.168.0.0/24"
echo ""
echo "
     -----------------------------
     -        Move around        -
     -----------------------------
     go up               ▲  or 'k'
     go down             ▼  or 'j'
     go to the top       'gg'
     go to the bottom    'G'
     select item         'enter'
     Quit                'q'

"
echo " --------------------------- "


