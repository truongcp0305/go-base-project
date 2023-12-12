pipeline {
    agent any

    environment {
        GO_VERSION = '1.19.8'
        APP_NAME = 'go-base-project'
        PORT = '1234'
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
                    //sh 'export PATH=$PATH:/usr/local/go/bin'
                    sh "go version"
                    sh "go build -o ${APP_NAME}"
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    sh "./${APP_NAME} &"
                }
            }
        }
    }

    post {
        always {
            script {
                echo "App deployed at http://localhost:${PORT}"
            }
        }
    }
}
