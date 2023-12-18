pipeline {
    agent any

    environment {
        GO_VERSION = '1.19.8'
        APP_NAME = 'go-base-project'
        PORT = '1234'
        PATH = "/usr/local/go/bin:$PATH"
        KUBECONFIG = 'C:/Users/ASUS/.kube/config'
    }

    stages {
        // stage('Checkout') {
        //     steps {
        //         script {
        //             git 'https://github.com/truongcp0305/go-base-project'
        //         }
        //     }
        // }

        stage('Check OS') {
            steps {
                script {
                    def os = checkOs()
                    echo "OS is  ${os}"
                }
            }
        }

        // stage('Build') {
        //     steps {
        //         script {
        //             //sh 'echo $PATH'
        //             //sh 'go version'
        //             //sh "go build -o ${APP_NAME}"
        //             bat 'go version'
        //             bat 'docker build -t go-base .'
        //             bat "docker tag go-base localhost:5000/go-base:${env.BUILD_NUMBER}"
        //             bat "docker push localhost:5000/go-base:${env.BUILD_NUMBER}"
        //         }
        //     }
        // }

    //     stage('Deploy') {
    //         steps {
    //             script {
    //                 withCredentials([usernamePassword(credentialsId: 'myregistrykey2', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
    //                     bat 'echo $PASSWORD | docker login -u $USERNAME --password-stdin localhost:5000'
    //                     bat "kubectl apply -f k8s/app_deployment.yaml"
    //                 }
    //                 //sh "./${APP_NAME} &"
    //             }
    //         }
    //     }
    // }

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
}
