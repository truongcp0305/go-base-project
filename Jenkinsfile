pipeline {
    agent any

    environment {
        GO_VERSION = '1.19.8'
        APP_NAME = 'go-base-project'
        PORT = '1234'
        PATH = '${PATH}'
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    git 'https://github.com/truongcp0305/go-base-project'
                }
            }
        }

        // stage("Set Path Golang"){
        //     steps {
        //         script {
        //             sh 'export PATH=$PATH:/usr/local/go/bin'
        //         }
        //     }
        // }

        stage('Build') {
            steps {
                script {
                    echo 'PATH is: $PATH'
                    // sh 'wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz'
                    // sh 'rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz'
                    sh 'export PATH=$PATH'
                    sh 'go version'
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
