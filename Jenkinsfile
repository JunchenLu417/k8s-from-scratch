pipeline {
	agent none  // No single default agent; jobs are assigned dynamically

    stages {
		stage('Launch Test...') {
			agent any  // Run on any available agent
            steps {
				echo '[Jenkins] launch the test set for k8s-from-scratch project :)'

				script {
					// Load the external Groovy script inside a script block
					parallelCreateContainers = load 'devops/containers/create.groovy'
					parallelRemoveContainers = load 'devops/containers/remove.groovy'
					checkEnv = load 'devops/env_setup.groovy'
				}
            }
        }

        stage('Execute Environment Setup Check') {
			agent any
			steps {
				script {
					def checkEnvStages = checkEnv.checkEnv(parallelCreateContainers, parallelRemoveContainers)
					for (stageBlock in checkEnvStages) {
						stageBlock()
					}
				}
			}
		}


    }
}