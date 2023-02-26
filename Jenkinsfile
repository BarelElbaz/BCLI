pipeline{
    agent any
    tools{
        go
    }
    stages{
        stage("Build"){
            steps{
                sh "go build"
            }
        }
        stage("Test"){
            steps{
                sh "go test"
            }
        }
    }
}