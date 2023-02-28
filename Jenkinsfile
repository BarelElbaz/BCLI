@Library('slib@main')_

pipeline{
    agent any

    tools{
        go '1.20.1'
    }

    stages{
        stage("Build"){
            steps{
                goBuild("bcli")
            }
        }
        stage("Test"){
            steps{
                goTest()
            }
        }
    }
}