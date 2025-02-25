
def checkEnv(parallelCreateContainers) {
    return [
        {
            stage('[env] phase2: create containers') {
                script {
                    parallel parallelCreateContainers
                }
            }
        }  // defined inside a closure
    ]
}

return this
