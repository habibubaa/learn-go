pipeline {
    
    agent any
    
    environment {
        imageName = "${env.deploymentName}"
        registryCredentials = "docker_credential_lokal"
        registry = "http://10.61.4.123:8087/"
        dockerImage = ''
        deploymentName = 'cekpn'
        namespaceKube = 'cekpn'
        containerPort = 80 
        targetPort = 30203
    }
    
    stages {
        stage('Checkout Source') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: 'main']], extensions: [],
                    userRemoteConfigs: [[credentialsId: 'git_lab_tes_ajz',
                    url: 'http://10.61.4.35/mbgenardo0517/learn_beego.git']]])
                }
            }
        stage('Building image') {
          steps{
            script {
              dockerImage = docker.build imageName
            }
          }
        }
        stage('Uploading to Nexus') {
         steps{  
             script {
                 docker.withRegistry(registry, registryCredentials ) {
                 dockerImage.push('latest')
              }
            }
          }
        }
         stage('Deploy To Kubernetes') {
             steps
             {
                sshagent(['ssh_key_kube']) {
                       script{
                           
                           sh "sed -i 's+nama-namespace: ganti-nama+name: ${namespaceKube}+g' kubernetes.yaml"
                           sh "sed -i 's+ganti-namespace: ganti-nama+namespace: ${namespaceKube}+g' kubernetes.yaml"
                           sh "sed -i 's+ganti-nama-app+${env.deploymentName}+g' kubernetes.yaml"
                           sh "sed -i 's+ganti-port-container+${containerPort}+g' kubernetes.yaml"
                           sh "sed -i 's+ganti-target-port+${targetPort}+g' kubernetes.yaml"
                           sh "sed -i 's+ganti-image-name+${imageName}+g' kubernetes.yaml"
                           sh "cat kubernetes.yaml"
                           try{
                            sh "ssh ubuntu@10.61.4.250 kubectl create secret docker-registry ajzdockerkey --namespace=${namespaceKube} --docker-server=10.61.4.123:8087 --docker-username=jazuli0701 --docker-password=Taik3ring"
                           }catch(error){
                               
                           }
                           try{
                                sh "ssh ubuntu@10.61.4.250 kubectl delete deployments ${env.deploymentName} --namespace=${env.namespaceKube}"
                               //sh "ssh ubuntu@10.61.4.250 kubectl delete ${namespaceKube} deployment ${env.deploymentName}" 
                               //sh "ssh ubuntu@10.61.4.250 kubectl apply -f /home/ubuntu/${env.JOB_NAME}/kubernetes.yaml --validate=false"
                           }catch(error){
                               //sh "ssh ubuntu@10.61.4.250 kubectl apply -f /home/ubuntu/${env.JOB_NAME}/kubernetes.yaml --validate=false"
                           }
                          
                           sh "cat kubernetes.yaml | ssh ubuntu@10.61.4.250 kubectl apply -f - --validate=false"
                       }   
                }
             }
         }
    }
}
