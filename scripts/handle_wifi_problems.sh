# Check internet and wifi

# watch dmesg live
dmesg -w

# check if interface cconfigured to be UP, and if it's down
ip l
ip l | grep $DEV

# list if device is found, and which drivrer is used
lspci -nnk | grep -A3 -i -E "wireless|network"

# scan with iw for networks
iw $DEV scan
# connect
iw -w $DEV "$BSSID"

# remove and reload driver
lsmod | grep $DRV
modprobe -r $DRV
modprobe $DRV

# Now try to reconnect
