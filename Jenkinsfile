pipeline {
	agent none

    stages {
		stage('[2] create containers') {
			steps {
				script {
					def agents = ['agent-1', 'agent-2', 'agent-0']  // Add more agents here if needed
                    def testStages = [:]  // Groovy Map (key-value pairs), parallel code blocks to run

                    for (agentName in agents) {
						testStages["[2] ${agentName} code"] = {
							node(agentName) {
								stage("Test on ${agentName}") {
									checkout([
                                        $class: 'GitSCM',
                                        branches: [[name: 'main']],
                                        userRemoteConfigs: [[url: 'https://github.com/JunchenLu417/k8s-from-scratch.git']]
                                    ])

                                    sh 'go version'
                                    sh 'go mod tidy || true'
                                    sh 'sudo go test -v github.com/JunchenLu417/k8s-from-scratch/init/test -run TestCreateContainer'
                                }
                            }
                        }
                    }

                    parallel testStages
                }
            }
        }
    }
}
