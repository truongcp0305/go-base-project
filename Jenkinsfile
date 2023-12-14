pipeline {
    agent any

    environment {
        GO_VERSION = '1.19.8'
        APP_NAME = 'go-base-project'
        PORT = '1234'
        PATH = "/usr/local/go/bin:$PATH"
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    echo "stage check out"
                    git 'https://github.com/truongcp0305/go-base-project'
                }
            }
        }

        stage('Build') {
            steps {
                script {
                    echo "stage build"
                    //sh 'echo $PATH'
                    sh 'go version'
                    //sh "go build -o ${APP_NAME}"
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
