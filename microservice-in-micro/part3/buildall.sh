#!/bin/bash
echo "start build auth ..."
export GOOS=linux
go build -o ./auth/auth ./auth/main.go ./auth/plugin.go
echo "auth build success."

#------------------------------
echo "start build inventory-srv ..."
go build -o ./inventory-srv/inventory-srv ./inventory-srv/main.go ./inventory-srv/plugin.go
echo "inventory-srv build success."

#------------------------------
echo "start build orders-srv ..."
go build -o ./orders-srv/orders-srv ./orders-srv/main.go ./orders-srv/plugin.go
echo "orders-srv build success."

#------------------------------

echo "start build orders-web ..."
go build -o ./orders-web/orders-web ./orders-web/main.go ./orders-web/plugin.go
echo "orders-web build success."

#------------------------------

echo "start build payment-srv ..."
go build -o ./payment-srv/payment-srv ./payment-srv/main.go ./payment-srv/plugin.go
echo "payment-srv build success."

#----------------------------------------------------------------------------

echo "start build payment-web ..."
go build -o ./payment-web/payment-web ./payment-web/main.go ./payment-web/plugin.go
echo "payment-web build success."

#----------------------------------------------------------------------------


echo "start build user-srv ..."
go build -o ./user-srv/user-srv ./user-srv/main.go ./user-srv/plugin.go
echo "user-srv build success."

#------------------------------


echo "start build user-web ..."
go build -o ./user-web/user-web ./user-web/main.go ./user-web/plugin.go
echo "user-web build success."

#------------------------------