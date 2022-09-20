pipeline {
    agent any
    tools {
        go '1.19'
    }
    stages {
        stage ('Initialize') {
            steps {
                sh '''
                    echo "GOPATH = ${GOPATH}"
                '''
            }
        }

        stage ('Build') {
            steps {
               sh 'go version'
               sh 'go get -u github.com/golang/dep/...'
               sh 'dep init'
               sh 'go build ./...'
            }
            post {
                success {
                    echo 'Build DONE'
                }
            }
        }
      stage ('test') {
            sh 'go vet ./...'
            sh 'go test -cover ./...'
      }
    }
}
