# Clone latest repo from github Open the project Root directory lenovo-preload-fabric
# Go to local folder inide network
cd network/local
# Downlaod fabric images and binaries required and Generate certificates using bootstrap file
./bootstrap.sh -r
# Create Channel add peers to channel
./admin.sh
  # throws error PKCS11 library path must be specified , run the following from terminal to set env variables available to connect to UKC server
  mkdir -p store/softhsm
  echo directories.tokendir=$PWD/store/softhsm > softhsm.conf
  export SOFTHSM2_CONF=$PWD/softhsm.conf
  softhsm2-util --init-token --slot 0 --label "ForFabric" --so-pin 1234 --pin 98765432
  # Set PKCS11 Library 
  export CRYPTO_PKCS11_LIB=/usr/lib/libekmpkcs11.so
  export CRYPTO_PKCS11_SLOT=0
  # Set CRYPTO_PKCS11_PIN which connects client to server
  export CRYPTO_PKCS11_PIN=Ch@iny@rd123  
  
# Go to api directory from present directory
cd ../.. && cd api
# Run the node app
node app.js


