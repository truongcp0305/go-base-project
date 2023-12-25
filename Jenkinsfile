    def checkOs(){
        if (isUnix()) {
            def uname = sh script: 'uname', returnStdout: true
            if (uname.startsWith("Darwin")) {
                return "Macos"
            }
            else {
                return "Linux"
            }
        }
        else {
            return "Windows"
        }
    }

pipeline {
    agent any

    environment {
        GO_VERSION = '1.19.8'
        APP_NAME = 'go-base'
        PORT = '1234'
        PATH = "/usr/local/go/bin:$PATH"
        NAMESPACE ='default'
        REGISTRY = 'localhost:5000'
    }

    stages {
        stage ('switch to minikube'){
            steps {
                script {
                    sh 'minikube ssh'
                }
            }
        }

        stage('Checkout') {
            steps {
                script {
                    git 'https://github.com/truongcp0305/go-base-project'
                }
            }
        }

        stage('Check OS') {
            steps {
                script {
                    def os = checkOs()
                    echo "OS is  ${os}"
                }
            }
        }

        stage('Build') {
            steps {
                script {
                    def os = checkOs()
                    
                    if (os == 'Windows'){
                        bat 'go version'
                        bat "docker build -t ${APP_NAME} ."
                        bat "docker tag ${APP_NAME} localhost:5000/${APP_NAME}:${env.BUILD_NUMBER}"
                        bat "docker push localhost:5000/${APP_NAME}:${env.BUILD_NUMBER}"
                    }else if (os == 'Linux'){
                        sh 'echo $PATH'
                        sh 'go version'
                        //sh 'eval $(minikube docker-env)'
                        //sh "docker build -t ${APP_NAME} ."
                        sh "minikube build -t ${APP_NAME}:${env.BUILD_NUMBER} ."
                        // sh "docker tag ${APP_NAME} ${APP_NAME}:${env.BUILD_NUMBER}"
                        // withCredentials([usernamePassword(credentialsId: 'myregistrykey', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]){
                        //     sh 'eval $(minikube docker-env)'
                        //     sh 'echo $PASSWORD | docker login -u $USERNAME --password-stdin localhost:5000'
                        //     sh "docker push ${REGISTRY}/${APP_NAME}:${env.BUILD_NUMBER}"
                        // }
                        //sh "minikube image load ${APP_NAME}:${env.BUILD_NUMBER} --daemon=true"
                    }else{
                        echo "OS not supported"
                    }
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    def os = checkOs()
                    if (os == "Windows"){
                        def p1 = '(Get-Content app_deployment.yaml) | ForEach-Object { $_ -replace "{BUILD_NUMBER}", '
                        def ver = "${env.BUILD_NUMBER} "
                        def p3 = '} | Set-Content app_deployment2.yaml'
                        writeFile file: "makefile.ps1", text: "${p1} ${ver} ${p3}" 
                        bat "powershell -ExecutionPolicy Bypass -File makefile.ps1"
                        withCredentials([usernamePassword(credentialsId: 'myregistrykey2', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                            bat 'echo $PASSWORD | docker login -u $USERNAME --password-stdin localhost:5000'
                            bat "kubectl apply -f app_deployment2.yaml"
                        }
                    }else if (os == "Linux"){
                        sh "sed 's/{BUILD_NUMBER}/${env.BUILD_NUMBER}/g' app_deployment.yaml > app_deployment2.yaml"
                        // sh 'eval $(minikube -p minikube docker-env)'
                        // withCredentials([usernamePassword(credentialsId: 'myregistrykey', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                        //     sh 'echo $PASSWORD | docker login -u $USERNAME --password-stdin localhost:5000'
                        // }
                        withCredentials([file(credentialsId: 'minikube', variable: 'KUBECONFIG')]) {
                            sh "kubectl --kubeconfig=${KUBECONFIG} --namespace=${NAMESPACE} apply -f app_deployment2.yaml"
                        }
                    }else{
                        echo "OS not supported"
                    }
                }
            }
        }
    }


    post {
        always {
            echo "Build # ${env.BUILD_NUMBER}" 
        }
        success{
            echo "App deployed at http://localhost:${PORT}"
        }
        failure{
            echo "pipeline execution failed"
        }
    }

}
