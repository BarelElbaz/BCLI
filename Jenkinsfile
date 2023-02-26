pipeline{
    agent any
    tools{
        go '1.20.1'
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