pipeline{
    agent any

    tools{
        go '1.20.1'
    }

    stages{
        stage("Build"){
            steps{
                echo(" Building bcli ")
                sh "go build -o bcli"
            }
        }
        stage("Test"){
            steps{
                echo(" Testing bcli")
                sh "go test"
            }
        }
    }
}