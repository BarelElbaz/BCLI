pipeline{
    agent any
    tools{
        go '1.19'
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