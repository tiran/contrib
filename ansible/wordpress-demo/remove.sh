kubectl delete rc frontend mysql
kubectl delete svc frontend mysql
kubectl delete pvc myclaim-db myclaim-web
kubectl delete pv pv5gdb pv5gweb

rm -rf /srv/kube-vol/wordpress/web/*
rm -rf /srv/kube-vol/wordpress/db/*

