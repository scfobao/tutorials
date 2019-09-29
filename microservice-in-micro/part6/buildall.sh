#!/bin/bash
typeset -l SVC
SVC=$1

PROJECTNAME="micro"
SERVICELIST=("config-grpc-srv" "auth" "inventory-srv" "user-srv" \
"user-web"  "orders-srv" "orders-web" \
"payment-srv" "payment-web" \
)

servicePubed(){

    FILENAME=service_pubed.txt
    rm $FILENAME
    
    docker service ls --filter name=$1 --format "{{.ID}}" > $FILENAME
    FILESIZE=$(stat -c%s "$FILENAME")
    if [ $FILESIZE != "0" ]; then
      return 1
    else
      return 0
    fi 
    
}

deployService(){

        srvname=$1 
        echo "start go build $srvname ..."    
        go build -o ./$srvname/$srvname ./$srvname/main.go ./$srvname/plugin.go
        echo "$srvname build success."

        docker build ./$srvname -t scfobao/${PROJECTNAME}_${srvname}:latest
        docker push scfobao/${PROJECTNAME}_${srvname}:latest
        # servicePubed ${PROJECTNAME}_${srvname}
        # if [ $? == 0 ]; then
            
        #     if [ "$srvname" == "config-grpc-srv" ]; then
        #          echo "$srvname with port publish"
        #         docker service create --name=${PROJECTNAME}_${srvname} --network $NETWORKNAME --publish 9600:9600 scfobao/${PROJECTNAME}_${srvname}:latest 
        #     else
        #         echo "$srvname without port publish"
        #         docker service create --name=${PROJECTNAME}_${srvname} --network $NETWORKNAME scfobao/${PROJECTNAME}_${srvname}:latest 
        #     fi
        # else       
        
        # docker service update  ${PROJECTNAME}_${srvname}
        # fi
  
}

deployAllService(){
    for element in ${SERVICELIST[@]}
    #也可以写成for element in ${array[*]}
    do
      deployService $element
    done
}

export GOOS=linux
export GOARCH=amd64

if [ $SVC == "all" ]; then   
    deployAllService
else    
    deployService $SVC
fi

