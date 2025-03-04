
def checkEnv(parallelCreateContainers, parallelRemoveContainers) {
    return [
        {
            stage('[env] phase1: create containers') {
                script {
                    parallel parallelCreateContainers
                }
            }
        },  // defined inside a closure
        {
            stage('[env] phase3: remove containers') {
                script {
                    parallel parallelRemoveContainers
                }
            }
        }
    ]
}

return this
