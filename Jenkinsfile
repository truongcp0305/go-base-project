pipeline {
    agent any

    environment {
        GO_VERSION = '1.19.8'
        APP_NAME = 'go-base-project'
        PORT = '1234'
        PATH = '/usr/local/go/bin:$PATH'
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

        stage('Build') {
            steps {
                script {
                    //sh 'echo $PATH'
                    //sh 'go version'
                    //sh "go build -o ${APP_NAME}"
                    bat 'kubectl config get-contexts'
                    //bat 'go version'
                   // bat 'docker build -t truong/go-base .'
                    //bat 'go build -o ${APP_NAME}'
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    //sh "./${APP_NAME} &"
                    bat 'kubectl apply -f k8s/app_deployment.yaml'
                    //bat 'go run main.go'
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
