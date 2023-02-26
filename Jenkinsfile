pipeline{

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