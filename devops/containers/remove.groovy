
// Assign jobs to agents, master/worker can have different jobs
def jobsEnvCheck = [
    "agent-0": "JobMaster",
    "agent-1": "JobWorker",
    "agent-2": "JobWorker"
]

// Convert jobs definition to parallel execution format
def parallelRemoveContainers = jobsEnvCheck.collectEntries { agentLabel, job ->
    ["${agentLabel}" : {
		stage("Stage: ${job} on ${agentLabel}") {
			node(agentLabel) {  // designate the node to execute the job

				// work that all the nodes should do
				checkout([
					$class: 'GitSCM',
					branches: [[name: 'main']],
					userRemoteConfigs: [[url: 'https://github.com/JunchenLu417/k8s-from-scratch.git']]
				])
				sh 'sudo go test -v -count=1 github.com/JunchenLu417/k8s-from-scratch/init/test -run TestRemoveContainer'
			}
		}
	}]
}

return parallelRemoveContainers
