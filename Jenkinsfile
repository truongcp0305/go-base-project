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
        KUBECONFIG = 'C:/Users/ASUS/.kube/config'
    }

    stages {
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
                        sh "docker build -t ${APP_NAME} ."
                        sh "docker tag ${APP_NAME} localhost:5000/${APP_NAME}:${env.BUILD_NUMBER}"
                        sh "docker push localhost:5000/${APP_NAME}:${env.BUILD_NUMBER}"
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
                        def v = "${env.BUILD_NUMBER} ^| Set-Content k8s\app_deployment2.yaml"
                        withCredentials([usernamePassword(credentialsId: 'myregistrykey2', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                            bat 'echo $PASSWORD | docker login -u $USERNAME --password-stdin localhost:5000'
                            bat 'echo (Get-Content k8s\app_deployment.yaml) ^| ForEach-Object { $_ -replace "{BUILD_NUMBER}", "> makefile.ps1'
                            bat "set /p=${v} <nul >> makefile.ps1" 
                            bat "powershell -ExecutionPolicy Bypass -File script.ps1"
                            bat "kubectl apply -f k8s/app_deployment2.yaml"
                        }
                    }else if (os == "Linux"){
                        withCredentials([usernamePassword(credentialsId: 'myregistrykey2', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                            sh 'echo $PASSWORD | docker login -u $USERNAME --password-stdin localhost:5000'
                            sh "sed 's/{BUILD_NUMBER}/${env.BUILD_NUMBER}/g' app_deployment.yaml > app_deployment2.yaml"
                            sh "kubectl apply -f k8s/app_deployment2.yaml"
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
