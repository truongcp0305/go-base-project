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
                    git 'https://github.com/truongcp0305/go-base-project'
                }
            }
        }

        stage('Build') {
            steps {
                script {
                    sh 'echo $PATH'
                    sh 'go version'
                    sh "go build -o ${APP_NAME}"
                }
            }
        }

        stage('Deploy') {
            steps {
                // script {
                //     sh "./${APP_NAME}"
                // }
                script{
                    def result = sh(script: "./${APP_NAME}", returnStatus: true)
                    if(result == 0) {
                        currentBuild.result = 'SUCCESS'
                    }else{
                        currentBuild.result = 'FAILURE'
                        error "Deploy failed"
                    }
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
