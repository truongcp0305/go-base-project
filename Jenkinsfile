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
                    bat 'go version'
                    bat 'docker build -t go-base .'
                    //bat 'docker push go-base'
                    //bat 'docker tag go-base truong/go-base'
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    bat 'sh "docker login -u truongcp0305 -p 123456"'
                    //sh "./${APP_NAME} &"
                    bat 'kubectl apply -f k8s/app_deployment.yaml'
                    bat 'kubectl port-forward deployment/go-base 1234:1234 &'
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
