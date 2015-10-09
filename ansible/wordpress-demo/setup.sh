#!/bin/sh
set -e

rm -rf /srv/kube-vol/wordpress/web/*
rm -rf /srv/kube-vol/wordpress/db/*

mkdir -p /srv/kube-vol/wordpress/web
mkdir -p /srv/kube-vol/wordpress/db

kubectl create -f pv-nfs-db.yaml
kubectl create -f claim-db.yaml
kubectl create -f mysql-custodia-rc.yaml
kubectl create -f mysql-service.yaml

sleep 5

kubectl create -f pv-nfs-web.yaml
kubectl create -f claim-web.yaml
kubectl create -f wordpress-custodia-rc.yaml
kubectl create -f wordpress-service.yaml

sleep 5

kubectl get pv,pvc
kubectl get rc
kubectl get po
kubectl get svc

